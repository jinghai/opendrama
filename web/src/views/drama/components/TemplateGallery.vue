<template>
  <div class="template-gallery">
    <div class="gallery-header">
      <h3>📋 剧本模板</h3>
      <el-button type="primary" @click="createTemplate">
        <el-icon><Plus /></el-icon> 创建模板
      </el-button>
    </div>
    
    <el-row :gutter="20">
      <el-col :span="8" v-for="template in templates" :key="template.id">
        <el-card shadow="hover" class="template-card" @click="useTemplate(template)">
          <div class="template-cover" :style="{ background: template.color }">
            <span class="template-category">{{ template.category }}</span>
          </div>
          <div class="template-info">
            <h4>{{ template.name }}</h4>
            <p>{{ template.description }}</p>
            <div class="template-meta">
              <el-tag size="small">{{ template.episodes }}集</el-tag>
              <el-tag size="small" type="info">{{ template.duration }}分钟/集</el-tag>
            </div>
          </div>
          <div class="template-actions">
            <el-button size="small" text @click.stop="previewTemplate(template)">预览</el-button>
            <el-button size="small" text type="primary" @click.stop="useTemplate(template)">使用</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
    
    <!-- 模板预览对话框 -->
    <el-dialog v-model="previewVisible" title="模板预览" width="800px">
      <div v-if="currentTemplate" class="preview-content">
        <h3>{{ currentTemplate.name }}</h3>
        <p class="desc">{{ currentTemplate.description }}</p>
        <el-divider />
        <h4>剧集结构</h4>
        <el-timeline>
          <el-timeline-item 
            v-for="(episode, index) in currentTemplate.episodes" 
            :key="index"
            :timestamp="episode.title"
            placement="top"
          >
            <p>{{ episode.description }}</p>
          </el-timeline-item>
        </el-timeline>
      </div>
      <template #footer>
        <el-button @click="previewVisible = false">关闭</el-button>
        <el-button type="primary" @click="useTemplate(currentTemplate)">使用此模板</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const router = useRouter()
const previewVisible = ref(false)
const currentTemplate = ref<any>(null)

const templates = ref([
  {
    id: '1',
    name: '甜宠短剧',
    category: '爱情',
    description: '甜蜜宠溺的都市爱情故事，适合年轻女性观众',
    episodes: 30,
    duration: 3,
    color: 'linear-gradient(135deg, #ff9a9e 0%, #fecfef 99%)',
    episodes: [
      { title: '第1集 意外相遇', description: '女主不小心撞到霸道总裁' },
      { title: '第2集 冤家路窄', description: '两人在同一公司上班' },
      { title: '第3集 渐生情愫', description: '日常相处中感情升温' }
    ]
  },
  {
    id: '2',
    name: '逆袭人生',
    category: '励志',
    description: '草根逆袭成CEO的励志故事',
    episodes: 50,
    duration: 5,
    color: 'linear-gradient(135deg, #a18cd1 0%, #fbc2eb 100%)',
    episodes: [
      { title: '第1集 人生低谷', description: '主角遭遇重大挫折' },
      { title: '第2集 获得机遇', description: '意外获得神秘帮助' }
    ]
  },
  {
    id: '3',
    name: '豪门恩怨',
    category: '家庭',
    description: '豪门家族的爱恨情仇',
    episodes: 40,
    duration: 4,
    color: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
    episodes: [
      { title: '第1集 意外身世', description: '揭示主角真实身份' }
    ]
  },
  {
    id: '4',
    name: '穿越时空',
    category: '奇幻',
    description: '穿越古代与现代的奇幻爱情',
    episodes: 25,
    duration: 3,
    color: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
    episodes: [
      { title: '第1集 意外穿越', description: '现代女意外回到古代' }
    ]
  },
  {
    id: '5',
    name: '职场奋斗',
    category: '职场',
    description: '职场新人的成长之路',
    episodes: 30,
    duration: 4,
    color: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
    episodes: [
      { title: '第1集 初入职场', description: '新人报道' }
    ]
  },
  {
    id: '6',
    name: '悬疑推理',
    category: '悬疑',
    description: '烧脑的悬疑推理故事',
    episodes: 20,
    duration: 5,
    color: 'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
    episodes: [
      { title: '第1集 神秘案件', description: '离奇案件发生' }
    ]
  }
])

const useTemplate = (template: any) => {
  ElMessage.success('使用模板: ' + template.name)
  router.push({
    path: '/dramas/create',
    query: { template: template.id }
  })
}

const previewTemplate = (template: any) => {
  currentTemplate.value = template
  previewVisible.value = true
}

const createTemplate = () => {
  ElMessage.info('创建模板功能开发中')
}
</script>

<style scoped lang="scss">
.template-gallery {
  .gallery-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    
    h3 {
      margin: 0;
    }
  }
  
  .template-card {
    cursor: pointer;
    margin-bottom: 20px;
    transition: transform 0.3s;
    
    &:hover {
      transform: translateY(-5px);
    }
    
    .template-cover {
      height: 100px;
      border-radius: 8px 8px 0 0;
      display: flex;
      align-items: flex-start;
      justify-content: flex-end;
      padding: 10px;
      
      .template-category {
        background: rgba(255, 255, 255, 0.9);
        padding: 4px 10px;
        border-radius: 12px;
        font-size: 12px;
        font-weight: 500;
      }
    }
    
    .template-info {
      padding: 15px;
      
      h4 {
        margin: 0 0 8px;
        font-size: 16px;
      }
      
      p {
        margin: 0 0 12px;
        font-size: 13px;
        color: #909399;
      }
      
      .template-meta {
        display: flex;
        gap: 8px;
      }
    }
    
    .template-actions {
      padding: 10px 15px;
      border-top: 1px solid #eee;
      display: flex;
      justify-content: flex-end;
    }
  }
  
  .preview-content {
    h3 {
      margin: 0 0 10px;
    }
    
    .desc {
      color: #909399;
      margin: 0 0 20px;
    }
  }
}
</style>
