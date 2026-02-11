# OpenDrama (开源短剧) - AI短剧创作平台

<div align="center">

**基于 Go + Vue3 的全栈 AI 短剧自动化生产平台**

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Vue Version](https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat&logo=vue.js)](https://vuejs.org)
[![License](https://img.shields.io/badge/License-MIT-lightgrey.svg)](https://opensource.org/licenses/MIT)
[![Open Source](https://img.shields.io/badge/开源-OpenDrama-blue.svg)](https://github.com/your-org/opendrama)

[功能特性](#功能特性) • [快速开始](#快速开始) • [开源贡献](#开源贡献)

[简体中文](README-CN.md) | [English](README.md)

</div>

---

## 📖 项目简介

OpenDrama 是一个基于 AI 的开源短剧自动化创作平台，实现从剧本生成、角色设计、分镜制作到视频合成的全流程自动化。

### 🎯 核心价值

- **🤖 AI 驱动**：使用大语言模型解析剧本，提取角色、场景和分镜信息
- **🎨 智能创作**：AI 绘图生成角色形象和场景背景
- **🎬 一站式体验**：完整的短剧制作工作流，从创意到成片一站式完成
- **🌍 开源协作**：开放源码，全球开发者共同参与

### 🛠️ 技术架构

采用**DDD 领域驱动设计**，清晰分层：

```
├── API层 (Gin HTTP)
├── 应用服务层 (Business Logic)
├── 领域层 (Domain Models)
└── 基础设施层 (Database, External Services)
```

---

## ✨ 功能特性

### 🎭 剧本管理
- ✅ 创建/编辑短剧剧本
- ✅ 剧本大纲管理
- ✅ 角色设定
- ✅ 分集管理

### 🎨 角色系统
- ✅ AI 生成角色形象
- ✅ 角色库管理（可复用）
- ✅ 批量角色生成
- ✅ 角色特征设定

### 🎬 分镜制作
- ✅ 自动生成分镜脚本
- ✅ 场景描述和镜头设计
- ✅ 分镜图片生成（文生图）
- ✅ 帧类型选择（首帧/关键帧/尾帧）

### 🎥 视频生成
- ✅ 图生视频自动生成
- ✅ 视频合成和剪辑
- ✅ 批量视频生成
- ✅ 音频提取和处理

### 🤖 AI集成
- ✅ 多AI服务商支持（NewAPI统一接口）
- ✅ 文本生成服务
- ✅ 图像生成服务
- ✅ 视频生成服务
- 🚧 TTS语音合成 (开发中)

### 📦 资源管理
- ✅ 素材库统一管理
- ✅ 本地存储支持
- ✅ 资源导入导出
- ✅ 任务进度追踪

---

## 🚀 快速开始

### 📋 环境要求

| 软件        | 版本要求 | 说明                 |
| ----------- | -------- | -------------------- |
| **Go**      | 1.23+    | 后端运行环境         |
| **Node.js** | 18+      | 前端构建环境         |
| **npm**     | 9+       | 包管理工具           |
| **FFmpeg**  | 4.0+     | 视频处理（**必需**） |
| **SQLite**  | 3.x      | 数据库（已内置）     |

### ⚙️ 配置文件

```yaml
# config.yaml
app:
  name: "OpenDrama"
  version: "1.0.0"
  language: "zh"

server:
  port: 5678
  host: "0.0.0.0"

ai:
  newapi:
    base_url: "https://your-newapi-instance.com"
    api_key: "your-api-key"
    enabled_providers: ["openai", "claude", "doubao"]
```

### 🎯 启动项目

```bash
# 克隆项目
git clone https://github.com/your-org/opendrama.git
cd opendrama

# 安装依赖
go mod download
cd web && npm install

# 启动开发环境
go run main.go
cd web && npm run dev
```

---

## 🤝 开源贡献

欢迎全球开发者参与 OpenDrama 建设！

### 🎯 贡献方向

1. **🎨 UI/UX 优化**
   - 界面设计和交互优化
   - 移动端适配

2. **🤖 AI 功能扩展**
   - 集成更多AI服务
   - TTS语音合成实现
   - 智能推荐算法

3. **🛠️ 架构优化**
   - NewAPI集成完善
   - 性能优化
   - 数据库设计改进

4. **📚 文档完善**
   - API文档编写
   - 使用教程制作
   - 多语言支持

### 📝 贡献流程

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交改动 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

---

## 🌟 开源协议

本项目采用 [MIT License](LICENSE)，允许自由使用、修改和分发。

---

## 📞 联系我们

- 📧 **Email**: opendrama@example.com
- 🐙 **GitHub**: https://github.com/your-org/opendrama
- 💬 **讨论**: [GitHub Discussions](https://github.com/your-org/opendrama/discussions)

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给一个 Star！**

让开源短剧创作更简单，让每个人都能成为剧作家。

Made with ❤️ by OpenDrama Community

</div>