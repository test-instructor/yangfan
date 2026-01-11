<template>
  <el-table
    highlight-current-row
    ref="exportParameterKey"
    stripe
    :height="height"
    :data="tableData"
    style="width: 98%"
    @cell-mouse-enter="handleCellMouseEnter"
    @cell-mouse-leave="handleCellMouseLeave"
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    @selection-change="handleSelectionChange"
    row-key="rowKeyID"
    :id="props.idName"
  >
    <el-table-column :reserve-selection="true" type="selection" width="40" />
    <el-table-column label="变量名" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="接收提取值后的变量名"
          @input="handleExtractData"
        />
      </template>
    </el-table-column>
    <el-table-column label="提取表达式" width="320">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.value"
          placeholder="提取表达式"
          @input="handleExtractData"
        />
      </template>
    </el-table-column>
    <el-table-column label="描述" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="提取值简要描述"
          @input="handleExtractData"
        />
      </template>
    </el-table-column>
    <el-table-column>
      <template #default="scope">
        <el-row :gutter="8">
          <el-button
            size="small"
            type="info"
            @click="handleAddRow"
            icon="add"
          />
          <el-button
            size="small"
            type="info"
            @click="handleCopyRow(scope.$index, scope.row)"
            icon="copy"
          />
          <el-button
            size="small"
            type="danger"
            @click="handleDeleteRow(scope.$index)"
            icon="delete"
          />
        </el-row>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>
  import { ref, computed, watch, onMounted, nextTick } from 'vue'

  // 定义Props，添加类型校验与默认值
  const props = defineProps({
    exportParameter: {
      type: Array,
      default: () => []
    },
    save: {
      type: Boolean,
      default: false
    },
    extract: {
      type: Array,
      default: () => []
    },
    heights: {
      type: Number,
      default: 400
    },
    idName: {
      type: String,
      default: 'extractTable'
    }
  })
  // 定义Emit事件
  const emit = defineEmits(['extractData', 'exportParameter'])

  // 响应式变量
  const exportParameterKey = ref(null)
  const tableData = ref([])
  const multipleSelection = ref([])
  const currentRow = ref('')
  const rowKeyID = ref(0) // 用于生成唯一行标识

  // 计算表格高度
  const height = computed(() => props.heights - 70)
  watch(
    () => props.extract,
    (newVal) => {
      if (newVal) {
        initTableData()
      }
    }
  )
  // 初始化表格数据
  const initTableData = () => {
    rowKeyID.value = 0 // 重置ID生成器
    tableData.value = []

    if (props.extract?.length) {
      // 处理传入的提取数据，补充唯一标识
      props.extract.forEach(item => {
        tableData.value.push({
          ...item,
          rowKeyID: rowKeyID.value++
        })
      })
    } else {
      // 默认初始化一行数据
      tableData.value.push({
        key: '',
        value: '',
        desc: '',
        rowKeyID: rowKeyID.value++
      })
    }
  }

  // 处理选中行回显
  const initSelectedRows = () => {
    if (props.exportParameter?.length && exportParameterKey.value) {
      nextTick(() => {
        // 筛选需要选中的行
        const rowsToSelect = tableData.value.filter(row =>
          props.exportParameter.includes(row.key)
        )

        multipleSelection.value = rowsToSelect
        // 执行选中操作
        rowsToSelect.forEach(row => {
          exportParameterKey.value.toggleRowSelection(row, true)
        })
      })
    }
  }

  // 提取数据并发射事件
  const handleExtractData = () => {
    // 仅发射 key、value 都不为空的有效数据
    const extract = ref([])

    tableData.value.forEach(item => {
      let extractData = {}
      const { key, value, desc } = item
      if (key != null && key !== '' && value != null && value !== '') {
        extractData = {
          key,
          value,
          desc
        }
        extract.value.push(extractData)
      }
    })

    // 这里向父组件传递已经过滤后的数据，避免空 key 被保存
    emit('extractData', extract.value)

    // 处理选中的导出参数
    const exportParams = multipleSelection.value.map(item => item.key)
    emit('exportParameter', exportParams)
  }

  // 选中变化事件
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
    handleExtractData()
  }

  // 鼠标悬浮事件
  const handleCellMouseEnter = (row) => {
    currentRow.value = row
  }
  const handleCellMouseLeave = () => {
    currentRow.value = ''
  }

  // 行操作方法
  const handleAddRow = () => {
    tableData.value.push({
      key: '',
      value: '',
      desc: '',
      rowKeyID: rowKeyID.value++
    })
    handleExtractData()
  }

  const handleCopyRow = (index, row) => {
    // 深拷贝行数据，避免引用关联
    const newRow = {
      ...row,
      rowKeyID: rowKeyID.value++ // 生成新的唯一标识
    }
    tableData.value.splice(index + 1, 0, newRow)
    handleExtractData()
  }

  const handleDeleteRow = (index) => {
    tableData.value.splice(index, 1)
    handleExtractData()

    // 确保表格至少保留一行
    if (tableData.value.length === 0) {
      handleAddRow()
    }
  }

  // 提取数据格式化
  const parseExtract = () => {
    const extract = { extract: [], desc: {} }

    tableData.value.forEach(item => {
      const { key, value, desc } = item
      if (key && value) {
        extract.extract.push({ [key]: value })
        extract.desc[key] = desc || '' // 避免desc为undefined
      }
    })

    return extract
  }

  // 监听表格数据变化，实时触发数据发射
  watch(tableData, handleExtractData, { deep: true })

  // 初始化
  onMounted(() => {
    initTableData()
    handleExtractData()
    initSelectedRows()
  })

  defineExpose({ parseExtract })
</script>

<style scoped></style>