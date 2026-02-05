<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.reports')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('autoReport.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('autoReport.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('autoReport.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #progress="{ record }">
          {{ formatProgress(record) }}
        </template>
        <template #actions="{ record }">
          <a-button type="text" @click="openDetail(record)">{{ t('autoReport.detail') }}</a-button>
        </template>
      </a-table>

      <div style="margin-top: 12px; display: flex; justify-content: flex-end">
        <a-pagination
          :current="page"
          :page-size="pageSize"
          :total="total"
          show-total
          show-jumper
          show-page-size
          :page-size-options="[10, 30, 50, 100]"
          @change="handlePageChange"
          @page-size-change="handlePageSizeChange"
        />
      </div>
    </a-card>
  </a-space>
</template>

<script setup>
import { Message } from '@arco-design/web-vue'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { getAutoReportList } from '@/services/appBridge'

const { t } = useI18n()
const router = useRouter()

const loading = ref(false)
const rows = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const searchForm = reactive({
  createdAtStart: '',
  createdAtEnd: ''
})

const columns = computed(() => [
  { title: t('autoReport.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('autoReport.name'), dataIndex: 'name' },
  { title: t('autoReport.caseType'), dataIndex: 'case_type', width: 120 },
  { title: t('autoReport.runMode'), dataIndex: 'run_mode', width: 120 },
  { title: t('autoReport.nodeName'), dataIndex: 'node_name' },
  { title: t('autoReport.status'), dataIndex: 'status', width: 100 },
  { title: t('autoReport.progress'), slotName: 'progress', width: 120 },
  { title: t('autoReport.actions'), slotName: 'actions', width: 120 }
])

const buildCreatedAtRange = () => {
  const start = String(searchForm.createdAtStart || '').trim()
  const end = String(searchForm.createdAtEnd || '').trim()
  if (!start || !end) return null
  return [start, end]
}

const loadList = async () => {
  loading.value = true
  try {
    const createdAtRange = buildCreatedAtRange()
    const query = {
      page: page.value,
      pageSize: pageSize.value
    }
    if (createdAtRange) query['createdAtRange[]'] = createdAtRange
    const res = await getAutoReportList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoReport.loadError'))
  } finally {
    loading.value = false
  }
}

const handleSearch = async () => {
  page.value = 1
  await loadList()
}

const handleReset = async () => {
  searchForm.createdAtStart = ''
  searchForm.createdAtEnd = ''
  page.value = 1
  await loadList()
}

const handlePageChange = async (p) => {
  page.value = p
  await loadList()
}

const handlePageSizeChange = async (s) => {
  pageSize.value = s
  page.value = 1
  await loadList()
}

const formatProgress = (row) => {
  const progress = row?.progress
  const totalApis = Number(progress?.total_apis || 0)
  const executed = Number(progress?.executed_apis || 0)
  if (!totalApis) return '-'
  return `${executed}/${totalApis}`
}

const openDetail = async (row) => {
  const id = row?.ID
  if (!id) return
  await router.push({ path: `/androidui/autoReport/${id}` })
}

onMounted(async () => {
  await loadList()
})
</script>
