<template>
  <el-select
    v-model="selectedId"
    filterable
    clearable
    :placeholder="selectedId === 0 ? '请选择配置' : '请选择配置'"
    @change="handleChange"
    :style="{ width: width }"
  >
    <el-option
      v-for="item in filteredOptions"
      :key="item.ID"
      :label="item.name"
      :value="item.ID"
    />
  </el-select>
</template>

<script setup>
  import { ref, watch, computed } from 'vue'
  import { getRunConfigList } from '@/api/platform/runconfig.js'

  // 接收外部传入的参数（新增width属性）
  const props = defineProps({
    modelValue: {
      type: Number,
      default: null
    },
    width: {
      type: String,
      default: '160px'  // 默认宽度160px
    }
  })

  // 向父组件emit选择结果
  const emit = defineEmits(['update:modelValue', 'change'])
  const selectedId = ref(props.modelValue === 0 ? null : props.modelValue)

  const options = ref([])

  const fetchOptions = async () => {
    try {
      const response = await getRunConfigList({
        page: 1,
        pageSize: 9999
      })
      if (response.code === 0) {
        options.value = response.data.list
      }
    } catch (error) {
      console.error('获取配置列表失败:', error)
    }
  }

  fetchOptions()

  watch(
    () => props.modelValue,
    (newVal) => {
      selectedId.value = newVal === 0 ? null : newVal
    }
  )

  const handleChange = (id) => {
    const emitValue = id === null ? 0 : id
    emit('update:modelValue', emitValue)
    if (id === null) {
      emit('change', null)
      return
    }
    const selectedItem = options.value.find(item => item.ID === id)
    emit('change', selectedItem ? { ID: selectedItem.ID, name: selectedItem.name } : null)
  }

  const filteredOptions = computed(() => {
    return options.value
  })
</script>

<style scoped>
  /* 可根据需要添加样式 */
</style>