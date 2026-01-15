<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" @keyup.enter="onSubmit">
        <el-form-item label="关键字" prop="keyword">
          <el-input v-model="searchInfo.keyword" placeholder="通道名称" clearable />
        </el-form-item>
        <el-form-item label="平台" prop="provider">
          <el-select v-model="searchInfo.provider" placeholder="全部" clearable style="width: 140px">
            <el-option label="飞书" value="feishu" />
            <el-option label="钉钉" value="dingtalk" />
            <el-option label="企业微信" value="wecom" />
          </el-select>
        </el-form-item>
        <el-form-item label="启用" prop="enabled">
          <el-select v-model="searchInfo.enabled" placeholder="全部" clearable style="width: 140px">
            <el-option label="启用" :value="true" />
            <el-option label="禁用" :value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="平台" prop="provider" width="120">
          <template #default="scope">{{ providerLabel(scope.row.provider) }}</template>
        </el-table-column>
        <el-table-column label="通道名称" prop="name" min-width="160" />
        <el-table-column label="发送规则" prop="send_rule" width="120">
          <template #default="scope">{{ ruleLabel(scope.row.send_rule) }}</template>
        </el-table-column>
        <el-table-column label="启用" prop="enabled" width="90">
          <template #default="scope">
            <el-tag :type="scope.row.enabled ? 'success' : 'info'">{{ scope.row.enabled ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Webhook" prop="webhook_url" min-width="260" show-overflow-tooltip />
        <el-table-column sortable label="更新时间" prop="UpdatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.UpdatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="openDialog(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="平台" prop="provider">
          <el-select v-model="formData.provider" placeholder="请选择平台" style="width: 100%">
            <el-option label="飞书" value="feishu" />
            <el-option label="钉钉" value="dingtalk" />
            <el-option label="企业微信" value="wecom" />
          </el-select>
        </el-form-item>

        <el-form-item label="通道名称" prop="name">
          <el-input v-model="formData.name" placeholder="例如：项目A" />
        </el-form-item>

        <el-form-item label="启用" prop="enabled">
          <el-switch v-model="formData.enabled" />
        </el-form-item>

        <el-form-item label="发送规则" prop="send_rule">
          <el-select v-model="formData.send_rule" placeholder="请选择发送规则" style="width: 100%">
            <el-option label="成功" value="success" />
            <el-option label="失败" value="fail" />
            <el-option label="始终" value="always" />
          </el-select>
        </el-form-item>

        <el-form-item label="Webhook URL" prop="webhook_url">
          <el-input v-model="formData.webhook_url" placeholder="机器人Webhook地址" />
        </el-form-item>

        <el-form-item label="Webhook Secret" prop="webhook_secret">
          <el-input v-model="formData.webhook_secret" placeholder="可选：飞书/钉钉签名密钥" show-password />
        </el-form-item>

        <el-form-item v-if="formData.provider !== 'feishu'" label="成功模板" prop="template_success">
          <el-input v-model="formData.template_success" type="textarea" :rows="5" placeholder="支持变量：{{title}} {{status}} {{env}} {{success}} {{fail}} {{time}} {{detail}}" />
        </el-form-item>

        <el-form-item v-if="formData.provider !== 'feishu'" label="失败模板" prop="template_fail">
          <el-input v-model="formData.template_fail" type="textarea" :rows="5" placeholder="支持变量：{{title}} {{status}} {{env}} {{success}} {{fail}} {{time}} {{detail}}" />
        </el-form-item>

        <el-form-item label="Web Base URL" prop="web_base_url">
          <el-input v-model="formData.web_base_url" placeholder="例如：https://test.example.com" />
        </el-form-item>

        <el-form-item label="扩展配置" prop="extra">
          <JsonEditor :heights="280" :jsons="formData.extra || {}" json-type="reportNotifyExtra" @jsonData="onExtraChange" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from '@/pinia'
import { formatDate } from '@/utils/format'
import JsonEditor from '@/components/platform/jsonEdit/index.vue'

import {
  createReportNotifyChannel,
  deleteReportNotifyChannel,
  deleteReportNotifyChannelByIds,
  updateReportNotifyChannel,
  getReportNotifyChannelList,
} from '@/api/projectmgr/reportNotify'

defineOptions({
  name: 'reportNotify'
})

const appStore = useAppStore()

const elSearchFormRef = ref()
const elFormRef = ref()
const multipleTable = ref()

const tableData = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const multipleSelection = ref([])

const searchInfo = reactive({
  keyword: '',
  provider: '',
  enabled: undefined,
})

const dialogFormVisible = ref(false)
const btnLoading = ref(false)
const type = ref('create')

const formData = reactive({
  ID: 0,
  provider: 'feishu',
  name: '',
  enabled: true,
  send_rule: 'always',
  webhook_url: '',
  webhook_secret: '',
  template_success: '',
  template_fail: '',
  web_base_url: '',
  extra: {},
})

const rule = reactive({
  provider: [{ required: true, message: '请选择平台', trigger: 'change' }],
  name: [{ required: true, message: '请输入通道名称', trigger: 'blur' }],
  send_rule: [{ required: true, message: '请选择发送规则', trigger: 'change' }],
  webhook_url: [{ required: true, message: '请输入Webhook URL', trigger: 'blur' }],
})

const providerLabel = (p) => {
  if (p === 'feishu') return '飞书'
  if (p === 'dingtalk') return '钉钉'
  if (p === 'wecom') return '企业微信'
  return p || '-'
}

const ruleLabel = (r) => {
  if (r === 'success') return '成功'
  if (r === 'fail') return '失败'
  if (r === 'always') return '始终'
  return r || '-'
}

const onExtraChange = (payload) => {
  if (payload?.isValid) {
    formData.extra = payload.data || {}
  }
}

const resetFormData = () => {
  formData.ID = 0
  formData.provider = 'feishu'
  formData.name = ''
  formData.enabled = true
  formData.send_rule = 'always'
  formData.webhook_url = ''
  formData.webhook_secret = ''
  formData.template_success = ''
  formData.template_fail = ''
  formData.web_base_url = ''
  formData.extra = {}
}

const openDialog = (row) => {
  resetFormData()
  if (row && row.ID) {
    type.value = 'update'
    Object.assign(formData, {
      ID: row.ID,
      provider: row.provider,
      name: row.name,
      enabled: row.enabled,
      send_rule: row.send_rule,
      webhook_url: row.webhook_url,
      webhook_secret: row.webhook_secret,
      template_success: row.template_success,
      template_fail: row.template_fail,
      web_base_url: row.web_base_url,
      extra: row.extra || {},
    })
  } else {
    type.value = 'create'
  }
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
}

const enterDialog = async () => {
  await elFormRef.value?.validate()
  btnLoading.value = true
  try {
    const payload = { ...formData }
    if (type.value === 'create') {
      const res = await createReportNotifyChannel(payload)
      if (res.code === 0) {
        ElMessage.success('创建成功')
        dialogFormVisible.value = false
        await getTableData()
      } else {
        ElMessage.error(res.msg || '创建失败')
      }
    } else {
      const res = await updateReportNotifyChannel(payload)
      if (res.code === 0) {
        ElMessage.success('更新成功')
        dialogFormVisible.value = false
        await getTableData()
      } else {
        ElMessage.error(res.msg || '更新失败')
      }
    }
  } finally {
    btnLoading.value = false
  }
}

const getTableData = async () => {
  const res = await getReportNotifyChannelList({
    page: page.value,
    pageSize: pageSize.value,
    keyword: searchInfo.keyword,
    provider: searchInfo.provider || undefined,
    enabled: searchInfo.enabled,
  })
  if (res.code === 0) {
    tableData.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

const onSubmit = async () => {
  page.value = 1
  await getTableData()
}

const onReset = async () => {
  searchInfo.keyword = ''
  searchInfo.provider = ''
  searchInfo.enabled = undefined
  page.value = 1
  await getTableData()
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const deleteRow = async (row) => {
  await ElMessageBox.confirm('确定删除该通道？', '提示', { type: 'warning' })
  const res = await deleteReportNotifyChannel({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    await getTableData()
  } else {
    ElMessage.error(res.msg || '删除失败')
  }
}

const onDelete = async () => {
  await ElMessageBox.confirm('确定批量删除选中通道？', '提示', { type: 'warning' })
  const ids = multipleSelection.value.map((i) => i.ID)
  const res = await deleteReportNotifyChannelByIds({ 'IDs[]': ids })
  if (res.code === 0) {
    ElMessage.success('批量删除成功')
    await getTableData()
  } else {
    ElMessage.error(res.msg || '批量删除失败')
  }
}

const handleSizeChange = async (val) => {
  pageSize.value = val
  await getTableData()
}

const handleCurrentChange = async (val) => {
  page.value = val
  await getTableData()
}

onMounted(async () => {
  await getTableData()
})
</script>

