<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    stripe
    :height="height"
    :data="paramsData"
    style="width: 100%"
    @cell-mouse-enter="handleCellMouseEnter"
    @cell-mouse-leave="handleCellMouseLeave"
  >
    <el-table-column label="请求Key" width="250">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="Key"
          @input="handleParamsData"
        />
      </template>
    </el-table-column>
    <el-table-column label="请求Value" width="340">
      <template #default="scope">
        <el-input
          v-model="scope.row.value"
          placeholder="Value"
          @input="handleParamsData"
        />
      </template>
    </el-table-column>
    <el-table-column label="描述" width="340">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="参数简要描述"
          @input="handleParamsData"
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
  import { ref, computed, onMounted, onUnmounted, watch } from 'vue'

  const props = defineProps({
    save: {
      type: Boolean,
      default: false
    },
    params: {
      type: Array,
      default: () => [] // 引用类型默认值用函数返回，避免共享引用
    },
    heights: {
      type: Number,
      default: 400 // 默认高度防止无值时表格异常
    }
  })

  const emit = defineEmits(['paramsData', 'request'])

  const paramsData = ref([]) // 表格数据源
  const currentRow = ref('') // 鼠标悬浮行
  const contentStyleObj = ref({ height: '', width: '98%' }) // 样式对象

  const height = computed(() => props.heights - 70)

  const initParamsData = () => {
    if (props.params?.length) {
      // 深拷贝避免修改父组件原始数据
      paramsData.value = props.params.map(item => ({ ...item }))
    } else {
      // 默认初始化一行空数据
      paramsData.value = [{ key: '', value: '', type: '', desc: '' }]
    }
    // 初始化后触发一次数据提交
    handleParamsData()
  }
  watch(
    () => props.params,
    (newVal) => {
      if (newVal) {
        initParamsData()
      }
    }
  )

  const handleParamsData = () => {
    // 过滤空行（key、value、desc均为空的行不提交）
    const emitData = paramsData.value.filter(row =>
      row.key !== '' || row.value !== '' || row.desc !== ''
    )
    // 发射处理后的数据给父组件
    emit('paramsData', emitData)
  }


  const handleCellMouseEnter = (row) => {
    currentRow.value = row
  }
  const handleCellMouseLeave = () => {
    currentRow.value = ''
  }

  const handleAddRow = () => {
    paramsData.value.push({ key: '', value: '', type: '', desc: '' })
    handleParamsData() // 添加后触发数据提交
  }

  const handleCopyRow = (index, row) => {
    const newRow = { ...row } // 深拷贝当前行
    paramsData.value.splice(index + 1, 0, newRow)
    handleParamsData()
  }

  const handleDeleteRow = (index) => {
    const deletedRow = paramsData.value[index]
    // 若有ID则发射删除请求
    if (deletedRow?.ID) emit('request', deletedRow.ID)

    // 删除行并提交数据
    paramsData.value.splice(index, 1)
    handleParamsData()

    // 空表时自动添加一行
    if (paramsData.value.length === 0) handleAddRow()
  }

  const handleResize = () => {
    contentStyleObj.value.height = `${height.value}px`
  }

  const parseParams = () => {
    const params = { params: {}, desc: {} }
    paramsData.value.forEach(row => {
      if (row.key !== '') {
        params.params[row.key] = row.value
        params.desc[row.key] = row.desc || '' // 避免desc为undefined
      }
    })
    return params
  }


  onMounted(() => {
    initParamsData()
    handleResize() // 初始设置样式
    window.addEventListener('resize', handleResize)
  })

  onUnmounted(() => {
    window.removeEventListener('resize', handleResize)
  })

  defineExpose({ parseParams })
</script>

<style scoped></style>