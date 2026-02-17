<template>
  <div class="export-dialog">
    <el-dialog v-model="visible" title="📤 导出设置" width="600px" @close="handleClose">
      <el-steps :active="step" finish-status="success" simple>
        <el-step title="选择格式" />
        <el-step title="导出设置" />
        <el-step title="开始导出" />
      </el-steps>
      
      <!-- 步骤1: 选择格式 -->
      <div v-if="step === 0" class="step-content">
        <h4>选择导出格式</h4>
        <el-row :gutter="20">
          <el-col :span="8" v-for="format in formats" :key="format.id">
            <div 
              class="format-card" 
              :class="{ selected: selectedFormat === format.id }"
              @click="selectedFormat = format.id"
            >
              <div class="format-icon">{{ format.icon }}</div>
              <h5>{{ format.name }}</h5>
              <p>{{ format.desc }}</p>
            </div>
          </el-col>
        </el-row>
      </div>
      
      <!-- 步骤2: 导出设置 -->
      <div v-if="step === 1" class="step-content">
        <h4>导出设置</h4>
        <el-form label-width="100px">
          <el-form-item label="文件名">
            <el-input v-model="exportSettings.filename" placeholder="请输入文件名" />
          </el-form-item>
          <el-form-item label="视频质量">
            <el-select v-model="exportSettings.quality">
              <el-option label="原画 (1080P)" value="1080p" />
              <el-option label="高清 (720P)" value="720p" />
              <el-option label="标清 (480P)" value="480p" />
            </el-select>
          </el-form-item>
          <el-form-item label="视频格式">
            <el-radio-group v-model="exportSettings.format">
              <el-radio label="mp4">MP4</el-radio>
              <el-radio label="mov">MOV</el-radio>
              <el-radio label="avi">AVI</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="包含字幕">
            <el-switch v-model="exportSettings.includeSubtitles" />
          </el-form-item>
          <el-form-item label="包含配音">
            <el-switch v-model="exportSettings.includeAudio" />
          </el-form-item>
        </el-form>
      </div>
      
      <!-- 步骤3: 开始导出 -->
      <div v-if="step === 2" class="step-content">
        <div v-if="!exporting" class="export-ready">
          <el-result icon="success" title="准备就绪" sub-title="确认设置后开始导出">
            <template #extra>
              <el-button type="primary" @click="startExport">开始导出</el-button>
              <el-button @click="visible = false">取消</el-button>
            </template>
          </el-result>
        </div>
        <div v-else class="exporting">
          <h4>正在导出...</h4>
          <el-progress :percentage="progress" :status="progressStatus" />
          <p class="export-tip">导出过程可能需要几分钟，请耐心等待</p>
        </div>
      </div>
      
      <template #footer>
        <el-button @click="handleClose">取消</el-button>
        <el-button v-if="step > 0" @click="step--">上一步</el-button>
        <el-button v-if="step < 2" type="primary" @click="step++">下一步</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps<{
  modelValue: boolean
}>()

const emit = defineEmits(['update:modelValue', 'export'])

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const step = ref(0)
const selectedFormat = ref('video')
const exporting = ref(false)
const progress = ref(0)

const exportSettings = ref({
  filename: '',
  quality: '1080p',
  format: 'mp4',
  includeSubtitles: true,
  includeAudio: true
})

const progressStatus = computed(() => {
  if (progress.value >= 100) return 'success'
  if (progress.value > 0) return 'warning'
  return undefined
})

const formats = ref([
  { id: 'video', name: '视频', icon: '🎬', desc: '完整视频文件' },
  { id: 'audio', name: '音频', icon: '🎵', desc: '纯音频文件' },
  { id: 'gif', name: 'GIF', icon: '🎞️', desc: '动态图片' },
  { id: 'subtitle', name: '字幕', icon: '📝', desc: 'SRT字幕文件' },
  { id: 'project', name: '工程', icon: '📁', desc: '可编辑项目文件' }
])

const startExport = () => {
  exporting.value = true
  progress.value = 0
  
  const timer = setInterval(() => {
    progress.value += 10
    if (progress.value >= 100) {
      clearInterval(timer)
      ElMessage.success('导出成功！')
      emit('export', exportSettings.value)
      handleClose()
    }
  }, 500)
}

const handleClose = () => {
  visible.value = false
  step.value = 0
  exporting.value = false
  progress.value = 0
}
</script>

<style scoped lang="scss">
.export-dialog {
  .step-content {
    padding: 30px 0;
    
    h4 {
      margin: 0 0 20px;
    }
  }
  
  .format-card {
    padding: 20px;
    border: 2px solid #eee;
    border-radius: 12px;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s;
    
    &:hover {
      border-color: #409eff;
    }
    
    &.selected {
      border-color: #409eff;
      background: rgba(64, 158, 255, 0.1);
    }
    
    .format-icon {
      font-size: 36px;
      margin-bottom: 10px;
    }
    
    h5 {
      margin: 0 0 8px;
    }
    
    p {
      margin: 0;
      font-size: 12px;
      color: #909399;
    }
  }
  
  .exporting {
    text-align: center;
    
    h4 {
      margin-bottom: 20px;
    }
    
    .export-tip {
      margin-top: 20px;
      color: #909399;
      font-size: 14px;
    }
  }
}
</style>
