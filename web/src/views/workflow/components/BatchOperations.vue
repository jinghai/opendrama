<template>
  <div class="batch-operations">
    <div class="header">
      <h3>📚 批量操作</h3>
      <el-button type="primary" @click="openDialog">
        <el-icon><Plus /></el-icon>
        新建任务
      </el-button>
    </div>
    
    <el-table :data="tasks" stripe>
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="type" label="类型">
        <template #default="{ row }">
          <el-tag>{{ typeMap[row.type] || row.type }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="statusType[row.status]">
            {{ statusMap[row.status] || row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="progress" label="进度" width="150">
        <template #default="{ row }">
          <el-progress :percentage="row.progress || 0" :status="row.status === 'failed' ? 'exception' : undefined" />
        </template>
      </el-table-column>
      <el-table-column prop="created_at" label="创建时间" width="180">
        <template #default="{ row }">
          {{ formatTime(row.created_at) }}
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template #default="{ row }">
          <el-button-group>
            <el-button size="small" @click="viewDetail(row)">详情</el-button>
            <el-button size="small" type="danger" text @click="cancelTask(row)" v-if="row.status === 'pending' || row.status === 'processing'">
              取消
            </el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
    
    <!-- 新建批量任务对话框 -->
    <el-dialog v-model="dialogVisible" title="新建批量任务" width="600px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="任务类型">
          <el-select v-model="form.type" placeholder="选择任务类型">
            <el-option label="批量生成图片" value="batch_image" />
            <el-option label="批量生成视频" value="batch_video" />
            <el-option label="批量配音" value="batch_tts" />
            <el-option label="批量提取角色" value="batch_character" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="任务描述">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="可选填写任务描述" />
        </el-form-item>
        
        <el-form-item label="数据来源">
          <el-radio-group v-model="form.source">
            <el-radio label="selected">选中项</el-radio>
            <el-radio label="all">全部</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="createTask">创建</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const dialogVisible = ref(false)
const tasks = ref<any[]>([])

const form = reactive({
  type: '',
  description: '',
  source: 'selected'
})

const typeMap: Record<string, string> = {
  batch_image: '批量图片',
  batch_video: '批量视频',
  batch_tts: '批量配音',
  batch_character: '批量角色'
}

const statusMap: Record<string, string> = {
  pending: '等待中',
  processing: '处理中',
  completed: '已完成',
  failed: '失败',
  cancelled: '已取消'
}

const statusType: Record<string, string> = {
  pending: 'info',
  processing: 'warning',
  completed: 'success',
  failed: 'danger',
  cancelled: 'info'
}

const formatTime = (time: string) => {
  if (!time) return '-'
  return new Date(time).toLocaleString()
}

const openDialog = () => {
  form.type = ''
  form.description = ''
  form.source = 'selected'
  dialogVisible.value = true
}

const createTask = () => {
  if (!form.type) {
    ElMessage.warning('请选择任务类型')
    return
  }
  ElMessage.success('任务创建成功')
  dialogVisible.value = false
}

const viewDetail = (row: any) => {
  ElMessage.info('详情功能开发中')
}

const cancelTask = (row: any) => {
  ElMessage.info('取消功能开发中')
}
</script>

<style scoped lang="scss">
.batch-operations {
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    
    h3 {
      margin: 0;
    }
  }
}
</style>
