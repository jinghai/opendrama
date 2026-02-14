package handlers

import (
	"net/http"

	"github.com/drama-generator/backend/pkg/ai/newapi"
	"github.com/gin-gonic/gin"
)

// NewAPIHandler NewAPI统一接口处理器
type NewAPIHandler struct {
	client *newapi.NewAPIClient
}

// NewNewAPIHandler 创建处理器
func NewNewAPIHandler(client *newapi.NewAPIClient) *NewAPIHandler {
	return &NewAPIHandler{client: client}
}

// GenerateTextRequest 文本生成请求
type GenerateTextRequest struct {
	Model       string `json:"model" binding:"required"`
	Prompt      string `json:"prompt" binding:"required"`
	SystemPrompt string `json:"system_prompt"`
}

// GenerateImageRequest 图像生成请求
type GenerateImageRequest struct {
	Model  string `json:"model" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
	Size   string `json:"size"`
	N      int    `json:"n"`
}

// GenerateText 统一文本生成
// @Summary 统一文本生成接口
// @Description 通过NewAPI统一接口生成文本
// @Tags NewAPI
// @Accept json
// @Produce json
// @Param request body GenerateTextRequest true "请求参数"
// @Success 200 {object} newapi.AIResponse
// @Router /api/v1/newapi/text [post]
func (h *NewAPIHandler) GenerateText(c *gin.Context) {
	var req GenerateTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aiReq := &newapi.AIRequest{
		Model:        req.Model,
		Prompt:       req.Prompt,
		SystemPrompt: req.SystemPrompt,
	}

	resp, err := h.client.GenerateText(c.Request.Context(), aiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GenerateImage 统一图像生成
// @Summary 统一图像生成接口
// @Description 通过NewAPI统一接口生成图像
// @Tags NewAPI
// @Accept json
// @Produce json
// @Param request body GenerateImageRequest true "请求参数"
// @Success 200 {object} newapi.AIResponse
// @Router /api/v1/newapi/image [post]
func (h *NewAPIHandler) GenerateImage(c *gin.Context) {
	var req GenerateImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aiReq := &newapi.AIRequest{
		Model:  req.Model,
		Prompt: req.Prompt,
	}

	resp, err := h.client.GenerateImage(c.Request.Context(), aiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetStats 获取服务商统计
// @Summary 获取统计信息
// @Description 获取各服务商的调用统计
// @Tags NewAPI
// @Produce json
// @Success 200 {object} map[string]newapi.ProviderStats
// @Router /api/v1/newapi/stats [get]
func (h *NewAPIHandler) GetStats(c *gin.Context) {
	stats := h.client.LoadBalancer.GetStats()
	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// UpdateConfig 更新配置
// @Summary 更新NewAPI配置
// @Description 更新负载均衡配置
// @Tags NewAPI
// @Accept json
// @Produce json
// @Param request body newapi.NewAPIConfig true "配置"
// @Success 200 {object} gin.H
// @Router /api/v1/newapi/config [put]
func (h *NewAPIHandler) UpdateConfig(c *gin.Context) {
	var config newapi.NewAPIConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.client.BaseURL = config.BaseURL
	h.client.APIKey = config.APIKey
	h.client.SetProviders(config.Providers)

	c.JSON(http.StatusOK, gin.H{"message": "config updated"})
}
