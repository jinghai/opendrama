# 开发计划

## 概述

本文档记录OpenDrama项目的开发计划和进度追踪。

## 开发进度总览

| 阶段 | 内容 | 状态 |
|------|------|------|
| 阶段一 | NewAPI架构升级 | ✅ 已完成 |
| 阶段二 | TTS语音合成 | ✅ 已完成 |
| 阶段三 | 前端优化 | 🔄 进行中 |
| 阶段四 | 功能完善 | ⏳ 待开始 |
| 阶段五 | 测试部署 | ⏳ 待开始 |

---

## 阶段一：架构升级 ✅ 已完成

- [x] NewAPI统一接口开发
- [x] 负载均衡器实现 (round-robin/weighted/least-cost)
- [x] 多服务商支持 (OpenAI/Claude/Gemini/通义千问/豆包)
- [x] 故障自动转移
- [x] 实时统计监控

**新增文件:**
- `pkg/ai/newapi/client.go` - NewAPI客户端
- `api/handlers/newapi.go` - API处理器
- `configs/newapi.example.yaml` - 配置示例

---

## 阶段二：TTS语音合成 ✅ 已完成

- [x] TTS客户端开发
- [x] Azure TTS集成
- [x] 阿里云TTS集成
- [x] 前端TTS界面
- [x] TTS路由和入口
- [x] 批量语音合成

**新增文件:**
- `pkg/ai/tts/client.go` - TTS客户端
- `api/handlers/tts.go` - TTS处理器
- `web/src/api/tts.ts` - 前端API
- `web/src/views/settings/TTSSettings.vue` - TTS设置页面

---

## 阶段三：前端优化 🔄 进行中

- [x] TTS设置入口
- [x] NewAPI设置页面
- [x] 品牌统一为OpenDrama
- [x] 首页仪表盘优化
- [x] QuickStats实时统计组件
- [x] QuickActions快捷操作
- [x] NewAPI API封装
- [x] 短剧卡片组件
- [x] 工作流进度组件
- [x] 批量操作组件
- [x] 图片画廊组件
- [x] 视频预览组件
- [x] 拖拽排序组件
- [ ] 响应式布局优化

---

## 阶段四：功能完善 ⏳ 待开始

- [ ] 用户管理系统
- [ ] 权限控制
- [ ] 项目分享功能
- [ ] 模板系统
- [ ] 导出功能 (PDF/视频)

---

## 阶段五：测试和部署 ⏳ 待开始

- [ ] 单元测试
- [ ] 集成测试
- [ ] E2E测试
- [ ] 部署文档
- [ ] Docker镜像优化
- [ ] 性能优化

---

## GitHub提交历史 (最近)

```
3a2f55b ⚙️ 更新配置文件 - OpenDrama品牌
c7f3c0d 🐳 更新Docker配置 - OpenDrama品牌
aa2abc2 📋 更新开发计划
27ae29e 📊 添加首页快速统计组件
65488f8 🔧 修复AppHeader品牌名称
48c04f3 🎨 首页仪表盘品牌更新
... 共15次提交
```

---

## 下一步任务

### 高优先级
1. 服务器部署测试 - 部署最新代码到服务器
2. API功能验证 - 测试NewAPI和TTS接口
3. 前端构建检查 - 确保前端可以正常构建

### 中优先级
1. 剧本编辑器改进
2. 分镜管理可视化
3. 批量操作功能

### 低优先级
1. 国际化支持
2. 主题定制
3. 插件系统

---

*最后更新: 2026-02-16*
