<template>
  <el-select
    v-model="selectedId"
    filterable
    clearable
    placeholder="请选择自动化步骤"
    @change="handleChange"
    :style="{ width: width }"
  >
    <el-option
      v-for="item in options"
      :key="item.ID"
      :label="item.name"
      :value="item.ID"
    />
  </el-select>
</template>

<script setup>
  import { ref, watch, onMounted } from 'vue'
  import { getAutoStepList } from '@/api/automation/autostep.js'

  const props = defineProps({
    modelValue: {
      type: Number,
      default: null
    },
    width: {
      type: String,
      default: '100%'
    }
  })

  const emit = defineEmits(['update:modelValue', 'change'])
  const selectedId = ref(props.modelValue || null)
  const options = ref([])

  const fetchOptions = async () => {
    try {
      const response = await getAutoStepList({
        page: 1,
        pageSize: 9999
      })
      if (response.code === 0) {
        options.value = response.data.list
      }
    } catch (error) {
      console.error('获取自动化步骤列表失败:', error)
    }
  }

  onMounted(() => {
    fetchOptions()
  })

  watch(
    () => props.modelValue,
    (newVal) => {
      selectedId.value = newVal || null
    }
  )

  const handleChange = (id) => {
    const emitValue = id || 0
    // 注意：后端接口如果是 uint 指针，可能需要 null，但这里通常传递 ID 值，0 在后端可能被处理为空
    // 但根据请求结构体 *uint，如果传 0 可能会有问题，前端最好传 null 或正确 ID
    // 此处保持与 RunConfig 类似逻辑，如果 id 为 null/undefined 则传 null
    emit('update:modelValue', id || null) 
    emit('change', id || null)
  }
</script>

<style scoped>
</style>