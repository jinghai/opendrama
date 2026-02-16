<template>
  <div class="newapi-settings">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>🔌 NewAPI统一接口设置</span>
          <el-button type="primary" @click="saveConfig" :loading="saving">
            💾 保存配置
          </el-button>
        </div>
      </template>
      
      <el-form :model="config" label-width="150px">
        <!-- 基础配置 -->
        <el-divider>基础配置</el-divider>
        
        <el-form-item label="NewAPI地址">
          <el-input v-model="config.base_url" placeholder="https://api.newapi.com" />
          <el-button size="small" @click="testConnection" :loading="testing" style="margin-left: 10px">
            🧪 测试连接
          </el-button>
        </el-form-item>
        
        <el-form-item label="API密钥">
          <el-input v-model="config.api_key" type="password" placeholder="请输入API密钥" show-password />
        </el-form-item>
        
        <!-- 负载均衡策略 -->
        <el-divider>负载均衡</el-divider>
        
        <el-form-item label="负载策略">
          <el-select v-model="config.load_balancer.strategy">
            <el-option label="轮询 (Round Robin)" value="round-robin" />
            <el-option label="加权轮询 (Weighted)" value="weighted" />
            <el-option label="最低成本 (Least Cost)" value="least-cost" />
            <el-option label="最低延迟 (Least Latency)" value="least-latency" />
          </el-select>
        </el-form-item>
        
        <!-- 服务商配置 -->
        <el-divider>服务商配置</el-divider>
        
        <div v-for="(provider, index) in config.load_balancer.providers" :key="index" class="provider-item">
          <el-card shadow="never" class="provider-card">
            <template #header>
              <div class="provider-header">
                <el-switch v-model="provider.enabled" />
                <span class="provider-name">{{ provider.name }}</span>
                <el-button type="danger" size="small" text @click="removeProvider(index)">
                  删除
                </el-button>
              </div>
            </template>
            
            <el-form-item label="名称">
              <el-input v-model="provider.name" />
            </el-form-item>
            
            <el-form-item label="API地址">
              <el-input v-model="provider.base_url" placeholder="https://api.xxx.com" />
            </el-form-item>
            
            <el-form-item label="API密钥">
              <el-input v-model="provider.api_key" type="password" show-password />
            </el-form-item>
            
            <el-form-item label="模型列表">
              <el-input v-model="provider.models_text" type="textarea" :rows="2" placeholder="gpt-4, gpt-3.5-turbo" @blur="updateModels(provider)" />
            </el-form-item>
            
            <el-row :gutter="20">
              <el-col :span="12">
                <el-form-item label="权重">
                  <el-input-number v-model="provider.weight" :min="1" :max="100" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="优先级">
                  <el-input-number v-model="provider.priority" :min="1" :max="100" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-card>
        </div>
        
        <el-button type="primary" plain @click="addProvider">
          ➕ 添加服务商
        </el-button>
      </el-form>
    </el-card>
    
    <!-- 统计信息 -->
    <el-card class="stats-card">
      <template #header>
        <span>📊 服务商统计</span>
        <el-button size="small" @click="loadStats">
          🔄 刷新
        </el-button>
      </template>
      
      <el-table :data="stats" stripe>
        <el-table-column prop="name" label="服务商" />
        <el-table-column prop="request_count" label="请求数" />
        <el-table-column prop="error_count" label="错误数" />
        <el-table-column prop="success_rate" label="成功率">
          <template #default="{ row }">
            {{ (row.success_rate * 100).toFixed(1) }}%
          </template>
        </el-table-column>
        <el-table-column prop="avg_latency" label="平均延迟">
          <template #default="{ row }">
            {{ (row.avg_latency / 1000000000).toFixed(2) }}s
          </template>
        </el-table-column>
        <el-table-column prop="last_used" label="最后使用">
          <template #default="{ row }">
            {{ row.last_used ? new Date(row.last_used).toLocaleString() : '-' }}
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getNewAPIConfig, updateNewAPIConfig, getNewAPIStats } from '@/api/newapi'

const saving = ref(false)
const stats = ref<any[]>([])

const config = reactive({
  base_url: 'https://api.newapi.com',
  api_key: '',
  load_balancer: {
    strategy: 'least-cost',
    providers: [
      {
        name: 'openai',
        base_url: 'https://api.openai.com',
        api_key: '',
        models: ['gpt-4', 'gpt-4-turbo', 'gpt-3.5-turbo'],
        models_text: 'gpt-4, gpt-4-turbo, gpt-3.5-turbo',
        priority: 10,
        weight: 10,
        enabled: true
      },
      {
        name: 'qwen',
        base_url: 'https://dashscope.aliyuncs.com/api/v1',
        api_key: '',
        models: ['qwen-turbo', 'qwen-plus', 'qwen-max'],
        models_text: 'qwen-turbo, qwen-plus, qwen-max',
        priority: 9,
        weight: 9,
        enabled: true
      },
      {
        name: 'doubao',
        base_url: 'https://ark.cn-beijing.volces.com/api/v3',
        api_key: '',
        models: ['doubao-pro-4k', 'doubao-pro-32k'],
        models_text: 'doubao-pro-4k, doubao-pro-32k',
        priority: 10,
        weight: 10,
        enabled: true
      }
    ]
  }
})

const updateModels = (provider: any) => {
  provider.models = provider.models_text.split(',').map((m: string) => m.trim()).filter((m: string) => m)
}

const addProvider = () => {
  config.load_balancer.providers.push({
    name: 'new-provider',
    base_url: 'https://api.xxx.com',
    api_key: '',
    models: [],
    models_text: '',
    priority: 5,
    weight: 5,
    enabled: true
  })
}

const removeProvider = (index: number) => {
  config.load_balancer.providers.splice(index, 1)
}

const saveConfig = async () => {
  saving.value = true
  try {
    // 更新模型列表
    config.load_balancer.providers.forEach(updateModels)
    
    await updateNewAPIConfig(config)
    
    ElMessage.success('配置保存成功')
  } catch (error: any) {
    ElMessage.error(error.message || '保存失败')
  } finally {
    saving.value = false
  }
}

const loadStats = async () => {
  try {
    const response = await getNewAPIStats()
    stats.value = Object.values(response.stats || {})
  } catch (error) {
    console.error('加载统计失败:', error)
  }
}

// 测试连接
const testing = ref(false)
const testResult = ref('')

const testConnection = async () => {
  testing.value = true
  testResult.value = ''
  try {
    // 模拟测试 - 实际应该调用后端API
    await new Promise(resolve => setTimeout(resolve, 1000))
    testResult.value = '测试功能需要配置API Key后使用'
    ElMessage.info('请先保存配置后再测试')
  } catch (error: any) {
    testResult.value = error.message || '测试失败'
  } finally {
    testing.value = false
  }
}

onMounted(() => {
  loadStats()
})
</script>

<style scoped lang="scss">
.newapi-settings {
  padding: 20px;
  
  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 18px;
    font-weight: 600;
  }
  
  .provider-item {
    margin-bottom: 20px;
    
    .provider-card {
      margin-top: 10px;
    }
    
    .provider-header {
      display: flex;
      align-items: center;
      gap: 10px;
      
      .provider-name {
        font-weight: 600;
      }
    }
  }
  
  .stats-card {
    margin-top: 20px;
  }
}
</style>
