package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/drama-generator/backend/pkg/ai/tts"
	"github.com/gin-gonic/gin"
)

// TTSHandler TTS处理器
type TTSHandler struct {
	service    *tts.TTSService
	outputPath string
}

// NewTTSHandler 创建TTS处理器
func NewTTSHandler(service *tts.TTSService, outputPath string) *TTSHandler {
	return &TTSHandler{
		service:    service,
		outputPath: outputPath,
	}
}

// GenerateRequest TTS生成请求
type GenerateRequest struct {
	Provider   string  `json:"provider" binding:"required"`
	Voice      string  `json:"voice" binding:"required"`
	Text       string  `json:"text" binding:"required"`
	Speed      float64 `json:"speed"`   // 0.5-2.0, 默认1.0
	Pitch      float64 `json:"pitch"`   // 0.5-2.0, 默认1.0
	Format     string  `json:"format"`  // mp3, wav
	SaveToFile bool    `json:"save_to_file"` // 是否保存到文件
}

// GenerateResponse TTS生成响应
type GenerateResponse struct {
	Success    bool   `json:"success"`
	AudioData  string `json:"audio_data,omitempty"`
	FilePath   string `json:"file_path,omitempty"`
	Duration   int    `json:"duration,omitempty"`
	Provider   string `json:"provider"`
	Voice      string `json:"voice"`
	Latency    string `json:"latency"`
	Error      string `json:"error,omitempty"`
}

// Generate TTS生成
// @Summary TTS语音合成
// @Description 将文本转换为语音
// @Tags TTS
// @Accept json
// @Produce json
// @Param request body GenerateRequest true "TTS请求"
// @Success 200 {object} GenerateResponse
// @Router /api/v1/tts/generate [post]
func (h *TTSHandler) Generate(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置默认值
	if req.Speed == 0 {
		req.Speed = 1.0
	}
	if req.Pitch == 0 {
		req.Pitch = 1.0
	}
	if req.Format == "" {
		req.Format = "mp3"
	}

	// 生成输出路径
	outputPath := ""
	if req.SaveToFile {
		outputPath = filepath.Join(h.outputPath, "tts", req.Provider, req.Voice, 
			string(rune('a'+int(time.Now().Unix()%26)))+".mp3")
	}

	resp, err := h.service.Generate(c.Request.Context(), req.Provider, req.Voice, req.Text, 
		req.Speed, req.Pitch, outputPath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GenerateResponse{
			Success: false,
			Provider: req.Provider,
			Voice:    req.Voice,
			Error:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GenerateResponse{
		Success:    resp.Success,
		AudioData:  resp.AudioData,
		FilePath:   resp.FilePath,
		Duration:   resp.Duration,
		Provider:   resp.Provider,
		Voice:      req.Voice,
		Latency:    resp.Latency.String(),
		Error:      resp.Error,
	})
}

// ListVoicesRequest 获取语音列表请求
type ListVoicesRequest struct {
	Provider string `json:"provider"`
}

// VoiceInfo 语音信息
type VoiceInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Language string `json:"language"`
	Gender   string `json:"gender"`
	Style    string `json:"style"`
	Provider string `json:"provider"`
}

// ListVoices 获取可用语音列表
// @Summary 获取可用语音列表
// @Description 获取指定提供商或所有提供商的可用语音
// @Tags TTS
// @Accept json
// @Produce json
// @Param provider query string false "提供商"
// @Success 200 {array} VoiceInfo
// @Router /api/v1/tts/voices [get]
func (h *TTSHandler) ListVoices(c *gin.Context) {
	provider := c.Query("provider")

	if provider != "" {
		voices, err := h.service.GetVoices(c.Request.Context(), provider)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		result := make([]VoiceInfo, len(voices))
		for i, v := range voices {
			result[i] = VoiceInfo(v)
		}
		c.JSON(http.StatusOK, result)
		return
	}

	// 获取所有提供商的语音
	allVoices := h.service.GetAllVoices(c.Request.Context())
	var result []VoiceInfo
	for provider, voices := range allVoices {
		for _, v := range voices {
			result = append(result, VoiceInfo{
				ID:       v.ID,
				Name:     v.Name,
				Language: v.Language,
				Gender:   v.Gender,
				Style:    v.Style,
				Provider: provider,
			})
		}
	}
	c.JSON(http.StatusOK, result)
}

// ListProviders 获取提供商列表
// @Summary 获取TTS提供商列表
// @Description 获取所有可用的TTS提供商
// @Tags TTS
// @Produce json
// @Success 200 {array} string
// @Router /api/v1/tts/providers [get]
func (h *TTSHandler) ListProviders(c *gin.Context) {
	providers := h.service.GetProviders()
	c.JSON(http.StatusOK, providers)
}

// BatchGenerateRequest 批量生成请求
type BatchGenerateRequest struct {
	Provider    string          `json:"provider" binding:"required"`
	Voice       string          `json:"voice" binding:"required"`
	Texts       []string        `json:"texts" binding:"required,min=1"`
	Speed       float64         `json:"speed"`
	Pitch       float64         `json:"pitch"`
	SaveToFile  bool            `json:"save_to_file"`
}

// BatchGenerateResponse 批量生成响应
type BatchGenerateResponse struct {
	Success   bool                  `json:"success"`
	Results   []GenerateResponse    `json:"results"`
	Total     int                   `json:"total"`
	Failed    int                   `json:"failed"`
}

// BatchGenerate 批量生成语音
// @Summary 批量TTS语音合成
// @Description 批量将多个文本转换为语音
// @Tags TTS
// @Accept json
// @Produce json
// @Param request body BatchGenerateRequest true "批量TTS请求"
// @Success 200 {object} BatchGenerateResponse
// @Router /api/v1/tts/batch [post]
func (h *TTSHandler) BatchGenerate(c *gin.Context) {
	var req BatchGenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Speed == 0 {
		req.Speed = 1.0
	}
	if req.Pitch == 0 {
		req.Pitch = 1.0
	}

	results := make([]GenerateResponse, 0, len(req.Texts))
	failed := 0

	for i, text := range req.Texts {
		outputPath := ""
		if req.SaveToFile {
			outputPath = filepath.Join(h.outputPath, "tts", req.Provider, req.Voice, 
				string(rune('a'+i%26))+".mp3")
		}

		resp, err := h.service.Generate(c.Request.Context(), req.Provider, req.Voice, text, 
			req.Speed, req.Pitch, outputPath)

		if err != nil {
			failed++
			results = append(results, GenerateResponse{
				Success: false,
				Provider: req.Provider,
				Voice:   req.Voice,
				Error:   err.Error(),
			})
			continue
		}

		results = append(results, GenerateResponse{
			Success:    resp.Success,
			AudioData:  resp.AudioData,
			FilePath:   resp.FilePath,
			Duration:   resp.Duration,
			Provider:   resp.Provider,
			Voice:      req.Voice,
			Latency:    resp.Latency.String(),
		})
	}

	c.JSON(http.StatusOK, BatchGenerateResponse{
		Success: failed == 0,
		Results: results,
		Total:   len(req.Texts),
		Failed:  failed,
	})
}
