<template>
  <div class="drama-card" @click="goToDrama">
    <div class="drama-cover">
      <img v-if="drama.cover_image" :src="drama.cover_image" :alt="drama.title" />
      <div v-else class="cover-placeholder">
        <el-icon :size="40"><VideoCamera /></el-icon>
      </div>
      <div class="cover-overlay">
        <el-tag v-if="drama.status" :type="statusType" size="small">
          {{ statusText }}
        </el-tag>
      </div>
    </div>
    <div class="drama-info">
      <h3 class="drama-title">{{ drama.title }}</h3>
      <p class="drama-desc">{{ drama.description || '暂无描述' }}</p>
      <div class="drama-meta">
        <span><el-icon><Collection /></el-icon> {{ drama.episodes?.length || 0 }} 集</span>
        <span><el-icon><User /></el-icon> {{ drama.characters?.length || 0 }} 角色</span>
      </div>
      <div class="drama-footer">
        <span class="update-time">{{ formatTime(drama.updated_at) }}</span>
        <el-dropdown trigger="click" @command="handleCommand">
          <el-button text size="small" @click.stop>
            <el-icon><MoreFilled /></el-icon>
          </el-button>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="edit">编辑</el-dropdown-item>
              <el-dropdown-item command="duplicate">复制</el-dropdown-item>
              <el-dropdown-item command="delete" divided>删除</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { VideoCamera, Collection, User, MoreFilled } from '@element-plus/icons-vue'
import { ElMessageBox, ElMessage } from 'element-plus'

const props = defineProps<{
  drama: any
}>()

const emit = defineEmits(['delete', 'duplicate'])

const router = useRouter()

const statusType = computed(() => {
  const map: Record<string, string> = {
    draft: 'info',
    in_progress: 'warning',
    completed: 'success',
    published: 'primary'
  }
  return map[props.drama.status] || 'info'
})

const statusText = computed(() => {
  const map: Record<string, string> = {
    draft: '草稿',
    in_progress: '进行中',
    completed: '已完成',
    published: '已发布'
  }
  return map[props.drama.status] || '未知'
})

const formatTime = (time: string) => {
  if (!time) return ''
  const date = new Date(time)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  if (days === 0) return '今天'
  if (days === 1) return '昨天'
  if (days < 7) return `${days}天前`
  return date.toLocaleDateString()
}

const goToDrama = () => {
  router.push(`/dramas/${props.drama.id}`)
}

const handleCommand = (command: string) => {
  switch (command) {
    case 'edit':
      router.push(`/dramas/${props.drama.id}`)
      break
    case 'duplicate':
      emit('duplicate', props.drama.id)
      ElMessage.success('复制功能开发中')
      break
    case 'delete':
      ElMessageBox.confirm('确定要删除这个短剧吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        emit('delete', props.drama.id)
      }).catch(() => {})
      break
  }
}
</script>

<style scoped lang="scss">
.drama-card {
  background: var(--bg-card);
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.3s;
  border: 1px solid var(--border-primary);
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  }
  
  .drama-cover {
    position: relative;
    height: 160px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    
    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    
    .cover-placeholder {
      width: 100%;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      opacity: 0.5;
    }
    
    .cover-overlay {
      position: absolute;
      top: 10px;
      right: 10px;
    }
  }
  
  .drama-info {
    padding: 15px;
    
    .drama-title {
      font-size: 16px;
      font-weight: 600;
      margin: 0 0 8px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
    
    .drama-desc {
      font-size: 13px;
      color: var(--text-secondary);
      margin: 0 0 12px;
      overflow: hidden;
      text-overflow: ellipsis;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
      height: 40px;
    }
    
    .drama-meta {
      display: flex;
      gap: 15px;
      font-size: 12px;
      color: var(--text-secondary);
      margin-bottom: 12px;
      
      span {
        display: flex;
        align-items: center;
        gap: 4px;
      }
    }
    
    .drama-footer {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding-top: 12px;
      border-top: 1px solid var(--border-primary);
      
      .update-time {
        font-size: 12px;
        color: var(--text-secondary);
      }
    }
  }
}
</style>
