<template>
  <div class="image-gallery">
    <el-row :gutter="20">
      <el-col :span="6" v-for="img in images" :key="img.id">
        <el-card shadow="hover" class="image-card" :class="{ selected: img.selected }" @click="selectImage(img)">
          <div class="image-wrapper">
            <img :src="img.url" :alt="img.name" />
            <div class="image-overlay">
              <el-checkbox v-model="img.selected" @click.stop />
              <div class="image-actions">
                <el-button size="small" circle @click.stop="previewImage(img)">
                  <el-icon><ZoomIn /></el-icon>
                </el-button>
                <el-button size="small" circle @click.stop="downloadImage(img)">
                  <el-icon><Download /></el-icon>
                </el-button>
                <el-button size="small" circle type="danger" @click.stop="deleteImage(img)">
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </div>
          </div>
          <div class="image-info">
            <p class="image-name">{{ img.name }}</p>
            <p class="image-time">{{ formatTime(img.created_at) }}</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
    
    <!-- 图片预览对话框 -->
    <el-dialog v-model="previewVisible" width="800px">
      <img :src="currentImage?.url" style="width: 100%" />
      <template #footer>
        <el-button @click="previewVisible = false">关闭</el-button>
        <el-button type="primary" @click="downloadImage(currentImage)">下载</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ZoomIn, Download, Delete } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const props = defineProps<{
  images: any[]
}>()

const emit = defineEmits(['select', 'delete'])

const previewVisible = ref(false)
const currentImage = ref<any>(null)

const formatTime = (time: string) => {
  if (!time) return ''
  return new Date(time).toLocaleString()
}

const selectImage = (img: any) => {
  img.selected = !img.selected
  emit('select', props.images.filter(i => i.selected))
}

const previewImage = (img: any) => {
  currentImage.value = img
  previewVisible.value = true
}

const downloadImage = (img: any) => {
  window.open(img.url, '_blank')
}

const deleteImage = (img: any) => {
  ElMessageBox.confirm('确定要删除这张图片吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    emit('delete', img.id)
    ElMessage.success('删除成功')
  }).catch(() => {})
}
</script>

<style scoped lang="scss">
.image-gallery {
  .image-card {
    margin-bottom: 20px;
    cursor: pointer;
    transition: all 0.3s;
    
    &.selected {
      border-color: #409eff;
    }
    
    .image-wrapper {
      position: relative;
      height: 150px;
      overflow: hidden;
      border-radius: 8px;
      
      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
      
      .image-overlay {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background: rgba(0, 0, 0, 0.5);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        padding: 10px;
        opacity: 0;
        transition: opacity 0.3s;
        
        &:hover {
          opacity: 1;
        }
        
        .image-actions {
          display: flex;
          justify-content: center;
          gap: 10px;
        }
      }
    }
    
    .image-info {
      padding: 10px 0;
      
      .image-name {
        margin: 0;
        font-size: 14px;
        font-weight: 500;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      
      .image-time {
        margin: 5px 0 0;
        font-size: 12px;
        color: #909399;
      }
    }
  }
}
</style>
