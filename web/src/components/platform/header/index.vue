<template>
  <el-table
    highlight-current-row
    ref="exportHeaderKey"
    :data="tableData"
    :height="height"
    style="width: 100%"
    :border="false"
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    @cell-mouse-enter="handleCellMouseEnter"
    @cell-mouse-leave="handleCellMouseLeave"
    @selection-change="handleSelectionChange"
    row-key="rowKeyID"
  >
    <el-table-column :reserve-selection="true" type="selection" width="55" />
    <el-table-column label="标签" width="300">
      <template #default="scope">
        <el-autocomplete
          v-model="scope.row.key"
          clearable
          :fetch-suggestions="handleQuerySearch"
          placeholder="头部标签"
          @input="handleTableDatas"
          @select="handleTableDatas"
        />
      </template>
    </el-table-column>
    <el-table-column label="内容" width="400">
      <template #default="scope">
        <el-input
          v-model="scope.row.value"
          clearable
          placeholder="头部内容"
          @input="handleTableDatas"
        />
      </template>
    </el-table-column>
    <el-table-column label="描述" width="220">
      <template #default="scope">
        <el-input
          v-model="scope.row.desc"
          clearable
          placeholder="头部信息简要描述"
          @input="handleTableDatas"
        />
      </template>
    </el-table-column>
    <el-table-column>
      <template #default="scope">
        <el-row size="small" :gutter="8">
          <el-button
            size="small"
            type="info"
            @click="handleAddRow"
            icon="add"
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
  import { getDict } from '@/utils/dictionary'

  const props = defineProps({
    save: {
      type: Boolean,
      default: false
    },
    header: {
      type: Array,
      default: () => [] // 修复原代码缺少默认值问题
    },
    exportHeader: {
      type: Array,
      default: () => []
    },
    heights: {
      type: Number,
      default: 400 // 避免无高度时表格异常
    }
  })

  const emit = defineEmits(['headerData', 'exportHeader', 'request'])

  const exportHeaderKey = ref(null) // 表格ref引用
  const tableData = ref([]) // 表格数据源
  const headerOptions = ref([]) // 自动完成下拉选项
  const multipleSelection = ref([]) // 选中的行数据
  const currentRow = ref('') // 鼠标悬浮行
  const rowKeyID = ref(0) // 行唯一标识生成器

  const height = computed(() => props.heights - 50)

  const initTableData = () => {
    tableData.value = []
    // 优先使用props传入的header数据
    if (props.header?.length) {
      props.header.forEach(item => {
        tableData.value.push({
          ...item,
          rowKeyID: rowKeyID.value++ // 补充唯一标识，避免选中异常
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

  // 6. 获取头部标签选项：优化异步逻辑+错误处理
  const getHeaderOptions = async () => {
    try {
      const res = await getDict('header')
      if (res && Array.isArray(res)) {
        headerOptions.value = res.map(item => ({ value: item.label }))
      }
    } catch (error) {
      console.error('获取头部标签选项失败：', error)
    }
  }

  const handleQuerySearch = (queryString, cb) => {
    const results = queryString
      ? headerOptions.value.filter(option =>
        option.value.toLowerCase().includes(queryString.toLowerCase())
      )
      : headerOptions.value
    cb(results)
  }

  const handleTableDatas = () => {
    // 过滤空行（key、value、desc均为空的行不提交）
    const validHeaderData = tableData.value.filter(row =>
      row.key !== ''
    )
    // 发射头部数据
    emit('headerData', validHeaderData)

    // 处理选中的导出头部
    const exportHeaderKeys = multipleSelection.value.map(item => item.key)
    emit('exportHeader', exportHeaderKeys)
  }

  const handleSelectionChange = (val) => {
    multipleSelection.value = val
    handleTableDatas()
  }

  const handleCellMouseEnter = (row) => {
    currentRow.value = row
  }
  const handleCellMouseLeave = () => {
    currentRow.value = ''
  }


  const handleAddRow = () => {
    tableData.value.push({
      key: '',
      value: '',
      desc: '',
      rowKeyID: rowKeyID.value++
    })
    handleTableDatas() // 添加后实时提交
  }

  const handleDeleteRow = (index) => {
    const deletedRow = tableData.value[index]
    // 若有ID则发射删除请求
    if (deletedRow?.ID) emit('request', deletedRow.ID)

    // 删除行并提交数据
    tableData.value.splice(index, 1)
    handleTableDatas()

    // 空表时自动添加一行
    if (tableData.value.length === 0) handleAddRow()
  }

  // 12. 头部信息格式化：修复逻辑，确保只处理有效数据
  const parseHeader = () => {
    const header = { header: {}, desc: {} }
    tableData.value.forEach(row => {
      if (row.key !== '' && row.value !== '') {
        header.header[row.key] = row.value
        header.desc[row.key] = row.desc || '' // 避免desc为undefined
      }
    })
    return header
  }

  watch(() => props.header, (newVal) => {
    if (newVal?.length) {
      initTableData()
      handleTableDatas()
    }
  }, { deep: true })

  const initSelectedRows = () => {
    if (props.exportHeader?.length && exportHeaderKey.value) {
      nextTick(() => {
        const selectedRows = tableData.value.filter(row =>
          props.exportHeader.includes(row.key)
        )
        multipleSelection.value = selectedRows
        // 触发表格选中状态更新
        selectedRows.forEach(row => {
          exportHeaderKey.value.toggleRowSelection(row, true)
        })
      })
    }
  }

  onMounted(() => {
    getHeaderOptions()
    initTableData()
    handleTableDatas()
    initSelectedRows()
  })

  // 16. 暴露方法供外部调用
  defineExpose({ parseHeader })
</script>

<style scoped></style>