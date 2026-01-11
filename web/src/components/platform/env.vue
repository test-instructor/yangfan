<template>
  <el-select
    v-model="internalValue"
    filterable
    clearable
    :placeholder="placeholder"
    :loading="isLoading"
    @change="handleSelectChange"
    :style="{ width: width }"
  >
    <el-option
      v-for="env in envList"
      :key="env.ID"
      :label="env.name"
      :value="env.ID"
    />
  </el-select>
</template>

<script setup>
  import { ref, watch, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import { getEnvList } from '@/api/platform/env'


  // 接收父组件传入的参数（新增width属性）
  const props = defineProps({
    modelValue: {
      type: Number,
      default: 0 // 默认值设为0，对应"请选择环境"
    },
    width: {
      type: String,
      default: '160px' // 默认宽度120px，支持px、%等单位
    }
  })

  // 向父组件触发事件
  const emit = defineEmits(['update:modelValue', 'change'])

  // 内部维护的选中值
  const internalValue = ref(props.modelValue === 0 ? null : props.modelValue)
  // 环境列表数据
  const envList = ref([])
  // 加载状态
  const isLoading = ref(false)

  // 占位符文案
  const placeholder = computed(() => {
    return props.modelValue === 0 ? '请选择环境' : '请选择环境'
  })

  // 获取环境列表数据
  const getEnvData = async () => {
    isLoading.value = true
    try {
      const res = await getEnvList({ page: 1, pageSize: 1000 })
      if (res.code === 0) {
        envList.value = res.data.list || []
      } else {
        ElMessage.error(`获取失败: ${res.msg || '未知错误'}`)
      }
    } catch (error) {
      console.error('接口请求异常:', error)
      ElMessage.error('网络异常，获取数据失败')
    } finally {
      isLoading.value = false
    }
  }

  // 初始化加载数据
  getEnvData()

  // 监听父组件传入的ID变化
  watch(
    () => props.modelValue,
    (newVal) => {
      internalValue.value = newVal === 0 ? null : newVal
    }
  )

  // 处理选择变化
  const handleSelectChange = (id) => {
    const emitId = id === null ? 0 : id
    emit('update:modelValue', emitId)

    if (id === null) {
      emit('change', null)
      return
    }
    const selectedEnv = envList.value.find(item => item.ID === id)
    emit('change', selectedEnv ? { ID: selectedEnv.ID, name: selectedEnv.name } : null)
  }
</script>

<style scoped>
</style>