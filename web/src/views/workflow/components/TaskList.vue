<template>
  <div class="task-list">
    <el-table :data="tasks" stripe>
      <el-table-column prop="name" label="任务名称" width="200" />
      <el-table-column prop="type" label="类型" width="100">
        <template #default="{ row }">
          <el-tag>{{ getTypeText(row.type) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)">{{ getStatusText(row.status) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="progress" label="进度" width="150">
        <template #default="{ row }">
          <el-progress :percentage="row.progress" :status="row.status === 'failed' ? 'exception' : undefined" />
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="创建时间" width="180" />
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button v-if="row.status === 'failed'" size="small" @click="retryTask(row)">重试</el-button>
          <el-button v-if="row.status === 'completed'" size="small" @click="viewResult(row)">查看结果</el-button>
          <el-button v-if="row.status !== 'completed' && row.status !== 'failed'" size="small" type="danger" text @click="cancelTask(row)">取消</el-button>
          <el-button size="small" type="danger" text @click="deleteTask(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination v-if="tasks.length > 0" layout="prev, pager, next" :total="tasks.length" style="margin-top: 20px; justify-content: center" />
  </div>
</template>

<script setup lang="ts">
import { ElMessage, ElMessageBox } from 'element-plus'

const props = defineProps<{ tasks: any[] }>()
const emit = defineEmits(['refresh'])

const getTypeText = (type: string) => {
  const map: Record<string, string> = { image: '图片生成', video: '视频生成', tts: '语音合成', audio: '音频提取' }
  return map[type] || type
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = { pending: '等待中', processing: '进行中', completed: '已完成', failed: '失败', cancelled: '已取消' }
  return map[status] || status
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = { pending: 'info', processing: 'warning', completed: 'success', failed: 'danger', cancelled: 'info' }
  return map[status] || 'info'
}

const retryTask = (task: any) => ElMessage.info('重试: ' + task.name)
const viewResult = (task: any) => ElMessage.info('查看结果: ' + task.name)
const cancelTask = (task: any) => ElMessage.info('取消: ' + task.name)
const deleteTask = (task: any) => {
  ElMessageBox.confirm('确定删除这个任务吗?', '提示', { type: 'warning' })
    .then(() => { emit('refresh') })
    .catch(() => {})
}
</script>

<style scoped></style>
