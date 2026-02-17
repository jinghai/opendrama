<template>
  <div class="storyboard-grid">
    <div 
      v-for="(scene, index) in scenes" 
      :key="scene.id"
      class="storyboard-card"
      :class="{ active: scene.id === activeSceneId }"
      @click="selectScene(scene)"
    >
      <div class="card-number">{{ index + 1 }}</div>
      
      <div class="card-image">
        <img v-if="scene.image" :src="scene.image" :alt="scene.title" />
        <div v-else class="image-placeholder">
          <el-icon :size="30"><Picture /></el-icon>
        </div>
      </div>
      
      <div class="card-content">
        <h4>{{ scene.title }}</h4>
        <p class="scene-desc">{{ scene.description }}</p>
        <div class="scene-meta">
          <el-tag size="small">{{ scene.location }}</el-tag>
          <el-tag size="small" :type="timeOfDayType(scene.timeOfDay)">
            {{ scene.timeOfDay }}
          </el-tag>
        </div>
        <div class="scene-dialogue" v-if="scene.dialogue">
          <p class="dialogue-text">"{{ scene.dialogue }}"</p>
          <p class="dialogue-speaker">— {{ scene.speaker }}</p>
        </div>
      </div>
      
      <div class="card-actions">
        <el-button size="small" circle @click.stop="editScene(scene)">
          <el-icon><Edit /></el-icon>
        </el-button>
        <el-button size="small" circle @click.stop="generateImage(scene)">
          <el-icon><PictureFilled /></el-icon>
        </el-button>
        <el-button size="small" circle type="danger" @click.stop="deleteScene(index)">
          <el-icon><Delete /></el-icon>
        </el-button>
      </div>
    </div>
    
    <!-- 添加场景按钮 -->
    <div class="add-card" @click="addScene">
      <el-icon :size="40"><Plus /></el-icon>
      <p>添加分镜</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Picture, Edit, PictureFilled, Delete, Plus } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const props = defineProps<{
  scenes: any[]
}>()

const emit = defineEmits(['select', 'edit', 'delete', 'add', 'generate-image'])

const activeSceneId = ref<string | null>(null)

const timeOfDayType = (time: string) => {
  const map: Record<string, string> = {
    '早晨': 'warning',
    '上午': 'success',
    '中午': 'danger',
    '下午': 'info',
    '傍晚': 'warning',
    '夜晚': 'primary'
  }
  return map[time] || 'info'
}

const selectScene = (scene: any) => {
  activeSceneId.value = scene.id
  emit('select', scene)
}

const editScene = (scene: any) => {
  emit('edit', scene)
}

const generateImage = (scene: any) => {
  emit('generate-image', scene)
}

const deleteScene = (index: number) => {
  ElMessageBox.confirm('确定要删除这个分镜吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    emit('delete', index)
    ElMessage.success('删除成功')
  }).catch(() => {})
}

const addScene = () => {
  emit('add')
}
</script>

<style scoped lang="scss">
.storyboard-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  padding: 20px;
  
  .storyboard-card {
    background: var(--bg-card);
    border-radius: 12px;
    overflow: hidden;
    border: 2px solid transparent;
    cursor: pointer;
    transition: all 0.3s;
    position: relative;
    
    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
      
      .card-actions {
        opacity: 1;
      }
    }
    
    &.active {
      border-color: #409eff;
    }
    
    .card-number {
      position: absolute;
      top: 10px;
      left: 10px;
      width: 30px;
      height: 30px;
      background: rgba(0, 0, 0, 0.6);
      color: white;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-weight: 600;
      z-index: 1;
    }
    
    .card-image {
      height: 160px;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
      
      .image-placeholder {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        opacity: 0.5;
      }
    }
    
    .card-content {
      padding: 15px;
      
      h4 {
        margin: 0 0 8px;
        font-size: 16px;
      }
      
      .scene-desc {
        margin: 0 0 10px;
        font-size: 13px;
        color: #909399;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
      }
      
      .scene-meta {
        display: flex;
        gap: 8px;
        margin-bottom: 10px;
      }
      
      .scene-dialogue {
        background: #f5f7fa;
        padding: 10px;
        border-radius: 8px;
        
        .dialogue-text {
          margin: 0;
          font-style: italic;
          font-size: 13px;
        }
        
        .dialogue-speaker {
          margin: 5px 0 0;
          text-align: right;
          font-size: 12px;
          color: #909399;
        }
      }
    }
    
    .card-actions {
      position: absolute;
      top: 10px;
      right: 10px;
      display: flex;
      gap: 5px;
      opacity: 0;
      transition: opacity 0.3s;
    }
  }
  
  .add-card {
    background: var(--bg-card);
    border: 2px dashed #dcdfe6;
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    min-height: 300px;
    cursor: pointer;
    transition: all 0.3s;
    
    &:hover {
      border-color: #409eff;
      background: rgba(64, 158, 255, 0.05);
    }
    
    p {
      margin: 10px 0 0;
      color: #909399;
    }
  }
}
</style>
