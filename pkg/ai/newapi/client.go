package newapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ==================== 统一AI客户端接口 ====================

// NewAPIClient 统一AI客户端
type NewAPIClient struct {
	BaseURL      string
	APIKey       string
	LoadBalancer *LoadBalancer
	HTTPClient   *http.Client
	mu           sync.RWMutex
}

// NewAPIConfig NewAPI配置
type NewAPIConfig struct {
	BaseURL      string           `yaml:"base_url" json:"base_url"`
	APIKey       string           `yaml:"api_key" json:"api_key"`
	LoadBalancer LoadBalancerConfig `yaml:"load_balancer" json:"load_balancer"`
}

// LoadBalancerConfig 负载均衡器配置
type LoadBalancerConfig struct {
	Strategy   string            `yaml:"strategy" json:"strategy"` // "round-robin", "weighted", "least-cost"
	Providers  []ProviderConfig  `yaml:"providers" json:"providers"`
}

// LoadBalancer 负载均衡器
type LoadBalancer struct {
	Strategy    string           
	Providers  []ProviderConfig 
	stats      map[string]*ProviderStats
	mu         sync.RWMutex
}

// ProviderConfig 服务商配置
type ProviderConfig struct {
	Name      string   `yaml:"name" json:"name"`
	BaseURL   string   `yaml:"base_url" json:"base_url"`
	APIKey    string   `yaml:"api_key" json:"api_key"`
	Models    []string `yaml:"models" json:"models"`
	Priority  int      `yaml:"priority" json:"priority"`
	Weight    int      `yaml:"weight" json:"weight"`
	Enabled   bool     `yaml:"enabled" json:"enabled"`
}

// ProviderStats 服务商统计
type ProviderStats struct {
	Name          string        `json:"name"`
	RequestCount int           `json:"request_count"`
	ErrorCount   int           `json:"error_count"`
	SuccessRate  float64       `json:"success_rate"`
	AvgLatency  time.Duration `json:"avg_latency"`
	LastUsed    time.Time     `json:"last_used"`
}

// ==================== 统一请求/响应 ====================

// AIRequest 统一AI请求
type AIRequest struct {
	Provider    string                 `json:"provider,omitempty"`
	Model       string                 `json:"model"`
	Prompt      string                 `json:"prompt"`
	SystemPrompt string               `json:"system_prompt,omitempty"`
	Images      []string              `json:"images,omitempty"`
	Options     map[string]interface{} `json:"options,omitempty"`
}

// AIResponse 统一AI响应
type AIResponse struct {
	Success bool                   `json:"success"`
	Content string                 `json:"content,omitempty"`
	Images  []string              `json:"images,omitempty"`
	Video   string                `json:"video,omitempty"`
	Error   string               `json:"error,omitempty"`
	Usage   map[string]interface{} `json:"usage,omitempty"`
	Latency time.Duration         `json:"latency"`
}

// ==================== 客户端方法 ====================

// NewClient 创建NewAPI客户端
func NewClient(baseURL, apiKey string) *NewAPIClient {
	return &NewAPIClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 120 * time.Second},
		LoadBalancer: &LoadBalancer{
			Strategy: "round-robin",
			stats:   make(map[string]*ProviderStats),
		},
	}
}

// SetProviders 设置服务商列表
func (c *NewAPIClient) SetProviders(providers []ProviderConfig) {
	c.LoadBalancer.mu.Lock()
	defer c.LoadBalancer.mu.Unlock()

	c.LoadBalancer.Providers = providers
	for _, p := range providers {
		c.LoadBalancer.stats[p.Name] = &ProviderStats{Name: p.Name}
	}
}

// GenerateText 统一文本生成
func (c *NewAPIClient) GenerateText(ctx context.Context, req *AIRequest) (*AIResponse, error) {
	start := time.Now()

	// 选择最优provider
	provider, err := c.LoadBalancer.SelectProvider("text")
	if err != nil {
		return &AIResponse{Success: false, Error: err.Error()}, err
	}

	// 构建请求
	url := fmt.Sprintf("%s/v1/chat/completions", provider.BaseURL)
	body := map[string]interface{}{
		"model": req.Model,
		"messages": []map[string]string{
			{"role": "system", "content": req.SystemPrompt},
			{"role": "user", "content": req.Prompt},
		},
	}

	// 发送请求
	resp, err := c.doRequest(ctx, url, provider.APIKey, body)
	if err != nil {
		c.LoadBalancer.RecordError(provider.Name)
		return &AIResponse{Success: false, Error: err.Error()}, err
	}

	c.LoadBalancer.RecordSuccess(provider.Name, time.Since(start))

	return &AIResponse{
		Success: true,
		Content: resp["content"].(string),
		Latency: time.Since(start),
	}, nil
}

// GenerateImage 统一图像生成
func (c *NewAPIClient) GenerateImage(ctx context.Context, req *AIRequest) (*AIResponse, error) {
	start := time.Now()

	provider, err := c.LoadBalancer.SelectProvider("image")
	if err != nil {
		return &AIResponse{Success: false, Error: err.Error()}, err
	}

	url := fmt.Sprintf("%s/v1/images/generations", provider.BaseURL)
	body := map[string]interface{}{
		"model": req.Model,
		"prompt": req.Prompt,
		"n": 1,
		"size": "1024x1024",
	}

	resp, err := c.doRequest(ctx, url, provider.APIKey, body)
	if err != nil {
		c.LoadBalancer.RecordError(provider.Name)
		return &AIResponse{Success: false, Error: err.Error()}, err
	}

	images := resp["images"].([]string)

	c.LoadBalancer.RecordSuccess(provider.Name, time.Since(start))

	return &AIResponse{
		Success: true,
		Images:  images,
		Latency: time.Since(start),
	}, nil
}

// ==================== 负载均衡核心 ====================

// SelectProvider 选择最优服务商
func (lb *LoadBalancer) SelectProvider(serviceType string) (*ProviderConfig, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	var candidates []ProviderConfig
	for _, p := range lb.Providers {
		if p.Enabled && lb.hasModel(p, serviceType) {
			candidates = append(candidates, p)
		}
	}

	if len(candidates) == 0 {
		return nil, fmt.Errorf("no available provider for %s", serviceType)
	}

	switch lb.Strategy {
	case "round-robin":
		return lb.roundRobinSelect(candidates), nil
	case "weighted":
		return lb.weightedSelect(candidates), nil
	case "least-cost":
		return lb.leastCostSelect(candidates), nil
	default:
		return &candidates[0], nil
	}
}

func (lb *LoadBalancer) hasModel(p ProviderConfig, serviceType string) bool {
	for _, m := range p.Models {
		if m == serviceType || m == "all" {
			return true
		}
	}
	return false
}

func (lb *LoadBalancer) roundRobinSelect(candidates []ProviderConfig) *ProviderConfig {
	// 简单轮询
	return &candidates[0]
}

func (lb *LoadBalancer) weightedSelect(candidates []ProviderConfig) *ProviderConfig {
	total := 0
	for _, p := range candidates {
		total += p.Weight
	}
	// 简化实现
	return &candidates[0]
}

func (lb *LoadBalancer) leastCostSelect(candidates []ProviderConfig) *ProviderConfig {
	// 选择成功率最高、延迟最低的
	var best *ProviderConfig
	var bestScore float64 = -1

	for i := range candidates {
		stats := lb.stats[candidates[i].Name]
		if stats == nil {
			return &candidates[i]
		}
		score := stats.SuccessRate / stats.AvgLatency.Seconds()
		if score > bestScore {
			bestScore = score
			best = &candidates[i]
		}
	}

	if best == nil {
		return &candidates[0]
	}
	return best
}

// RecordSuccess 记录成功
func (lb *LoadBalancer) RecordSuccess(name string, latency time.Duration) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if stats, ok := lb.stats[name]; ok {
		stats.RequestCount++
		stats.LastUsed = time.Now()
		stats.AvgLatency = (stats.AvgLatency*time.Duration(stats.RequestCount-1) + latency) / time.Duration(stats.RequestCount)
		stats.SuccessRate = float64(stats.RequestCount-stats.ErrorCount) / float64(stats.RequestCount)
	}
}

// RecordError 记录错误
func (lb *LoadBalancer) RecordError(name string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if stats, ok := lb.stats[name]; ok {
		stats.ErrorCount++
		stats.SuccessRate = float64(stats.RequestCount-stats.ErrorCount) / float64(stats.RequestCount)
	}
}

// ==================== HTTP请求 ====================

func (c *NewAPIClient) doRequest(ctx context.Context, url, apiKey string, body interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, io.NopCloser(bytes.NewBuffer(jsonBody)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetStats 获取统计信息
func (lb *LoadBalancer) GetStats() map[string]*ProviderStats {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	result := make(map[string]*ProviderStats)
	for k, v := range lb.stats {
		result[k] = v
	}
	return result
}
