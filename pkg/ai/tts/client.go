package tts

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ==================== TTS客户端接口 ====================

// TTSClient TTS客户端接口
type TTSClient interface {
	Generate(ctx context.Context, req *TTSRequest) (*TTSResponse, error)
	GetVoices(ctx context.Context) ([]Voice, error)
}

// TTSRequest TTS请求
type TTSRequest struct {
	Text     string            `json:"text"`
	Voice    string            `json:"voice"`
	Speed    float64           `json:"speed"`    // 0.5 - 2.0
	Pitch    float64           `json:"pitch"`    // 0.5 - 2.0
	Format   string            `json:"format"`   // mp3, wav, ogg
	OutputPath string          `json:"output_path"`
	Options  map[string]interface{} `json:"options"`
}

// TTSResponse TTS响应
type TTSResponse struct {
	Success    bool   `json:"success"`
	AudioData  string `json:"audio_data,omitempty"` // base64 encoded
	FilePath   string `json:"file_path,omitempty"`
	Duration   int    `json:"duration,omitempty"` // milliseconds
	Error      string `json:"error,omitempty"`
	Provider   string `json:"provider"`
	Latency    time.Duration `json:"latency"`
}

// Voice 语音角色
type Voice struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Language    string `json:"language"`
	Gender      string `json:"gender"`
	Style       string `json:"style"`
	Provider    string `json:"provider"`
}

// AzureTTSClient Azure TTS客户端
type AzureTTSClient struct {
	APIKey     string
	Region     string
	Endpoint   string
	HTTPClient *http.Client
}

// NewAzureTTSClient 创建Azure TTS客户端
func NewAzureTTSClient(apiKey, region string) *AzureTTSClient {
	return &AzureTTSClient{
		APIKey:     apiKey,
		Region:     region,
		Endpoint:   fmt.Sprintf("https://%s.tts.speech.microsoft.com", region),
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
	}
}

// AzureTTSRequest Azure TTS请求
type AzureTTSRequest struct {
	Input     []struct {
		Text string `json:"text"`
	} `json:"input"`
	Voice struct {
		Language     string `json:"language"`
		Name         string `json:"name"`
		Role         string `json:"role"`
		Neural       string `json:"@type"`
		Style        string `json:"style"`
		StyleDegree  string `json:"styleDegree"`
	} `json:"voice"`
	OutputFormat struct {
		Container string `json:"container"`
		Format    string `json:"format"`
		Profile   string `json:"profile"`
		Sample    string `json:"sample"`
	} `json:"outputFormat"`
	Synthesis struct {
		BatchMode string `json:"batchMode"`
	} `json:"synthesis"`
}

// Generate Azure TTS生成
func (c *AzureTTSClient) Generate(ctx context.Context, req *TTSRequest) (*TTSResponse, error) {
	start := time.Now()

	// 构建请求
	azureReq := AzureTTSRequest{}
	azureReq.Input = []struct {
		Text string `json:"text"`
	}{{Text: req.Text}}
	
	azureReq.Voice.Language = "zh-CN"
	azureReq.Voice.Name = req.Voice
	azureReq.Voice.Neural = "Neural"
	azureReq.Voice.Style = "neutral"
	azureReq.Voice.StyleDegree = "1.0"
	
	azureReq.OutputFormat.Container = "audio"
	azureReq.OutputFormat.Format = "raw-24khz-16bit-mono-pcm"
	azureReq.OutputFormat.Profile = "audio-24khz-96kbitrate-mono-mp3"
	azureReq.OutputFormat.Sample = "24000"
	
	azureReq.Synthesis.BatchMode = "CRIS"

	jsonBody, _ := json.Marshal(azureReq)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", 
		c.Endpoint+"/cognitiveservices/v1", 
		io.NopCloser(bytes.NewBuffer(jsonBody)))
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "azure"}, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Ocp-Apim-Subscription-Key", c.APIKey)
	httpReq.Header.Set("X-Microsoft-OutputFormat", "audio-24khz-96kbitrate-mono-mp3")

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "azure"}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return &TTSResponse{Success: false, Error: fmt.Sprintf("Azure TTS error: %s", string(body)), Provider: "azure"}, fmt.Errorf("Azure TTS error: %d", resp.StatusCode)
	}

	// 读取音频数据
	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "azure"}, err
	}

	// 保存到文件
	var outputPath string
	if req.OutputPath != "" {
		outputPath = req.OutputPath
	} else {
		outputPath = filepath.Join("data", "tts", fmt.Sprintf("%d.mp3", time.Now().Unix()))
	}

	// 确保目录存在
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	
	// 写入文件
	if err := os.WriteFile(outputPath, audioData, 0644); err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "azure"}, err
	}

	return &TTSResponse{
		Success:    true,
		FilePath:   outputPath,
		AudioData:  base64.StdEncoding.EncodeToString(audioData),
		Provider:   "azure",
		Latency:    time.Since(start),
	}, nil
}

// GetVoices 获取可用语音列表
func (c *AzureTTSClient) GetVoices(ctx context.Context) ([]Voice, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", 
		c.Endpoint+"/cognitiveservices/voices/list", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var voices []struct {
		ID             string `json:"id"`
		LocalName      string `json:"localName"`
		Language       string `json:"language"`
		Gender         string `json:"gender"`
		VoiceType      string `json:"voiceType"`
		LanguageStatus string `json:"languageStatus"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&voices); err != nil {
		return nil, err
	}

	result := make([]Voice, 0, len(voices))
	for _, v := range voices {
		result = append(result, Voice{
			ID:       v.ID,
			Name:     v.LocalName,
			Language: v.Language,
			Gender:   v.Gender,
			Style:    v.VoiceType,
			Provider: "azure",
		})
	}

	return result, nil
}

// ==================== 阿里云TTS客户端 ====================

// AlibabaTTSClient 阿里云TTS客户端
type AlibabaTTSClient struct {
	APIKey     string
	AppKey     string
	Endpoint   string
	HTTPClient *http.Client
}

// NewAlibabaTTSClient 创建阿里云TTS客户端
func NewAlibabaTTSClient(apiKey, appKey string) *AlibabaTTSClient {
	return &AlibabaTTSClient{
		APIKey:     apiKey,
		AppKey:     appKey,
		Endpoint:   "https://nls-gateway-cn-shanghai.aliyuncs.com/stream/v1/tts",
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
	}
}

// Generate 阿里云TTS生成
func (c *AlibabaTTSClient) Generate(ctx context.Context, req *TTSRequest) (*TTSResponse, error) {
	start := time.Now()

	// 构建请求参数
	params := map[string]string{
		"appkey":    c.AppKey,
		"token":     c.APIKey,
		"text":      req.Text,
		"format":    "mp3",
		"voice":     req.Voice,
		"speech_rate": fmt.Sprintf("%d", int(req.Speed*100-100)),
		"pitch_rate": fmt.Sprintf("%d", int(req.Pitch*100-100)),
	}

	// 构建URL
	url := c.Endpoint + "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}

	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "alibaba"}, err
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "alibaba"}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return &TTSResponse{Success: false, Error: fmt.Sprintf("Alibaba TTS error: %s", string(body)), Provider: "alibaba"}, fmt.Errorf("Alibaba TTS error: %d", resp.StatusCode)
	}

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return &TTSResponse{Success: false, Error: err.Error(), Provider: "alibaba"}, err
	}

	outputPath := filepath.Join("data", "tts", fmt.Sprintf("%d.mp3", time.Now().Unix()))
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	os.WriteFile(outputPath, audioData, 0644)

	return &TTSResponse{
		Success:    true,
		FilePath:   outputPath,
		AudioData:  base64.StdEncoding.EncodeToString(audioData),
		Provider:   "alibaba",
		Latency:    time.Since(start),
	}, nil
}

// GetVoices 获取可用语音列表
func (c *AlibabaTTSClient) GetVoices(ctx context.Context) ([]Voice, error) {
	// 预定义常用语音
	return []Voice{
		{ID: "xiaoyun", Name: "云小蜜", Language: "zh-CN", Gender: "female", Style: "chat", Provider: "alibaba"},
		{ID: "xiaogang", Name: "钢小蛋", Language: "zh-CN", Gender: "male", Style: "chat", Provider: "alibaba"},
		{ID: "ruoxi", Name: "若雪", Language: "zh-CN", Gender: "female", Style: "story", Provider: "alibaba"},
		{ID: "shanshan", Name: "杉杉", Language: "zh-CN", Gender: "female", Style: "young", Provider: "alibaba"},
	}, nil
}

// ==================== 统一TTS服务 ====================

// TTSService 统一TTS服务
type TTSService struct {
	Clients    map[string]TTSClient
	HTTPClient *http.Client
}

// NewTTSService 创建TTS服务
func NewTTSService() *TTSService {
	return &TTSService{
		Clients:    make(map[string]TTSClient),
		HTTPClient: &http.Client{Timeout: 120 * time.Second},
	}
}

// RegisterClient 注册TTS客户端
func (s *TTSService) RegisterClient(provider string, client TTSClient) {
	s.Clients[provider] = client
}

// Generate 生成语音
func (s *TTSService) Generate(ctx context.Context, provider, voice, text string, speed, pitch float64, outputPath string) (*TTSResponse, error) {
	client, ok := s.Clients[provider]
	if !ok {
		errMsg := fmt.Sprintf("provider %s not found", provider)
		return &TTSResponse{Success: false, Error: errMsg, Provider: provider}, fmt.Errorf(errMsg)
	}

	req := &TTSRequest{
		Text:       text,
		Voice:      voice,
		Speed:      speed,
		Pitch:      pitch,
		OutputPath: outputPath,
	}

	return client.Generate(ctx, req)
}

// GetVoices 获取所有可用语音
func (s *TTSService) GetVoices(ctx context.Context, provider string) ([]Voice, error) {
	client, ok := s.Clients[provider]
	if !ok {
		return nil, fmt.Errorf("provider %s not found", provider)
	}
	return client.GetVoices(ctx)
}

// GetAllVoices 获取所有提供商的语音
func (s *TTSService) GetAllVoices(ctx context.Context) map[string][]Voice {
	result := make(map[string][]Voice)
	for provider, client := range s.Clients {
		voices, _ := client.GetVoices(ctx)
		result[provider] = voices
	}
	return result
}

// GetProviders 获取所有提供商
func (s *TTSService) GetProviders() []string {
	providers := make([]string, 0, len(s.Clients))
	for k := range s.Clients {
		providers = append(providers, k)
	}
	return providers
}
