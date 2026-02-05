<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.testSteps')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('autoCaseStep.name')">
          <a-input v-model="searchForm.name" allow-clear :placeholder="t('autoCaseStep.namePlaceholder')" />
        </a-form-item>
        <a-form-item :label="t('autoCaseStep.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('autoCaseStep.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('autoCaseStep.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('autoCaseStep.create') }}</a-button>
      </a-space>

      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openStepManager(record)">{{ t('autoCaseStep.manageSteps') }}</a-button>
            <a-button type="text" @click="openEdit(record)">{{ t('autoCaseStep.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('autoCaseStep.delete') }}</a-button>
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
        <a-form-item field="name" :label="t('autoCaseStep.name')" :rules="[{ required: true, message: t('autoCaseStep.nameRequired') }]">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item field="loops" :label="t('autoCaseStep.loops')">
          <a-input-number v-model="form.loops" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="retry" :label="t('autoCaseStep.retry')">
          <a-input-number v-model="form.retry" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="ignore_popup" :label="t('autoCaseStep.ignorePopup')">
          <a-switch v-model="form.ignore_popup" />
        </a-form-item>
        <a-form-item field="variables" :label="t('autoCaseStep.variables')">
          <a-textarea v-model="form.variables" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="parameters" :label="t('autoCaseStep.parameters')">
          <a-textarea v-model="form.parameters" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="setup_hooks" :label="t('autoCaseStep.setupHooks')">
          <a-textarea v-model="form.setup_hooks" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="teardown_hooks" :label="t('autoCaseStep.teardownHooks')">
          <a-textarea v-model="form.teardown_hooks" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="extract" :label="t('autoCaseStep.extract')">
          <a-textarea v-model="form.extract" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="validate" :label="t('autoCaseStep.validate')">
          <a-textarea v-model="form.validate" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
        <a-form-item field="export" :label="t('autoCaseStep.export')">
          <a-textarea v-model="form.export" :auto-size="{ minRows: 4, maxRows: 12 }" />
        </a-form-item>
      </a-form>

      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('autoCaseStep.save') }}</a-button>
          <a-button @click="drawerVisible = false">{{ t('autoCaseStep.cancel') }}</a-button>
        </a-space>
      </template>
    </a-drawer>

    <a-modal v-model:visible="stepMgrVisible" :title="t('autoCaseStep.manageSteps')" width="860px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-space>
          <a-button type="primary" :loading="stepMgrLoading" @click="openStepPicker">{{ t('autoCaseStep.addStep') }}</a-button>
          <a-button :loading="stepMgrSaving" @click="saveStepOrder">{{ t('autoCaseStep.saveOrder') }}</a-button>
        </a-space>
        <a-table :data="stepApis" :columns="stepMgrColumns" :loading="stepMgrLoading" row-key="ID" :pagination="false">
          <template #order="{ rowIndex }">
            {{ rowIndex + 1 }}
          </template>
          <template #mgrActions="{ record, rowIndex }">
            <a-space>
              <a-button type="text" :disabled="rowIndex === 0" @click="moveStep(rowIndex, -1)">{{ t('autoCaseStep.up') }}</a-button>
              <a-button type="text" :disabled="rowIndex === stepApis.length - 1" @click="moveStep(rowIndex, 1)">{{ t('autoCaseStep.down') }}</a-button>
              <a-button type="text" status="danger" @click="removeStep(record)">{{ t('autoCaseStep.remove') }}</a-button>
            </a-space>
          </template>
        </a-table>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="stepMgrVisible = false">{{ t('autoCaseStep.close') }}</a-button>
        </a-space>
      </template>
    </a-modal>

    <a-modal v-model:visible="stepPickerVisible" :title="t('autoCaseStep.pickStep')" width="900px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-form :model="stepPickerSearch" layout="inline">
          <a-form-item :label="t('autoStep.name')">
            <a-input v-model="stepPickerSearch.name" allow-clear />
          </a-form-item>
          <a-form-item>
            <a-space>
              <a-button type="primary" :loading="stepPickerLoading" @click="loadStepPickerList">{{ t('common.search') }}</a-button>
              <a-button @click="resetStepPicker">{{ t('autoStep.reset') }}</a-button>
            </a-space>
          </a-form-item>
        </a-form>
        <a-table :data="stepPickerRows" :columns="stepPickerColumns" :loading="stepPickerLoading" row-key="ID" :pagination="false">
          <template #pick="{ record }">
            <a-button type="text" @click="addStepToGroup(record)">{{ t('autoCaseStep.add') }}</a-button>
          </template>
        </a-table>
        <div style="display: flex; justify-content: flex-end">
          <a-pagination
            :current="stepPickerPage"
            :page-size="stepPickerPageSize"
            :total="stepPickerTotal"
            show-total
            show-jumper
            show-page-size
            :page-size-options="[10, 30, 50, 100]"
            @change="handleStepPickerPageChange"
            @page-size-change="handleStepPickerPageSizeChange"
          />
        </div>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="stepPickerVisible = false">{{ t('autoCaseStep.close') }}</a-button>
        </a-space>
      </template>
    </a-modal>
  </a-space>
</template>

<script setup>
import { Message, Modal } from '@arco-design/web-vue'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  addAutoCaseStepApi,
  createAutoCaseStep,
  deleteAutoCaseStep,
  deleteAutoCaseStepApi,
  findAutoCaseStepApis,
  getAutoCaseStepList,
  getAutoStepList,
  sortAutoCaseStepApis,
  updateAutoCaseStep
} from '@/services/appBridge'

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
  variables: '{}',
  parameters: '{}',
  setup_hooks: '[]',
  teardown_hooks: '[]',
  extract: '{}',
  validate: '[]',
  export: '[]'
})

const columns = computed(() => [
  { title: t('autoCaseStep.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('autoCaseStep.name'), dataIndex: 'name' },
  { title: t('autoCaseStep.envName'), dataIndex: 'envName' },
  { title: t('autoCaseStep.configName'), dataIndex: 'configName' },
  { title: t('autoCaseStep.loops'), dataIndex: 'loops', width: 80 },
  { title: t('autoCaseStep.actions'), slotName: 'actions', width: 180 }
])

const stepMgrVisible = ref(false)
const stepMgrLoading = ref(false)
const stepMgrSaving = ref(false)
const currentGroupId = ref(0)
const stepApis = ref([])

const stepMgrColumns = computed(() => [
  { title: t('autoCaseStep.order'), slotName: 'order', width: 80 },
  { title: t('autoStep.name'), dataIndex: 'name' },
  { title: 'ID', dataIndex: 'ID', width: 90 },
  { title: t('autoCaseStep.actions'), slotName: 'mgrActions', width: 220 }
])

const stepPickerVisible = ref(false)
const stepPickerLoading = ref(false)
const stepPickerRows = ref([])
const stepPickerTotal = ref(0)
const stepPickerPage = ref(1)
const stepPickerPageSize = ref(10)
const stepPickerSearch = reactive({ name: '' })

const stepPickerColumns = computed(() => [
  { title: 'ID', dataIndex: 'ID', width: 90 },
  { title: t('autoStep.name'), dataIndex: 'name' },
  { title: t('autoStep.actionCount'), width: 120, render: ({ record }) => (Array.isArray(record?.android?.actions) ? record.android.actions.length : 0) },
  { title: t('autoCaseStep.actions'), slotName: 'pick', width: 120 }
])

const drawerTitle = computed(() => (drawerMode.value === 'edit' ? t('autoCaseStep.edit') : t('autoCaseStep.create')))

const parseJSON = (text, labelKey) => {
  const s = String(text || '').trim()
  if (!s) return null
  try {
    return JSON.parse(s)
  } catch (e) {
    throw new Error(t('autoCaseStep.jsonInvalid', { field: t(labelKey) }))
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
    const res = await getAutoCaseStepList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.loadError'))
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
      variables: parseJSON(form.variables, 'autoCaseStep.variables'),
      parameters: parseJSON(form.parameters, 'autoCaseStep.parameters'),
      setup_hooks: parseJSON(form.setup_hooks, 'autoCaseStep.setupHooks'),
      teardown_hooks: parseJSON(form.teardown_hooks, 'autoCaseStep.teardownHooks'),
      extract: parseJSON(form.extract, 'autoCaseStep.extract'),
      validate: parseJSON(form.validate, 'autoCaseStep.validate'),
      export: parseJSON(form.export, 'autoCaseStep.export')
    }
    if (drawerMode.value === 'edit') {
      await updateAutoCaseStep(payload)
      Message.success(t('autoCaseStep.updateSuccess'))
    } else {
      await createAutoCaseStep(payload)
      Message.success(t('autoCaseStep.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('autoCaseStep.deleteConfirmTitle'),
    content: t('autoCaseStep.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteAutoCaseStep(id)
        Message.success(t('autoCaseStep.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('autoCaseStep.deleteError'))
      }
    }
  })
}

const openStepManager = async (row) => {
  const id = Number(row?.ID || 0)
  if (!id) return
  currentGroupId.value = id
  stepMgrVisible.value = true
  await refreshStepManager()
}

const refreshStepManager = async () => {
  if (!currentGroupId.value) return
  stepMgrLoading.value = true
  try {
    const list = await findAutoCaseStepApis(currentGroupId.value)
    stepApis.value = Array.isArray(list) ? list : []
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.loadStepsError'))
  } finally {
    stepMgrLoading.value = false
  }
}

const moveStep = (index, delta) => {
  const target = index + delta
  if (target < 0 || target >= stepApis.value.length) return
  const next = stepApis.value.slice()
  const [item] = next.splice(index, 1)
  next.splice(target, 0, item)
  stepApis.value = next
}

const saveStepOrder = async () => {
  if (!stepApis.value.length) return
  stepMgrSaving.value = true
  try {
    const data = stepApis.value.map((s, i) => ({ id: Number(s?.ID || 0), sort: i + 1 })).filter((x) => x.id)
    await sortAutoCaseStepApis(data)
    Message.success(t('autoCaseStep.saveOrderSuccess'))
    await refreshStepManager()
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.saveOrderError'))
  } finally {
    stepMgrSaving.value = false
  }
}

const removeStep = (record) => {
  const id = Number(record?.ID || 0)
  if (!id) return
  Modal.confirm({
    title: t('autoCaseStep.removeConfirmTitle'),
    content: t('autoCaseStep.removeConfirmContent'),
    onOk: async () => {
      try {
        await deleteAutoCaseStepApi(id)
        Message.success(t('autoCaseStep.removeSuccess'))
        await refreshStepManager()
      } catch (e) {
        Message.error(e?.message || t('autoCaseStep.removeError'))
      }
    }
  })
}

const openStepPicker = async () => {
  stepPickerVisible.value = true
  stepPickerPage.value = 1
  await loadStepPickerList()
}

const loadStepPickerList = async () => {
  stepPickerLoading.value = true
  try {
    const res = await getAutoStepList({
      page: stepPickerPage.value,
      pageSize: stepPickerPageSize.value,
      type: currentType.value,
      name: stepPickerSearch.name || undefined
    })
    const list = res?.list || res?.List || []
    stepPickerRows.value = Array.isArray(list) ? list : []
    stepPickerTotal.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.pickLoadError'))
  } finally {
    stepPickerLoading.value = false
  }
}

const resetStepPicker = async () => {
  stepPickerSearch.name = ''
  stepPickerPage.value = 1
  await loadStepPickerList()
}

const handleStepPickerPageChange = async (p) => {
  stepPickerPage.value = p
  await loadStepPickerList()
}

const handleStepPickerPageSizeChange = async (s) => {
  stepPickerPageSize.value = s
  stepPickerPage.value = 1
  await loadStepPickerList()
}

const addStepToGroup = async (record) => {
  const apiId = Number(record?.ID || 0)
  if (!apiId || !currentGroupId.value) return
  try {
    await addAutoCaseStepApi(currentGroupId.value, apiId, stepApis.value.length + 1)
    Message.success(t('autoCaseStep.addSuccess'))
    await refreshStepManager()
  } catch (e) {
    Message.error(e?.message || t('autoCaseStep.addError'))
  }
}

onMounted(async () => {
  await loadList()
})
</script>
