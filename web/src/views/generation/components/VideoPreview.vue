<template>
  <div class="video-preview">
    <div class="preview-main">
      <video 
        ref="videoRef" 
        :src="currentVideo?.url" 
        controls
        class="preview-video"
        @timeupdate="onTimeUpdate"
        @loadedmetadata="onLoaded"
      />
    </div>
    
    <div class="preview-sidebar">
      <div class="sidebar-header">
        <h4>视频信息</h4>
      </div>
      
      <div class="video-info">
        <div class="info-item">
          <span class="label">名称：</span>
          <span class="value">{{ currentVideo?.name || '-' }}</span>
        </div>
        <div class="info-item">
          <span class="label">时长：</span>
          <span class="value">{{ formatDuration(currentVideo?.duration) }}</span>
        </div>
        <div class="info-item">
          <span class="label">分辨率：</span>
          <span class="value">{{ currentVideo?.resolution || '-' }}</span>
        </div>
        <div class="info-item">
          <span class="label">大小：</span>
          <span class="value">{{ formatSize(currentVideo?.size) }}</span>
        </div>
        <div class="info-item">
          <span class="label">创建时间：</span>
          <span class="value">{{ formatTime(currentVideo?.created_at) }}</span>
        </div>
      </div>
      
      <el-divider />
      
      <div class="timeline-bar">
        <div class="timeline-header">
          <span>时间轴</span>
          <span>{{ formatTime(currentTime) }} / {{ formatDuration(totalDuration) }}</span>
        </div>
        <div class="timeline-track" @click="seekTo">
          <div class="track-progress" :style="{ width: progress + '%' }"></div>
        </div>
      </div>
      
      <el-divider />
      
      <div class="video-actions">
        <el-button type="primary" @click="downloadVideo">
          <el-icon><Download /></el-icon>
          下载视频
        </el-button>
        <el-button @click="previewVisible = true">
          <el-icon><FullScreen /></el-icon>
          全屏预览
        </el-button>
      </div>
    </div>
    
    <!-- 视频列表 -->
    <div class="video-list">
      <h4>所有视频 ({{ videos.length }})</h4>
      <el-scrollbar height="200px">
        <div class="list-items">
          <div 
            v-for="video in videos" 
            :key="video.id"
            class="video-item"
            :class="{ active: video.id === currentVideo?.id }"
            @click="selectVideo(video)"
          >
            <div class="video-thumb">
              <img v-if="video.thumbnail" :src="video.thumbnail" />
              <div v-else class="thumb-placeholder">
                <el-icon><VideoPlay /></el-icon>
              </div>
              <div class="video-duration">{{ formatDuration(video.duration) }}</div>
            </div>
            <div class="video-meta">
              <p class="video-name">{{ video.name }}</p>
            </div>
          </div>
        </div>
      </el-scrollbar>
    </div>
    
    <!-- 全屏预览对话框 -->
    <el-dialog v-model="previewVisible" fullscreen>
      <video 
        :src="currentVideo?.url" 
        controls 
        autoplay
        class="fullscreen-video"
      />
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Download, FullScreen, VideoPlay } from '@element-plus/icons-vue'

const props = defineProps<{
  videos: any[]
}>()

const emit = defineEmits(['select'])

const videoRef = ref<HTMLVideoElement>()
const currentVideo = ref<any>(props.videos[0])
const currentTime = ref(0)
const totalDuration = ref(0)
const previewVisible = ref(false)

const progress = computed(() => {
  if (totalDuration.value === 0) return 0
  return (currentTime.value / totalDuration.value) * 100
})

const formatTime = (seconds: number) => {
  if (!seconds) return '00:00'
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

const formatDuration = (seconds: number) => {
  if (!seconds) return '00:00'
  const hours = Math.floor(seconds / 3600)
  const mins = Math.floor((seconds % 3600) / 60)
  const secs = Math.floor(seconds % 60)
  if (hours > 0) {
    return `${hours}:${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
  }
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

const formatSize = (bytes: number) => {
  if (!bytes) return '-'
  const units = ['B', 'KB', 'MB', 'GB']
  let size = bytes
  let unitIndex = 0
  while (size >= 1024 && unitIndex < units.length - 1) {
    size /= 1024
    unitIndex++
  }
  return `${size.toFixed(2)} ${units[unitIndex]}`
}

const onTimeUpdate = () => {
  if (videoRef.value) {
    currentTime.value = videoRef.value.currentTime
  }
}

const onLoaded = () => {
  if (videoRef.value) {
    totalDuration.value = videoRef.value.duration
  }
}

const seekTo = (e: MouseEvent) => {
  if (videoRef.value) {
    const rect = (e.target as HTMLElement).getBoundingClientRect()
    const percent = (e.clientX - rect.left) / rect.width
    videoRef.value.currentTime = percent * totalDuration.value
  }
}

const selectVideo = (video: any) => {
  currentVideo.value = video
  emit('select', video)
}

const downloadVideo = () => {
  if (currentVideo.value?.url) {
    window.open(currentVideo.value.url, '_blank')
  }
}
</script>

<style scoped lang="scss">
.video-preview {
  display: grid;
  grid-template-columns: 1fr 300px;
  grid-template-rows: auto auto;
  gap: 20px;
  
  .preview-main {
    grid-row: 1;
    grid-column: 1;
    
    .preview-video {
      width: 100%;
      max-height: 500px;
      background: #000;
      border-radius: 8px;
    }
  }
  
  .preview-sidebar {
    grid-row: 1;
    grid-column: 2;
    background: var(--bg-card);
    border-radius: 8px;
    padding: 20px;
    
    .sidebar-header h4 {
      margin: 0 0 15px;
    }
    
    .video-info {
      .info-item {
        display: flex;
        margin-bottom: 10px;
        font-size: 13px;
        
        .label {
          color: #909399;
          width: 70px;
        }
        
        .value {
          flex: 1;
          word-break: break-all;
        }
      }
    }
    
    .timeline-bar {
      .timeline-header {
        display: flex;
        justify-content: space-between;
        font-size: 12px;
        color: #909399;
        margin-bottom: 8px;
      }
      
      .timeline-track {
        height: 8px;
        background: #f5f7fa;
        border-radius: 4px;
        cursor: pointer;
        overflow: hidden;
        
        .track-progress {
          height: 100%;
          background: #409eff;
          border-radius: 4px;
          transition: width 0.1s;
        }
      }
    }
    
    .video-actions {
      display: flex;
      gap: 10px;
      
      .el-button {
        flex: 1;
      }
    }
  }
  
  .video-list {
    grid-row: 2;
    grid-column: 1 / -1;
    background: var(--bg-card);
    border-radius: 8px;
    padding: 20px;
    
    h4 {
      margin: 0 0 15px;
    }
    
    .list-items {
      display: flex;
      gap: 15px;
      
      .video-item {
        width: 150px;
        cursor: pointer;
        border: 2px solid transparent;
        border-radius: 8px;
        overflow: hidden;
        transition: all 0.2s;
        
        &:hover, &.active {
          border-color: #409eff;
        }
        
        .video-thumb {
          position: relative;
          height: 80px;
          background: #000;
          
          img {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
          
          .thumb-placeholder {
            width: 100%;
            height: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-size: 24px;
          }
          
          .video-duration {
            position: absolute;
            bottom: 5px;
            right: 5px;
            background: rgba(0, 0, 0, 0.7);
            color: white;
            font-size: 11px;
            padding: 2px 5px;
            border-radius: 3px;
          }
        }
        
        .video-meta {
          padding: 8px;
          
          .video-name {
            margin: 0;
            font-size: 12px;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
          }
        }
      }
    }
  }
  
  .fullscreen-video {
    width: 100%;
    height: 100%;
  }
}
</style>
