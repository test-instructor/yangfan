<template>
  <div>
    <el-button type="primary" @click="openDialog" style="margin-right: 15px;">数据仓库字段</el-button>

    <el-dialog
      title="数据仓库字段"
      v-model="showDialog"
      width="800px"
      destroy-on-close
    >
      <el-table
        :data="fieldList"
        border
        size="small"
        style="width: 100%"
        height="500px"
      >
        <el-table-column label="字段" prop="field" align="left" />
        <el-table-column label="类型" prop="type" align="left" width="100" />
        <el-table-column label="值" prop="value" align="left">
           <template #default="scope">
             {{ scope.row.value }}
           </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="120">
          <template #default="scope">
            <el-button
              icon="copy-document"
              type="text"
              @click="copyField(scope.row.field)"
              size="small"
            >
              复制
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div v-if="!currentType" class="text-center text-gray-500 py-4" style="text-align: center; color: #909399; padding-top: 20px; padding-bottom: 20px;">
        请先在数据仓库配置中选择数据类型
      </div>
      <div v-else-if="fieldList.length === 0" class="text-center text-gray-500 py-4" style="text-align: center; color: #909399; padding-top: 20px; padding-bottom: 20px;">
        该类型暂无字段数据
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getDataCategoryTypeList } from '@/api/datawarehouse/dataCategoryManagement'

const props = defineProps({
  currentType: {
    type: String,
    default: ''
  }
})

const showDialog = ref(false)
const typeOptions = ref([])

const openDialog = () => {
  showDialog.value = true
}

const fieldList = computed(() => {
  if (!props.currentType) return []
  const found = typeOptions.value.find(item => item.type === props.currentType)
  if (!found || !found.value) return []
  
  return Object.keys(found.value).map(key => ({
    field: key,
    value: found.value[key],
    type: typeof found.value[key]
  }))
})

const getDataCategory = async () => {
  const res = await getDataCategoryTypeList()
  if (res.code === 0) {
    typeOptions.value = res.data || []
  }
}

const copyField = (field) => {
  const text = `$${field}`
  navigator.clipboard.writeText(text)
    .then(() => ElMessage({ type: 'success', message: '字段已复制' }))
    .catch(() => ElMessage({ type: 'error', message: '复制失败' }))
}

onMounted(() => {
  getDataCategory()
})
</script>
