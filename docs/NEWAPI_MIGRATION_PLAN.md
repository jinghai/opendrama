# NewAPI 架构升级方案

## 🎯 升级目标

将当前分散的AI服务商对接架构升级为统一的NewAPI架构，实现：
- 统一API接口
- 智能负载均衡
- 成本优化
- 高可用性
- 简化配置

## 📊 现状分析

### 当前架构问题
```go
// 当前代码中的硬编码逻辑
switch req.Provider {
case "gemini", "google":
    endpoint = "/v1beta/models/{model}:generateContent"
case "openai":
    endpoint = "/chat/completions"
case "chatfire":
    endpoint = "/chat/completions"
case "doubao", "volcengine", "volces":
    endpoint = "/contents/generations/tasks"
}
```

### 配置复杂度高
- 需要为每个厂商配置API Key
- 需要维护不同的Endpoint格式
- 新增厂商需要修改多处代码

## 🚀 NewAPI架构方案

### 1. 配置文件简化
```yaml
# 当前配置
ai:
  default_text_provider: "openai"
  default_image_provider: "openai"
  default_video_provider: "doubao"
  # 需要为每个厂商配置...

# NewAPI配置
ai:
  newapi:
    base_url: "https://your-newapi-instance.com"
    api_key: "sk-xxxxxxxxx"
    enabled_providers: ["openai", "claude", "doubao", "gemini", "azure"]
    load_balancing: true
    cost_optimization: true
```

### 2. 数据库模型重构
```go
// 简化AI配置模型
type AIServiceConfig struct {
    ID            uint       `json:"id"`
    ServiceType   string     `json:"service_type"` // text, image, video, tts
    NewAPIBaseURL string     `json:"newapi_base_url"`
    NewAPIKey    string     `json:"newapi_key"`
    Model         string     `json:"model"`        // 单一模型名称
    Priority      int        `json:"priority"`
    Settings      string     `json:"settings"`    // JSON配置
}

// 移除Provider字段，统一使用NewAPI
```

### 3. API服务层重构
```go
type NewAPIService struct {
    baseURL    string
    apiKey     string
    httpClient *http.Client
}

// 统一的AI请求接口
func (s *NewAPIService) GenerateText(request *TextRequest) (*TextResponse, error)
func (s *NewAPIService) GenerateImage(request *ImageRequest) (*ImageResponse, error) 
func (s *NewAPIService) GenerateVideo(request *VideoRequest) (*VideoResponse, error)
func (s *NewAPIService) GenerateSpeech(request *TTSRequest) (*AudioResponse, error)
```

## 📋 实施步骤

### 阶段一: NewAPI集成准备 (1周)
1. **搭建NewAPI实例**
   - 部署NewAPI服务
   - 配置各厂商API Key
   - 测试通道可用性

2. **创建NewAPI客户端**
   - 实现NewAPI Go客户端
   - 统一请求格式
   - 错误处理和重试

### 阶段二: 后端架构重构 (2-3周)
1. **重构AIService**
   - 移除分散provider逻辑
   - 统一NewAPI调用
   - 保持现有API兼容性

2. **扩展配置支持**
   - 添加TTS服务配置
   - 添加NewAPI配置字段
   - 数据库迁移脚本

3. **前端配置界面**
   - 简化AI配置页面
   - 添加TTS配置
   - NewAPI状态监控

### 阶段三: 功能完善 (2周)
1. **智能负载均衡**
   - 自动通道选择
   - 成本实时监控
   - 故障自动切换

2. **监控和统计**
   - 使用统计面板
   - 成本分析报表
   - 通道性能监控

## 🎯 预期效果

### 开发效率提升
- **配置简化**: 从N个厂商配置 → 1个NewAPI配置
- **代码简化**: 移除大量provider判断逻辑
- **维护成本**: 降低80%的配置维护工作量

### 系统稳定性
- **高可用**: 多厂商自动故障转移
- **成本优化**: 智能选择最优价格通道
- **监控完善**: 实时监控所有AI服务状态

### 扩展性
- **新增厂商**: 在NewAPI中配置即可，无需修改代码
- **新功能**: 统一接口更容易添加新功能
- **API演进**: 跟随NewAPI功能升级

## 🛠️ 技术实现要点

### 1. NewAPI部署选项
- **自托管**: 完全控制，数据安全
- **托管服务**: 快速上线，减少运维
- **混合部署**: 核心服务自托管，补充使用托管

### 2. 迁移策略
- **渐进式迁移**: 保持现有API兼容
- **A/B测试**: 新旧架构并行运行
- **平滑切换**: 验证稳定后切换

### 3. 监控指标
- **响应时间**: 各通道响应时间对比
- **成功率**: 请求成功率统计
- **成本分析**: 按厂商、按功能成本分析
- **错误率**: 错误类型和频率统计

---

**注意**: 这是核心架构升级，建议优先实施，为后续功能扩展打下坚实基础。