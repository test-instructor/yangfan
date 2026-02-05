<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.runConfig')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('runConfig.name')">
          <a-input v-model="searchForm.name" allow-clear :placeholder="t('runConfig.namePlaceholder')" />
        </a-form-item>
        <a-form-item :label="t('runConfig.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('runConfig.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('runConfig.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('runConfig.create') }}</a-button>
      </a-space>

      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openEdit(record)">{{ t('runConfig.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('runConfig.delete') }}</a-button>
          </a-space>
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

    <a-drawer v-model:visible="drawerVisible" :title="drawerTitle" width="720px" :mask-closable="false">
      <a-form ref="formRef" :model="form" layout="vertical">
        <a-form-item field="name" :label="t('runConfig.name')" :rules="[{ required: true, message: t('runConfig.nameRequired') }]">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item field="timeout" :label="t('runConfig.timeout')">
          <a-input-number v-model="form.timeout" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="allow_redirects" :label="t('runConfig.allowRedirects')">
          <a-switch v-model="form.allow_redirects" />
        </a-form-item>
        <a-form-item field="verify" :label="t('runConfig.verify')">
          <a-switch v-model="form.verify" />
        </a-form-item>
        <a-form-item field="preparatorySteps" :label="t('runConfig.preparatorySteps')">
          <a-input-number v-model="form.preparatorySteps" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="setup_case_id" :label="t('runConfig.setupCaseId')">
          <a-input-number v-model="form.setup_case_id" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="retry" :label="t('runConfig.retry')">
          <a-input-number v-model="form.retry" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>

        <a-form-item field="androidDeviceOptionsId" :label="t('runConfig.androidDeviceOptions')">
          <a-select v-model="form.androidDeviceOptionsId" allow-clear :loading="deviceLoading" :placeholder="t('runConfig.androidDeviceOptionsPlaceholder')">
            <a-option v-for="d in devices" :key="d.ID" :value="d.ID">
              {{ d.name || d.serial }}
            </a-option>
          </a-select>
        </a-form-item>

        <a-form-item field="header_temp" :label="t('runConfig.headerTemp')">
          <a-textarea v-model="form.header_temp" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="variable_temp" :label="t('runConfig.variableTemp')">
          <a-textarea v-model="form.variable_temp" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="parameters_temp" :label="t('runConfig.parametersTemp')">
          <a-textarea v-model="form.parameters_temp" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="data_warehouse_temp" :label="t('runConfig.dataWarehouseTemp')">
          <a-textarea v-model="form.data_warehouse_temp" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
      </a-form>
      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('runConfig.save') }}</a-button>
          <a-button @click="drawerVisible = false">{{ t('runConfig.cancel') }}</a-button>
        </a-space>
      </template>
    </a-drawer>
  </a-space>
</template>

<script setup>
import { Message, Modal } from '@arco-design/web-vue'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { createRunConfig, deleteRunConfig, getAndroidDeviceOptionsList, getRunConfigList, updateRunConfig } from '@/services/appBridge'

const { t } = useI18n()
const route = useRoute()
const currentType = computed(() => String(route.query?.type || 'android'))

const loading = ref(false)
const rows = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const searchForm = reactive({
  name: '',
  createdAtStart: '',
  createdAtEnd: ''
})

const devices = ref([])
const deviceLoading = ref(false)

const drawerVisible = ref(false)
const saving = ref(false)
const drawerMode = ref('create')
const formRef = ref()

const form = reactive({
  ID: undefined,
  name: '',
  timeout: 30,
  allow_redirects: true,
  verify: false,
  preparatorySteps: 0,
  setup_case_id: 0,
  retry: 0,
  androidDeviceOptionsId: undefined,
  header_temp: '{}',
  variable_temp: '{}',
  parameters_temp: '{}',
  data_warehouse_temp: '{}'
})

const columns = computed(() => [
  { title: t('runConfig.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('runConfig.name'), dataIndex: 'name' },
  { title: t('runConfig.baseUrl'), dataIndex: 'base_url' },
  { title: t('runConfig.preparatorySteps'), dataIndex: 'preparatorySteps' },
  { title: t('runConfig.setupCaseId'), dataIndex: 'setup_case_id' },
  { title: t('runConfig.actions'), slotName: 'actions', width: 180 }
])

const drawerTitle = computed(() => (drawerMode.value === 'edit' ? t('runConfig.edit') : t('runConfig.create')))

const parseMaybeJSON = (text, labelKey) => {
  const s = String(text || '').trim()
  if (!s) return null
  try {
    return JSON.parse(s)
  } catch (e) {
    throw new Error(t('runConfig.jsonInvalid', { field: t(labelKey) }))
  }
}

const loadDevices = async () => {
  deviceLoading.value = true
  try {
    const res = await getAndroidDeviceOptionsList({ page: 1, pageSize: 1000 })
    const list = res?.list || res?.List || []
    devices.value = Array.isArray(list) ? list : []
  } finally {
    deviceLoading.value = false
  }
}

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
      pageSize: pageSize.value,
      type: currentType.value,
      name: searchForm.name || undefined
    }
    if (createdAtRange) {
      query['createdAtRange[]'] = createdAtRange
    }
    const res = await getRunConfigList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('runConfig.loadError'))
  } finally {
    loading.value = false
  }
}

const handleSearch = async () => {
  page.value = 1
  await loadList()
}

const handleReset = async () => {
  searchForm.name = ''
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

const resetForm = () => {
  form.ID = undefined
  form.name = ''
  form.timeout = 30
  form.allow_redirects = true
  form.verify = false
  form.preparatorySteps = 0
  form.setup_case_id = 0
  form.retry = 0
  form.androidDeviceOptionsId = undefined
  form.header_temp = '{}'
  form.variable_temp = '{}'
  form.parameters_temp = '{}'
  form.data_warehouse_temp = '{}'
}

const openCreate = async () => {
  resetForm()
  drawerMode.value = 'create'
  drawerVisible.value = true
  if (!devices.value.length) await loadDevices()
}

const openEdit = async (row) => {
  resetForm()
  drawerMode.value = 'edit'
  form.ID = row?.ID
  form.name = row?.name || ''
  form.timeout = Number(row?.timeout || 30)
  form.allow_redirects = Boolean(row?.allow_redirects)
  form.verify = Boolean(row?.verify)
  form.preparatorySteps = Number(row?.preparatorySteps || 0)
  form.setup_case_id = Number(row?.setup_case_id || 0)
  form.retry = Number(row?.retry || 0)
  form.androidDeviceOptionsId = row?.androidDeviceOptionsId ?? row?.android_device_options_id
  form.header_temp = JSON.stringify(row?.header_temp || row?.headerTemp || {}, null, 2)
  form.variable_temp = JSON.stringify(row?.variable_temp || row?.variableTemp || {}, null, 2)
  form.parameters_temp = JSON.stringify(row?.parameters_temp || row?.parametersTemp || {}, null, 2)
  form.data_warehouse_temp = JSON.stringify(row?.data_warehouse_temp || row?.dataWarehouseTemp || {}, null, 2)
  drawerVisible.value = true
  if (!devices.value.length) await loadDevices()
}

const submit = async () => {
  const invalid = await formRef.value?.validate?.()
  if (invalid) return
  saving.value = true
  try {
    const payload = {
      ID: form.ID,
      name: form.name,
      timeout: Number(form.timeout || 0),
      allow_redirects: Boolean(form.allow_redirects),
      verify: Boolean(form.verify),
      preparatorySteps: Number(form.preparatorySteps || 0),
      setup_case_id: Number(form.setup_case_id || 0),
      retry: Number(form.retry || 0),
      type: currentType.value,
      androidDeviceOptionsId: form.androidDeviceOptionsId || undefined,
      header_temp: parseMaybeJSON(form.header_temp, 'runConfig.headerTemp'),
      variable_temp: parseMaybeJSON(form.variable_temp, 'runConfig.variableTemp'),
      parameters_temp: parseMaybeJSON(form.parameters_temp, 'runConfig.parametersTemp'),
      data_warehouse_temp: parseMaybeJSON(form.data_warehouse_temp, 'runConfig.dataWarehouseTemp')
    }
    if (drawerMode.value === 'edit') {
      await updateRunConfig(payload)
      Message.success(t('runConfig.updateSuccess'))
    } else {
      await createRunConfig(payload)
      Message.success(t('runConfig.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('runConfig.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('runConfig.deleteConfirmTitle'),
    content: t('runConfig.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteRunConfig(id)
        Message.success(t('runConfig.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('runConfig.deleteError'))
      }
    }
  })
}

onMounted(async () => {
  await loadList()
})
</script>
