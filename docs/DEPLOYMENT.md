# 服务器部署指南

## 快速部署命令

在服务器上执行以下命令：

```bash
# 1. 进入项目目录
cd /root/opendrama

# 2. 拉取最新代码
git pull origin main

# 3. 重新构建并启动
docker-compose up -d --build

# 4. 查看日志
docker logs -f opendrama

# 5. 测试API
curl http://localhost:5678/health
```

## 测试命令

部署完成后测试以下接口：

```bash
# 健康检查
curl http://localhost:5678/health

# NewAPI统计
curl http://localhost:5678/api/v1/newapi/stats

# TTS语音列表
curl http://localhost:5678/api/v1/tts/voices

# 短剧列表
curl http://localhost:5678/api/v1/dramas
```

## 前端测试

```bash
# 进入前端目录
cd /root/opendrama/web

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 或构建生产版本
npm run build
```

## Docker常用命令

```bash
# 查看容器状态
docker ps

# 查看日志
docker logs opendrama

# 重启服务
docker-compose restart

# 停止服务
docker-compose down

# 重新构建
docker-compose build --no-cache
```
