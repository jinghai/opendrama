<template>
  <div class="audio-manager">
    <div class="toolbar">
      <el-input v-model="searchText" placeholder="搜索音频..." prefix-icon="Search" clearable style="width: 300px" />
      <el-button type="primary" @click="uploadAudio" style="margin-left: auto">
        <el-icon><Upload /></el-icon> 上传音频
      </el-button>
    </div>

    <el-table :data="filteredAudios" stripe>
      <el-table-column prop="name" label="音频名称" width="200" />
      <el-table-column prop="type" label="类型" width="100">
        <template #default="{ row }">
          <el-tag>{{ getTypeText(row.type) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="duration" label="时长" width="100">
        <template #default="{ row }">
          {{ formatDuration(row.duration) }}
        </template>
      </el-table-column>
      <el-table-column prop="size" label="大小" width="100">
        <template #default="{ row }">
          {{ formatSize(row.size) }}
        </template>
      </el-table-column>
      <el-table-column prop="createdAt" label="上传时间" width="180" />
      <el-table-column label="操作" width="250">
        <template #default="{ row }">
          <el-button size="small" @click="playAudio(row)">播放</el-button>
          <el-button size="small" @click="downloadAudio(row)">下载</el-button>
          <el-button size="small" type="danger" text @click="deleteAudio(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 音频播放器 -->
    <el-dialog v-model="playerVisible" title="播放音频" width="500px">
      <div class="player-content">
        <h4>{{ currentAudio?.name }}</h4>
        <audio ref="audioPlayer" :src="currentAudio?.url" controls style="width: 100%" />
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Upload } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchText = ref('')
const playerVisible = ref(false)
const currentAudio = ref<any>(null)
const audioPlayer = ref<HTMLAudioElement>()

const audios = ref([
  { id: '1', name: '背景音乐1', type: 'bgm', duration: 180, size: 2048000, url: '', createdAt: '2026-02-15 10:30' },
  { id: '2', name: '配音文件', type: 'tts', duration: 60, size: 512000, url: '', createdAt: '2026-02-16 14:20' },
  { id: '3', name: '音效-门铃声', type: 'effect', duration: 5, size: 51200, url: '', createdAt: '2026-02-17 09:15' }
])

const filteredAudios = computed(() => {
  return audios.value.filter(a => !searchText.value || a.name.includes(searchText.value))
})

const getTypeText = (type: string) => {
  const map: Record<string, string> = { bgm: '背景音乐', tts: '配音', effect: '音效', voice: '原声' }
  return map[type] || type
}

const formatDuration = (seconds: number) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

const formatSize = (bytes: number) => {
  return (bytes / 1024).toFixed(1) + ' KB'
}

const uploadAudio = () => ElMessage.info('上传音频功能开发中')
const playAudio = (audio: any) => { currentAudio.value = audio; playerVisible.value = true }
const downloadAudio = (audio: any) => ElMessage.info('下载: ' + audio.name)
const deleteAudio = (audio: any) => {
  ElMessageBox.confirm('确定删除这个音频吗?', '提示', { type: 'warning' })
    .then(() => { audios.value = audios.value.filter(x => x.id !== audio.id); ElMessage.success('删除成功') })
    .catch(() => {})
}
</script>

<style scoped>
.toolbar { display: flex; align-items: center; margin-bottom: 20px; }
.player-content { text-align: center; }
.player-content h4 { margin-bottom: 15px; }
</style>
