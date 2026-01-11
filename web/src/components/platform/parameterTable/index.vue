<template>
  <div class="parameter-table">
    <div class="actions" style="margin-bottom: 10px;">
      <el-button type="primary" icon="plus" @click="openAddColumnDialog">添加列</el-button>
      <el-button type="primary" icon="plus" @click="addRow">添加行</el-button>
    </div>
    
    <el-table :data="tableData" style="width: 100%" border>
      <el-table-column
        v-for="(col, index) in columns"
        :key="col.key"
        :prop="col.key"
        :min-width="150"
      >
        <template #header>
           <div style="display: flex; align-items: center; justify-content: space-between;">
             <span style="cursor: pointer;" @click="openEditColumnDialog(index)">{{ col.key }}</span>
             <div>
               <el-button type="primary" icon="edit" circle size="small" @click="openEditColumnDialog(index)" />
               <el-button type="danger" icon="delete" circle size="small" @click="deleteColumn(index)" />
             </div>
           </div>
        </template>
        <template #default="scope">
          <el-input
            v-if="col.type === 'string'"
            v-model="scope.row[col.key]"
            @input="emitData"
            placeholder="请输入"
          />
          <el-switch
            v-else-if="col.type === 'boolean'"
            v-model="scope.row[col.key]"
            @change="emitData"
            active-text="true"
            inactive-text="false"
          />
          <el-input-number
            v-else
            v-model="scope.row[col.key]"
            @change="emitData"
            style="width: 100%"
            placeholder="请输入"
          />
        </template>
      </el-table-column>
      
      <el-table-column label="操作" width="80" fixed="right">
        <template #default="scope">
          <el-button type="danger" icon="delete" circle size="small" @click="deleteRow(scope.$index)" />
        </template>
      </el-table-column>
    </el-table>

    <!-- Add/Edit Column Dialog -->
    <el-dialog v-model="dialogVisible" :title="isEditMode ? '编辑列' : '添加列'" width="400px" append-to-body>
      <el-form :model="currentColumn" label-width="80px" ref="columnFormRef">
        <el-form-item label="Key" prop="key">
           <el-input v-model="currentColumn.key" placeholder="请输入Key (英文)" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
           <el-select v-model="currentColumn.type" placeholder="请选择类型" style="width: 100%" :disabled="isEditMode">
             <el-option label="字符串" value="string" />
             <el-option label="数字" value="number" />
             <el-option label="布尔" value="boolean" />
           </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="saveColumn">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  jsons: {
    type: Object,
    default: () => ({})
  },
  parametersTemp: {
    type: Object,
    default: () => ({})
  }
})

const emit = defineEmits(['jsonData', 'tempData'])

const tableData = ref([])
const columns = ref([])
const dialogVisible = ref(false)
const isEditMode = ref(false)
const editingIndex = ref(-1)
const currentColumn = ref({
  key: '',
  type: 'string'
})

// Initialize data from props
const initData = () => {
  // Priority to parametersTemp
  const temp = props.parametersTemp && props.parametersTemp.parameters_temp_data 
    ? props.parametersTemp.parameters_temp_data 
    : props.parametersTemp

  if (temp && temp.columns && temp.columns.length > 0) {
    columns.value = temp.columns
    tableData.value = temp.data || []
    return
  }

  if (!props.jsons || Object.keys(props.jsons).length === 0) {
    tableData.value = []
    columns.value = []
    return
  }

  const keys = Object.keys(props.jsons)
  const maxRows = Math.max(...keys.map(k => (Array.isArray(props.jsons[k]) ? props.jsons[k].length : 0)))
  
  // Infer columns
  columns.value = keys.map(key => {
    const values = props.jsons[key]
    let type = 'string'
    // Simple type inference based on the first non-null value
    if (Array.isArray(values) && values.length > 0) {
        const firstVal = values.find(v => v !== null && v !== undefined)
        if (typeof firstVal === 'number') {
            type = 'number'
        } else if (typeof firstVal === 'boolean') {
            type = 'boolean'
        }
    }
    return {
        key,
        type
    }
  })

  // Construct rows
  const rows = []
  for (let i = 0; i < maxRows; i++) {
    const row = {}
    keys.forEach(key => {
        const val = props.jsons[key][i]
        row[key] = val
    })
    rows.push(row)
  }
  tableData.value = rows
}

onMounted(() => {
  initData()
})

watch(() => props.parametersTemp, (newVal) => {
    const temp = newVal && newVal.parameters_temp_data ? newVal.parameters_temp_data : newVal
    if (temp && temp.columns && temp.columns.length > 0) {
         // Check if we need to update (simple check to avoid loop)
         if (JSON.stringify(temp.columns) !== JSON.stringify(columns.value) || 
             JSON.stringify(temp.data) !== JSON.stringify(tableData.value)) {
             columns.value = temp.columns
             tableData.value = temp.data || []
         }
    }
}, { deep: true })

watch(() => props.jsons, (newVal) => {
    // Only use jsons if parametersTemp is empty/invalid and tableData is empty
    // This prevents jsons overwrite if we are operating on temp data
    const temp = props.parametersTemp && props.parametersTemp.parameters_temp_data 
        ? props.parametersTemp.parameters_temp_data 
        : props.parametersTemp
        
    if ((!temp || !temp.columns) && tableData.value.length === 0 && Object.keys(newVal || {}).length > 0) {
        initData()
    }
}, { deep: true })


const openAddColumnDialog = () => {
  isEditMode.value = false
  currentColumn.value = { key: '', type: 'string' }
  dialogVisible.value = true
}

const openEditColumnDialog = (index) => {
  isEditMode.value = true
  editingIndex.value = index
  // Copy current values
  currentColumn.value = { ...columns.value[index] }
  dialogVisible.value = true
}

const saveColumn = () => {
  if (!currentColumn.value.key) {
    ElMessage.warning('请输入Key')
    return
  }
  // Check for english key
  if (!/^[a-zA-Z0-9_]+$/.test(currentColumn.value.key)) {
      ElMessage.warning('Key只能包含英文、数字和下划线')
      return
  }

  if (isEditMode.value) {
     // Edit existing column
     const oldKey = columns.value[editingIndex.value].key
     const newKey = currentColumn.value.key
     
     // If key changed, check for duplicate
     if (oldKey !== newKey && columns.value.find(c => c.key === newKey)) {
         ElMessage.warning('Key已存在')
         return
     }

     // Update column definition
     // Note: type is not updated as per requirement (disabled in UI)
     columns.value[editingIndex.value].key = newKey
     
     // Update data rows if key changed
     if (oldKey !== newKey) {
         tableData.value.forEach(row => {
             row[newKey] = row[oldKey]
             delete row[oldKey]
         })
     }

  } else {
      // Add new column
      if (columns.value.find(c => c.key === currentColumn.value.key)) {
        ElMessage.warning('Key已存在')
        return
      }

      columns.value.push({ ...currentColumn.value })
      // Initialize the new column in existing rows
      tableData.value.forEach(row => {
        if (currentColumn.value.type === 'number') {
            row[currentColumn.value.key] = 0
        } else if (currentColumn.value.type === 'boolean') {
            row[currentColumn.value.key] = false
        } else {
            row[currentColumn.value.key] = ''
        }
      })
  }
  
  dialogVisible.value = false
  emitData()
}

const addRow = () => {
  const row = {}
  columns.value.forEach(col => {
    if (col.type === 'number') {
        row[col.key] = 0
    } else if (col.type === 'boolean') {
        row[col.key] = false
    } else {
        row[col.key] = ''
    }
  })
  tableData.value.push(row)
  emitData()
}

const deleteColumn = (index) => {
  const col = columns.value[index]
  columns.value.splice(index, 1)
  // Remove data from rows
  tableData.value.forEach(row => {
    delete row[col.key]
  })
  emitData()
}

const deleteRow = (index) => {
  tableData.value.splice(index, 1)
  emitData()
}

const emitData = () => {
  const result = {}
  
  if (columns.value.length > 0) {
    const keys = columns.value.map(col => col.key)
    const combinedKey = keys.join('-')
    
    const values = tableData.value.map(row => {
        return columns.value.map(col => {
            let val = row[col.key]
            if (col.type === 'number') {
                return Number(val)
            }
            if (col.type === 'boolean') {
                return Boolean(val)
            }
            return String(val || '')
        })
    })
    result[combinedKey] = values
  }
  
  // Emit backend format
  emit('jsonData', {
    data: result,
    isValid: true,
    error: null
  })

  // Emit frontend temp format
  emit('tempData', {
    parameters_temp_data: {
        columns: columns.value,
        data: tableData.value
    }
  })
}

</script>

<style scoped>
.parameter-table {
  padding: 10px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
}
</style>