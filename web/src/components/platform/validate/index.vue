<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    stripe
    :height="height"
    :data="tableData"
    style="width: 98%"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
    :id="props.idName"
  >
    <el-table-column fixed label="字段" width="200">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.check"
          placeholder="字段"
          @input="validateDatas"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="类型" width="180">
      <template #default="scope">
        <el-select
          filterable
          v-model="scope.row.assert"
          placeholder="请选择类型"
          @change="validateDatas"
        >
          <el-option
            v-for="item in validateOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
    </el-table-column>

    <el-table-column label="期望类型" width="120">
      <template #default="scope">
        <el-select v-model="scope.row.type" @change="validateDatas">
          <el-option
            v-for="item in dataTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
    </el-table-column>

    <el-table-column label="期望数据" width="250">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.expect"
          placeholder="期望数据"
          @input="validateDatas"
        />
      </template>
    </el-table-column>

    <el-table-column label="描述" width="200">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.msg"
          placeholder="描述"
          @input="validateDatas"
        />
      </template>
    </el-table-column>

    <el-table-column>
      <template #default="scope">
        <el-row>
          <el-button
            size="small"
            type="info"
            @click="handleAdd(scope.$index, scope.row)"
            icon="add"
          />
          <el-button
            size="small"
            type="info"
            @click="handleCopy(scope.$index, scope.row)"
            icon="copy"
          />
          <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index)"
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

  // 定义props
  const props = defineProps({
    save: Boolean,
    validate: {
      type: Array,
      default: () => []
    },
    heights: {
      type: Number,
      default: 400
    },
    idName: {
      type: String,
      default: 'validateTable'
    }
  })

  // 定义emits
  const emit = defineEmits(['jsonData', 'requestSkipData'])

  // 响应式变量
  const tableData = ref([])
  const validateOptions = ref([])
  const dataTypeOptions = ref([])
  const currentRow = ref('')

  // 计算属性
  const height = computed(() => props.heights - 70)

  // 常量定义
  const ASSERT_INT_KEY = [
    'equals', 'less_than', 'less_or_equals', 'greater_than',
    'greater_or_equals', 'not_equal', 'length_equals',
    'length_greater_than', 'length_greater_or_equals',
    'length_less_than', 'length_less_or_equals', 'contains'
  ]

  // 初始化表格数据
  const initTableData = () => {
    if (props.validate && props.validate.length > 0) {
      tableData.value = [...props.validate]
    } else {
      tableData.value = [{
        expect: '',
        check: '',
        assert: 'equals',
        msg: '',
        expectTemp: '',
        type: 'String'
      }]
    }
  }
  watch(
    () => props.validate,
    (newVal) => {
      initTableData()
    },
    { deep: true }
  );

  // 获取验证类型选项
  const getValidateOptions = async () => {
    try {
      const res = await getDict('assert')
      if (res && Array.isArray(res)) {
        validateOptions.value = res.map(item => ({
          label: item.label,
          value: item.value
        }))
      }
    } catch (error) {
      console.error('获取验证类型选项失败:', error)
    }
  }

  // 获取数据类型选项
  const getDataTypeOptions = async () => {
    try {
      const res = await getDict('fieldType')
      if (res && Array.isArray(res)) {
        dataTypeOptions.value = res.map(item => ({
          label: item.label,
          value: item.value
        }))
      }
    } catch (error) {
      console.error('获取数据类型选项失败:', error)
    }
  }

  // 验证数据并发射事件
  const validateDatas = () => {
    const emitdata = []

    tableData.value.forEach(row => {
      if (row.check) { // 只处理有值的行
        const emitdataDict = {
          expect: row.expect,
          check: row.check,
          msg: row.msg,
          assert: row.assert,
          type: row.type
        }

        // 类型转换处理
        const value = row.expectTemp
        emitdata.push(emitdataDict)
      }
    })

    emit('jsonData', emitdata)
    emit('requestSkipData', emitdata)
  }

  // 鼠标进入单元格
  const cellMouseEnter = (row) => {
    currentRow.value = row
  }

  // 鼠标离开单元格
  const cellMouseLeave = () => {
    currentRow.value = ''
  }

  // 添加新行
  const handleAdd = () => {
    tableData.value.push({
      expect: '',
      check: '',
      assert: 'equals',
      msg: '',
      expectTemp: '',
      type: 'String'
    })
    validateDatas()
  }

  // 复制行
  const handleCopy = (index, row) => {
    tableData.value.splice(index + 1, 0, {
      ...row,
      // 深拷贝防止引用问题
      expect: JSON.parse(JSON.stringify(row.expect)),
      expectTemp: JSON.parse(JSON.stringify(row.expectTemp))
    })
    validateDatas()
  }

  // 删除行
  const handleDelete = (index) => {
    tableData.value.splice(index, 1)
    validateDatas()

    // 确保至少保留一行
    if (tableData.value.length === 0) {
      handleAdd()
    }
  }

  // 解析验证规则
  const parseValidate = () => {
    const validate = { validate: [] }

    tableData.value.forEach(content => {
      if (content.check) {
        try {
          const expect = parseType(content.msg, content.expect)
          validate.validate.push({
            [content.assert]: [content.check, expect]
          })
        } catch (error) {
          console.error('解析验证规则失败:', error)
        }
      }
    })

    return validate
  }

  // 类型解析辅助函数
  const parseType = (msg, expect) => {
    // 根据实际需求实现类型解析逻辑
    return expect
  }

  // 初始化
  onMounted(() => {
    initTableData()
    getValidateOptions()
    getDataTypeOptions()
  })

  // 监听表格数据变化
  watch(tableData, validateDatas, { deep: true })

  // 暴露方法供外部使用
  defineExpose({
    parseValidate
  })
</script>

<style scoped></style>