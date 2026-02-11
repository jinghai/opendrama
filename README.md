# OpenDrama (开源短剧) - AI短剧创作平台

<div align="center">

**基于 Go + Vue3 的全栈 AI 短剧自动化生产平台**

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Vue Version](https://img.shields.io/badge/Vue-3.x-4FC08D?style=flat&logo=vue.js)](https://vuejs.org)
[![License](https://img.shields.io/badge/License-MIT-lightgrey.svg)](https://opensource.org/licenses/MIT)
[![Open Source](https://img.shields.io/badge/开源-OpenDrama-blue.svg)](https://github.com/jinghai/opendrama)

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
- ✅ 多AI服务商支持（当前NewAPI升级中）
- ✅ 文本生成服务
- ✅ 图像生成服务
- ✅ 视频生成服务

### 🚧 TTS语音合成系统 (开发中)
- 🔄 **架构设计**: 统一TTS接口，支持多服务商
- 🔄 **语音角色**: 音色配置和角色绑定
- 🔄 **批量处理**: 批量配音生成功能
- 🔄 **前端界面**: 配音工作室和音频编辑器

### 📦 资源管理
- ✅ 素材库统一管理
- ✅ 本地存储支持
- ✅ 资源导入导出
- ✅ 任务进度追踪

---

## 🚀 快速开始

```bash
# 克隆项目
git clone https://github.com/jinghai/opendrama.git

# 安装依赖
cd opendrama
go mod download
cd web && npm install

# 启动开发环境
# 终端1: 启动后端服务
go run main.go

# 终端2: 启动前端开发服务器
cd web && npm run dev

# 访问应用
# 前端: http://localhost:3012
# 后端API: http://localhost:5678/api/v1
```

## 🌐 项目链接

- **🏠 GitHub仓库**: https://github.com/jinghai/opendrama
- **🌐 官方网站**: https://opendrama.ai (开发中)
- **📖 文档**: 在线文档和API参考

## 🎯 开源协议

MIT License - 允许自由使用、修改和分发

---

## 🔮 开发路线图

### 阶段一：核心架构升级 (进行中)
1. **🔧 NewAPI集成** - 统一AI服务接口
   - 移除分散的provider逻辑
   - 实现负载均衡和成本优化
   - 支持多厂商统一管理

2. **🚧 TTS语音合成** - 完整配音系统
   - 多TTS服务商支持(Azure/阿里云/腾讯云)
   - 语音角色和音色配置
   - 批量配音生成

### 阶段二：功能完善 (计划中)
1. **🎨 前端界面优化**
   - 剧本编辑器改进
   - 可视化时间轴编辑
   - 拖拽式分镜管理

2. **📱 用户体验提升**
   - 一键生成完整短剧
   - 模板系统和预设
   - 实时预览功能

3. **🔧 系统完善**
   - 用户管理系统
   - 权限控制和分享功能
   - 性能监控和优化

### 阶段三：生态建设 (未来)
1. **🌍 官网和文档**
   - GitHub Pages部署官网
   - 完整API文档
   - 使用教程和示例

2. **📝 社区建设**
   - 开源贡献指南
   - Discrod/Telegram交流群
   - 开发者生态建设

---

**🎉 开源让AI短剧创作更简单！**

Made with ❤️ by OpenDrama Community
