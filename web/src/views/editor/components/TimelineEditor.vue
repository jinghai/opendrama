<template>
  <div class="timeline-editor">
    <!-- 时间轴头部 -->
    <div class="timeline-header">
      <div class="header-left">
        <h3>📐 时间轴编辑</h3>
      </div>
      <div class="header-right">
        <el-button @click="zoomOut" size="small">➖ 缩小</el-button>
        <span class="zoom-level">{{ Math.round(zoom * 100) }}%</span>
        <el-button @click="zoomIn" size="small">➕ 放大</el-button>
        <el-divider direction="vertical" />
        <el-button type="primary" @click="addClip">
          <el-icon><Plus /></el-icon> 添加片段
        </el-button>
      </div>
    </div>
    
    <!-- 时间轴主体 -->
    <div class="timeline-body">
      <!-- 轨道列表 -->
      <div class="tracks-container">
        <div class="track" v-for="track in tracks" :key="track.id">
          <div class="track-header">
            <el-icon><VideoCamera /></el-icon>
            <span>{{ track.name }}</span>
            <el-switch v-model="track.visible" size="small" />
            <el-button text size="small" @click="deleteTrack(track.id)">删除</el-button>
          </div>
          <div 
            class="track-content"
            :style="{ width: timelineWidth + 'px' }"
            @click="onTimelineClick"
          >
            <!-- 播放头 -->
            <div class="playhead" :style="{ left: playheadPosition + 'px' }"></div>
            
            <!-- 视频片段 -->
            <div 
              v-for="clip in track.clips" 
              :key="clip.id"
              class="clip"
              :class="{ selected: clip.id === selectedClipId }"
              :style="{ 
                left: clip.startTime * zoom + 'px',
                width: (clip.endTime - clip.startTime) * zoom + 'px',
                backgroundColor: getClipColor(clip.type)
              }"
              @click.stop="selectClip(clip)"
              @dblclick="editClip(clip)"
            >
              <div class="clip-content">
                <span class="clip-name">{{ clip.name }}</span>
                <span class="clip-duration">{{ formatDuration(clip.endTime - clip.startTime) }}</span>
              </div>
              <!-- 缩放手柄 -->
              <div class="clip-handle left" @mousedown.stop="startResize($event, clip, 'left')"></div>
              <div class="clip-handle right" @mousedown.stop="startResize($event, clip, 'right')"></div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 时间刻度 -->
      <div class="time-ruler">
        <div 
          v-for="tick in timeTicks" 
          :key="tick.time"
          class="tick"
          :style="{ left: tick.time * zoom + 'px' }"
        >
          <span class="tick-label">{{ formatTime(tick.time) }}</span>
        </div>
      </div>
    </div>
    
    <!-- 底部控制栏 -->
    <div class="timeline-footer">
      <el-button-group>
        <el-button :type="isPlaying ? 'danger' : 'primary'" @click="togglePlay">
          <el-icon><VideoPlay v-if="!isPlaying" /><VideoPause v-else /></el-icon>
          {{ isPlaying ? '暂停' : '播放' }}
        </el-button>
        <el-button @click="stop">
          <el-icon><VideoCamera /></el-icon>
          停止
        </el-button>
      </el-button-group>
      
      <div class="time-display">
        <span class="current-time">{{ formatTime(currentTime) }}</span>
        <span class="separator">/</span>
        <span class="total-time">{{ formatTime(totalDuration) }}</span>
      </div>
      
      <el-slider 
        v-model="currentTime" 
        :max="totalDuration" 
        :format-tooltip="formatTime"
        class="time-slider"
        @input="onSeek"
      />
    </div>
    
    <!-- 添加片段对话框 -->
    <el-dialog v-model="dialogVisible" title="添加视频片段" width="500px">
      <el-form :model="clipForm" label-width="80px">
        <el-form-item label="片段名称">
          <el-input v-model="clipForm.name" placeholder="请输入片段名称" />
        </el-form-item>
        <el-form-item label="开始时间">
          <el-input-number v-model="clipForm.startTime" :min="0" :max="totalDuration" />
        </el-form-item>
        <el-form-item label="结束时间">
          <el-input-number v-model="clipForm.endTime" :min="0" :max="totalDuration" />
        </el-form-item>
        <el-form-item label="片段类型">
          <el-select v-model="clipForm.type">
            <el-option label="视频" value="video" />
            <el-option label="图片" value="image" />
            <el-option label="音频" value="audio" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmAddClip">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import { Plus, VideoCamera, VideoPlay, VideoPause } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const props = defineProps<{
  duration?: number
}>()

const totalDuration = ref(props.duration || 60) // 默认60秒
const currentTime = ref(0)
const isPlaying = ref(false)
const zoom = ref(10) // 每秒10像素
const selectedClipId = ref<string | null>(null)
const dialogVisible = ref(false)

const tracks = ref([
  {
    id: '1',
    name: '视频轨道',
    visible: true,
    clips: [
      { id: 'c1', name: '片段1', startTime: 0, endTime: 10, type: 'video' },
      { id: 'c2', name: '片段2', startTime: 12, endTime: 25, type: 'video' }
    ]
  },
  {
    id: '2',
    name: '音频轨道',
    visible: true,
    clips: [
      { id: 'c3', name: '配音1', startTime: 0, endTime: 15, type: 'audio' }
    ]
  }
])

const clipForm = reactive({
  name: '',
  startTime: 0,
  endTime: 10,
  type: 'video'
})

const timelineWidth = computed(() => totalDuration.value * zoom.value)

const playheadPosition = computed(() => currentTime.value * zoom.value)

const timeTicks = computed(() => {
  const ticks = []
  const interval = zoom.value < 5 ? 10 : 5
  for (let i = 0; i <= totalDuration.value; i += interval) {
    ticks.push({ time: i })
  }
  return ticks
})

const formatTime = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60)
  return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

const formatDuration = (seconds: number) => {
  return formatTime(seconds)
}

const getClipColor = (type: string) => {
  const colors: Record<string, string> = {
    video: '#409eff',
    image: '#67c23a',
    audio: '#e6a23c'
  }
  return colors[type] || '#909399'
}

const zoomIn = () => {
  zoom.value = Math.min(zoom.value + 5, 50)
}

const zoomOut = () => {
  zoom.value = Math.max(zoom.value - 5, 5)
}

const togglePlay = () => {
  isPlaying.value = !isPlaying.value
  if (isPlaying.value) {
    startPlayback()
  }
}

let playTimer: number | null = null

const startPlayback = () => {
  playTimer = window.setInterval(() => {
    if (currentTime.value >= totalDuration.value) {
      currentTime.value = 0
      isPlaying.value = false
      if (playTimer) clearInterval(playTimer)
    } else {
      currentTime.value += 0.1
    }
  }, 100)
}

const stop = () => {
  isPlaying.value = false
  currentTime.value = 0
  if (playTimer) clearInterval(playTimer)
}

const onSeek = (val: number) => {
  currentTime.value = val
}

const onTimelineClick = (e: MouseEvent) => {
  const rect = (e.currentTarget as HTMLElement).getBoundingClientRect()
  const x = e.clientX - rect.left
  currentTime.value = Math.max(0, Math.min(x / zoom.value, totalDuration.value))
}

const selectClip = (clip: any) => {
  selectedClipId.value = clip.id
}

const editClip = (clip: any) => {
  ElMessage.info('编辑片段: ' + clip.name)
}

const addClip = () => {
  clipForm.name = ''
  clipForm.startTime = 0
  clipForm.endTime = 10
  dialogVisible.value = true
}

const confirmAddClip = () => {
  if (!clipForm.name) {
    ElMessage.warning('请输入片段名称')
    return
  }
  tracks.value[0].clips.push({
    id: 'c' + Date.now(),
    name: clipForm.name,
    startTime: clipForm.startTime,
    endTime: clipForm.endTime,
    type: clipForm.type
  })
  dialogVisible.value = false
  ElMessage.success('添加成功')
}

const deleteTrack = (trackId: string) => {
  const index = tracks.value.findIndex(t => t.id === trackId)
  if (index > -1) {
    tracks.value.splice(index, 1)
  }
}

const startResize = (e: MouseEvent, clip: any, direction: string) => {
  ElMessage.info('调整大小功能开发中')
}
</script>

<style scoped lang="scss">
.timeline-editor {
  background: var(--bg-card);
  border-radius: 8px;
  padding: 20px;
  
  .timeline-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    
    .zoom-level {
      margin: 0 10px;
      font-size: 14px;
      color: #909399;
    }
  }
  
  .timeline-body {
    border: 1px solid var(--border-primary);
    border-radius: 8px;
    overflow: hidden;
    
    .tracks-container {
      .track {
        border-bottom: 1px solid var(--border-primary);
        
        .track-header {
          display: flex;
          align-items: center;
          gap: 10px;
          padding: 10px 15px;
          background: #f5f7fa;
          font-weight: 500;
        }
        
        .track-content {
          position: relative;
          height: 60px;
          background: #fff;
          
          .playhead {
            position: absolute;
            top: 0;
            bottom: 0;
            width: 2px;
            background: #f56c6c;
            z-index: 10;
            cursor: ew-resize;
            
            &::before {
              content: '';
              position: absolute;
              top: -5px;
              left: -5px;
              width: 0;
              height: 0;
              border-left: 6px solid transparent;
              border-right: 6px solid transparent;
              border-top: 8px solid #f56c6c;
            }
          }
          
          .clip {
            position: absolute;
            top: 10px;
            height: 40px;
            border-radius: 6px;
            cursor: pointer;
            transition: all 0.2s;
            
            &:hover {
              transform: scale(1.02);
              box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
            }
            
            &.selected {
              outline: 2px solid #fff;
              box-shadow: 0 0 0 2px #409eff;
            }
            
            .clip-content {
              padding: 5px 10px;
              color: white;
              font-size: 12px;
              height: 100%;
              display: flex;
              flex-direction: column;
              justify-content: center;
              
              .clip-name {
                font-weight: 500;
              }
              
              .clip-duration {
                opacity: 0.8;
              }
            }
            
            .clip-handle {
              position: absolute;
              top: 0;
              bottom: 0;
              width: 8px;
              cursor: ew-resize;
              
              &.left {
                left: 0;
                background: rgba(255, 255, 255, 0.3);
              }
              
              &.right {
                right: 0;
                background: rgba(255, 255, 255, 0.3);
              }
            }
          }
        }
      }
    }
    
    .time-ruler {
      position: relative;
      height: 30px;
      background: #f5f7fa;
      border-top: 1px solid var(--border-primary);
      
      .tick {
        position: absolute;
        top: 0;
        height: 100%;
        border-left: 1px solid #dcdfe6;
        
        .tick-label {
          position: absolute;
          top: 5px;
          left: 5px;
          font-size: 11px;
          color: #909399;
          white-space: nowrap;
        }
      }
    }
  }
  
  .timeline-footer {
    display: flex;
    align-items: center;
    gap: 20px;
    margin-top: 20px;
    padding: 15px;
    background: #f5f7fa;
    border-radius: 8px;
    
    .time-display {
      font-size: 14px;
      font-family: monospace;
      
      .current-time {
        color: #409eff;
        font-weight: 600;
      }
      
      .separator {
        margin: 0 5px;
        color: #909399;
      }
      
      .total-time {
        color: #606266;
      }
    }
    
    .time-slider {
      flex: 1;
    }
  }
}
</style>
