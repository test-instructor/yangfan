<template>
  <el-card class="box-card" shadow="hover">
    <template #header>
      <div class="flex justify-between items-center flex-wrap gap-2">
        <span class="text-lg font-bold">运行概况</span>
        <el-tag :type="statusType" size="large" effect="dark">{{ statusText }}</el-tag>
      </div>
    </template>
    <el-descriptions :column="columnCount" border>
      <el-descriptions-item label="运行名称">{{ detail.name }}</el-descriptions-item>
      <el-descriptions-item label="报告ID">{{ detail.ID }}</el-descriptions-item>
      <el-descriptions-item label="任务名称">{{ detail.case_type }}</el-descriptions-item>
      <el-descriptions-item label="任务类型">{{ detail.case_type }}</el-descriptions-item>
      <el-descriptions-item label="运行模式">{{ detail.run_mode }}</el-descriptions-item>
      <el-descriptions-item label="触发人">-</el-descriptions-item>
      <el-descriptions-item label="开始时间">{{ formatDate(detail.time?.start_at) }}</el-descriptions-item>
      <el-descriptions-item label="总耗时">{{ formatDuration(detail.time?.duration) }}</el-descriptions-item>
      <el-descriptions-item label="执行节点">{{ detail.hostname }}</el-descriptions-item>
      <el-descriptions-item label="运行环境">{{ detail.env_name }}</el-descriptions-item>
    </el-descriptions>
  </el-card>
</template>

<script setup>
import { computed } from 'vue'
import { formatDate } from '@/utils/format'
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core'

const props = defineProps({
  detail: {
    type: Object,
    required: true,
    default: () => ({})
  }
})

const breakpoints = useBreakpoints(breakpointsTailwind)

const columnCount = computed(() => {
  if (breakpoints.xl.value) return 4
  if (breakpoints.lg.value) return 3
  if (breakpoints.md.value) return 2
  return 1
})

const statusText = computed(() => {
  const s = props.detail.status
  if (s === 1) return '运行中'
  if (s === 2) return '失败'
  if (s === 3) return '成功'
  return '未知'
})

const statusType = computed(() => {
  const s = props.detail.status
  if (s === 1) return 'primary'
  if (s === 2) return 'danger'
  if (s === 3) return 'success'
  return 'info'
})

const formatDuration = (seconds) => {
  if (seconds === null || seconds === undefined || Number.isNaN(Number(seconds))) return '-'
  return Number(seconds).toFixed(2) + 's'
}
</script>
