# OpenDrama (开源短剧)

<p align="center">
  <img src="drama.png" alt="OpenDrama Logo" width="200"/>
</p>

<p align="center">
  <a href="https://github.com/jinghai/opendrama/releases/latest">
    <img src="https://img.shields.io/github/release/jinghai/opendrama?style=flat-square" alt="版本"/>
  </a>
  <a href="https://github.com/jinghai/opendrama/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/jinghai/opendrama?style=flat-square" alt="许可证"/>
  </a>
  <a href="https://gitter.im/opendrama/community">
    <img src="https://img.shields.io/gitter/room/opendrama/community?style=flat-square" alt="Gitter"/>
  </a>
</p>

---

## 📖 项目简介

**OpenDrama (开源短剧)** 是一款基于AI的短剧创作平台，致力于让每个人都能成为剧作家。

我们相信，AI技术应该 democratize（民主化）创意表达。通过 OpenDrama，创作者可以：

- 🤖 使用AI生成剧本、分镜、角色
- 🎬 自动生成视频和配音
- 👥 轻松管理角色和场景库
- 📝 一站式完成短剧创作全流程

---

## ✨ 核心特性

### 🎭 智能创作
- **AI剧本生成** - 输入主题，AI自动生成完整剧本
- **角色智能提取** - 从剧本自动识别并创建角色
- **分镜自动生成** - AI分析剧本，生成详细分镜脚本

### 🎬 视频生成
- **图生视频** - 将静态图像转换为动态视频
- **多模型支持** - 支持多种AI视频生成服务
- **批量处理** - 高效处理大量素材

### 🎙️ 语音合成 (TTS)
- **多语言支持** - 支持中文、英文等多种语言
- **丰富音色** - 多种语音风格可选
- **情感表达** - 支持语速、音调调节

### 👥 资源管理
- **角色库** - 统一管理所有角色形象
- **道具库** - 丰富的道具资源管理
- **场景库** - 场景背景素材管理

---

## 🚀 快速开始

### 环境要求

- **Go** 1.20+
- **Node.js** 18+
- **PostgreSQL** 14+ 或 **SQLite**
- **FFmpeg** (用于音视频处理)

### 安装部署

```bash
# 1. 克隆项目
git clone https://github.com/jinghai/opendrama.git
cd opendrama

# 2. 配置后端
cp configs/config.example.yaml configs/config.yaml
# 编辑配置文件，填入API Key

# 3. 启动后端
go run main.go

# 4. 启动前端
cd web
npm install
npm run dev
```

访问 http://localhost:3012 即可开始使用。

### Docker部署

```bash
# 使用Docker Compose一键部署
docker-compose up -d
```

---

## 📁 项目结构

```
opendrama/
├── api/                    # API路由和处理器
│   ├── handlers/          # 请求处理器
│   ├── routes/            # 路由配置
│   └── middlewares/       # 中间件
├── application/            # 应用层逻辑
│   └── services/          # 业务服务
├── domain/                # 领域模型
│   └── models/            # 数据模型
├── infrastructure/        # 基础设施
│   ├── database/          # 数据库
│   ├── storage/           # 存储
│   └── external/          # 外部工具
├── pkg/                   # 公共包
│   ├── ai/               # AI客户端
│   │   ├── newapi/       # NewAPI统一接口
│   │   └── tts/          # TTS语音合成
│   ├── config/           # 配置管理
│   └── logger/           # 日志
├── web/                   # Vue.js前端
│   ├── src/
│   │   ├── api/          # API调用
│   │   ├── components/   # 组件
│   │   ├── views/        # 页面视图
│   │   └── stores/       # 状态管理
│   └── package.json
├── configs/               # 配置文件
├── docs/                 # 文档
└── main.go              # 后端入口
```

---

## 🔧 配置说明

### AI服务配置

在 `configs/config.yaml` 中配置AI服务：

```yaml
ai:
  # OpenAI配置
  openai:
    api_key: "your-openai-api-key"
    base_url: "https://api.openai.com"
  
  # Google Gemini配置
  gemini:
    api_key: "your-gemini-api-key"
  
  # 火山引擎(豆包)配置
  volcengine:
    api_key: "your-ark-api-key"
    space_id: "your-space-id"

# NewAPI统一接口配置
newapi:
  base_url: "https://api.newapi.com"
  api_key: "your-newapi-key"
  load_balancer:
    strategy: "least-cost"
    providers:
      - name: "openai"
        enabled: true
      - name: "qwen"
        enabled: true

# TTS服务配置
tts:
  azure:
    api_key: "your-azure-tts-key"
    region: "eastus"
  alibaba:
    api_key: "your-alibaba-key"
    app_key: "your-app-key"
```

---

## 📡 API接口

### 短剧管理

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/dramas` | GET | 获取短剧列表 |
| `/api/v1/dramas` | POST | 创建短剧 |
| `/api/v1/dramas/:id` | GET | 获取短剧详情 |
| `/api/v1/dramas/:id` | PUT | 更新短剧 |
| `/api/v1/dramas/:id` | DELETE | 删除短剧 |

### AI生成

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/generation/characters` | POST | 生成角色 |
| `/api/v1/images` | POST | 生成图片 |
| `/api/v1/videos` | POST | 生成视频 |

### NewAPI统一接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/newapi/text` | POST | 文本生成 |
| `/api/v1/newapi/image` | POST | 图像生成 |
| `/api/v1/newapi/stats` | GET | 获取统计 |

### TTS语音合成

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/v1/tts/generate` | POST | 生成语音 |
| `/api/v1/tts/voices` | GET | 获取语音列表 |
| `/api/v1/tts/providers` | GET | 获取提供商列表 |
| `/api/v1/tts/batch` | POST | 批量生成 |

---

## 🛠️ 开发计划

### 阶段一：架构升级 (进行中)
- [x] NewAPI统一接口开发
- [x] 负载均衡器实现
- [ ] 多服务商集成测试

### 阶段二：TTS语音合成 (进行中)
- [x] TTS客户端开发
- [x] Azure TTS集成
- [x] 阿里云TTS集成
- [ ] 前端界面完善
- [ ] API Key配置

### 阶段三：前端优化
- [ ] 剧本编辑器改进
- [ ] 可视化分镜管理
- [ ] 拖拽排序功能
- [ ] 批量操作支持

### 阶段四：功能完善
- [ ] 用户管理系统
- [ ] 权限控制
- [ ] 项目分享功能
- [ ] 模板系统

---

## 🤝 贡献指南

欢迎贡献代码！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解如何参与贡献。

### 提交问题

如果您发现bug或有功能建议，请提交 Issue。

### 提交代码

1. Fork 本仓库
2. 创建分支 (`git checkout -b feature/xxx`)
3. 提交更改 (`git commit -m 'Add xxx'`)
4. 推送分支 (`git push origin feature/xxx`)
5. 创建 Pull Request

---

## 📄 许可证

本项目采用 **MIT** 许可证 - 详见 [LICENSE](LICENSE) 文件。

---

## 🙏 致谢

- [Vue.js](https://vuejs.org/) - 前端框架
- [Gin](https://gin-gonic.com/) - Go Web框架
- [Element Plus](https://element-plus.org/) - UI组件库
- [OpenAI](https://openai.com/) - AI服务
- [火山引擎](https://www.volcengine.com/) - 国产AI服务

---

## 📱 联系方式

- **GitHub**: https://github.com/jinghai/opendrama
- **问题反馈**: https://github.com/jinghai/opendrama/issues
- **邮箱**: contact@opendrama.ai

---

<p align="center">
  <strong>OpenDrama (开源短剧)</strong> - 让每个人都能成为剧作家 🎭🎬
</p>
