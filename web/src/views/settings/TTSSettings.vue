<template>
  <div class="tts-settings">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>🎙️ TTS语音合成设置</span>
        </div>
      </template>
      
      <el-form :model="form" label-width="120px">
        <el-form-item label="选择提供商">
          <el-select v-model="form.provider" placeholder="选择TTS提供商" @change="loadVoices">
            <el-option label="Azure TTS" value="azure" />
            <el-option label="阿里云TTS" value="alibaba" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="选择语音">
          <el-select v-model="form.voice" placeholder="选择语音角色" filterable>
            <el-option
              v-for="voice in voices"
              :key="voice.id"
              :label="`${voice.name} (${voice.gender})`"
              :value="voice.id"
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="语速">
          <el-slider v-model="form.speed" :min="0.5" :max="2" :step="0.1" show-stops />
          <span class="speed-value">{{ form.speed }}x</span>
        </el-form-item>
        
        <el-form-item label="音调">
          <el-slider v-model="form.pitch" :min="0.5" :max="2" :step="0.1" show-stops />
          <span class="pitch-value">{{ form.pitch }}</span>
        </el-form-item>
        
        <el-form-item label="输入文本">
          <el-input
            v-model="form.text"
            type="textarea"
            :rows="4"
            placeholder="请输入需要转换为语音的文本..."
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="generateSpeech" :loading="loading">
            🎙️ 生成语音
          </el-button>
          <el-button @click="testVoice">测试语音</el-button>
        </el-form-item>
      </el-form>
      
      <!-- 生成结果 -->
      <el-divider v-if="audioUrl" />
      
      <div v-if="audioUrl" class="result-section">
        <h4>🎉 生成结果</h4>
        <audio :src="audioUrl" controls class="audio-player" />
        <el-button type="success" @click="downloadAudio">
          ⬇️ 下载音频
        </el-button>
      </div>
    </el-card>
    
    <!-- 批量生成 -->
    <el-card class="batch-card">
      <template #header>
        <div class="card-header">
          <span>📚 批量语音生成</span>
        </div>
      </template>
      
      <el-form :model="batchForm" label-width="120px">
        <el-form-item label="提供商">
          <el-select v-model="batchForm.provider" @change="loadVoices">
            <el-option label="Azure TTS" value="azure" />
            <el-option label="阿里云TTS" value="alibaba" />
          </el-select>
        </el-form-item>
        
        <el-form-item label="语音">
          <el-select v-model="batchForm.voice" filterable>
            <el-option
              v-for="voice in voices"
              :key="voice.id"
              :label="voice.name"
              :value="voice.id"
            />
          </el-select>
        </el-form-item>
        
        <el-form-item label="文本列表">
          <el-input
            v-model="batchForm.textsText"
            type="textarea"
            :rows="6"
            placeholder="每行一个文本片段..."
          />
        </el-form-item>
        
        <el-form-item>
          <el-button type="primary" @click="batchGenerate" :loading="batchLoading">
            📚 批量生成
          </el-button>
        </el-form-item>
      </el-form>
      
      <!-- 批量结果 -->
      <el-divider v-if="batchResults.length" />
      
      <div v-if="batchResults.length" class="batch-results">
        <h4>📊 批量生成结果</h4>
        <el-table :data="batchResults" stripe>
          <el-table-column prop="index" label="序号" width="60" />
          <el-table-column prop="text" label="文本" show-overflow-tooltip />
          <el-table-column prop="success" label="状态" width="80">
            <template #default="{ row }">
              <el-tag :type="row.success ? 'success' : 'danger'">
                {{ row.success ? '成功' : '失败' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="duration" label="时长" width="80" />
        </el-table>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { generateTTS, getTTSVoices, getTTSProviders, batchGenerateTTS } from '@/api/tts'

const loading = ref(false)
const batchLoading = ref(false)
const audioUrl = ref('')
const voices = ref<any[]>([])

const form = reactive({
  provider: 'azure',
  voice: 'zh-CN-XiaoxiaoNeural',
  speed: 1.0,
  pitch: 1.0,
  text: ''
})

const batchForm = reactive({
  provider: 'azure',
  voice: 'zh-CN-XiaoxiaoNeural',
  textsText: '',
})

const batchResults = ref<any[]>([])

// 加载提供商列表
onMounted(async () => {
  try {
    const providers = await getTTSProviders()
    if (providers.length > 0) {
      form.provider = providers[0]
      batchForm.provider = providers[0]
    }
    await loadVoices()
  } catch (error) {
    console.error('加载提供商失败:', error)
  }
})

// 加载语音列表
const loadVoices = async () => {
  try {
    voices.value = await getTTSVoices(form.provider)
    if (voices.value.length > 0) {
      form.voice = voices.value[0].id
      batchForm.voice = voices.value[0].id
    }
  } catch (error) {
    console.error('加载语音列表失败:', error)
    ElMessage.error('加载语音列表失败')
  }
}

// 生成语音
const generateSpeech = async () => {
  if (!form.text.trim()) {
    ElMessage.warning('请输入文本内容')
    return
  }
  
  loading.value = true
  try {
    const result = await generateTTS({
      provider: form.provider,
      voice: form.voice,
      text: form.text,
      speed: form.speed,
      pitch: form.pitch,
      save_to_file: true
    })
    
    if (result.success) {
      audioUrl.value = result.file_path
      ElMessage.success('语音生成成功!')
    } else {
      ElMessage.error(result.error || '生成失败')
    }
  } catch (error: any) {
    ElMessage.error(error.message || '生成失败')
  } finally {
    loading.value = false
  }
}

// 测试语音
const testVoice = async () => {
  form.text = '你好!这是一个测试语音。'
  await generateSpeech()
}

// 下载音频
const downloadAudio = () => {
  if (audioUrl.value) {
    window.open(audioUrl.value, '_blank')
  }
}

// 批量生成
const batchGenerate = async () => {
  const texts = batchForm.textsText.split('\n').filter(t => t.trim())
  if (texts.length === 0) {
    ElMessage.warning('请输入至少一个文本')
    return
  }
  
  batchLoading.value = true
  try {
    const result = await batchGenerateTTS({
      provider: batchForm.provider,
      voice: batchForm.voice,
      texts: texts,
      speed: form.speed,
      pitch: form.pitch,
      save_to_file: true
    })
    
    batchResults.value = result.results.map((r: any, i: number) => ({
      index: i + 1,
      text: texts[i],
      success: r.success,
      duration: r.duration || '-',
      file_path: r.file_path
    }))
    
    ElMessage.success(`批量生成完成: ${result.total - result.failed}/${result.total} 成功`)
  } catch (error: any) {
    ElMessage.error(error.message || '批量生成失败')
  } finally {
    batchLoading.value = false
  }
}
</script>

<style scoped lang="scss">
.tts-settings {
  padding: 20px;
  
  .card-header {
    font-size: 18px;
    font-weight: 600;
  }
  
  .speed-value,
  .pitch-value {
    margin-left: 10px;
    color: #409eff;
  }
  
  .result-section {
    text-align: center;
    
    .audio-player {
      width: 100%;
      margin: 20px 0;
    }
  }
  
  .batch-card {
    margin-top: 20px;
  }
  
  .batch-results {
    margin-top: 20px;
  }
}
</style>
