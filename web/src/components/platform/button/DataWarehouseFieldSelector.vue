<template>
  <div>
    <el-button type="primary" @click="openDialog" style="margin-right: 15px;">数据仓库字段</el-button>

    <el-dialog
      title="数据仓库字段"
      v-model="showDialog"
      width="800px"
      destroy-on-close
    >
      <div style="display: flex; align-items: center; gap: 12px; margin-bottom: 12px;">
        <div style="width: 64px; color: #606266;">数据类型</div>
        <el-select
          v-if="!typeFromParent"
          v-model="selectedType"
          placeholder="请选择数据类型"
          filterable
          clearable
          style="flex: 1;"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.type"
            :label="item.type"
            :value="item.type"
          />
        </el-select>
        <el-select v-else :model-value="typeFromParent" disabled style="flex: 1;">
          <el-option :label="typeFromParent" :value="typeFromParent" />
        </el-select>
      </div>
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
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
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
const selectedType = ref('')

const typeFromParent = computed(() => (props.currentType || '').trim())
const effectiveType = computed(() => typeFromParent.value || selectedType.value)

const openDialog = () => {
  showDialog.value = true
  if (!typeFromParent.value) {
    selectedType.value = ''
  }
  if (typeOptions.value.length === 0) {
    getDataCategory()
  }
}

const fieldList = computed(() => {
  if (!effectiveType.value) return []
  const found = typeOptions.value.find(item => item.type === effectiveType.value)
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

watch(
  () => props.currentType,
  () => {
    if (typeFromParent.value) {
      selectedType.value = ''
    }
  }
)
</script>
