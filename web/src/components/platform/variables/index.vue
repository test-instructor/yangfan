<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    stripe
    :height="height"
    :data="tableData"
    style="width: 98%"
    @cell-mouse-enter="handleCellMouseEnter"
    @cell-mouse-leave="handleCellMouseLeave"
  >
    <el-table-column label="变量名" width="220">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="Key"
          @input="handleVariablesData"
        />
      </template>
    </el-table-column>


    <el-table-column label="类型" width="120">
      <template #default="scope">
        <el-select
          v-model="scope.row.type"
          @change="handleVariablesData"
          placeholder="请选择类型"
        >
          <el-option
            v-for="item in dataTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
    </el-table-column>

    <el-table-column label="变量值" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.valueTemp"
          placeholder="Value"
          @input="handleVariablesData"></el-input>
      </template>
    </el-table-column>

    <el-table-column label="内容" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="变量简要描述"
          @input="handleVariablesData"
        />
      </template>
    </el-table-column>

    <el-table-column>
      <template #default="scope">
        <el-row :gutter="8">
          <el-button
            size="small"
            type="info"
            @click="handleAddRow(scope.$index)"
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
  import { ref, computed, watch, onMounted } from 'vue'
  import { getDict } from '@/utils/dictionary'

  const props = defineProps({
    save: {
      type: Boolean,
      default: false
    },
    variables: {
      type: Array,
      default: () => [] // 避免引用类型默认值共享问题
    },
    heights: {
      type: Number,
      default: 400 // 提供默认高度，防止无值时异常
    }
  })

  const emit = defineEmits(['variablesData', 'request'])

  const tableData = ref([]) // 表格数据源
  const dataTypeOptions = ref([]) // 类型选择下拉选项
  const currentRow = ref('') // 鼠标悬浮行

  const height = computed(() => props.heights - 50)
  watch(() => props.variables, (newVal) => {
    if (newVal?.length) {
      initTableData()
    }
  }, { deep: true })
  const initTableData = () => {
    if (props.variables?.length) {
      // 深拷贝避免父组件数据被直接修改
      tableData.value = props.variables.map(item => ({ ...item }))
    } else {
      // 默认初始化一行空数据
      tableData.value = [{
        key: '',
        value: '',
        valueTemp: '',
        type: 'String',
        desc: ''
      }]
    }
    // 初始化后触发一次数据处理
    handleVariablesData()
  }

  const getDataTypeOptions = async () => {
    try {
      const res = await getDict('fieldType')
      if (res && Array.isArray(res)) {
        // 简化数据映射，避免冗余变量
        dataTypeOptions.value = res.map(item => ({
          label: item.label,
          value: item.value
        }))
      }
    } catch (error) {
      console.error('获取变量类型选项失败：', error)
      // 可在此处添加错误提示（如ElMessage）
    }
  }

  const handleVariablesData = () => {
    const emitData = []

    tableData.value.forEach(row => {
      // 只处理关键字段非空的行（避免空数据提交）
      if (row.key || row.valueTemp || row.desc) {
        const baseData = {
          type: row.type || 'String', // 默认类型为1（String）
          desc: row.desc || '',
          key: row.key || '',
          valueTemp: row.valueTemp || '',
          value: row.valueTemp || '' // 初始值与valueTemp一致
        }

        // 优化类型转换逻辑（简化判断+兼容引用类型）
        if (row.valueTemp) {
          baseData.value = convertValueType(baseData.type, row.valueTemp)
        }

        emitData.push(baseData)
      }
    })

    emit('variablesData', emitData)
  }

  const convertValueType = (type, value) => {
    // 引用类型（含$）直接返回，不做转换
    if (value.includes('$')) return value

    switch (type) {
      case 1: // String
        return value
      case 2: // Integer
        return parseInt(value, 10) // 指定进制，避免异常
      case 3: // Float
        return parseFloat(value)
      case 4: // Boolean
        return value === 'True' ? true : value === 'False' ? false : value
      case 5: // List
      case 6: // Dict
        try {
          return JSON.parse(value)
        } catch (error) {
          console.warn('JSON解析失败，返回原始值：', error)
          return value
        }
      default:
        return value
    }
  }

  const handleCellMouseEnter = (row) => {
    currentRow.value = row
  }
  const handleCellMouseLeave = () => {
    currentRow.value = ''
  }

  const handleAddRow = (index) => {
    tableData.value.splice(index + 1, 0, {
      key: '',
      value: '',
      valueTemp: '',
      type: 'String',
      desc: ''
    })
    handleVariablesData()
  }

  const handleCopyRow = (index, row) => {
    const newRow = { ...row } // 深拷贝当前行数据
    tableData.value.splice(index + 1, 0, newRow)
    handleVariablesData()
  }

  const handleDeleteRow = (index) => {
    const deletedRow = tableData.value[index]
    // 发射删除请求（若有ID字段）
    if (deletedRow.ID) emit('request', deletedRow.ID)

    // 删除行数据
    tableData.value.splice(index, 1)
    handleVariablesData()

    // 确保表格至少保留一行
    if (tableData.value.length === 0) {
      handleAddRow(-1) // 插入到第一行
    }
  }

  // 初始化：获取选项+初始化表格
  onMounted(() => {
    getDataTypeOptions()
    initTableData()
  })

  // 监听表格数据变化（深度监听，确保子属性变更触发）
  watch(tableData, handleVariablesData, { deep: true })
</script>

<style scoped></style>