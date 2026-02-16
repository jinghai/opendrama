<template>
  <div class="drag-sort-list">
    <div class="list-header">
      <el-checkbox 
        :model-value="isAllSelected" 
        @change="handleSelectAll"
        indeterminate
      >
        已选 {{ selectedCount }} 项
      </el-checkbox>
      <div class="header-actions">
        <el-button size="small" @click="moveUp" :disabled="selectedCount === 0">
          <el-icon><Top /></el-icon>
          上移
        </el-button>
        <el-button size="small" @click="moveDown" :disabled="selectedCount === 0">
          <el-icon><Bottom /></el-icon>
          下移
        </el-button>
        <el-button size="small" type="danger" @click="deleteSelected" :disabled="selectedCount === 0">
          <el-icon><Delete /></el-icon>
          批量删除
        </el-button>
      </div>
    </div>
    
    <el-scrollbar height="400px">
      <draggable 
        v-model="localItems" 
        item-key="id"
        handle=".drag-handle"
        @end="onDragEnd"
        class="sort-list"
      >
        <template #item="{ element, index }">
          <div 
            class="sort-item" 
            :class="{ selected: element.selected }"
            @click="toggleSelect(element)"
          >
            <div class="drag-handle">
              <el-icon><Rank /></el-icon>
            </div>
            <div class="item-index">{{ index + 1 }}</div>
            <div class="item-content">
              <div class="item-icon" v-if="element.type === 'image'">
                <el-icon><Picture /></el-icon>
              </div>
              <div class="item-icon video" v-else-if="element.type === 'video'">
                <el-icon><VideoPlay /></el-icon>
              </div>
              <div class="item-icon audio" v-else-if="element.type === 'audio'">
                <el-icon><Microphone /></el-icon>
              </div>
              <div class="item-info">
                <p class="item-name">{{ element.name }}</p>
                <p class="item-desc">{{ element.description || element.prompt || '无描述' }}</p>
              </div>
            </div>
            <div class="item-actions">
              <el-button size="small" text @click.stop="editItem(element)">编辑</el-button>
              <el-button size="small" text type="danger" @click.stop="deleteItem(element)">删除</el-button>
            </div>
          </div>
        </template>
      </draggable>
    </el-scrollbar>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Top, Bottom, Delete, Rank, Picture, VideoPlay, Microphone } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import draggable from 'vuedraggable'

const props = defineProps<{
  items: any[]
}>()

const emit = defineEmits(['update:items', 'delete', 'edit', 'reorder'])

const localItems = ref([...props.items])

watch(() => props.items, (newItems) => {
  localItems.value = [...newItems]
}, { deep: true })

const isAllSelected = computed(() => {
  return localItems.value.length > 0 && localItems.value.every(i => i.selected)
})

const selectedCount = computed(() => {
  return localItems.value.filter(i => i.selected).length
})

const handleSelectAll = (val: boolean) => {
  localItems.value.forEach(i => i.selected = val)
}

const toggleSelect = (item: any) => {
  item.selected = !item.selected
}

const onDragEnd = () => {
  emit('update:items', localItems.value)
  emit('reorder', localItems.value)
}

const moveUp = () => {
  const selected = localItems.value.filter(i => i.selected)
  if (selected.length === 0) return
  
  selected.forEach(item => {
    const index = localItems.value.findIndex(i => i.id === item.id)
    if (index > 0) {
      const temp = localItems.value[index]
      localItems.value[index] = localItems.value[index - 1]
      localItems.value[index - 1] = temp
    }
  })
  emit('update:items', localItems.value)
}

const moveDown = () => {
  const selected = localItems.value.filter(i => i.selected)
  if (selected.length === 0) return
  
  selected.reverse().forEach(item => {
    const index = localItems.value.findIndex(i => i.id === item.id)
    if (index < localItems.value.length - 1) {
      const temp = localItems.value[index]
      localItems.value[index] = localItems.value[index + 1]
      localItems.value[index + 1] = temp
    }
  })
  emit('update:items', localItems.value)
}

const deleteSelected = () => {
  ElMessageBox.confirm(`确定要删除选中的 ${selectedCount.value} 项吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    const toDelete = localItems.value.filter(i => i.selected)
    emit('delete', toDelete.map(i => i.id))
    ElMessage.success(`已删除 ${toDelete.length} 项`)
  }).catch(() => {})
}

const editItem = (item: any) => {
  emit('edit', item)
}

const deleteItem = (item: any) => {
  ElMessageBox.confirm(`确定要删除「${item.name}」吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    emit('delete', [item.id])
    ElMessage.success('删除成功')
  }).catch(() => {})
}
</script>

<style scoped lang="scss">
.drag-sort-list {
  .list-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
    padding: 10px;
    background: var(--bg-card);
    border-radius: 8px;
    
    .header-actions {
      display: flex;
      gap: 10px;
    }
  }
  
  .sort-list {
    .sort-item {
      display: flex;
      align-items: center;
      padding: 12px;
      margin-bottom: 10px;
      background: var(--bg-card);
      border: 1px solid var(--border-primary);
      border-radius: 8px;
      cursor: pointer;
      transition: all 0.2s;
      
      &:hover {
        border-color: #409eff;
      }
      
      &.selected {
        background: rgba(64, 158, 255, 0.1);
        border-color: #409eff;
      }
      
      .drag-handle {
        cursor: move;
        padding: 5px;
        color: #909399;
        
        &:hover {
          color: #409eff;
        }
      }
      
      .item-index {
        width: 30px;
        text-align: center;
        font-weight: 600;
        color: #909399;
      }
      
      .item-content {
        flex: 1;
        display: flex;
        align-items: center;
        gap: 15px;
        
        .item-icon {
          width: 50px;
          height: 50px;
          display: flex;
          align-items: center;
          justify-content: center;
          background: #f5f7fa;
          border-radius: 8px;
          font-size: 24px;
          color: #409eff;
          
          &.video {
            color: #67c23a;
          }
          
          &.audio {
            color: #e6a23c;
          }
        }
        
        .item-info {
          flex: 1;
          
          .item-name {
            margin: 0;
            font-weight: 500;
          }
          
          .item-desc {
            margin: 5px 0 0;
            font-size: 12px;
            color: #909399;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            max-width: 300px;
          }
        }
      }
      
      .item-actions {
        display: flex;
        gap: 5px;
      }
    }
  }
}
</style>
