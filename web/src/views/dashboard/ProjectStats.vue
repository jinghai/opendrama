<template>
  <div class="project-stats">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #667eea, #764ba2)">
            <el-icon :size="30"><Document /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.dramas }}</h3>
            <p>短剧总数</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #f093fb, #f5576c)">
            <el-icon :size="30"><User /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.characters }}</h3>
            <p>角色总数</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #4facfe, #00f2fe)">
            <el-icon :size="30"><Picture /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.images }}</h3>
            <p>生成图片</p>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: linear-gradient(135deg, #43e97b, #38f9d7)">
            <el-icon :size="30"><VideoPlay /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.videos }}</h3>
            <p>生成视频</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card>
          <template #header><h4>📈 最近7天活动</h4></template>
          <div class="chart-placeholder">
            <el-empty description="图表开发中" />
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card>
          <template #header><h4>🏆 热门模板</h4></template>
          <el-table :data="hotTemplates" stripe>
            <el-table-column prop="name" label="模板名称" />
            <el-table-column prop="usage" label="使用次数" width="100" />
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Document, User, Picture, VideoPlay } from '@element-plus/icons-vue'

const stats = ref({ dramas: 12, characters: 48, images: 256, videos: 34 })
const hotTemplates = ref([
  { name: '甜宠短剧', usage: 156 },
  { name: '逆袭人生', usage: 98 },
  { name: '豪门恩怨', usage: 76 },
  { name: '穿越时空', usage: 54 }
])
</script>

<style scoped lang="scss">
.stat-card {
  display: flex;
  align-items: center;
  padding: 20px;
  .stat-icon {
    width: 60px; height: 60px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    margin-right: 15px;
  }
  .stat-content {
    h3 { font-size: 28px; font-weight: 600; margin: 0; }
    p { margin: 5px 0 0; color: #909399; font-size: 14px; }
  }
}
.chart-placeholder { height: 200px; display: flex; align-items: center; justify-content: center; }
</style>
