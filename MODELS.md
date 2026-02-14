# AI模型使用记录

## 🤖 OpenDrama项目支持的AI模型

### 已集成的AI服务商

#### 🎭 **文本生成 (Text Generation)**
- **OpenAI**: GPT-4, GPT-3.5
- **Gemini**: Gemini Pro
- **Claude**: Claude-3, Claude-3.5
- **ChatFire**: 国产文本模型

#### 🎨 **图像生成 (Image Generation)**  
- **OpenAI**: DALL-E 3
- **Gemini**: Imagen 2
- **Volcengine**: 国产图像模型
- **Midjourney**: 通过API集成

#### 🎥 **视频生成 (Video Generation)**
- **OpenAI**: Sora (预览版)
- **RunwayML**: Runway Gen-2
- **Pika Labs**: Pika Gen-1
- **Minimax**: 国产视频模型
- **ChatFire**: 国产视频模型

#### 🎬 **其他AI服务**
- **NewAPI**: 统一接口服务 (计划中)

---

## 🔧 TTS语音合成 (待开发)

### 计划集成的TTS服务商
1. **Azure TTS** (中文语音优秀)
   - 中文支持度最高
   - 音色丰富，情感表达能力好
   - 企业级稳定性

2. **阿里云TTS** 
   - 国产化程度高
   - 价格优惠，性价比好
   - 中文发音标准

3. **腾讯云TTS**
   - 国内访问速度快
   - 支持多种音色
   - 价格适中

4. **OpenAI TTS** 
   - 音质优秀，多语言支持
   - 基于GPT技术，语音自然

5. **ElevenLabs**
   - 开源TTS引擎
   - 支持本地部署
   - 音色可定制
   - 免费开源选项

---

## 📋 当前使用的模型

### 🔥 **OpenDrama项目AI配置**
- **文本生成**: 优先使用Gemini Pro
- **图像生成**: 主要使用Gemini Imagen 2
- **视频生成**: 接入多种模型，根据需求选择
- **TTS语音**: 计划中，将集成Azure TTS

### 🎯 **大模型规划**
- **短期目标**: 集成Azure TTS，完成语音合成功能
- **中期目标**: 接入更多OpenAI模型，提升生成质量
- **长期目标**: 自研或开源模型，降低成本

---

## 📊 技术架构优化

### 🔍 NewAPI架构 (进行中)
- **统一接口**: 所有AI请求通过NewAPI路由
- **负载均衡**: 自动选择最优可用通道
- **成本优化**: 比价和用量监控
- **高可用性**: 多厂商故障转移

---

*最后更新: 2026-02-12*