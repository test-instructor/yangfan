<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.elementAction')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('autoStep.name')">
          <a-input v-model="searchForm.name" allow-clear :placeholder="t('autoStep.namePlaceholder')" />
        </a-form-item>
        <a-form-item :label="t('autoStep.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('autoStep.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('autoStep.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('autoStep.create') }}</a-button>
      </a-space>

      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #actionCount="{ record }">
          {{ getAndroidActionsCount(record) }}
        </template>
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openEdit(record)">{{ t('autoStep.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('autoStep.delete') }}</a-button>
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

    <a-drawer v-model:visible="drawerVisible" :title="drawerTitle" width="760px" :mask-closable="false">
      <a-form ref="formRef" :model="form" layout="vertical">
        <a-form-item field="name" :label="t('autoStep.name')" :rules="[{ required: true, message: t('autoStep.nameRequired') }]">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item field="loops" :label="t('autoStep.loops')">
          <a-input-number v-model="form.loops" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="retry" :label="t('autoStep.retry')">
          <a-input-number v-model="form.retry" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="ignore_popup" :label="t('autoStep.ignorePopup')">
          <a-switch v-model="form.ignore_popup" />
        </a-form-item>

        <a-form-item field="android" :label="t('autoStep.androidConfig')">
          <a-textarea v-model="form.android" :auto-size="{ minRows: 6, maxRows: 18 }" />
        </a-form-item>

        <a-form-item field="variables" :label="t('autoStep.variables')">
          <a-textarea v-model="form.variables" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="parameters" :label="t('autoStep.parameters')">
          <a-textarea v-model="form.parameters" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="setup_hooks" :label="t('autoStep.setupHooks')">
          <a-textarea v-model="form.setup_hooks" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="teardown_hooks" :label="t('autoStep.teardownHooks')">
          <a-textarea v-model="form.teardown_hooks" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="extract" :label="t('autoStep.extract')">
          <a-textarea v-model="form.extract" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="validate" :label="t('autoStep.validate')">
          <a-textarea v-model="form.validate" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="export" :label="t('autoStep.export')">
          <a-textarea v-model="form.export" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
      </a-form>

      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('autoStep.save') }}</a-button>
          <a-button @click="drawerVisible = false">{{ t('autoStep.cancel') }}</a-button>
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
import { createAutoStep, deleteAutoStep, getAutoStepList, updateAutoStep } from '@/services/appBridge'

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

const drawerVisible = ref(false)
const saving = ref(false)
const drawerMode = ref('create')
const formRef = ref()

const form = reactive({
  ID: undefined,
  name: '',
  loops: 1,
  retry: 0,
  ignore_popup: false,
  android: '{\n  \"actions\": []\n}',
  variables: '{}',
  parameters: '{}',
  setup_hooks: '[]',
  teardown_hooks: '[]',
  extract: '{}',
  validate: '[]',
  export: '[]'
})

const columns = computed(() => [
  { title: t('autoStep.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('autoStep.name'), dataIndex: 'name' },
  { title: t('autoStep.actionCount'), slotName: 'actionCount', width: 120 },
  { title: t('autoStep.actions'), slotName: 'actions', width: 180 }
])

const drawerTitle = computed(() => (drawerMode.value === 'edit' ? t('autoStep.edit') : t('autoStep.create')))

const parseJSON = (text, labelKey) => {
  const s = String(text || '').trim()
  if (!s) return null
  try {
    return JSON.parse(s)
  } catch (e) {
    throw new Error(t('autoStep.jsonInvalid', { field: t(labelKey) }))
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
    if (createdAtRange) query['createdAtRange[]'] = createdAtRange
    const res = await getAutoStepList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoStep.loadError'))
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
  form.loops = 1
  form.retry = 0
  form.ignore_popup = false
  form.android = '{\n  \"actions\": []\n}'
  form.variables = '{}'
  form.parameters = '{}'
  form.setup_hooks = '[]'
  form.teardown_hooks = '[]'
  form.extract = '{}'
  form.validate = '[]'
  form.export = '[]'
}

const openCreate = () => {
  resetForm()
  drawerMode.value = 'create'
  drawerVisible.value = true
}

const openEdit = (row) => {
  resetForm()
  drawerMode.value = 'edit'
  form.ID = row?.ID
  form.name = row?.name || ''
  form.loops = Number(row?.loops || 1)
  form.retry = Number(row?.retry || 0)
  form.ignore_popup = Boolean(row?.ignore_popup)
  form.android = JSON.stringify(row?.android || { actions: [] }, null, 2)
  form.variables = JSON.stringify(row?.variables || {}, null, 2)
  form.parameters = JSON.stringify(row?.parameters || {}, null, 2)
  form.setup_hooks = JSON.stringify(row?.setup_hooks || row?.setupHooks || [], null, 2)
  form.teardown_hooks = JSON.stringify(row?.teardown_hooks || row?.teardownHooks || [], null, 2)
  form.extract = JSON.stringify(row?.extract || {}, null, 2)
  form.validate = JSON.stringify(row?.validate || row?.validators || [], null, 2)
  form.export = JSON.stringify(row?.export || row?.step_export || [], null, 2)
  drawerVisible.value = true
}

const submit = async () => {
  const invalid = await formRef.value?.validate?.()
  if (invalid) return
  saving.value = true
  try {
    const payload = {
      ID: form.ID,
      name: form.name,
      loops: Number(form.loops || 0),
      retry: Number(form.retry || 0),
      ignore_popup: Boolean(form.ignore_popup),
      type: currentType.value,
      android: parseJSON(form.android, 'autoStep.androidConfig'),
      variables: parseJSON(form.variables, 'autoStep.variables'),
      parameters: parseJSON(form.parameters, 'autoStep.parameters'),
      setup_hooks: parseJSON(form.setup_hooks, 'autoStep.setupHooks'),
      teardown_hooks: parseJSON(form.teardown_hooks, 'autoStep.teardownHooks'),
      extract: parseJSON(form.extract, 'autoStep.extract'),
      validate: parseJSON(form.validate, 'autoStep.validate'),
      export: parseJSON(form.export, 'autoStep.export')
    }
    if (drawerMode.value === 'edit') {
      await updateAutoStep(payload)
      Message.success(t('autoStep.updateSuccess'))
    } else {
      await createAutoStep(payload)
      Message.success(t('autoStep.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('autoStep.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('autoStep.deleteConfirmTitle'),
    content: t('autoStep.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteAutoStep(id)
        Message.success(t('autoStep.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('autoStep.deleteError'))
      }
    }
  })
}

const getAndroidActionsCount = (row) => {
  const actions = row?.android?.actions
  return Array.isArray(actions) ? actions.length : 0
}

onMounted(async () => {
  await loadList()
})
</script>
