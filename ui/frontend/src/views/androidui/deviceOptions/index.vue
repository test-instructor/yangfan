<template>
  <a-space direction="vertical" fill size="16">
    <a-card :title="t('androidDeviceOptions.title')" :bordered="false">
      <a-form :model="searchForm" layout="inline">
        <a-form-item :label="t('androidDeviceOptions.name')">
          <a-input v-model="searchForm.name" :placeholder="t('androidDeviceOptions.namePlaceholder')" allow-clear />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" :loading="loading" @click="handleSearch">{{ t('common.search') }}</a-button>
            <a-button @click="handleReset">{{ t('androidDeviceOptions.reset') }}</a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </a-card>

    <a-card :bordered="false">
      <a-space style="margin-bottom: 12px">
        <a-button type="primary" @click="openCreate">{{ t('androidDeviceOptions.create') }}</a-button>
      </a-space>

      <a-table
        :data="rows"
        :columns="columns"
        :loading="loading"
        row-key="ID"
        :pagination="false"
      >
        <template #logOn="{ record }">
          <a-tag :color="record?.logOn ? 'green' : 'gray'">{{ record?.logOn ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-tag>
        </template>
        <template #ignorePopup="{ record }">
          <a-tag :color="record?.ignorePopup ? 'green' : 'gray'">{{ record?.ignorePopup ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-tag>
        </template>
        <template #composite="{ record }">
          <a-tag :color="record?.composite ? 'green' : 'gray'">{{ record?.composite ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-tag>
        </template>
        <template #uia2="{ record }">
          <a-tag :color="record?.uia2 ? 'green' : 'gray'">{{ record?.uia2 ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-tag>
        </template>
        <template #actions="{ record }">
          <a-space>
            <a-button type="text" @click="openDetail(record)">{{ t('androidDeviceOptions.detail') }}</a-button>
            <a-button type="text" @click="openEdit(record)">{{ t('androidDeviceOptions.edit') }}</a-button>
            <a-button type="text" status="danger" @click="confirmDelete(record)">{{ t('androidDeviceOptions.delete') }}</a-button>
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

    <a-drawer v-model:visible="drawerVisible" :title="drawerTitle" width="520px" :mask-closable="false">
      <a-form ref="formRef" :model="form" layout="vertical">
        <a-form-item field="name" :label="t('androidDeviceOptions.name')">
          <a-input v-model="form.name" allow-clear />
        </a-form-item>
        <a-form-item field="serial" :label="t('androidDeviceOptions.serial')" :rules="[{ required: true, message: t('androidDeviceOptions.serialRequired') }]">
          <a-input v-model="form.serial" />
        </a-form-item>
        <a-form-item field="logOn" :label="t('androidDeviceOptions.logOn')">
          <a-switch v-model="form.logOn" />
        </a-form-item>
        <a-form-item field="ignorePopup" :label="t('androidDeviceOptions.ignorePopup')">
          <a-switch v-model="form.ignorePopup" />
        </a-form-item>
        <a-form-item field="adbServerHost" :label="t('androidDeviceOptions.adbServerHost')">
          <a-input v-model="form.adbServerHost" />
        </a-form-item>
        <a-form-item field="adbServerPort" :label="t('androidDeviceOptions.adbServerPort')">
          <a-input-number v-model="form.adbServerPort" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="composite" :label="t('androidDeviceOptions.composite')">
          <a-switch v-model="form.composite" />
        </a-form-item>
        <a-form-item field="uia2" :label="t('androidDeviceOptions.uia2')">
          <a-switch v-model="form.uia2" />
        </a-form-item>
        <a-form-item field="uia2Ip" :label="t('androidDeviceOptions.uia2Ip')">
          <a-input v-model="form.uia2Ip" />
        </a-form-item>
        <a-form-item field="uia2Port" :label="t('androidDeviceOptions.uia2Port')">
          <a-input-number v-model="form.uia2Port" :min="0" :precision="0" style="width: 100%" />
        </a-form-item>
        <a-form-item field="uia2ServerPackageName" :label="t('androidDeviceOptions.uia2ServerPackageName')">
          <a-input v-model="form.uia2ServerPackageName" />
        </a-form-item>
        <a-form-item field="uia2ServerTestPackageName" :label="t('androidDeviceOptions.uia2ServerTestPackageName')">
          <a-input v-model="form.uia2ServerTestPackageName" />
        </a-form-item>
      </a-form>
      <template #footer>
        <a-space>
          <a-button type="primary" :loading="saving" @click="submit">{{ t('androidDeviceOptions.save') }}</a-button>
          <a-button @click="closeDrawer">{{ t('androidDeviceOptions.cancel') }}</a-button>
        </a-space>
      </template>
    </a-drawer>

    <a-drawer v-model:visible="detailVisible" :title="t('androidDeviceOptions.detail')" width="520px" :footer="false">
      <a-descriptions :column="1" bordered>
        <a-descriptions-item :label="t('androidDeviceOptions.name')">{{ detailRow?.name }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.serial')">{{ detailRow?.serial }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.logOn')">{{ detailRow?.logOn ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.ignorePopup')">{{ detailRow?.ignorePopup ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.adbServerHost')">{{ detailRow?.adbServerHost }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.adbServerPort')">{{ detailRow?.adbServerPort }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.composite')">{{ detailRow?.composite ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.uia2')">{{ detailRow?.uia2 ? t('androidDeviceOptions.yes') : t('androidDeviceOptions.no') }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.uia2Ip')">{{ detailRow?.uia2Ip }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.uia2Port')">{{ detailRow?.uia2Port }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.uia2ServerPackageName')">{{ detailRow?.uia2ServerPackageName }}</a-descriptions-item>
        <a-descriptions-item :label="t('androidDeviceOptions.uia2ServerTestPackageName')">{{ detailRow?.uia2ServerTestPackageName }}</a-descriptions-item>
      </a-descriptions>
    </a-drawer>
  </a-space>
</template>

<script setup>
import { Message, Modal } from '@arco-design/web-vue'
import { computed, onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import {
  createAndroidDeviceOptions,
  deleteAndroidDeviceOptions,
  getAndroidDeviceOptionsList,
  updateAndroidDeviceOptions
} from '@/services/appBridge'

const { t } = useI18n()

const loading = ref(false)
const rows = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const searchForm = reactive({
  name: ''
})

const drawerVisible = ref(false)
const saving = ref(false)
const drawerMode = ref('create')
const formRef = ref()

const form = reactive({
  ID: undefined,
  name: '',
  serial: '',
  logOn: false,
  ignorePopup: false,
  adbServerHost: '',
  adbServerPort: 5037,
  composite: false,
  uia2: false,
  uia2Ip: '',
  uia2Port: 0,
  uia2ServerPackageName: '',
  uia2ServerTestPackageName: ''
})

const detailVisible = ref(false)
const detailRow = ref(null)

const columns = computed(() => [
  { title: t('androidDeviceOptions.name'), dataIndex: 'name' },
  { title: t('androidDeviceOptions.serial'), dataIndex: 'serial' },
  { title: t('androidDeviceOptions.logOn'), slotName: 'logOn' },
  { title: t('androidDeviceOptions.ignorePopup'), slotName: 'ignorePopup' },
  { title: t('androidDeviceOptions.adbServerHost'), dataIndex: 'adbServerHost' },
  { title: t('androidDeviceOptions.adbServerPort'), dataIndex: 'adbServerPort' },
  { title: t('androidDeviceOptions.composite'), slotName: 'composite' },
  { title: t('androidDeviceOptions.uia2'), slotName: 'uia2' },
  { title: t('androidDeviceOptions.actions'), slotName: 'actions', width: 240 }
])

const drawerTitle = computed(() => {
  return drawerMode.value === 'edit' ? t('androidDeviceOptions.edit') : t('androidDeviceOptions.create')
})

const toNumberOrNull = (v) => {
  if (v === '' || v === null || v === undefined) return null
  const n = Number(v)
  return Number.isFinite(n) ? n : null
}

const loadList = async () => {
  loading.value = true
  try {
    const res = await getAndroidDeviceOptionsList({
      page: page.value,
      pageSize: pageSize.value,
      name: searchForm.name || undefined
    })
    const list = res?.list || res?.List || []
    rows.value = Array.isArray(list) ? list : []
    total.value = Number(res?.total || res?.Total || 0)
  } catch (e) {
    Message.error(e?.message || t('androidDeviceOptions.loadError'))
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
  form.serial = ''
  form.logOn = false
  form.ignorePopup = false
  form.adbServerHost = ''
  form.adbServerPort = 5037
  form.composite = false
  form.uia2 = false
  form.uia2Ip = ''
  form.uia2Port = 0
  form.uia2ServerPackageName = ''
  form.uia2ServerTestPackageName = ''
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
  form.serial = row?.serial || ''
  form.logOn = Boolean(row?.logOn)
  form.ignorePopup = Boolean(row?.ignorePopup)
  form.adbServerHost = row?.adbServerHost || ''
  form.adbServerPort = row?.adbServerPort ?? 5037
  form.composite = Boolean(row?.composite)
  form.uia2 = Boolean(row?.uia2)
  form.uia2Ip = row?.uia2Ip || ''
  form.uia2Port = row?.uia2Port ?? 0
  form.uia2ServerPackageName = row?.uia2ServerPackageName || ''
  form.uia2ServerTestPackageName = row?.uia2ServerTestPackageName || ''
  drawerVisible.value = true
}

const openDetail = (row) => {
  detailRow.value = row || null
  detailVisible.value = true
}

const closeDrawer = () => {
  drawerVisible.value = false
}

const submit = async () => {
  const valid = await formRef.value?.validate?.()
  if (valid) return

  saving.value = true
  try {
    const payload = {
      ID: form.ID,
      name: form.name || undefined,
      serial: form.serial,
      logOn: Boolean(form.logOn),
      ignorePopup: Boolean(form.ignorePopup),
      adbServerHost: form.adbServerHost || undefined,
      adbServerPort: toNumberOrNull(form.adbServerPort),
      composite: Boolean(form.composite),
      uia2: Boolean(form.uia2),
      uia2Ip: form.uia2Ip || undefined,
      uia2Port: toNumberOrNull(form.uia2Port),
      uia2ServerPackageName: form.uia2ServerPackageName || undefined,
      uia2ServerTestPackageName: form.uia2ServerTestPackageName || undefined
    }
    if (drawerMode.value === 'edit') {
      await updateAndroidDeviceOptions(payload)
      Message.success(t('androidDeviceOptions.updateSuccess'))
    } else {
      await createAndroidDeviceOptions(payload)
      Message.success(t('androidDeviceOptions.createSuccess'))
    }
    drawerVisible.value = false
    await loadList()
  } catch (e) {
    Message.error(e?.message || t('androidDeviceOptions.saveError'))
  } finally {
    saving.value = false
  }
}

const confirmDelete = (row) => {
  const id = row?.ID
  if (!id) return
  Modal.confirm({
    title: t('androidDeviceOptions.deleteConfirmTitle'),
    content: t('androidDeviceOptions.deleteConfirmContent'),
    onOk: async () => {
      try {
        await deleteAndroidDeviceOptions(id)
        Message.success(t('androidDeviceOptions.deleteSuccess'))
        await loadList()
      } catch (e) {
        Message.error(e?.message || t('androidDeviceOptions.deleteError'))
      }
    }
  })
}

onMounted(async () => {
  await loadList()
})
</script>
