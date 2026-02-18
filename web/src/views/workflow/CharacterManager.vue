<template>
  <div class="character-manager">
    <!-- 搜索和筛选 -->
    <div class="toolbar">
      <el-input v-model="searchText" placeholder="搜索角色..." prefix-icon="Search" clearable style="width: 300px" />
      <el-select v-model="filterRole" placeholder="筛选角色" clearable style="width: 150px; margin-left: 10px">
        <el-option label="主角" value="protagonist" />
        <el-option label="配角" value="supporting" />
        <el-option label="反派" value="antagonist" />
        <el-option label="客串" value="cameo" />
      </el-select>
      <el-button type="primary" @click="createCharacter" style="margin-left: auto">
        <el-icon><Plus /></el-icon> 创建角色
      </el-button>
    </div>

    <!-- 角色列表 -->
    <div class="character-grid">
      <el-card v-for="char in filteredCharacters" :key="char.id" shadow="hover" class="character-card">
        <div class="character-avatar">
          <img v-if="char.image" :src="char.image" :alt="char.name" />
          <div v-else class="avatar-placeholder">{{ char.name?.charAt(0) }}</div>
        </div>
        <div class="character-info">
          <h4>{{ char.name }}</h4>
          <el-tag size="small" :type="getRoleType(char.role)">{{ getRoleText(char.role) }}</el-tag>
          <p class="character-desc">{{ char.description || '暂无描述' }}</p>
          <div class="character-actions">
            <el-button size="small" @click="editCharacter(char)">编辑</el-button>
            <el-button size="small" type="primary" @click="generateImage(char)">生成图片</el-button>
          </div>
        </div>
      </el-card>
    </div>

    <!-- 创建/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑角色' : '创建角色'" width="600px">
      <el-form :model="characterForm" label-width="80px">
        <el-form-item label="角色名">
          <el-input v-model="characterForm.name" placeholder="请输入角色名" />
        </el-form-item>
        <el-form-item label="角色类型">
          <el-select v-model="characterForm.role">
            <el-option label="主角" value="protagonist" />
            <el-option label="配角" value="supporting" />
            <el-option label="反派" value="antagonist" />
            <el-option label="客串" value="cameo" />
          </el-select>
        </el-form-item>
        <el-form-item label="角色描述">
          <el-input v-model="characterForm.description" type="textarea" :rows="3" placeholder="请输入角色描述" />
        </el-form-item>
        <el-form-item label="性格特点">
          <el-input v-model="characterForm.personality" type="textarea" :rows="2" placeholder="请输入性格特点" />
        </el-form-item>
        <el-form-item label="外观描述">
          <el-input v-model="characterForm.appearance" type="textarea" :rows="2" placeholder="请输入外观描述" />
        </el-form-item>
        <el-form-item label="角色图片">
          <el-upload :show-file-list="false" :on-change="handleImageUpload">
            <el-button>上传图片</el-button>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveCharacter">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const searchText = ref('')
const filterRole = ref('')
const dialogVisible = ref(false)
const isEdit = ref(false)

const characters = ref([
  { id: '1', name: '林悦', role: 'protagonist', description: '年轻漂亮的女主', image: '', personality: '坚强独立', appearance: '长发披肩，气质优雅' },
  { id: '2', name: '陆霆', role: 'antagonist', description: '霸道总裁', image: '', personality: '冷酷无情', appearance: '西装革履' },
  { id: '3', name: '苏晴', role: 'supporting', description: '女主闺蜜', image: '', personality: '活泼开朗', appearance: '短发可爱' }
])

const characterForm = ref({
  id: '', name: '', role: 'protagonist', description: '', personality: '', appearance: '', image: ''
})

const filteredCharacters = computed(() => {
  return characters.value.filter(char => {
    const matchSearch = !searchText.value || char.name.includes(searchText.value) || char.description?.includes(searchText.value)
    const matchRole = !filterRole.value || char.role === filterRole.value
    return matchSearch && matchRole
  })
})

const getRoleType = (role: string) => {
  const map: Record<string, string> = { protagonist: 'danger', supporting: 'success', antagonist: 'warning', cameo: 'info' }
  return map[role] || 'info'
}

const getRoleText = (role: string) => {
  const map: Record<string, string> = { protagonist: '主角', supporting: '配角', antagonist: '反派', cameo: '客串' }
  return map[role] || role
}

const createCharacter = () => {
  isEdit.value = false
  characterForm.value = { id: '', name: '', role: 'protagonist', description: '', personality: '', appearance: '', image: '' }
  dialogVisible.value = true
}

const editCharacter = (char: any) => {
  isEdit.value = true
  characterForm.value = { ...char }
  dialogVisible.value = true
}

const saveCharacter = () => {
  if (!characterForm.value.name) {
    ElMessage.warning('请输入角色名')
    return
  }
  ElMessage.success(isEdit.value ? '编辑成功' : '创建成功')
  dialogVisible.value = false
}

const generateImage = (char: any) => {
  ElMessage.info('生成图片功能: ' + char.name)
}

const handleImageUpload = (file: any) => {
  ElMessage.info('上传图片功能开发中')
}
</script>

<style scoped lang="scss">
.character-manager {
  .toolbar {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
  }

  .character-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 20px;

    .character-card {
      .character-avatar {
        width: 80px;
        height: 80px;
        border-radius: 50%;
        overflow: hidden;
        margin: 0 auto 15px;

        img { width: 100%; height: 100%; object-fit: cover; }
        .avatar-placeholder {
          width: 100%; height: 100%;
          background: linear-gradient(135deg, #667eea, #764ba2);
          color: white;
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: 32px;
          font-weight: 600;
        }
      }

      .character-info {
        text-align: center;
        h4 { margin: 0 0 8px; }
        .character-desc {
          margin: 10px 0;
          font-size: 13px;
          color: #909399;
          display: -webkit-box;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
          overflow: hidden;
        }
        .character-actions {
          display: flex;
          gap: 10px;
          justify-content: center;
          margin-top: 15px;
        }
      }
    }
  }
}
</style>
