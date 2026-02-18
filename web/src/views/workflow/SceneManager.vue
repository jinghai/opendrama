<template>
  <div class="scene-manager">
    <div class="toolbar">
      <el-input v-model="searchText" placeholder="搜索场景..." prefix-icon="Search" clearable style="width: 300px" />
      <el-button type="primary" @click="createScene" style="margin-left: auto">
        <el-icon><Plus /></el-icon> 添加场景
      </el-button>
    </div>

    <el-table :data="filteredScenes" stripe>
      <el-table-column prop="name" label="场景名称" width="150" />
      <el-table-column prop="location" label="地点" width="120" />
      <el-table-column prop="timeOfDay" label="时间段" width="100">
        <template #default="{ row }">
          <el-tag :type="getTimeType(row.timeOfDay)">{{ row.timeOfDay }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" show-overflow-tooltip />
      <el-table-column prop="image" label="背景图" width="100">
        <template #default="{ row }">
          <el-image v-if="row.image" :src="row.image" style="width: 60px; height: 40px" fit="cover" />
          <span v-else class="no-image">未生成</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="editScene(row)">编辑</el-button>
          <el-button size="small" type="primary" @click="generateImage(row)">生成背景</el-button>
          <el-button size="small" type="danger" text @click="deleteScene(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑场景' : '添加场景'" width="500px">
      <el-form :model="sceneForm" label-width="80px">
        <el-form-item label="场景名">
          <el-input v-model="sceneForm.name" />
        </el-form-item>
        <el-form-item label="地点">
          <el-select v-model="sceneForm.location">
            <el-option label="室内" value="indoor" />
            <el-option label="室外" value="outdoor" />
            <el-option label="车内" value="car" />
            <el-option label="办公室" value="office" />
            <el-option label="家中" value="home" />
          </el-select>
        </el-form-item>
        <el-form-item label="时间段">
          <el-select v-model="sceneForm.timeOfDay">
            <el-option label="早晨" value="早晨" />
            <el-option label="上午" value="上午" />
            <el-option label="中午" value="中午" />
            <el-option label="下午" value="下午" />
            <el-option label="傍晚" value="傍晚" />
            <el-option label="夜晚" value="夜晚" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="sceneForm.description" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveScene">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchText = ref('')
const dialogVisible = ref(false)
const isEdit = ref(false)

const scenes = ref([
  { id: '1', name: '公司会议室', location: 'office', timeOfDay: '上午', description: '宽敞明亮的会议室', image: '' },
  { id: '2', name: '咖啡厅', location: 'indoor', timeOfDay: '下午', description: '温馨的咖啡厅', image: '' },
  { id: '3', name: '家中客厅', location: 'home', timeOfDay: '夜晚', description: '豪华的别墅客厅', image: '' }
])

const sceneForm = ref({ id: '', name: '', location: 'indoor', timeOfDay: '上午', description: '' })

const filteredScenes = computed(() => {
  return scenes.value.filter(s => !searchText.value || s.name.includes(searchText.value))
})

const getTimeType = (time: string) => {
  const map: Record<string, string> = { '早晨': 'warning', '上午': 'success', '中午': 'danger', '下午': 'info', '傍晚': 'warning', '夜晚': 'primary' }
  return map[time] || 'info'
}

const createScene = () => { isEdit.value = false; sceneForm.value = { id: '', name: '', location: 'indoor', timeOfDay: '上午', description: '' }; dialogVisible.value = true }
const editScene = (s: any) => { isEdit.value = true; sceneForm.value = { ...s }; dialogVisible.value = true }
const saveScene = () => { ElMessage.success(isEdit.value ? '编辑成功' : '添加成功'); dialogVisible.value = false }
const generateImage = (s: any) => ElMessage.info('生成背景: ' + s.name)
const deleteScene = (s: any) => {
  ElMessageBox.confirm('确定删除这个场景吗?', '提示', { type: 'warning' })
    .then(() => { scenes.value = scenes.value.filter(x => x.id !== s.id); ElMessage.success('删除成功') })
    .catch(() => {})
}
</script>

<style scoped>
.toolbar { display: flex; align-items: center; margin-bottom: 20px; }
.no-image { color: #909399; font-size: 12px; }
</style>
