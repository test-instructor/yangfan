<template>
  <a-card :title="t('menu.androidui.reportDetail')" :bordered="false">
    <a-spin v-if="loading" :tip="t('common.loading')" />
    <a-result v-else-if="errorText" status="error" :title="t('autoReport.detailLoadError')" :subtitle="errorText" />
    <a-descriptions v-else :column="1" bordered>
      <a-descriptions-item :label="t('androidui.reportId')">{{ id }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.name')">{{ report?.name }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.caseType')">{{ report?.case_type }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.runMode')">{{ report?.run_mode }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.nodeName')">{{ report?.node_name }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.status')">{{ report?.status }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.progress')">{{ formatProgress(report) }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.describe')">{{ report?.describe }}</a-descriptions-item>
      <a-descriptions-item :label="t('autoReport.raw')">
        <a-textarea :model-value="rawText" readonly :auto-size="{ minRows: 8, maxRows: 18 }" />
      </a-descriptions-item>
    </a-descriptions>
  </a-card>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { findAutoReport } from '@/services/appBridge'

const route = useRoute()
const { t } = useI18n()

const id = computed(() => String(route.params?.id || ''))

const loading = ref(false)
const report = ref(null)
const errorText = ref('')

const rawText = computed(() => {
  try {
    return JSON.stringify(report.value || {}, null, 2)
  } catch {
    return ''
  }
})

const formatProgress = (r) => {
  const progress = r?.progress
  const totalApis = Number(progress?.total_apis || 0)
  const executed = Number(progress?.executed_apis || 0)
  if (!totalApis) return '-'
  return `${executed}/${totalApis}`
}

onMounted(async () => {
  const parsed = Number(id.value)
  if (!parsed) return
  loading.value = true
  try {
    report.value = await findAutoReport(parsed)
  } catch (e) {
    errorText.value = e?.message || ''
  } finally {
    loading.value = false
  }
})
</script>
