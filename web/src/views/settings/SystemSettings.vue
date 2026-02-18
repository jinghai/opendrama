<template>
  <div class="settings-panel">
    <el-tabs v-model="activeTab">
      <el-tab-pane label="基本设置" name="basic">
        <el-form label-width="120px">
          <el-form-item label="应用名称">
            <el-input v-model="settings.appName" />
          </el-form-item>
          <el-form-item label="语言">
            <el-select v-model="settings.language">
              <el-option label="中文" value="zh" />
              <el-option label="English" value="en" />
            </el-select>
          </el-form-item>
          <el-form-item label="主题">
            <el-radio-group v-model="settings.theme">
              <el-radio label="light">浅色</el-radio>
              <el-radio label="dark">深色</el-radio>
              <el-radio label="auto">跟随系统</el-radio>
            </el-radio-group>
          </el-form-item>
          <el-form-item label="自动保存">
            <el-switch v-model="settings.autoSave" />
          </el-form-item>
          <el-form-item label="保存间隔">
            <el-input-number v-model="settings.autoSaveInterval" :min="1" :max="60" /> 分钟
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="AI服务" name="ai">
        <el-form label-width="120px">
          <el-form-item label="默认文本模型">
            <el-select v-model="settings.defaultTextModel">
              <el-option label="GPT-4" value="gpt-4" />
              <el-option label="GPT-3.5" value="gpt-3.5" />
              <el-option label="Claude" value="claude" />
              <el-option label="通义千问" value="qwen" />
            </el-select>
          </el-form-item>
          <el-form-item label="默认图像模型">
            <el-select v-model="settings.defaultImageModel">
              <el-option label="DALL-E 3" value="dalle-3" />
              <el-option label="Midjourney" value="midjourney" />
              <el-option label="Stable Diffusion" value="sd" />
            </el-select>
          </el-form-item>
          <el-form-item label="AI生图数量">
            <el-input-number v-model="settings.defaultImageCount" :min="1" :max="10" />
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="导出设置" name="export">
        <el-form label-width="120px">
          <el-form-item label="默认格式">
            <el-select v-model="settings.defaultFormat">
              <el-option label="MP4" value="mp4" />
              <el-option label="MOV" value="mov" />
              <el-option label="AVI" value="avi" />
            </el-select>
          </el-form-item>
          <el-form-item label="默认质量">
            <el-select v-model="settings.defaultQuality">
              <el-option label="原画 (1080P)" value="1080p" />
              <el-option label="高清 (720P)" value="720p" />
              <el-option label="标清 (480P)" value="480p" />
            </el-select>
          </el-form-item>
          <el-form-item label="自动导出">
            <el-switch v-model="settings.autoExport" />
          </el-form-item>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="存储" name="storage">
        <el-form label-width="120px">
          <el-form-item label="存储位置">
            <el-input v-model="settings.storagePath" />
          </el-form-item>
          <el-form-item label="最大存储">
            <el-input-number v-model="settings.maxStorage" :min="1" :max="1000" /> GB
          </el-form-item>
          <el-form-item label="自动清理">
            <el-switch v-model="settings.autoCleanup" />
          </el-form-item>
          <el-form-item label="清理周期">
            <el-select v-model="settings.cleanupPeriod" :disabled="!settings.autoCleanup">
              <el-option label="每天" value="daily" />
              <el-option label="每周" value="weekly" />
              <el-option label="每月" value="monthly" />
            </el-select>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>

    <div class="form-footer">
      <el-button @click="resetSettings">重置</el-button>
      <el-button type="primary" @click="saveSettings">保存设置</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const activeTab = ref('basic')
const settings = reactive({
  appName: 'OpenDrama',
  language: 'zh',
  theme: 'light',
  autoSave: true,
  autoSaveInterval: 5,
  defaultTextModel: 'gpt-4',
  defaultImageModel: 'dalle-3',
  defaultImageCount: 4,
  defaultFormat: 'mp4',
  defaultQuality: '1080p',
  autoExport: false,
  storagePath: './data/storage',
  maxStorage: 100,
  autoCleanup: false,
  cleanupPeriod: 'weekly'
})

const saveSettings = () => ElMessage.success('设置已保存')
const resetSettings = () => ElMessage.info('重置设置')
</script>

<style scoped>
.settings-panel { padding: 20px; }
.form-footer { margin-top: 30px; text-align: right; }
</style>
