<template>
  <el-drawer
    v-model="visible"
    title="接口详情"
    size="80%"
    direction="rtl"
    destroy-on-close
    :with-header="true"
    @opened="handleOpened"
  >
    <div v-if="currentRecord" class="h-full flex flex-col">
      <div class="mb-4">
        <!-- 顶部横向 HTTP 耗时图表 -->
        <div class="flex items-center mb-4">
          <div class="flex-1">
            <div ref="httpstatChartRef" style="width: 100%; height: 300px;"></div>
          </div>
        </div>

        <!-- 基本信息描述 -->
        <el-descriptions border>
          <el-descriptions-item label="接口名称">
            {{ currentRecord.data?.req_resps?.request?.url || '-' }}
          </el-descriptions-item>
          <el-descriptions-item label="Method">
            <el-tag>{{ currentRecord.data?.req_resps?.request?.method || '-' }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="状态码">
            <el-tag :type="(currentRecord.data?.req_resps?.response?.status_code || 0) < 400 ? 'success' : 'danger'">
              {{ currentRecord.data?.req_resps?.response?.status_code ?? '-' }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="耗时">
            {{ currentRecord.elapsed_ms }} ms
          </el-descriptions-item>
        </el-descriptions>
      </div>

      <el-tabs v-model="activeTab" type="border-card" class="flex-1 overflow-hidden flex flex-col">
        <el-tab-pane label="请求内容" name="request" class="h-full overflow-auto">
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

        <el-tab-pane label="响应内容" name="response" class="h-full overflow-auto">
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

        <el-tab-pane label="断言信息" name="assert" class="h-full overflow-auto">
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

        <el-tab-pane label="参数提取" name="extract" class="h-full overflow-auto">
          <el-table :data="objectToTableData(currentRecord.export_vars)" border>
            <el-table-column prop="key" label="Variable Name" />
            <el-table-column prop="value" label="Value" />
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref, computed, nextTick, watch, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  record: {
    type: Object,
    default: null
  }
})

const emit = defineEmits(['update:modelValue'])

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const currentRecord = computed(() => props.record)
const activeTab = ref('request')
const requestSubTab = ref('headers')
const responseSubTab = ref('headers')
const httpstatChartRef = ref(null)

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

const httpstatTableData = computed(() => {
  if (!currentRecord.value || !currentRecord.value.httpstat) return []
  const hs = currentRecord.value.httpstat
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

  let chart = echarts.getInstanceByDom(httpstatChartRef.value)
  if (!chart) {
    chart = echarts.init(httpstatChartRef.value)
  }

  // Calculate start times for waterfall effect
  let currentStart = 0
  const startTimes = data.map(d => {
    const start = currentStart
    currentStart += d.duration
    return start
  })

  // Colors for different phases
  const colors = {
    'DNS 解析': '#5470C6',
    'TCP 连接': '#91CC75',
    'TLS 握手': '#FAC858',
    '服务端处理': '#EE6666',
    '数据传输': '#73C0DE'
  }

  chart.setOption({
    title: {
      text: 'HTTP 请求耗时分布',
      left: 'center',
      textStyle: { fontSize: 14 }
    },
    toolbox: {
      feature: {
        saveAsImage: { title: '保存图片' }
      }
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'shadow' },
      formatter: function (params) {
        let tar = params[1] // The second series is the visible one
        return tar.name + '<br/>' + tar.seriesName + ' : ' + tar.value + ' ms'
      }
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'value',
      name: '耗时 (ms)'
    },
    yAxis: {
      type: 'category',
      splitLine: { show: false },
      data: data.map(d => d.phase)
    },
    series: [
      {
        name: 'Placeholder',
        type: 'bar',
        stack: 'Total',
        itemStyle: {
          borderColor: 'transparent',
          color: 'transparent'
        },
        emphasis: {
          itemStyle: {
            borderColor: 'transparent',
            color: 'transparent'
          }
        },
        data: startTimes
      },
      {
        name: '耗时',
        type: 'bar',
        stack: 'Total',
        label: {
          show: true,
          position: 'right',
          formatter: '{c} ms'
        },
        data: data.map(d => ({
          value: d.duration,
          itemStyle: {
            color: colors[d.phase] || '#5470C6'
          }
        }))
      }
    ]
  })
}

const handleOpened = () => {
  nextTick(() => {
    renderHttpstatChart()
  })
}

const handleResize = () => {
  if (httpstatChartRef.value) {
    const chart = echarts.getInstanceByDom(httpstatChartRef.value)
    if (chart) chart.resize()
  }
}

window.addEventListener('resize', handleResize)

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  if (httpstatChartRef.value) {
    const chart = echarts.getInstanceByDom(httpstatChartRef.value)
    if (chart) chart.dispose()
  }
})

watch(() => props.record, () => {
  activeTab.value = 'request'
  requestSubTab.value = 'headers'
  responseSubTab.value = 'headers'
})
</script>

<style scoped>
.code-block {
  background-color: #f4f4f5;
  padding: 10px;
  border-radius: 4px;
  overflow-x: auto;
  font-family: monospace;
  white-space: pre-wrap;
}
</style>
