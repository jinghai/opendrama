<template>
  <div class="prop-manager">
    <div class="toolbar">
      <el-input v-model="searchText" placeholder="搜索道具..." prefix-icon="Search" clearable style="width: 300px" />
      <el-button type="primary" @click="createProp" style="margin-left: auto">
        <el-icon><Plus /></el-icon> 添加道具
      </el-button>
    </div>

    <el-table :data="filteredProps" stripe>
      <el-table-column prop="name" label="道具名称" width="150" />
      <el-table-column prop="category" label="类别" width="120">
        <template #default="{ row }">
          <el-tag>{{ getCategoryText(row.category) }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" show-overflow-tooltip />
      <el-table-column prop="image" label="图片" width="100">
        <template #default="{ row }">
          <el-image v-if="row.image" :src="row.image" style="width: 60px; height: 60px" fit="cover" />
          <span v-else>-</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template #default="{ row }">
          <el-button size="small" @click="editProp(row)">编辑</el-button>
          <el-button size="small" type="primary" @click="generateImage(row)">生成图片</el-button>
          <el-button size="small" type="danger" text @click="deleteProp(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑道具' : '添加道具'" width="500px">
      <el-form :model="propForm" label-width="80px">
        <el-form-item label="道具名">
          <el-input v-model="propForm.name" />
        </el-form-item>
        <el-form-item label="类别">
          <el-select v-model="propForm.category">
            <el-option label="服装" value="clothing" />
            <el-option label="饰品" value="accessory" />
            <el-option label="道具" value="prop" />
            <el-option label="场景" value="scene" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="propForm.description" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveProp">保存</el-button>
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

const props = ref([
  { id: '1', name: '钻戒', category: 'accessory', description: '奢华钻戒', image: '' },
  { id: '2', name: '礼服', category: 'clothing', description: '白色晚礼服', image: '' }
])

const propForm = ref({ id: '', name: '', category: 'prop', description: '' })

const filteredProps = computed(() => {
  return props.value.filter(p => !searchText.value || p.name.includes(searchText.value))
})

const getCategoryText = (cat: string) => {
  const map: Record<string, string> = { clothing: '服装', accessory: '饰品', prop: '道具', scene: '场景' }
  return map[cat] || cat
}

const createProp = () => { isEdit.value = false; propForm.value = { id: '', name: '', category: 'prop', description: '' }; dialogVisible.value = true }
const editProp = (p: any) => { isEdit.value = true; propForm.value = { ...p }; dialogVisible.value = true }
const saveProp = () => { ElMessage.success(isEdit.value ? '编辑成功' : '添加成功'); dialogVisible.value = false }
const generateImage = (p: any) => ElMessage.info('生成图片: ' + p.name)
const deleteProp = (p: any) => {
  ElMessageBox.confirm('确定删除这个道具吗?', '提示', { type: 'warning' })
    .then(() => { props.value = props.value.filter(x => x.id !== p.id); ElMessage.success('删除成功') })
    .catch(() => {})
}
</script>

<style scoped>
.toolbar { display: flex; align-items: center; margin-bottom: 20px; }
</style>
