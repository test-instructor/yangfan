<template>
  <div class="auto-report-detail" v-loading="loading">
    <!-- Header -->
    <div class="flex justify-between items-start mb-6">
      <div>
        <h1 class="text-2xl font-bold text-blue-600 mb-2">{{ detail.name || '报告名称' }}</h1>
        <div class="text-gray-500 dark:text-gray-400 text-sm">
          任务 ID: {{ detail.ID }} | 创建任务时间: {{ formatDate(detail.CreatedAt) }}
        </div>
      </div>
      <div class="flex gap-2">
        <div class="text-sm text-gray-600 dark:text-gray-300 flex items-center">执行模式: <el-tag class="ml-1">{{ detail.run_mode }}</el-tag></div>
        <div class="text-sm text-gray-600 dark:text-gray-300 flex items-center">环境: <el-tag type="info" class="ml-1">{{ detail.env_name || '测试环境' }}</el-tag></div>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-4 gap-4 mb-6">
      <!-- Card 1: Use Cases -->
      <el-card shadow="hover" class="summary-card rounded-lg">
          <div class="flex justify-between items-start">
            <div>
            <div class="text-gray-500 dark:text-gray-400 text-sm mb-1">用例执行概要</div>
            <div class="text-3xl font-bold mb-2">{{ detail.stat?.testcases?.total || 0 }} <span class="text-sm font-normal text-gray-500 dark:text-gray-400">个用例</span></div>
            <div class="flex items-center text-sm">
              <span class="w-3 h-3 rounded-full bg-green-500 mr-1"></span>
              <span class="mr-3">{{ detail.stat?.testcases?.success || 0 }} 成功</span>
              <span class="w-3 h-3 rounded-full bg-red-500 mr-1"></span>
              <span class="mr-3">{{ detail.stat?.testcases?.fail || 0 }} 失败</span>
              <span class="w-3 h-3 rounded-full bg-gray-400 mr-1"></span>
              <span>{{ detail.stat?.testcases?.skip || 0 }} 跳过</span>
            </div>
          </div>
          <div class="p-3 bg-blue-50 dark:bg-blue-900/30 rounded-lg text-blue-500 dark:text-blue-300">
            <el-icon :size="24"><Document /></el-icon>
          </div>
        </div>
      </el-card>

      <!-- Card 2: Interfaces -->
      <el-card shadow="hover" class="summary-card rounded-lg">
        <div class="flex justify-between items-start">
          <div>
             <div class="text-gray-500 dark:text-gray-400 text-sm mb-1">接口执行概要</div>
             <div class="text-3xl font-bold mb-2">{{ apiStats.total }} <span class="text-sm font-normal text-gray-500 dark:text-gray-400">个接口</span></div>
             <div class="flex items-center text-sm">
               <span class="w-3 h-3 rounded-full bg-green-500 mr-1"></span>
               <span class="mr-3">{{ apiStats.success }} 成功</span>
               <span class="w-3 h-3 rounded-full bg-red-500 mr-1"></span>
               <span class="mr-3">{{ apiStats.fail }} 失败</span>
               <span class="w-3 h-3 rounded-full bg-gray-400 mr-1"></span>
               <span>{{ apiStats.skip }} 跳过</span>
             </div>
          </div>
          <div class="p-3 bg-green-50 dark:bg-green-900/30 rounded-lg text-green-500 dark:text-green-300">
             <el-icon :size="24"><Sort /></el-icon>
          </div>
        </div>
      </el-card>

      <!-- Card 3: Time -->
      <el-card shadow="hover" class="summary-card rounded-lg">
         <div class="flex justify-between items-start">
           <div>
             <div class="text-gray-500 dark:text-gray-400 text-sm mb-1">执行时间</div>
             <div class="text-3xl font-bold mb-2">{{ detail.time?.duration ? detail.time.duration.toFixed(2) : 0 }} <span class="text-sm font-normal text-gray-500 dark:text-gray-400">秒</span></div>
             <div class="text-sm text-gray-500 dark:text-gray-400">
               开始时间: {{ formatDate(detail.time?.start_at) }}
             </div>
           </div>
           <div class="p-3 bg-orange-50 dark:bg-orange-900/30 rounded-lg text-orange-500 dark:text-orange-300">
              <el-icon :size="24"><Timer /></el-icon>
           </div>
         </div>
      </el-card>

  <!-- Card 4: Status -->
      <el-card shadow="hover" class="summary-card rounded-lg">
         <div class="flex justify-between items-start">
            <div class="flex-1 min-w-0 mr-4">
              <div class="text-gray-500 dark:text-gray-400 text-sm mb-1">执行状态</div>
              <div class="mb-2">
                <el-tag :type="statusType" effect="light" size="large" class="!text-base px-4 py-1 rounded-full flex items-center w-fit">
                   <el-icon class="mr-1">
                     <CircleCheckFilled v-if="detail.status === 3" />
                     <CircleCloseFilled v-else-if="detail.status === 2" />
                     <Refresh v-else-if="detail.status === 1" />
                     <Clock v-else />
                   </el-icon>
                   {{ statusText }}
                </el-tag>
              </div>
              <div class="text-sm text-gray-500 dark:text-gray-400 truncate w-full" :title="detail.node_name">
                运行节点: {{ detail.node_name || '-' }}
              </div>
              <div v-if="showProgress" class="mt-3 w-full">
                 <div class="flex justify-between text-xs text-gray-500 dark:text-gray-400 mb-1">
                    <span>执行进度({{progressText}}）</span>
                 </div>
                 <el-progress :percentage="progressPercentage" :status="progressStatus" :stroke-width="16" :text-inside="true"
                              striped
                              striped-flow />
              </div>
              <div v-if="!isFinished" class="mt-2">
                 <el-button type="primary" link icon="Refresh" @click="refreshData" :loading="loading">刷新</el-button>
              </div>
            </div>
            <div class="p-3 bg-green-50 dark:bg-green-900/30 rounded-lg text-green-500 dark:text-green-300">
               <el-icon :size="24"><CircleCheck /></el-icon>
            </div>
         </div>
      </el-card>
    </div>

    <!-- Charts -->
    <div class="grid grid-cols-3 gap-4 mb-6">
      <!-- Donut Chart -->
      <el-card shadow="hover" class="col-span-1 rounded-lg">
        <template #header>
          <div class="font-bold">执行结果统计</div>
        </template>
        <div class="flex flex-col items-center justify-center h-64 relative">
           <div ref="resultChartRef" style="width: 100%; height: 100%;"></div>
           <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-center pointer-events-none">
              <!-- Optional center text -->
           </div>
           <div class="flex justify-center gap-4 mt-[-20px]">
              <div class="flex items-center"><span class="w-3 h-3 rounded-full bg-green-500 mr-1"></span>成功</div>
              <div class="flex items-center"><span class="w-3 h-3 rounded-full bg-red-500 mr-1"></span>失败</div>
              <div class="flex items-center"><span class="w-3 h-3 rounded-full bg-gray-400 mr-1"></span>跳过</div>
           </div>
        </div>
      </el-card>

      <!-- Bar Chart -->
      <el-card shadow="hover" class="col-span-2 rounded-lg">
        <template #header>
          <div class="font-bold">接口执行时间分布</div>
        </template>
        <div class="h-64">
           <div ref="timeChartRef" style="width: 100%; height: 100%;"></div>
        </div>
      </el-card>
    </div>

    <!-- Platform Info -->
    <el-card shadow="hover" class="mb-6 rounded-lg">
      <template #header>
        <div class="font-bold">平台信息</div>
      </template>
      <div class="flex flex-wrap gap-8 text-sm text-gray-700 dark:text-gray-300">
        <div class="flex items-center">
          <span class="w-2 h-2 rounded-full bg-blue-500 mr-2"></span>
          Go 版本: {{ platformInfo.go_version || '-' }}
        </div>
        <div class="flex items-center">
          <span class="w-2 h-2 rounded-full bg-blue-500 mr-2"></span>
          HTTPRunner 版本: {{ platformInfo.httprunner_version || '-' }}
        </div>
        <div class="flex items-center">
          <span class="w-2 h-2 rounded-full bg-blue-500 mr-2"></span>
          平台: {{ platformInfo.platform || '-' }}
        </div>
      </div>
    </el-card>

    <!-- Test Case Details -->
    <div class="mb-4">
      <h2 class="text-xl font-bold mb-4">测试用例详情</h2>
      
      <div v-for="(item, index) in filteredDetails" :key="item.ID" class="mb-4">
        <!-- Case Header -->
        <el-card shadow="hover" :body-style="{ padding: '0px' }" class="overflow-hidden rounded-lg">
           <div class="p-4 flex items-center justify-between bg-white dark:bg-slate-800 cursor-pointer hover:bg-gray-50 dark:hover:bg-slate-700/50 transition-colors" @click="toggleExpand(item.ID)">
              <div class="flex items-center gap-3">
                 <span class="font-bold text-lg">用例{{ index + 1 }}</span>
                 <el-tag v-if="item.skip" type="info" effect="light">跳过</el-tag>
                 <el-tag v-else :type="item.success ? 'success' : 'danger'" effect="light">{{ item.success ? '成功' : '失败' }}</el-tag>
                 <span class="font-medium ml-2 text-gray-700 dark:text-gray-200">{{ item.name }}</span>
              </div>
              <div class="flex items-center gap-6 text-gray-500 dark:text-gray-400 text-sm">
                 <div class="flex items-center"><el-icon class="mr-1"><Timer /></el-icon> {{ item.time?.duration?.toFixed(2) }}s</div>
                 <div class="flex items-center"><el-icon class="mr-1"><List /></el-icon> {{ apiCount(item) }}个接口</div>
                 <el-icon :class="{'transform rotate-180': isExpanded(item.ID)}" class="transition-transform duration-300"><ArrowDown /></el-icon>
              </div>
           </div>

           <!-- Expanded Content -->
           <div v-show="isExpanded(item.ID)" class="border-t border-gray-100 dark:border-gray-700 bg-gray-50 dark:bg-slate-900/40 p-4">
              <!-- Config Vars -->
              <div class="bg-blue-50 dark:bg-blue-900/30 rounded p-3 mb-4 text-sm" v-if="hasCaseVars(item)">
                 <div class="font-bold text-blue-700 dark:text-blue-300 mb-2">配置变量</div>
                 <div class="grid grid-cols-2 gap-x-4 gap-y-1">
                    <div v-for="(val, key) in getCaseVars(item)" :key="key" class="flex">
                       <span class="text-gray-600 dark:text-gray-300 w-40 shrink-0 text-right mr-2">{{ key }}:</span>
                       <span class="font-medium truncate text-gray-800 dark:text-gray-100" :title="val">{{ val }}</span>
                    </div>
                 </div>
              </div>

              <!-- Steps -->
              <div class="mb-3">
                 <el-collapse>
                    <el-collapse-item 
                      v-for="group in item.groupedRecords" 
                      :key="group.stepName" 
                      :name="group.stepName"
                    >
                       <template #title>
                          <div class="flex items-center gap-2 w-full">
                             <span class="font-bold text-gray-700 dark:text-gray-200">{{ group.stepName }}</span>
                             <el-tag size="small" effect="plain">{{ group.records.length }}个接口</el-tag>
                          </div>
                       </template>

                       <!-- Interface List -->
                       <div class="bg-white dark:bg-slate-800 rounded border border-gray-200 dark:border-gray-700 divide-y divide-gray-100 dark:divide-gray-700">
                          <div v-for="record in group.records" :key="record.index" class="p-3 flex items-center justify-between hover:bg-gray-50 dark:hover:bg-slate-700/50 transition-colors">
                             <div class="flex items-center gap-3 overflow-hidden flex-1 mr-4">
                                <el-tag :type="getMethodColor(getRequestMethod(record))" class="w-16 text-center font-bold" effect="light">{{ getRequestMethod(record) }}</el-tag>
                                <div class="flex flex-col overflow-hidden">
                                   <div class="font-medium truncate text-gray-800 dark:text-gray-100" :title="getRequestDesc(record)">{{ getRequestDesc(record) }}</div>
                                   <div class="text-xs text-gray-400 dark:text-gray-500 flex items-center gap-1 truncate" :title="getRequestUrl(record)">
                                      <el-icon><Link /></el-icon> {{ getRequestUrl(record) }}
                                   </div>
                                </div>
                             </div>
                             
                             <div class="flex items-center gap-6 shrink-0">
                                <el-tag v-if="record.skip" type="info" size="small" effect="light">跳过</el-tag>
                                <el-tag v-else :type="record.success ? 'success' : 'danger'" size="small" effect="light">{{ record.success ? '成功' : '失败' }}</el-tag>
                                <div class="text-sm text-gray-500 dark:text-gray-400 w-24 text-right flex justify-end items-center"><el-icon class="mr-1"><Timer /></el-icon>{{ record.elapsed_ms }} ms</div>
                                <el-button
                                  v-if="['request', 'api'].includes(record.step_type) || getRequestMethod(record) || record.skip"
                                  link 
                                  type="primary" 
                                  @click.stop="showDetail(record)"
                                >
                                  查看详情 <el-icon class="ml-1"><ArrowRight /></el-icon>
                                </el-button>
                             </div>
                          </div>
                       </div>
                    </el-collapse-item>
                 </el-collapse>
              </div>
           </div>
        </el-card>
      </div>
    </div>

    <!-- Module 4: 接口详情抽屉 -->
    <el-drawer
      v-model="dialogVisible"
      title="接口详情"
      size="80%"
      direction="rtl"
      destroy-on-close
      :with-header="true"
    >
      <div v-if="currentRecord" class="h-full flex flex-col">
        <div class="mb-4">
          <!-- 顶部横向 HTTP 耗时图表 -->
          <div class="flex items-center mb-4">
            <div class="flex-1">
              <div ref="httpstatChartRef" style="width: 100%; height: 100px;"></div>
            </div>
          </div>

          <!-- 基本信息描述 -->
          <el-descriptions border>
            <el-descriptions-item label="接口名称">
              {{ currentRecord.data?.req_resps?.request?.url || '-' }}
            </el-descriptions-item>
            <el-descriptions-item label="Method">
              <el-tag>{{ currentRecord.data?.req_resps?.request?.method || (currentRecord.skip ? 'SKIP' : '-') }}</el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="状态码">
              <el-tag v-if="currentRecord.skip" type="info">SKIPPED</el-tag>
              <el-tag v-else :type="(currentRecord.data?.req_resps?.response?.status_code || 0) < 400 ? 'success' : 'danger'">
                {{ currentRecord.data?.req_resps?.response?.status_code ?? '-' }}
              </el-tag>
            </el-descriptions-item>
            <el-descriptions-item label="耗时">
              {{ currentRecord.elapsed_ms }} ms
            </el-descriptions-item>
          </el-descriptions>
        </div>

        <el-tabs v-model="activeTab" type="border-card" class="flex-1 overflow-hidden flex flex-col">
          <el-tab-pane v-if="currentRecord.skip" label="跳过原因" name="skip" class="h-full overflow-auto">
            <div class="p-4 bg-orange-50 dark:bg-orange-900/20 text-orange-700 dark:text-orange-300 rounded border border-orange-200 dark:border-orange-800/50">
               <div class="font-bold mb-2 flex items-center">
                  <el-icon class="mr-1"><Warning /></el-icon> 跳过原因
               </div>
               <pre class="whitespace-pre-wrap font-mono text-sm">{{ currentRecord.attachments }}</pre>
            </div>
          </el-tab-pane>
          <el-tab-pane v-if="!currentRecord.success && !currentRecord.skip && currentRecord.attachments" label="失败原因" name="error" class="h-full overflow-auto">
            <pre class="code-block">{{ currentRecord.attachments }}</pre>
          </el-tab-pane>
          <el-tab-pane label="请求内容" name="request" class="h-full overflow-auto" v-if="!currentRecord.skip">
            <el-tabs v-model="requestSubTab" type="card" class="h-full">
              <el-tab-pane label="Headers" name="headers" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Request Headers</h4>
                <el-table :data="objectToTableData(currentRecord.data?.req_resps?.request?.headers)" border size="small">
                  <el-table-column prop="key" label="Key" width="220" />
                  <el-table-column prop="value" label="Value" />
                </el-table>
              </el-tab-pane>
              <el-tab-pane label="Body" name="body" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Request Body</h4>
                <pre class="code-block">{{ formatJson(currentRecord.data?.req_resps?.request?.body || currentRecord.data?.req_resps?.request?.json) }}</pre>
              </el-tab-pane>
              <el-tab-pane label="Raw" name="raw" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Raw Request</h4>
                <pre class="code-block">{{ formatJson(currentRecord.data?.req_resps?.request) }}</pre>
              </el-tab-pane>
            </el-tabs>
          </el-tab-pane>

          <el-tab-pane label="响应内容" name="response" class="h-full overflow-auto" v-if="!currentRecord.skip">
            <el-tabs v-model="responseSubTab" type="card" class="h-full">
              <el-tab-pane label="Headers" name="headers" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Response Headers</h4>
                <el-table :data="objectToTableData(currentRecord.data?.req_resps?.response?.headers)" border size="small">
                  <el-table-column prop="key" label="Key" width="220" />
                  <el-table-column prop="value" label="Value" />
                </el-table>
              </el-tab-pane>
              <el-tab-pane label="Body" name="body" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Response Body</h4>
                <pre class="code-block">{{ formatJson(currentRecord.data?.req_resps?.response?.body) }}</pre>
              </el-tab-pane>
              <el-tab-pane label="Raw" name="raw" class="h-full overflow-auto">
                <h4 class="font-bold my-2">Raw Response</h4>
                <pre class="code-block">{{ formatJson(currentRecord.data?.req_resps?.response) }}</pre>
              </el-tab-pane>
            </el-tabs>
          </el-tab-pane>

          <el-tab-pane label="断言信息" name="assert" class="h-full overflow-auto" v-if="!currentRecord.skip">
            <el-table :data="currentRecord.data?.validators" border>
              <el-table-column prop="check" label="Check" />
              <el-table-column prop="assert" label="Assert" />
              <el-table-column prop="expect" label="Expect" />
              <el-table-column prop="check_value" label="Actual" />
              <el-table-column prop="check_result" label="Result">
                <template #default="scope">
                  <el-tag :type="scope.row.check_result === 'pass' ? 'success' : 'danger'">
                    {{ scope.row.check_result }}
                  </el-tag>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>

          <el-tab-pane label="参数提取" name="extract" class="h-full overflow-auto" v-if="!currentRecord.skip">
            <el-table :data="objectToTableData(currentRecord.export_vars)" border>
              <el-table-column prop="key" label="Variable Name" />
              <el-table-column prop="value" label="Value" />
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed, nextTick, watch } from 'vue'
import { useRoute } from 'vue-router'
import { findAutoReport } from '@/api/automation/autoreport'
import { formatDate } from '@/utils/format'
import * as echarts from 'echarts'
import { ElMessage } from 'element-plus'
import { 
  Document, Sort, Timer, CircleCheck, CircleCheckFilled, CircleCloseFilled, 
  List, ArrowDown, ArrowRight, Link, Calendar, Search, Refresh, Clock, Warning
} from '@element-plus/icons-vue'

const route = useRoute()
const detail = ref({})
const searchText = ref('')
const loading = ref(false)
// 抽屉控制
const dialogVisible = ref(false)
const activeTab = ref('request')
const currentRecord = ref(null)
// Expanded keys for custom list
const expandedRowKeys = ref([])

// 内部子 Tab：请求/响应的 header/body/raw
const requestSubTab = ref('headers')
const responseSubTab = ref('headers')

const resultChartRef = ref(null)
const timeChartRef = ref(null)
const httpstatChartRef = ref(null)
const timer = ref(null)

const startTimer = () => {
    stopTimer()
    timer.value = setTimeout(() => {
        loadData(true)
    }, 5000)
}

const stopTimer = () => {
    if (timer.value) {
        clearTimeout(timer.value)
        timer.value = null
    }
}

// 工具方法：解析步骤名称，例如 "步骤1 - GET接口测试1"
const parseStepName = (name) => {
    if (!name) {
        return { step: '', desc: '' }
    }
    const parts = name.split(' - ')
    return {
        step: parts[0] || '',
        desc: parts.slice(1).join(' - ')
    }
}

// 将 records 按步骤前缀分组
const groupRecordsByStep = (records = []) => {
    const groupsMap = {}
    records.forEach((r) => {
        const { step, desc } = parseStepName(r.name)
        const stepName = step || '未命名步骤'
        if (!groupsMap[stepName]) {
            groupsMap[stepName] = {
                stepName,
                records: []
            }
        }
        // 在记录上挂载解析后的描述
        groupsMap[stepName].records.push({
            ...r,
            _stepDesc: desc || r.name // Fallback to name if no dash
        })
    })
    return Object.values(groupsMap)
}

// Load Data
const loadData = async (isAuto = false) => {
    if (!route.params.id) return
    if (!isAuto) loading.value = true
    try {
        const res = await findAutoReport({ ID: route.params.id })
        if (res.code === 0) {
            const data = res.data
            // Pre-process groups
            if (data.details) {
                data.details.forEach(item => {
                    item.groupedRecords = groupRecordsByStep(item.records)
                })
            }
            detail.value = data
            initCharts()

            if (detail.value.status === 0 || detail.value.status === 1) {
                startTimer()
            } else {
                stopTimer()
            }
        }
    } finally {
        if (!isAuto) loading.value = false
    }
}

const refreshData = () => {
    loadData()
    ElMessage.success('刷新成功')
}

onMounted(() => {
    loadData()
})

onUnmounted(() => {
    stopTimer()
})

// Computed
const isFinished = computed(() => {
    const s = detail.value.status
    return s === 2 || s === 3
})

const showProgress = computed(() => {
    return !isFinished.value && detail.value.progress
})

const progressPercentage = computed(() => {
    if (!detail.value.progress) return 0
    const { executed_apis, total_apis } = detail.value.progress
    if (!total_apis) return 0
    return Math.floor((executed_apis / total_apis) * 100)
})

const progressText = computed(() => {
    if (!detail.value.progress) return ''
    const { executed_apis, total_apis } = detail.value.progress
    return `${executed_apis}/${total_apis}`
})

const progressStatus = computed(() => {
    if (progressPercentage.value === 100) return 'success'
    return ''
})

const statusText = computed(() => {
    const s = detail.value.status
    if (s === 1) return '运行中'
    if (s === 2) return '失败'
    if (s === 3) return '成功'
    return '未知'
})

const statusType = computed(() => {
    const s = detail.value.status
    if (s === 1) return 'primary'
    if (s === 2) return 'danger'
    if (s === 3) return 'success'
    return 'info'
})

const apiStats = computed(() => {
    const api = detail.value.stat?.teststepapi || {}
    return {
        total: api.total || 0,
        success: api.success || 0,
        fail: api.fail || 0,
        skip: api.skip || 0
    }
})

const platformInfo = computed(() => {
    if (!detail.value.platform) return {}
    try {
        if (typeof detail.value.platform === 'string') {
            return JSON.parse(detail.value.platform)
        }
        return detail.value.platform
    } catch (e) {
        return {}
    }
})

const filteredDetails = computed(() => {
    if (!detail.value.details) return []
    let list = detail.value.details
    if (searchText.value) {
        const lower = searchText.value.toLowerCase()
        list = list.filter(item => {
             const matchCase = item.name.toLowerCase().includes(lower)
             if (matchCase) return true
             // Search inside records
             if (item.records && item.records.some(r => r.name.toLowerCase().includes(lower))) {
                 return true
             }
             return false
        })
    }
    return list
})

// Logic for custom expanded list
const isExpanded = (id) => expandedRowKeys.value.includes(id)
const toggleExpand = (id) => {
    if (isExpanded(id)) {
        expandedRowKeys.value = expandedRowKeys.value.filter(key => key !== id)
    } else {
        expandedRowKeys.value.push(id)
    }
}

const apiCount = (item) => {
    if (!item.records) return 0
    return item.records.length
}

const getCaseVars = (item) => {
    // Try to find variables in in_out or other places
    if (item.in_out && item.in_out.vars) {
        return item.in_out.vars
    }
    // Fallback if structure is different
    return null
}

const hasCaseVars = (item) => {
    const vars = getCaseVars(item)
    return vars && Object.keys(vars).length > 0
}

const getRequestDesc = (record) => {
    // Use the pre-parsed desc if available
    if (record._stepDesc) return record._stepDesc
    const { desc } = parseStepName(record.name)
    return desc || record.name
}

// Formatters
const formatDuration = (seconds) => {
    if (!seconds) return '-'
    return seconds.toFixed(2) + 's'
}

const formatJson = (data) => {
    try {
        if (typeof data === 'string') return JSON.stringify(JSON.parse(data), null, 2)
        return JSON.stringify(data, null, 2)
    } catch (e) {
        return data
    }
}

const objectToTableData = (obj) => {
    if (!obj) return []
    return Object.keys(obj).map(key => ({
        key,
        value: typeof obj[key] === 'object' ? JSON.stringify(obj[key]) : obj[key]
    }))
}

const getRequest = (record) => record?.data?.req_resps?.request || {}
const getResponse = (record) => record?.data?.req_resps?.response || {}

const getRequestMethod = (record) => getRequest(record).method || ''
const getRequestUrl = (record) => getRequest(record).url || ''

const getMethodColor = (method) => {
    if (!method) return 'info'
    const m = method.toUpperCase()
    if (m === 'GET') return ''
    if (m === 'POST') return 'success'
    if (m === 'PUT') return 'warning'
    if (m === 'DELETE') return 'danger'
    return 'info'
}


// httpstat 数据（用于 ECharts）
const httpstatTableData = computed(() => {
    if (!currentRecord.value || !currentRecord.value.httpstat) return []
    const hs = currentRecord.value.httpstat
    // 固定顺序，便于图例和堆叠展示
    const order = [
        { key: 'DNSLookup', label: 'DNS 解析' },
        { key: 'TCPConnection', label: 'TCP 连接' },
        { key: 'TLSHandshake', label: 'TLS 握手' },
        { key: 'ServerProcessing', label: '服务端处理' },
        { key: 'ContentTransfer', label: '数据传输' }
    ]
    return order.map((item) => ({
        phase: item.label,
        key: item.key,
        duration: hs[item.key] || 0
    }))
})

const renderHttpstatChart = () => {
    if (!httpstatChartRef.value || !currentRecord.value || !currentRecord.value.httpstat) return
    const data = httpstatTableData.value
    if (!data.length) return

    const total = currentRecord.value.httpstat.Total || data.reduce((sum, d) => sum + d.duration, 0)

    let chart = echarts.getInstanceByDom(httpstatChartRef.value)
    if (!chart) {
        chart = echarts.init(httpstatChartRef.value)
    }

    const colors = ['#5470C6', '#91CC75', '#FAC858', '#EE6666', '#73C0DE']

    chart.setOption({
        tooltip: {
            trigger: 'item',
            formatter: (p) => `${p.seriesName}: ${p.value} ms`
        },
        legend: {
            top: 0,
            data: data.map((d) => d.phase)
        },
        grid: {
            left: 40,
            right: 20,
            top: 30,
            bottom: 30
        },
        xAxis: {
            type: 'value',
            name: 'ms',
            min: 0,
            max: total,
            boundaryGap: [0, 0],
            axisLine: { show: false }
        },
        yAxis: {
            type: 'category',
            data: ['耗时'],
            axisLine: { show: false },
            axisTick: { show: false },
            axisLabel: { show: false }
        },
        series: data.map((d, idx) => ({
            name: d.phase,
            type: 'bar',
            stack: 'total',
            barWidth: 18,
            itemStyle: { color: colors[idx % colors.length] },
            data: [d.duration]
        }))
    })
}

// Actions
const showDetail = (record) => {
    currentRecord.value = record
    if (record.skip) {
        activeTab.value = 'skip'
    } else {
        activeTab.value = (!record.success && record.attachments) ? 'error' : 'request'
    }
    requestSubTab.value = 'headers'
    responseSubTab.value = 'headers'
    dialogVisible.value = true
    nextTick(() => {
        renderHttpstatChart()
    })
}

const handleExpandChange = (row, expandedRows) => {
    // Not used in new custom list
}
const handleRowClick = (row) => {
    // Not used in new custom list
}

// Charts
const initCharts = () => {
    nextTick(() => {
        renderResultChart()
        renderTimeChart()
    })
}

const renderResultChart = () => {
    if (!resultChartRef.value) return
    const chart = echarts.init(resultChartRef.value)
    const success = detail.value.stat?.testcases?.success || 0
    const fail = detail.value.stat?.testcases?.fail || 0
    const skip = detail.value.stat?.testcases?.skip || 0
    
    // Donut Chart
    const option = {
        tooltip: {
            trigger: 'item'
        },
        legend: {
             show: false
        },
        series: [
            {
                name: '执行结果',
                type: 'pie',
                radius: ['50%', '70%'],
                avoidLabelOverlap: false,
                label: {
                    show: false,
                    position: 'center'
                },
                emphasis: {
                    label: {
                        show: true,
                        fontSize: '20',
                        fontWeight: 'bold'
                    }
                },
                labelLine: {
                    show: false
                },
                data: [
                    { value: success, name: '成功', itemStyle: { color: '#67C23A' } },
                    { value: fail, name: '失败', itemStyle: { color: '#F56C6C' } },
                    { value: skip, name: '跳过', itemStyle: { color: '#909399' } }
                ]
            }
        ]
    }
    chart.setOption(option)
}

const renderTimeChart = () => {
    if (!timeChartRef.value) return
    const chart = echarts.init(timeChartRef.value)
    
    // Prepare data: get all steps from all cases
    let steps = []
    if (detail.value.details) {
        detail.value.details.forEach(d => {
             if (d.records) {
                 d.records.forEach(r => {
                     // Only include requests
                     if (['request', 'api'].includes(r.step_type) || getRequestMethod(r)) {
                         steps.push({
                             name: r.name,
                             duration: r.elapsed_ms,
                             stepDesc: r._stepDesc || r.name
                         })
                     }
                 })
             }
        })
    }
    
    // Limit to top 20 or similar to avoid overcrowding? Or just show all?
    // The image shows a bar chart with interface names.
    // If too many, maybe scrollable or top N. Let's show first 10-15 for now.
    const chartData = steps.slice(0, 20) 
    
    const option = {
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                type: 'shadow'
            }
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: [
            {
                type: 'category',
                data: chartData.map(s => s.stepDesc.length > 10 ? s.stepDesc.substring(0,10)+'...' : s.stepDesc),
                axisTick: {
                    alignWithLabel: true
                },
                axisLabel: {
                    rotate: 45
                }
            }
        ],
        yAxis: [
            {
                type: 'value',
                name: '执行时间 (ms)'
            }
        ],
        series: [
            {
                name: '耗时',
                type: 'bar',
                barWidth: '60%',
                data: chartData.map((s, index) => ({
                    value: s.duration,
                    itemStyle: {
                        color: index % 2 === 0 ? '#5470C6' : '#91CC75' // Alternate colors or single color
                    }
                }))
            }
        ]
    }
    chart.setOption(option)
    
    // Handle resize
    window.addEventListener('resize', () => {
        chart.resize()
    })
}

</script>

<style scoped>
.auto-report-detail {
    padding: 20px;
}
.code-block {
    background-color: #f4f4f5;
    padding: 10px;
    border-radius: 4px;
    overflow-x: auto;
    font-family: monospace;
    white-space: pre-wrap;
}
</style>
