<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('menu.androidui.scheduledTasks')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('timerTask.configName')">
          <a-input v-model="searchForm.configName" allow-clear />
        </a-form-item>
        <a-form-item :label="t('timerTask.envName')">
          <a-input v-model="searchForm.envName" allow-clear />
        </a-form-item>
        <a-form-item :label="t('timerTask.messageName')">
          <a-input v-model="searchForm.messageName" allow-clear />
        </a-form-item>
        <a-form-item :label="t('timerTask.createdAtStart')">
          <a-input v-model="searchForm.createdAtStart" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item :label="t('timerTask.createdAtEnd')">
          <a-input v-model="searchForm.createdAtEnd" allow-clear placeholder="YYYY-MM-DD HH:mm:ss" />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('timerTask.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('timerTask.create') }}</a-button>
      </a-space>

      <a-table :data="rows" :columns="columns" :loading="loading" row-key="ID" :pagination="false">
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openTaskCases(record)">{{ t('timerTask.manageCases') }}</a-button>
            <a-button type="text" @click="openEdit(record)">{{ t('timerTask.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('timerTask.delete') }}</a-button>
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
        <a-form-item field="name" :label="t('timerTask.name')" :rules="[{ required: true, message: t('timerTask.nameRequired') }]">
          <a-input v-model="form.name" />
        </a-form-item>
        <a-form-item field="runTime" :label="t('timerTask.runTime')" :rules="[{ required: true, message: t('timerTask.runTimeRequired') }]">
          <a-input v-model="form.runTime" :placeholder="t('timerTask.runTimePlaceholder')" />
        </a-form-item>
        <a-form-item field="status" :label="t('timerTask.status')">
          <a-switch v-model="form.status" />
        </a-form-item>
        <a-form-item field="runnerNodeName" :label="t('timerTask.runnerNodeName')">
          <a-input v-model="form.runnerNodeName" allow-clear />
        </a-form-item>
        <a-form-item field="runNumber" :label="t('timerTask.runNumber')">
          <a-input-number v-model="form.runNumber" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="failfast" :label="t('timerTask.failfast')">
          <a-switch v-model="form.failfast" />
        </a-form-item>
        <a-form-item field="configID" :label="t('timerTask.configId')">
          <a-input-number v-model="form.configID" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="tag" :label="t('timerTask.tag')">
          <a-input v-model="form.tag" allow-clear :placeholder="t('timerTask.tagPlaceholder')" />
        </a-form-item>
        <a-form-item field="envID" :label="t('timerTask.envId')">
          <a-input-number v-model="form.envID" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="notifyEnabled" :label="t('timerTask.notifyEnabled')">
          <a-switch v-model="form.notifyEnabled" />
        </a-form-item>
        <a-form-item field="notifyRule" :label="t('timerTask.notifyRule')">
          <a-input v-model="form.notifyRule" allow-clear />
        </a-form-item>
        <a-form-item field="messageID" :label="t('timerTask.messageId')">
          <a-input-number v-model="form.messageID" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="describe" :label="t('timerTask.describe')">
          <a-textarea v-model="form.describe" :auto-size="{ minRows: 3, maxRows: 8 }" />
        </a-form-item>
      </a-form>

      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('timerTask.save') }}</a-button>
          <a-button @click="drawerVisible = false">{{ t('timerTask.cancel') }}</a-button>
        </a-space>
      </template>
    </a-drawer>

    <a-modal v-model:visible="taskCasesVisible" :title="t('timerTask.manageCases')" width="920px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-space>
          <a-button type="primary" :loading="taskCasesLoading" @click="openCasePicker">{{ t('timerTask.addCase') }}</a-button>
          <a-button :loading="taskCasesSaving" @click="saveTaskCaseOrder">{{ t('timerTask.saveOrder') }}</a-button>
        </a-space>
        <a-table :data="taskCases" :columns="taskCasesColumns" :loading="taskCasesLoading" row-key="parentId" :pagination="false">
          <template #order="{ rowIndex }">
            {{ rowIndex + 1 }}
          </template>
          <template #mgrActions="{ record, rowIndex }">
            <a-space>
              <a-button type="text" :disabled="rowIndex === 0" @click="moveTaskCase(rowIndex, -1)">{{ t('timerTask.up') }}</a-button>
              <a-button type="text" :disabled="rowIndex === taskCases.length - 1" @click="moveTaskCase(rowIndex, 1)">{{ t('timerTask.down') }}</a-button>
              <a-button type="text" status="danger" @click="removeTaskCase(record)">{{ t('timerTask.remove') }}</a-button>
            </a-space>
          </template>
        </a-table>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="taskCasesVisible = false">{{ t('timerTask.close') }}</a-button>
        </a-space>
      </template>
    </a-modal>

    <a-modal v-model:visible="casePickerVisible" :title="t('timerTask.pickCase')" width="900px" :mask-closable="false">
      <a-space direction="vertical" fill size="12">
        <a-form :model="casePickerSearch" layout="inline">
          <a-form-item :label="t('autoCase.caseName')">
            <a-input v-model="casePickerSearch.caseName" allow-clear />
          </a-form-item>
          <a-form-item>
            <a-space>
              <a-button type="primary" :loading="casePickerLoading" @click="loadCasePickerList">{{ t('common.search') }}</a-button>
              <a-button @click="resetCasePicker">{{ t('autoCase.reset') }}</a-button>
            </a-space>
          </a-form-item>
        </a-form>
        <a-table :data="casePickerRows" :columns="casePickerColumns" :loading="casePickerLoading" row-key="ID" :pagination="false">
          <template #pick="{ record }">
            <a-button type="text" @click="addCaseToTask(record)">{{ t('timerTask.add') }}</a-button>
          </template>
        </a-table>
        <div style="display: flex; justify-content: flex-end">
          <a-pagination
            :current="casePickerPage"
            :page-size="casePickerPageSize"
            :total="casePickerTotal"
            show-total
            show-jumper
            show-page-size
            :page-size-options="[10, 30, 50, 100]"
            @change="handleCasePickerPageChange"
            @page-size-change="handleCasePickerPageSizeChange"
          />
        </div>
      </a-space>
      <template #footer>
        <a-space>
          <a-button @click="casePickerVisible = false">{{ t('timerTask.close') }}</a-button>
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
  addTimerTaskCase,
  createTimerTask,
  deleteTimerTask,
  deleteTimerTaskCaseRef,
  getAutoCaseList,
  getTimerTaskCases,
  getTimerTaskList,
  sortTimerTaskCases,
  updateTimerTask
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
  configName: '',
  envName: '',
  messageName: '',
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
  runTime: '',
  status: true,
  runnerNodeName: '',
  runNumber: 1,
  failfast: false,
  configID: 0,
  tag: '',
  envID: 0,
  notifyEnabled: false,
  notifyRule: '',
  messageID: 0,
  describe: ''
})

const columns = computed(() => [
  { title: t('timerTask.createdAt'), dataIndex: 'CreatedAt' },
  { title: t('timerTask.name'), dataIndex: 'name' },
  { title: t('timerTask.runTime'), dataIndex: 'runTime' },
  { title: t('timerTask.nextRunTime'), dataIndex: 'nextRunTime' },
  { title: t('timerTask.status'), dataIndex: 'status', width: 90 },
  { title: t('timerTask.runNumber'), dataIndex: 'runNumber', width: 100 },
  { title: t('timerTask.runnerNodeName'), dataIndex: 'runnerNodeName' },
  { title: t('timerTask.failfast'), dataIndex: 'failfast', width: 100 },
  { title: t('timerTask.configName'), dataIndex: 'configName' },
  { title: t('timerTask.tag'), dataIndex: 'tag' },
  { title: t('timerTask.envName'), dataIndex: 'envName' },
  { title: t('timerTask.messageName'), dataIndex: 'messageName' },
  { title: t('timerTask.notifyEnabled'), dataIndex: 'notifyEnabled', width: 120 },
  { title: t('timerTask.notifyRule'), dataIndex: 'notifyRule' },
  { title: t('timerTask.describe'), dataIndex: 'describe' },
  { title: t('timerTask.actions'), slotName: 'actions', width: 180 }
])

const drawerTitle = computed(() => (drawerMode.value === 'edit' ? t('timerTask.edit') : t('timerTask.create')))

const taskCasesVisible = ref(false)
const taskCasesLoading = ref(false)
const taskCasesSaving = ref(false)
const currentTaskId = ref(0)
const taskCases = ref([])

const taskCasesColumns = computed(() => [
  { title: t('timerTask.order'), slotName: 'order', width: 80 },
  { title: t('autoCase.caseName'), dataIndex: 'caseName' },
  { title: t('autoCase.envName'), dataIndex: 'envName' },
  { title: t('autoCase.configName'), dataIndex: 'configName' },
  { title: t('timerTask.actions'), slotName: 'mgrActions', width: 240 }
])

const casePickerVisible = ref(false)
const casePickerLoading = ref(false)
const casePickerRows = ref([])
const casePickerTotal = ref(0)
const casePickerPage = ref(1)
const casePickerPageSize = ref(10)
const casePickerSearch = reactive({ caseName: '' })

const casePickerColumns = computed(() => [
  { title: 'ID', dataIndex: 'ID', width: 90 },
  { title: t('autoCase.caseName'), dataIndex: 'caseName' },
  { title: t('autoCase.status'), dataIndex: 'status', width: 120 },
  { title: t('timerTask.actions'), slotName: 'pick', width: 120 }
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
      pageSize: pageSize.value,
      type: currentType.value,
      configName: searchForm.configName || undefined,
      envName: searchForm.envName || undefined,
      messageName: searchForm.messageName || undefined
    }
    if (createdAtRange) query['createdAtRange[]'] = createdAtRange
    const res = await getTimerTaskList(query)
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('timerTask.loadError'))
  } finally {
    loading.value = false
  }
}

const handleSearch = async () => {
  page.value = 1
  await loadList()
}

const handleReset = async () => {
  searchForm.configName = ''
  searchForm.envName = ''
  searchForm.messageName = ''
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
  form.runTime = ''
  form.status = true
  form.runnerNodeName = ''
  form.runNumber = 1
  form.failfast = false
  form.configID = 0
  form.tag = ''
  form.envID = 0
  form.notifyEnabled = false
  form.notifyRule = ''
  form.messageID = 0
  form.describe = ''
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
  form.runTime = row?.runTime || ''
  form.status = Boolean(row?.status)
  form.runnerNodeName = row?.runnerNodeName || ''
  form.runNumber = Number(row?.runNumber || 1)
  form.failfast = Boolean(row?.failfast)
  form.configID = Number(row?.configID || 0)
  form.tag = Array.isArray(row?.tag) ? row.tag.join(',') : (row?.tag || '')
  form.envID = Number(row?.envID || 0)
  form.notifyEnabled = Boolean(row?.notifyEnabled)
  form.notifyRule = row?.notifyRule || ''
  form.messageID = Number(row?.messageID || 0)
  form.describe = row?.describe || ''
  drawerVisible.value = true
}

const submit = async () => {
  const invalid = await formRef.value?.validate?.()
  if (invalid) return
  saving.value = true
  try {
    const tags = String(form.tag || '')
      .split(',')
      .map((s) => s.trim())
      .filter(Boolean)
    const payload = {
      ID: form.ID,
      type: currentType.value,
      name: form.name,
      runTime: form.runTime,
      status: Boolean(form.status),
      runnerNodeName: form.runnerNodeName || '',
      runNumber: Number(form.runNumber || 0),
      failfast: Boolean(form.failfast),
      configID: Number(form.configID || 0),
      tag: tags,
      envID: Number(form.envID || 0),
      notifyEnabled: Boolean(form.notifyEnabled),
      notifyRule: form.notifyRule || '',
      messageID: Number(form.messageID || 0),
      describe: form.describe || ''
    }
    if (drawerMode.value === 'edit') {
      await updateTimerTask(payload)
      Message.success(t('timerTask.updateSuccess'))
    } else {
      await createTimerTask(payload)
      Message.success(t('timerTask.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('timerTask.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('timerTask.deleteConfirmTitle'),
    content: t('timerTask.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteTimerTask(id)
        Message.success(t('timerTask.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('timerTask.deleteError'))
      }
    }
  })
}

const openTaskCases = async (row) => {
  const id = Number(row?.ID || 0)
  if (!id) return
  currentTaskId.value = id
  taskCasesVisible.value = true
  await refreshTaskCases()
}

const refreshTaskCases = async () => {
  if (!currentTaskId.value) return
  taskCasesLoading.value = true
  try {
    const list = await getTimerTaskCases(currentTaskId.value)
    taskCases.value = Array.isArray(list) ? list : []
  } catch (e) {
    Message.error(e?.message || t('timerTask.loadCasesError'))
  } finally {
    taskCasesLoading.value = false
  }
}

const moveTaskCase = (index, delta) => {
  const target = index + delta
  if (target < 0 || target >= taskCases.value.length) return
  const next = taskCases.value.slice()
  const [item] = next.splice(index, 1)
  next.splice(target, 0, item)
  taskCases.value = next
}

const saveTaskCaseOrder = async () => {
  if (!currentTaskId.value || !taskCases.value.length) return
  taskCasesSaving.value = true
  try {
    const data = taskCases.value
      .map((c, i) => ({ id: Number(c?.parentId || 0), sort: i + 1 }))
      .filter((x) => x.id)
    await sortTimerTaskCases(currentTaskId.value, data)
    Message.success(t('timerTask.saveOrderSuccess'))
    await refreshTaskCases()
  } catch (e) {
    Message.error(e?.message || t('timerTask.saveOrderError'))
  } finally {
    taskCasesSaving.value = false
  }
}

const removeTaskCase = (record) => {
  const refId = Number(record?.parentId || 0)
  if (!refId) return
  Modal.confirm({
    title: t('timerTask.removeConfirmTitle'),
    content: t('timerTask.removeConfirmContent'),
    onOk: async () => {
      try {
        await deleteTimerTaskCaseRef(refId)
        Message.success(t('timerTask.removeSuccess'))
        await refreshTaskCases()
      } catch (e) {
        Message.error(e?.message || t('timerTask.removeError'))
      }
    }
  })
}

const openCasePicker = async () => {
  casePickerVisible.value = true
  casePickerPage.value = 1
  await loadCasePickerList()
}

const loadCasePickerList = async () => {
  casePickerLoading.value = true
  try {
    const res = await getAutoCaseList({
      page: casePickerPage.value,
      pageSize: casePickerPageSize.value,
      type: currentType.value,
      caseName: casePickerSearch.caseName || undefined
    })
    const list = res?.list || res?.List || []
    casePickerRows.value = Array.isArray(list) ? list : []
    casePickerTotal.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('timerTask.pickLoadError'))
  } finally {
    casePickerLoading.value = false
  }
}

const resetCasePicker = async () => {
  casePickerSearch.caseName = ''
  casePickerPage.value = 1
  await loadCasePickerList()
}

const handleCasePickerPageChange = async (p) => {
  casePickerPage.value = p
  await loadCasePickerList()
}

const handleCasePickerPageSizeChange = async (s) => {
  casePickerPageSize.value = s
  casePickerPage.value = 1
  await loadCasePickerList()
}

const addCaseToTask = async (record) => {
  const caseId = Number(record?.ID || 0)
  if (!currentTaskId.value || !caseId) return
  try {
    await addTimerTaskCase(currentTaskId.value, caseId)
    Message.success(t('timerTask.addSuccess'))
    await refreshTaskCases()
  } catch (e) {
    Message.error(e?.message || t('timerTask.addError'))
  }
}

onMounted(async () => {
  await loadList()
})
</script>
