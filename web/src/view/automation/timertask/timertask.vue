<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
               @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
          </template>

          <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>

        <el-form-item label="运行配置" prop="configName">
          <el-input v-model="searchInfo.configName" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="环境名称" prop="envName">
          <el-input v-model="searchInfo.envName" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="消息名称" prop="messageName">
          <el-input v-model="searchInfo.messageName" placeholder="搜索条件" />
        </el-form-item>


        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开
          </el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">
          删除
        </el-button>

        <el-button type="primary" style="margin-left: 10px;" @click="openTagDialog">标签管理</el-button>
        <el-button type="primary" style="margin-left: 10px;" @click="openCIDialog">CI调用</el-button>

      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="任务名称" prop="name" width="120" />

        <el-table-column align="left" label="运行时间" prop="runTime" width="120" />

        <el-table-column align="left" label="下次运行时间" prop="nextRunTime" width="180">
          <template #default="scope">{{ formatDate(scope.row.nextRunTime) }}</template>
        </el-table-column>
        <el-table-column align="left" label="运行状态" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>
        <el-table-column align="left" label="运行次数" prop="runNumber" width="120" />

        <el-table-column align="left" label="运行节点" prop="runnerNodeName" width="180" />

        <el-table-column align="left" label="运行配置" prop="configName" width="120" />

        <el-table-column label="标签" prop="tag" width="200">
          <template #default="scope">
            <span>{{ mapTagIdsToNames(scope.row.tag).join(', ') }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="环境名称" prop="envName" width="120" />

        <el-table-column align="left" label="通知通道" prop="messageName" width="140" />

        <el-table-column align="left" label="发送消息" prop="notifyEnabled" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.notifyEnabled) }}</template>
        </el-table-column>

        <el-table-column align="left" label="发送规则" prop="notifyRule" width="120" />

        <el-table-column align="left" label="备注" prop="describe" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="320">
          <template #default="scope">
            <div style="display: flex; align-items: center;">
              <Runner case_type="task" :case_id="scope.row.ID" style="margin-right: 10px;" />
              <el-button type="primary" link class="table-button" @click="openCaseDetail(scope.row)">
                <el-icon style="margin-right: 5px">
                  <Document />
                </el-icon>
                用例详情
              </el-button>
              <el-button type="primary" link icon="edit" class="table-button" @click="updateTimerTaskFunc(scope.row)">
                编辑
              </el-button>
              <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </div>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false"
               :before-close="closeDialog">
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
        <el-form-item label="任务名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item label="运行时间:" prop="runTime">
          <Cron v-model="formData.runTime" />
        </el-form-item>
        <el-form-item label="运行状态:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="运行节点:" prop="runnerNodeName">
          <el-select v-model="formData.runnerNodeName" filterable clearable placeholder="选择节点" style="width: 100%">
            <el-option label="不指定" value="" />
            <el-option
              v-for="opt in runnerNodeOptions"
              :key="opt.nodeName"
              :label="opt.alias ? `${opt.alias}(${opt.nodeName})` : opt.nodeName"
              :value="opt.nodeName"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="运行次数:" prop="runNumber">
          <el-input v-model.number="formData.runNumber" :clearable="false" placeholder="请输入运行次数" />
        </el-form-item>
        <el-form-item label="运行配置:" prop="configName">
          <RunConfig v-model="formData.configID" @change="handleConfigSelect" />
        </el-form-item>
        <el-form-item label="标签:" prop="tag">
          <el-select v-model="formData.tag" multiple filterable placeholder="选择标签">
            <el-option
              v-for="opt in tagOptions"
              :key="opt.ID || opt.id"
              :label="opt.name"
              :value="opt.ID || opt.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="环境名称:" prop="envName">
          <Env v-model="formData.envID" @change="handleEnvChange" />
        </el-form-item>
        <el-form-item label="发送消息:" prop="notifyEnabled">
          <el-switch v-model="formData.notifyEnabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="发送规则:" prop="notifyRule">
          <el-select v-model="formData.notifyRule" placeholder="请选择发送规则" style="width: 100%"
                     :disabled="!formData.notifyEnabled">
            <el-option label="总是发送" value="always" />
            <el-option label="仅成功发送" value="success" />
            <el-option label="仅失败发送" value="fail" />
          </el-select>
        </el-form-item>
        <el-form-item label="通知通道:" prop="messageID">
          <el-select v-model="formData.messageID" filterable clearable placeholder="选择通道" style="width: 100%"
                     :disabled="!formData.notifyEnabled" @change="handleNotifyChannelChange">
            <el-option
              v-for="opt in notifyChannelOptions"
              :key="opt.ID"
              :label="`${providerLabel(opt.provider)}-${opt.name}`"
              :value="opt.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:" prop="describe">
          <el-input v-model="formData.describe" :clearable="false" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
    </el-drawer>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true"
               :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="任务名称">
          {{ detailForm.name }}
        </el-descriptions-item>
        <el-descriptions-item label="运行时间">
          {{ detailForm.runTime }}
        </el-descriptions-item>
        <el-descriptions-item label="下次运行时间">
          {{ formatDate(detailForm.nextRunTime) }}
        </el-descriptions-item>
        <el-descriptions-item label="运行状态">
          {{ formatBoolean(detailForm.status) }}
        </el-descriptions-item>
        <el-descriptions-item label="运行次数">
          {{ detailForm.runNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="运行节点">
          {{ detailForm.runnerNodeName || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="运行配置">
          {{ detailForm.configName }}
        </el-descriptions-item>
        <el-descriptions-item label="标签">
          {{ mapTagIdsToNames(detailForm.tag).join(', ') }}
        </el-descriptions-item>
        <el-descriptions-item label="环境名称">
          {{ detailForm.envName }}
        </el-descriptions-item>
        <el-descriptions-item label="通知通道">
          {{ detailForm.messageName }}
        </el-descriptions-item>
        <el-descriptions-item label="发送消息">
          {{ formatBoolean(detailForm.notifyEnabled) }}
        </el-descriptions-item>
        <el-descriptions-item label="发送规则">
          {{ detailForm.notifyRule }}
        </el-descriptions-item>
        <el-descriptions-item label="备注">
          {{ detailForm.describe }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <TagManager v-model="tagDialogVisible" @changed="onTagChanged" />

    <el-dialog
      v-model="ciDialogVisible"
      title="CI调用"
      width="760px"
      append-to-body
      destroy-on-close
    >
      <el-form label-position="right" label-width="110px">
        <el-form-item label="任务标签">
          <el-select v-model="ciForm.tagId" filterable placeholder="选择标签" style="width: 100%">
            <el-option
              v-for="opt in tagOptions"
              :key="opt.ID || opt.id"
              :label="opt.name"
              :value="opt.ID || opt.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="运行环境">
          <Env v-model="ciForm.envId" @change="handleCIEnvChange" style="width: 100%" />
        </el-form-item>
        <el-form-item label="运行节点">
          <el-select v-model="ciForm.nodeName" filterable clearable placeholder="选择节点" style="width: 100%">
            <el-option label="不指定" value="" />
            <el-option
              v-for="opt in ciRunnerNodeOptions"
              :key="opt.nodeName"
              :label="opt.alias ? `${opt.alias}(${opt.nodeName})` : opt.nodeName"
              :value="opt.nodeName"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="发送消息">
          <el-switch v-model="ciForm.notifyEnabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
                     inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="发送规则">
          <el-select v-model="ciForm.notifyRule" placeholder="请选择发送规则" style="width: 100%"
                     :disabled="!ciForm.notifyEnabled">
            <el-option label="总是发送" value="always" />
            <el-option label="仅成功发送" value="success" />
            <el-option label="仅失败发送" value="fail" />
          </el-select>
        </el-form-item>
        <el-form-item label="通知通道">
          <el-select v-model="ciForm.notifyChannelIds" multiple filterable clearable placeholder="选择通道" style="width: 100%"
                     :disabled="!ciForm.notifyEnabled">
            <el-option
              v-for="opt in notifyChannelOptions"
              :key="opt.ID"
              :label="`${providerLabel(opt.provider)}-${opt.name}`"
              :value="opt.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="返回方式">
          <el-radio-group v-model="ciForm.responseMode">
            <el-radio-button label="sync">同步返回</el-radio-button>
            <el-radio-button label="callback">回调</el-radio-button>
            <el-radio-button label="webhook">Webhook</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="ciForm.responseMode === 'callback'" label="回调URL">
          <el-input v-model="ciForm.callbackUrl" placeholder="https://your-ci/callback" />
        </el-form-item>
        <el-form-item v-if="ciForm.responseMode === 'webhook'" label="消息类型">
          <el-input :model-value="selectedWebhookProviderLabel" readonly />
        </el-form-item>
        <el-form-item v-if="ciForm.responseMode === 'webhook'" label="通道名称(发送规则)">
          <el-select
            v-model="ciForm.webhookChannelId"
            filterable
            clearable
            placeholder="选择通道"
            style="width: 100%"
          >
            <el-option
              v-for="opt in webhookChannelOptions"
              :key="opt.ID"
              :label="`${opt.name}(${sendRuleLabel(opt.send_rule)})`"
              :value="opt.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="CI鉴权标识">
          <el-input v-model="ciForm.ciUUID" placeholder="" />
        </el-form-item>
        <el-form-item label="CI鉴权密钥">
          <el-input v-model="ciForm.ciSecret" placeholder="" />
        </el-form-item>
      </el-form>

      <el-tabs v-model="ciPreviewTab">
        <el-tab-pane label="GET" name="get">
          <el-input :model-value="ciGetURL" type="textarea" :rows="4" readonly />
          <div style="margin-top: 10px; text-align: right;">
            <el-button type="primary" @click="copyToClipboard(ciGetURL)">复制</el-button>
          </div>
        </el-tab-pane>
        <el-tab-pane label="POST(JSON)" name="post">
          <el-input :model-value="ciPostCurl" type="textarea" :rows="8" readonly />
          <div style="margin-top: 10px; text-align: right;">
            <el-button type="primary" @click="copyToClipboard(ciPostCurl)">复制</el-button>
          </div>
        </el-tab-pane>
      </el-tabs>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="ciDialogVisible = false">关闭</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer destroy-on-close size="1200px" v-model="caseDetailVisible" :show-close="true" :before-close="closeCaseDetail" title="任务-用例详情">
      <TaskCaseDetail :taskID="currentTask.ID" :taskName="currentTask.name || ''" />
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createTimerTask,
    deleteTimerTask,
    deleteTimerTaskByIds,
    updateTimerTask,
    findTimerTask,
    getTimerTaskList
  } from '@/api/automation/timertask'
  import { getTagList } from '@/api/automation/tag'
  import { getReportNotifyChannelList } from '@/api/projectmgr/reportNotify'
  import { getRunnerNodeList } from '@/api/platform/runnernode'
  import TagManager from '@/components/automation/TagManager.vue'
  import TaskCaseDetail from './components/TaskCaseDetail.vue'
  
  import Cron from '@/components/cron/index.vue'
  import RunConfig from '@/components/platform/runConfig.vue'
  import Env from '@/components/platform/env.vue'
  import Runner from '@/components/platform/Runner.vue'

  // 全量引入格式化工具 请按需保留
  import {
    getDictFunc,
    formatDate,
    formatBoolean,
    filterDict,
    filterDataSource,
    returnArrImg,
    onDownloadFile
  } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive, computed, onMounted } from 'vue'
  import { useAppStore } from '@/pinia'
  import { useUserStore } from '@/pinia/modules/user'


  defineOptions({
    name: 'TimerTask'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()
  const userStore = useUserStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    runTime: '',
    nextRunTime: new Date(),
    status: false,
    runNumber: 0,
    configName: '',
    configID: null,
    tag: [],
    envName: '',
    envID: null,
    messageName: '',
    messageID: null,
    notifyEnabled: false,
    notifyRule: 'always',
    describe: '',
    runnerNodeName: ''
  })


  // 验证规则
  const rule = reactive({})

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      page.value = 1
      if (searchInfo.value.status === '') {
        searchInfo.value.status = null
      }
      getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getTimerTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  const tagOptions = ref([])
  const tagNameMap = ref({})
  const notifyChannelOptions = ref([])
  const runnerNodeOptions = ref([])
  const ciRunnerNodeOptions = ref([])

  const loadTagOptions = async () => {
    const res = await getTagList({ page: 1, pageSize: 1000 })
    if (res.code === 0) {
      const list = res.data.list || []
      tagOptions.value = list
      const map = {}
      list.forEach(item => { map[item.ID || item.id] = item.name })
      tagNameMap.value = map
    }
  }

  const loadNotifyChannelOptions = async () => {
    const res = await getReportNotifyChannelList({ page: 1, pageSize: 1000 })
    if (res.code === 0) {
      notifyChannelOptions.value = res.data.list || []
    }
  }

  const loadRunnerNodeOptions = async () => {
    const res = await getRunnerNodeList({ page: 1, pageSize: 1000, runContents: ['timer', 'all'] })
    if (res.code === 0) {
      const list = res.data.list || []
      runnerNodeOptions.value = list.filter((x) => x.nodeName)
    }
  }

  const loadCIRunnerNodeOptions = async () => {
    const res = await getRunnerNodeList({ page: 1, pageSize: 1000, runContents: ['runner', 'all'] })
    if (res.code === 0) {
      const list = res.data.list || []
      ciRunnerNodeOptions.value = list.filter((x) => x.nodeName)
    }
  }

  const providerLabel = (p) => {
    if (p === 'feishu') return '飞书'
    if (p === 'dingtalk') return '钉钉'
    if (p === 'wecom') return '企业微信'
    return p || ''
  }

  const handleNotifyChannelChange = (id) => {
    if (!id) {
      formData.value.messageName = ''
      formData.value.messageID = null
      return
    }
    const opt = notifyChannelOptions.value.find((x) => x.ID === id)
    if (opt) {
      formData.value.messageName = opt.name
    }
  }

  const normalizeTagArray = (raw) => {
    if (!raw) return []
    if (Array.isArray(raw)) {
      const ids = []
      raw.forEach(it => {
        if (typeof it === 'number') ids.push(it)
        else if (typeof it === 'string') {
          const entry = Object.entries(tagNameMap.value).find(([, name]) => name === it)
          if (entry) ids.push(Number(entry[0]))
        } else if (it && typeof it === 'object') {
          if (it.ID) ids.push(it.ID)
          else if (it.id) ids.push(it.id)
        }
      })
      return ids
    }
    return []
  }

  const mapTagIdsToNames = (raw) => {
    const ids = normalizeTagArray(raw)
    return ids.map(id => tagNameMap.value[id]).filter(Boolean)
  }

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {
    await loadTagOptions()
    await loadNotifyChannelOptions()
  }

  onMounted(() => {
    getTableData()
    loadRunnerNodeOptions()
    setOptions()
  })

  const ciDialogVisible = ref(false)
  const ciPreviewTab = ref('get')
  const ciForm = reactive({
    tagId: null,
    envId: null,
    nodeName: '',
    notifyEnabled: false,
    notifyRule: 'always',
    notifyChannelIds: [],
    responseMode: 'sync',
    callbackUrl: '',
    ciUUID: '',
    ciSecret: '',
    webhookChannelId: null,
    webhookType: 'feishu',
    webhookUrl: '',
    webhookSecret: ''
  })

  const ciEndpoint = computed(() => {
    const basePath = import.meta.env.VITE_BASE_PATH || window.location.origin
    const baseApi = import.meta.env.VITE_BASE_API || ''
    return `${basePath}${baseApi}/open/runner/run`
  })

  const ciProjectId = computed(() => userStore.userInfo?.projectId || 0)

  const ciGetURL = computed(() => {
    const params = new URLSearchParams()
    params.set('projectId', String(ciProjectId.value))
    params.set('uuid', ciForm.ciUUID || '')
    params.set('secret', ciForm.ciSecret || '')
    params.set('case_type', 'tag')
    params.set('case_id', ciForm.tagId ? String(ciForm.tagId) : '')
    params.set('env_id', ciForm.envId ? String(ciForm.envId) : '')
    if (ciForm.nodeName) {
      params.set('node_name', ciForm.nodeName)
    }
    params.set('notify_enabled', String(!!ciForm.notifyEnabled))
    params.set('notify_rule', ciForm.notifyRule || '')
    if (Array.isArray(ciForm.notifyChannelIds)) {
      ciForm.notifyChannelIds.forEach((id) => {
        params.append('notify_channel_ids', String(id))
      })
    }
    params.set('response_mode', ciForm.responseMode)
    if (ciForm.responseMode === 'callback') {
      params.set('callback_url', ciForm.callbackUrl || '')
    }
    if (ciForm.responseMode === 'webhook') {
      params.set('webhook_channel_id', ciForm.webhookChannelId ? String(ciForm.webhookChannelId) : '')
    }
    return `${ciEndpoint.value}?${params.toString()}`
  })

  const ciPostCurl = computed(() => {
    const payload = {
      projectId: ciProjectId.value,
      uuid: ciForm.ciUUID || '',
      secret: ciForm.ciSecret || '',
      case_type: 'tag',
      case_id: ciForm.tagId || 0,
      env_id: ciForm.envId || 0,
      node_name: ciForm.nodeName || '',
      notify_enabled: !!ciForm.notifyEnabled,
      notify_rule: ciForm.notifyRule || '',
      notify_channel_ids: Array.isArray(ciForm.notifyChannelIds) ? ciForm.notifyChannelIds : [],
      response_mode: ciForm.responseMode,
      callback_url: ciForm.responseMode === 'callback' ? (ciForm.callbackUrl || '') : '',
      webhook_channel_id: ciForm.responseMode === 'webhook' ? (ciForm.webhookChannelId || 0) : 0
    }
    return `curl -X POST '${ciEndpoint.value}' \\\n  -H 'Content-Type: application/json' \\\n  -d '${JSON.stringify(payload)}'`
  })

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text)
      .then(() => ElMessage({ type: 'success', message: '已复制' }))
      .catch(() => ElMessage({ type: 'error', message: '复制失败' }))
  }

  const openCIDialog = async () => {
    await loadTagOptions()
    await loadNotifyChannelOptions()
    await loadCIRunnerNodeOptions()
    ciDialogVisible.value = true
    ciPreviewTab.value = 'get'
  }

  const handleCIEnvChange = (env) => {
    if (env && typeof env === 'object') {
      ciForm.envId = env.ID || env.id || ciForm.envId
    }
  }

  const sendRuleLabel = (rule) => {
    if (rule === 'always') return '总是发送'
    if (rule === 'success') return '仅成功发送'
    if (rule === 'fail') return '仅失败发送'
    return rule || ''
  }

  const webhookChannelOptions = computed(() => {
    const supported = new Set(['feishu', 'wecom', 'dingtalk', 'custom'])
    return (notifyChannelOptions.value || []).filter((x) => {
      if (!x) return false
      if (x.enabled === false) return false
      const provider = (x.provider || '').toLowerCase()
      if (!supported.has(provider)) return false
      if (!x.webhook_url) return false
      return true
    })
  })

  const selectedWebhookChannel = computed(() => {
    if (!ciForm.webhookChannelId) return null
    return webhookChannelOptions.value.find((x) => x.ID === ciForm.webhookChannelId) || null
  })

  const selectedWebhookProviderLabel = computed(() => {
    const opt = selectedWebhookChannel.value
    return opt ? providerLabel(opt.provider) : ''
  })


  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteTimerTaskFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
      const res = await deleteTimerTaskByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateTimerTaskFunc = async (row) => {
    const res = await findTimerTask({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      formData.value.tag = normalizeTagArray(res.data.tag)
      if (!formData.value.runnerNodeName) formData.value.runnerNodeName = ''
      if (formData.value.notifyEnabled == null) formData.value.notifyEnabled = false
      if (!formData.value.notifyRule) formData.value.notifyRule = 'always'
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteTimerTaskFunc = async (row) => {
    const res = await deleteTimerTask({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  const handleConfigSelect = (config) => {
    if (config) {
      formData.value.configName = config.name
    } else {
      formData.value.configName = null
    }
  }

  const handleEnvChange = (env) => {
    if (env) {
      formData.value.envName = env.name
    } else {
      formData.value.envName = null
    }
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      name: '',
      runTime: '',
      nextRunTime: new Date(),
      status: false,
      runNumber: 0,
      configName: '',
      configID: null,
      tag: [],
      envName: '',
      envID: null,
      messageName: '',
      messageID: null,
      notifyEnabled: false,
      notifyRule: 'always',
      describe: '',
      runnerNodeName: ''
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createTimerTask(formData.value)
          break
        case 'update':
          res = await updateTimerTask(formData.value)
          break
        default:
          res = await createTimerTask(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
    })
  }

  const detailForm = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)


  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }


  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findTimerTask({ ID: row.ID })
    if (res.code === 0) {
      detailForm.value = res.data
      openDetailShow()
    }
  }


  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
  }

  const openTagDialog = () => {
    tagDialogVisible.value = true
  }

  const tagDialogVisible = ref(false)
  const onTagChanged = () => { loadTagOptions() }

  const caseDetailVisible = ref(false)
  const currentTask = ref({})
  const openCaseDetail = async (row) => {
    const res = await findTimerTask({ ID: row.ID })
    if (res.code === 0) {
      currentTask.value = res.data
      caseDetailVisible.value = true
    }
  }
  const closeCaseDetail = () => {
    caseDetailVisible.value = false
    currentTask.value = {}
  }


</script>

<style>

</style>
