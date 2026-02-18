<template>
  <div class="task-center">
    <el-tabs v-model="activeTab">
      <el-tab-pane label="所有任务" name="all">
        <task-list :tasks="allTasks" @refresh="loadTasks" />
      </el-tab-pane>
      <el-tab-pane label="进行中" name="processing">
        <task-list :tasks="processingTasks" @refresh="loadTasks" />
      </el-tab-pane>
      <el-tab-pane label="已完成" name="completed">
        <task-list :tasks="completedTasks" @refresh="loadTasks" />
      </el-tab-pane>
      <el-tab-pane label="失败" name="failed">
        <task-list :tasks="failedTasks" @refresh="loadTasks" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import TaskList from './components/TaskList.vue'

const activeTab = ref('all')
const tasks = ref([
  { id: '1', name: '生成图片-角色1', type: 'image', status: 'completed', progress: 100, createdAt: '2026-02-18 10:00', completedAt: '2026-02-18 10:05' },
  { id: '2', name: '生成视频-第3集', type: 'video', status: 'processing', progress: 65, createdAt: '2026-02-18 11:00' },
  { id: '3', name: 'TTS配音', type: 'tts', status: 'pending', progress: 0, createdAt: '2026-02-18 11:30' },
  { id: '4', name: '提取音频', type: 'audio', status: 'failed', progress: 0, createdAt: '2026-02-18 09:00', error: '网络超时' }
])

const allTasks = computed(() => tasks.value)
const processingTasks = computed(() => tasks.value.filter(t => t.status === 'processing' || t.status === 'pending'))
const completedTasks = computed(() => tasks.value.filter(t => t.status === 'completed'))
const failedTasks = computed(() => tasks.value.filter(t => t.status === 'failed'))

const loadTasks = () => { console.log('刷新任务列表') }
</script>
