<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.testCases')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('autoCase.caseName')">
          <a-input v-model="searchForm.caseName" allow-clear :placeholder="t('autoCase.caseNamePlaceholder')" />
        </a-form-item>
        <a-form-item :label="t('autoCase.status')">
          <a-select v-model="searchForm.status" allow-clear style="width: 180px">
            <a-option v-for="s in statusOptions" :key="s" :value="s">{{ s }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item :label="t('autoCase.configName')">
          <a-input v-model="searchForm.configName" allow-clear />
        </a-form-item>
        <a-form-item :label="t('autoCase.envName')">
          <a-input v-model="searchForm.envName" allow-clear />
        </a-form-item>
        <a-form-item :label="t('autoCase.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('autoCase.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('autoCase.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('autoCase.create') }}</a-button>
      </a-space>

      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openCaseSteps(record)">{{ t('autoCase.manageSteps') }}</a-button>
            <a-button type="text" @click="openEdit(record)">{{ t('autoCase.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('autoCase.delete') }}</a-button>
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

    <a-drawer v-model:visible="drawerVisible" :title="drawerTitle" width="680px" :mask-closable="false">
      <a-form ref="formRef" :model="form" layout="vertical">
        <a-form-item field="caseName" :label="t('autoCase.caseName')" :rules="[{ required: true, message: t('autoCase.caseNameRequired') }]">
          <a-input v-model="form.caseName" />
        </a-form-item>
        <a-form-item field="runNumber" :label="t('autoCase.runNumber')">
          <a-input-number v-model="form.runNumber" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="status" :label="t('autoCase.status')" :rules="[{ required: true, message: t('autoCase.statusRequired') }]">
          <a-select v-model="form.status" allow-clear>
            <a-option v-for="s in statusOptions" :key="s" :value="s">{{ s }}</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="configID" :label="t('autoCase.configId')">
          <a-input-number v-model="form.configID" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="envID" :label="t('autoCase.envId')">
          <a-input-number v-model="form.envID" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="desc" :label="t('autoCase.desc')">
          <a-textarea v-model="form.desc" :auto-size="{ minRows: 3, maxRows: 8 }" />
        </a-form-item>
      </a-form>

      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('autoCase.save') }}</a-button>
          <a-button @click="drawerVisible = false">{{ t('autoCase.cancel') }}</a-button>
        </a-space>
      </template>
    </a-drawer>

    <a-modal v-model:visible="caseStepsVisible" :title="t('autoCase.manageSteps')" width="920px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-space>
          <a-button type="primary" :loading="caseStepsLoading" @click="openCaseStepPicker">{{ t('autoCase.addStep') }}</a-button>
          <a-button :loading="caseStepsSaving" @click="saveCaseStepOrder">{{ t('autoCase.saveOrder') }}</a-button>
        </a-space>
        <a-table :data="caseSteps" :columns="caseStepsColumns" :loading="caseStepsLoading" row-key="parentId" :pagination="false">
          <template #order="{ rowIndex }">
            {{ rowIndex + 1 }}
          </template>
          <template #isConfig="{ record }">
            <a-switch :model-value="Boolean(record?.isConfig)" @change="(v) => updateCaseStepConfig(record, v, record?.isStepConfig)" />
          </template>
          <template #isStepConfig="{ record }">
            <a-switch :model-value="Boolean(record?.isStepConfig)" @change="(v) => updateCaseStepConfig(record, record?.isConfig, v)" />
          </template>
          <template #mgrActions="{ record, rowIndex }">
            <a-space>
              <a-button type="text" :disabled="rowIndex === 0" @click="moveCaseStep(rowIndex, -1)">{{ t('autoCase.up') }}</a-button>
              <a-button type="text" :disabled="rowIndex === caseSteps.length - 1" @click="moveCaseStep(rowIndex, 1)">{{ t('autoCase.down') }}</a-button>
              <a-button type="text" status="danger" @click="removeCaseStep(record)">{{ t('autoCase.remove') }}</a-button>
            </a-space>
          </template>
        </a-table>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="caseStepsVisible = false">{{ t('autoCase.close') }}</a-button>
        </a-space>
      </template>
    </a-modal>

    <a-modal v-model:visible="caseStepPickerVisible" :title="t('autoCase.pickStep')" width="900px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-form :model="caseStepPickerSearch" layout="inline">
          <a-form-item :label="t('autoCaseStep.name')">
            <a-input v-model="caseStepPickerSearch.name" allow-clear />
          </a-form-item>
          <a-form-item>
            <a-space>
              <a-button type="primary" :loading="caseStepPickerLoading" @click="loadCaseStepPickerList">{{ t('common.search') }}</a-button>
              <a-button @click="resetCaseStepPicker">{{ t('autoCaseStep.reset') }}</a-button>
            </a-space>
          </a-form-item>
        </a-form>
        <a-table :data="caseStepPickerRows" :columns="caseStepPickerColumns" :loading="caseStepPickerLoading" row-key="ID" :pagination="false">
          <template #pick="{ record }">
            <a-button type="text" @click="addStepToCase(record)">{{ t('autoCase.add') }}</a-button>
          </template>
        </a-table>
        <div style="display: flex; justify-content: flex-end">
          <a-pagination
            :current="caseStepPickerPage"
            :page-size="caseStepPickerPageSize"
            :total="caseStepPickerTotal"
            show-total
            show-jumper
            show-page-size
            :page-size-options="[10, 30, 50, 100]"
            @change="handleCaseStepPickerPageChange"
            @page-size-change="handleCaseStepPickerPageSizeChange"
          />
        </div>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="caseStepPickerVisible = false">{{ t('autoCase.close') }}</a-button>
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
  addAutoCaseStep,
  createAutoCase,
  deleteAutoCase,
  deleteAutoCaseStepRef,
  getAutoCaseList,
  getAutoCaseStepList,
  getAutoCaseSteps,
  setAutoCaseStepConfig,
  sortAutoCaseSteps,
  updateAutoCase
} from '@/services/appBridge'

const { t } = useI18n()
const route = useRoute()
const currentType = computed(() => String(route.query?.type || 'android'))

const statusOptions = ['测试中', '待评审', '评审不通过', '已发布', '禁用', '已废弃']

const loading = ref(false)
const rows = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const searchForm = reactive({
  caseName: '',
  status: '',
  configName: '',
  envName: '',
  createdAtStart: '',
  createdAtEnd: ''
})

const drawerVisible = ref(false)
const saving = ref(false)
const drawerMode = ref('create')
const formRef = ref()

const form = reactive({
  ID: undefined,
  caseName: '',
  runNumber: 1,
  status: '待评审',
  configID: 0,
  envID: 0,
  desc: ''
})

const columns = computed(() => [
  { title: t('autoCase.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('autoCase.caseName'), dataIndex: 'caseName' },
  { title: t('autoCase.type'), dataIndex: 'type', width: 100 },
  { title: t('autoCase.runNumber'), dataIndex: 'runNumber', width: 100 },
  { title: t('autoCase.status'), dataIndex: 'status', width: 120 },
  { title: t('autoCase.envName'), dataIndex: 'envName' },
  { title: t('autoCase.configName'), dataIndex: 'configName' },
  { title: t('autoCase.actions'), slotName: 'actions', width: 180 }
])

const caseStepsVisible = ref(false)
const caseStepsLoading = ref(false)
const caseStepsSaving = ref(false)
const currentCaseId = ref(0)
const caseSteps = ref([])

const caseStepsColumns = computed(() => [
  { title: t('autoCase.order'), slotName: 'order', width: 80 },
  { title: t('autoCaseStep.name'), dataIndex: 'name' },
  { title: t('autoCaseStep.envName'), dataIndex: 'envName' },
  { title: t('autoCaseStep.configName'), dataIndex: 'configName' },
  { title: t('autoCase.isConfig'), slotName: 'isConfig', width: 120 },
  { title: t('autoCase.isStepConfig'), slotName: 'isStepConfig', width: 150 },
  { title: t('autoCase.actions'), slotName: 'mgrActions', width: 240 }
])

const caseStepPickerVisible = ref(false)
const caseStepPickerLoading = ref(false)
const caseStepPickerRows = ref([])
const caseStepPickerTotal = ref(0)
const caseStepPickerPage = ref(1)
const caseStepPickerPageSize = ref(10)
const caseStepPickerSearch = reactive({ name: '' })

const caseStepPickerColumns = computed(() => [
  { title: 'ID', dataIndex: 'ID', width: 90 },
  { title: t('autoCaseStep.name'), dataIndex: 'name' },
  { title: t('autoCaseStep.envName'), dataIndex: 'envName' },
  { title: t('autoCaseStep.configName'), dataIndex: 'configName' },
  { title: t('autoCase.actions'), slotName: 'pick', width: 120 }
])

const drawerTitle = computed(() => (drawerMode.value === 'edit' ? t('autoCase.edit') : t('autoCase.create')))

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
      caseName: searchForm.caseName || undefined,
      status: searchForm.status || undefined,
      configName: searchForm.configName || undefined,
      envName: searchForm.envName || undefined
    }
    if (createdAtRange) query['createdAtRange[]'] = createdAtRange
    const res = await getAutoCaseList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoCase.loadError'))
  } finally {
    loading.value = false
  }
}

const handleSearch = async () => {
  page.value = 1
  await loadList()
}

const handleReset = async () => {
  searchForm.caseName = ''
  searchForm.status = ''
  searchForm.configName = ''
  searchForm.envName = ''
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
  form.caseName = ''
  form.runNumber = 1
  form.status = '待评审'
  form.configID = 0
  form.envID = 0
  form.desc = ''
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
  form.caseName = row?.caseName || ''
  form.runNumber = Number(row?.runNumber || 1)
  form.status = row?.status || '待评审'
  form.configID = Number(row?.configID || 0)
  form.envID = Number(row?.envID || 0)
  form.desc = row?.desc || ''
  drawerVisible.value = true
}

const submit = async () => {
  const invalid = await formRef.value?.validate?.()
  if (invalid) return
  saving.value = true
  try {
    const payload = {
      ID: form.ID,
      type: currentType.value,
      caseName: form.caseName,
      runNumber: Number(form.runNumber || 0),
      status: form.status,
      configID: Number(form.configID || 0),
      envID: Number(form.envID || 0),
      desc: form.desc || ''
    }
    if (drawerMode.value === 'edit') {
      await updateAutoCase(payload)
      Message.success(t('autoCase.updateSuccess'))
    } else {
      await createAutoCase(payload)
      Message.success(t('autoCase.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('autoCase.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('autoCase.deleteConfirmTitle'),
    content: t('autoCase.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteAutoCase(id)
        Message.success(t('autoCase.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('autoCase.deleteError'))
      }
    }
  })
}

const openCaseSteps = async (row) => {
  const id = Number(row?.ID || 0)
  if (!id) return
  currentCaseId.value = id
  caseStepsVisible.value = true
  await refreshCaseSteps()
}

const refreshCaseSteps = async () => {
  if (!currentCaseId.value) return
  caseStepsLoading.value = true
  try {
    const list = await getAutoCaseSteps(currentCaseId.value)
    caseSteps.value = Array.isArray(list) ? list : []
  } catch (e) {
    Message.error(e?.message || t('autoCase.loadStepsError'))
  } finally {
    caseStepsLoading.value = false
  }
}

const moveCaseStep = (index, delta) => {
  const target = index + delta
  if (target < 0 || target >= caseSteps.value.length) return
  const next = caseSteps.value.slice()
  const [item] = next.splice(index, 1)
  next.splice(target, 0, item)
  caseSteps.value = next
}

const saveCaseStepOrder = async () => {
  if (!currentCaseId.value || !caseSteps.value.length) return
  caseStepsSaving.value = true
  try {
    const data = caseSteps.value
      .map((s, i) => ({ id: Number(s?.parentId || 0), sort: i + 1 }))
      .filter((x) => x.id)
    await sortAutoCaseSteps(currentCaseId.value, data)
    Message.success(t('autoCase.saveOrderSuccess'))
    await refreshCaseSteps()
  } catch (e) {
    Message.error(e?.message || t('autoCase.saveOrderError'))
  } finally {
    caseStepsSaving.value = false
  }
}

const removeCaseStep = (record) => {
  const refId = Number(record?.parentId || 0)
  if (!refId) return
  Modal.confirm({
    title: t('autoCase.removeConfirmTitle'),
    content: t('autoCase.removeConfirmContent'),
    onOk: async () => {
      try {
        await deleteAutoCaseStepRef(refId)
        Message.success(t('autoCase.removeSuccess'))
        await refreshCaseSteps()
      } catch (e) {
        Message.error(e?.message || t('autoCase.removeError'))
      }
    }
  })
}

const updateCaseStepConfig = async (record, isConfig, isStepConfig) => {
  const refId = Number(record?.parentId || 0)
  if (!refId) return
  try {
    await setAutoCaseStepConfig(refId, Boolean(isConfig), Boolean(isStepConfig))
    record.isConfig = Boolean(isConfig)
    record.isStepConfig = Boolean(isStepConfig)
  } catch (e) {
    Message.error(e?.message || t('autoCase.setConfigError'))
  }
}

const openCaseStepPicker = async () => {
  caseStepPickerVisible.value = true
  caseStepPickerPage.value = 1
  await loadCaseStepPickerList()
}

const loadCaseStepPickerList = async () => {
  caseStepPickerLoading.value = true
  try {
    const res = await getAutoCaseStepList({
      page: caseStepPickerPage.value,
      pageSize: caseStepPickerPageSize.value,
      type: currentType.value,
      name: caseStepPickerSearch.name || undefined
    })
    const list = res?.list || res?.List || []
    caseStepPickerRows.value = Array.isArray(list) ? list : []
    caseStepPickerTotal.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('autoCase.pickLoadError'))
  } finally {
    caseStepPickerLoading.value = false
  }
}

const resetCaseStepPicker = async () => {
  caseStepPickerSearch.name = ''
  caseStepPickerPage.value = 1
  await loadCaseStepPickerList()
}

const handleCaseStepPickerPageChange = async (p) => {
  caseStepPickerPage.value = p
  await loadCaseStepPickerList()
}

const handleCaseStepPickerPageSizeChange = async (s) => {
  caseStepPickerPageSize.value = s
  caseStepPickerPage.value = 1
  await loadCaseStepPickerList()
}

const addStepToCase = async (record) => {
  const stepId = Number(record?.ID || 0)
  if (!currentCaseId.value || !stepId) return
  try {
    await addAutoCaseStep(currentCaseId.value, stepId)
    Message.success(t('autoCase.addSuccess'))
    await refreshCaseSteps()
  } catch (e) {
    Message.error(e?.message || t('autoCase.addError'))
  }
}

onMounted(async () => {
  await loadList()
})
</script>
