<template>
  <div class="dashboard-container">
    <el-container>
      <el-header class="header">
        <div class="header-content">
          <h2>{{ $t('dashboard.title') }}</h2>
          <LanguageSwitcher />
        </div>
      </el-header>
      
      <el-main>
        <div class="welcome-section">
          <h1>🎬 OpenDrama</h1>
          <p>开源短剧 · AI创作 · 人人皆是剧作家</p>
        </div>
        
        <!-- 快速统计组件 -->
        <QuickStats />
                <el-icon :size="40" color="#67c23a"><Picture /></el-icon>
                <h3>0</h3>
                <p>{{ $t('dashboard.stats.images') }}</p>
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="6">
            <el-card shadow="hover">
              <div class="stat-item">
                <el-icon :size="40" color="#e6a23c"><VideoPlay /></el-icon>
                <h3>0</h3>
                <p>{{ $t('dashboard.stats.videos') }}</p>
              </div>
            </el-card>
          </el-col>
          
          <el-col :span="6">
            <el-card shadow="hover">
              <div class="stat-item">
                <el-icon :size="40" color="#f56c6c"><Clock /></el-icon>
                <h3>0</h3>
                <p>{{ $t('dashboard.stats.tasks') }}</p>
              </div>
            </el-card>
          </el-col>
        </el-row>
        
        <!-- 快捷操作组件 -->
        <QuickActions @open-tts="openTTSDialog" />
            <el-col :span="8">
              <el-card shadow="hover" class="action-card" @click="goToDramas">
                <el-icon :size="50" color="#409eff"><Plus /></el-icon>
                <h3>{{ $t('dashboard.actions.newProject') }}</h3>
                <p>{{ $t('dashboard.actions.newProjectDesc') }}</p>
              </el-card>
            </el-col>
            
            <el-col :span="8">
              <el-card shadow="hover" class="action-card" @click="goToDramas">
                <el-icon :size="50" color="#67c23a"><FolderOpened /></el-icon>
                <h3>{{ $t('dashboard.actions.myProjects') }}</h3>
                <p>{{ $t('dashboard.actions.myProjectsDesc') }}</p>
              </el-card>
            </el-col>
            
          </el-row>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { Document, Picture, VideoPlay, Clock, Plus, FolderOpened, Setting } from '@element-plus/icons-vue'
import LanguageSwitcher from '@/components/LanguageSwitcher.vue'
import QuickStats from './components/QuickStats.vue'
import QuickActions from './components/QuickActions.vue'

const router = useRouter()

const goToDramas = () => {
  router.push('/dramas')
}

const goToSettings = () => {
  router.push('/settings/ai-config')
}

const openTTSDialog = () => {
  // Emit event to parent to open TTS dialog
  const layout = document.querySelector('.app-layout')
  if (layout) {
    // Try to find and click the TTS button
    const ttsBtn = layout.querySelector('button:has(.el-icon-microphone)')
    if (ttsBtn) {
      (ttsBtn as HTMLElement).click()
    }
  }
}
</script>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background: #f5f7fa;
}

.header {
  background: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
}

.header-content h2 {
  margin: 0;
  color: #409eff;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.welcome-section {
  text-align: center;
  padding: 40px 0;
}

.welcome-section h1 {
  font-size: 36px;
  margin-bottom: 10px;
  color: #333;
}

.welcome-section p {
  font-size: 18px;
  color: #666;
}

.stats-row {
  margin-bottom: 40px;
}

.stat-item {
  text-align: center;
  padding: 20px 0;
}

.stat-item h3 {
  font-size: 32px;
  margin: 10px 0;
  color: #333;
}

.stat-item p {
  color: #666;
  margin: 0;
}

.quick-actions h2 {
  margin-bottom: 20px;
  color: #333;
}

.action-card {
  cursor: pointer;
  text-align: center;
  padding: 30px 20px;
  transition: all 0.3s;
}

.action-card:hover {
  transform: translateY(-5px);
}

.action-card h3 {
  margin: 15px 0 10px;
  color: #333;
}

.action-card p {
  color: #666;
  margin: 0;
}
</style>
