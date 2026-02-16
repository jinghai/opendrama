<template>
  <div class="quick-actions">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover" class="action-card" @click="goTo('/dramas/create')">
          <div class="action-icon" style="background: #409eff">
            <el-icon :size="30"><Plus /></el-icon>
          </div>
          <div class="action-content">
            <h3>创建短剧</h3>
            <p>开始新的短剧项目</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="action-card" @click="goTo('/settings/newapi')">
          <div class="action-icon" style="background: #67c23a">
            <el-icon :size="30"><Connection /></el-icon>
          </div>
          <div class="action-content">
            <h3>NewAPI配置</h3>
            <p>管理AI服务商</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="action-card" @click="showTTSDialog">
          <div class="action-icon" style="background: #e6a23c">
            <el-icon :size="30"><Microphone /></el-icon>
          </div>
          <div class="action-content">
            <h3>TTS配音</h3>
            <p>语音合成设置</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="6">
        <el-card shadow="hover" class="action-card" @click="goTo('/settings/ai-config')">
          <div class="action-icon" style="background: #f56c6c">
            <el-icon :size="30"><Setting /></el-icon>
          </div>
          <div class="action-content">
            <h3>AI配置</h3>
            <p>配置AI服务</p>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { Plus, Connection, Microphone, Setting } from '@element-plus/icons-vue'

const router = useRouter()
const emit = defineEmits(['open-tts'])

const goTo = (path: string) => {
  router.push(path)
}

const showTTSDialog = () => {
  emit('open-tts')
}
</script>

<style scoped lang="scss">
.quick-actions {
  margin: 20px 0;
  
  .action-card {
    cursor: pointer;
    display: flex;
    align-items: center;
    padding: 20px;
    transition: transform 0.2s;
    
    &:hover {
      transform: translateY(-5px);
    }
    
    .action-icon {
      width: 60px;
      height: 60px;
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      margin-right: 15px;
    }
    
    .action-content {
      h3 {
        font-size: 16px;
        font-weight: 600;
        margin: 0;
        color: #303133;
      }
      
      p {
        margin: 5px 0 0;
        font-size: 12px;
        color: #909399;
      }
    }
  }
}
</style>
