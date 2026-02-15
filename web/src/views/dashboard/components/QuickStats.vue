<template>
  <div class="quick-stats">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: #409eff">
            <el-icon :size="30"><Document /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.dramas || 0 }}</h3>
            <p>短剧总数</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: #67c23a">
            <el-icon :size="30"><User /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.characters || 0 }}</h3>
            <p>角色总数</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: #e6a23c">
            <el-icon :size="30"><Picture /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.images || 0 }}</h3>
            <p>生成图片</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-icon" style="background: #f56c6c">
            <el-icon :size="30"><VideoPlay /></el-icon>
          </div>
          <div class="stat-content">
            <h3>{{ stats.videos || 0 }}</h3>
            <p>生成视频</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Document, User, Picture, VideoPlay } from '@element-plus/icons-vue'
import { getDramaStats } from '@/api/drama'

const stats = ref({
  dramas: 0,
  characters: 0,
  images: 0,
  videos: 0
})

onMounted(async () => {
  try {
    const data = await getDramaStats()
    stats.value = {
      dramas: data.total_dramas || 0,
      characters: data.total_characters || 0,
      images: data.total_images || 0,
      videos: data.total_videos || 0
    }
  } catch (error) {
    console.error('Failed to load stats:', error)
  }
})
</script>

<style scoped lang="scss">
.quick-stats {
  margin: 20px 0;
  
  .stat-card {
    display: flex;
    align-items: center;
    padding: 20px;
    
    .stat-icon {
      width: 60px;
      height: 60px;
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      margin-right: 15px;
    }
    
    .stat-content {
      h3 {
        font-size: 28px;
        font-weight: 600;
        margin: 0;
        color: #303133;
      }
      
      p {
        margin: 5px 0 0;
        font-size: 14px;
        color: #909399;
      }
    }
  }
}
</style>
