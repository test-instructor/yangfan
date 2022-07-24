<template>
  <div style="display: flex;">
    <div class="dashboard-line-box">
      <div id="caseDetail">
        <el-table
            border
            :data="testCaseSimple"
            :cell-style="{ textAlign: 'center' }"
            :show-header="false">
          <el-table-column property="label" label="label" width="120"/>
          <el-table-column property="name" label="label" width="278"/>
        </el-table>
      </div>
      <div id="testcases">
      </div>
      <div id="testSteps">
      </div>
    </div>
    <div style="width:960px;margin-left:20px;">
      <el-table
          id="reportDataId"
          ref="reportDataId"
          :data="reportData.details"
          height="98%"
      >

        <el-table-column label="运行状态" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.success?'success':'info'">{{ scope.row.success ? '成功' : '失败' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column property="name" label="用例名称" width="320"></el-table-column>
        <el-table-column label="运行时间" :formatter="runTime" width="160"></el-table-column>
        <el-table-column label="运行运行时长" width="120">
          <template #default="scope">
            {{ Number(scope.row.time.duration).toFixed(3) }}
          </template>
        </el-table-column>
        <el-table-column property="stat.successes" label="成功用例" width="100"></el-table-column>
        <el-table-column property="stat.failures" label="失败用例" width="100"></el-table-column>
        <el-table-column width="79">
          <template #default="scope">
            <el-button type="text" @click="toggleExpand(scope.row)">
              <span>{{ scope.row.ID === currentIndex ? '收起' : '展开' }}</span>
            </el-button>
          </template>
        </el-table-column>
        <el-table-column
            type="expand"
            width="1">
          <template #default="scope">
            <el-table
                style="width: 960px;padding-left: 20px"
                ref="apiTableData"
                id="apiTableData"
                :data="scope.row.records"
                :show-header="false"
            >
              <el-table-column
                  width="100"
              >
                <template #default="scope">
                  <el-tag :type="scope.row.success?'success':'info'">{{ scope.row.success ? '成功' : '失败' }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column
                  min-width="550"
                  align="center"
              >
                <template #default="scope">
                  <div class="block" :class="`block_${scope.row.data.req_resps.request.method.toLowerCase()}`">
                <span class="block-method block_method_color"
                      :class="`block_method_${scope.row.data.req_resps.request.method.toLowerCase()}`">
                  {{ scope.row.data.req_resps.request.method }}
                </span>
                    <span class="block-method block_url">{{ scope.row.data.req_resps.request.url }}</span>
                    <span class="block-summary-description">{{ scope.row.name }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column min-width="60">
                <template #default="scope">
                  <el-button type="text" @click="openDrawer(scope.row)">
                    <span>
                      详情
                    </span>
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-drawer v-if="drawer" v-model="drawer" :with-header="false" size="45%" title="请求详情">
      <div id="requestTimeEl"></div>

      <div
          style="margin:20px;height: 750px;overflow:auto;padding-right: 10px"
      >
        <div class="tableDetail">
          <div>
            <el-button
                type="info"
                @click="requestFunc"
            >
              {{ requestTable ? "收起" : "展开" }} Request 详情
            </el-button>
          </div>
          <br/>
          <el-table
              border
              :data="activeRow.requestData"
              v-show="requestTable"
              class="tableDetail"
          >
            <el-table-column
                width="120"
                align="center"
                prop="key"
                label="key"
            >
            </el-table-column>
            <el-table-column
                align="center"
                label="value"
            >
              <template #default="scope">
                <span v-if="!scope.row.isTable || JSON.stringify(scope.row.value) === '{}'">{{ scope.row.value }}</span>
                <tableKeyValue
                    :tableData="scope.row.value"
                    v-if="scope.row.isTable"
                ></tableKeyValue>

              </template>
            </el-table-column>
          </el-table>
        </div>

        <div class="tableDetail">
          <div>
            <el-button
                type="info"
                @click="responseFunc"
            >
              {{ responseTable ? "收起" : "展开" }} Response 详情
            </el-button>
          </div>
          <br/>
          <el-table
              border
              :data="activeRow.responseData"
              v-show="responseTable"
          >
            <el-table-column
                width="120"
                align="center"
                prop="key"
                label="key"
            >
            </el-table-column>
            <el-table-column
                align="center"
                label="value"
            >
              <template #default="scope">

                <span v-if="!scope.row.isTable">{{ scope.row.value }}</span>
                <tableKeyValue
                    :tableData="scope.row.value"
                    v-if="scope.row.isTable"
                ></tableKeyValue>

              </template>
            </el-table-column>
          </el-table>
        </div>

        <div
            v-if="activeRow.validators"
            class="tableDetail"
        >
          <div>
            <el-button
                type="info"
                @click="validatorsFunc"
            >
              {{ validatorsTable ? "收起" : "展开" }} 断言结果
            </el-button>
          </div>
          <br/>

          <el-table
              border
              :data="activeRow.validators"
              v-show="validatorsTable"
          >
            <el-table-column
                align="center"
                prop="check"
                label="状态"
            >
              <template #default="scope">
                <el-tag :type="scope.row.check_result==='pass'?'success':'danger'">
                  {{
                    scope.row.check_result === 'pass' ? '成功' : '失败'
                  }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column
                align="center"
                prop="check"
                label="断言字段"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="assert"
                label="断言类型"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="expect"
                label="期望结果"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="check_value"
                label="实际结果"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="msg"
                label="描述"
            >
            </el-table-column>

          </el-table>
        </div>

        <div
            v-if="activeRow.exportVars.length>0"
            class="tableDetail"
        >
          <div>
            <el-button
                type="info"
                @click="exportFunc"
            >
              {{ exportTable ? "收起" : "展开" }} 提取参数详情
            </el-button>
          </div>
          <br/>
          <el-table
              border
              :data="activeRow.exportVars"
              v-show="exportTable"
          >
            <el-table-column
                width="148"
                align="center"
                prop="key"
                label="key"
            >
            </el-table-column>
            <el-table-column
                align="center"
                label="value"
            >
              <template #default="scope">
                <span>{{ scope.row.value }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </el-drawer>
  </div>

</template>


<script setup>

import tableKeyValue from "@/view/interface/Reports/tableKeyValue.vue"

import {
  findReport
} from '@/api/report'

name = "ReportDetail"
import {onBeforeMount, onMounted, onUpdated, ref, watch} from 'vue'
import 'echarts/theme/macarons'
// import * as echarts from 'echarts';
import * as echarts from 'echarts/core';
import {
  TooltipComponent,
  LegendComponent,
  GridComponent,
  TitleComponent
} from 'echarts/components';
import {PieChart} from 'echarts/charts';
import {LabelLayout} from 'echarts/features';
import {BarChart} from 'echarts/charts';
import {CanvasRenderer} from 'echarts/renderers';
import { useRoute } from "vue-router";
import {getCurrentInstance} from "vue";

const route = useRoute()
echarts.use([
  TooltipComponent,
  LegendComponent,
  PieChart,
  CanvasRenderer,
  BarChart,
  LabelLayout,
  GridComponent,
  TitleComponent
]);

let pieOption;
const reportData = ref({})
const testCasesData = ref([])
const testStepsData = ref([])
const testCaseSimple = ref([])
let reportID = 1
const validatorsTable = ref(true)
const responseTable = ref(true)
const requestTable = ref(true)
const exportTable = ref(true)

const validatorsFunc = () => {
  validatorsTable.value = !validatorsTable.value
}
const responseFunc = () => {
  responseTable.value = !responseTable.value
}
const requestFunc = () => {
  requestTable.value = !requestTable.value
}
const exportFunc = () => {
  exportTable.value = !exportTable.value
}

let currentInstance
const currentIndex = ref(0);

const drawer = ref(false)
const activeRow = ref({})
let requestTimeOption;
requestTimeOption = {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      // Use axis to trigger tooltip
      type: 'shadow' // 'shadow' as default; can also be 'line' or 'shadow'
    }
  },
  legend: {},
  grid: {
    left: '3%',
    bottom: '2%',
    containLabel: true
  },
  xAxis: {
    type: 'value'
  },
  yAxis: {
    type: 'category',
    data: ['响应时间']
  },
}
const tableDatas = ref()
const tableKeyToValue = (data) => {
  let tableData = []
  for (let k in data) {
    let tableJson = {key: k, value: data[k]}
    tableData.push(tableJson)
  }
  tableDatas.value = tableData
  return tableData
}

const openDrawer = (row) => {
  drawer.value = true
  validatorsTable.value = true
  responseTable.value = true
  requestTable.value = true
  let requestData = []
  let responseData = []
  {
    requestData.push({key: "url", value: row.data.req_resps.request.url})
    requestData.push({key: "method", value: row.data.req_resps.request.method})
    requestData.push({key: "headers", value: row.data.req_resps.request.headers, isTable: true})
    requestData.push({key: "body", value: row.data.req_resps.request.body, isTable: true})
    requestData.push({key: "data", value: row.data.req_resps.request.data, isTable: true})
    requestData.push({key: "params", value: row.data.req_resps.request.params, isTable: true})
    responseData.push({key: "status_code", value: row.data.req_resps.response.status_code})
    responseData.push({key: "body", value: row.data.req_resps.response.body})
    responseData.push({key: "cookies", value: row.data.req_resps.response.cookies, isTable: true})
    responseData.push({key: "headers", value: row.data.req_resps.response.headers, isTable: true})
  }
  let export_vars = tableKeyToValue(row.export_vars)
  activeRow.value = {
    requestData: requestData,
    responseData: responseData,
    validators: row.data.validators,
    exportVars: export_vars,
  }
  let series = [];
  series = [
    {
      name: 'DNS 解析',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [row.httpstat.DNSLookup]
    },
    {
      name: 'TCP 连接',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [row.httpstat.TCPConnection]
    },
    {
      name: 'TLS 握手',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [row.httpstat.TLSHandshake]
    },
    {
      name: '服务端处理',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [row.httpstat.ServerProcessing]
    },
    {
      name: '数据传输',
      type: 'bar',
      stack: 'total',
      label: {
        show: true
      },
      emphasis: {
        focus: 'series'
      },
      data: [row.httpstat.ContentTransfer]
    }
  ]
  requestTimeOption.series = series
  setTimeout(() => {
    const requestTimeChartDom = document.getElementById('requestTimeEl');
    const requestTimeChart = echarts.init(requestTimeChartDom, null, {
      renderer: 'canvas',
      useDirtyRect: false
    });
    requestTimeChart.setOption(requestTimeOption);
  }, 100)
}


const getTestCaseDetailFunc = async (testCaseID) => {
  const res = await findReport({ID: testCaseID})
  if (res.code === 0) {
    reportData.value = res.data.reapicase
    return true
  }
}


const runTime = (row, column) => {
  let dt = new Date(row.time.start_at)
  return dt.getFullYear() + '-' + (dt.getMonth() + 1) + '-' + dt.getDate() + ' ' + dt.getHours() + ':' + dt.getMinutes() + ':' + dt.getSeconds()
}

const tableDdata = ref([])
const initData = async () => {
  testStepsData.value = []
  testStepsData.value = []
  testCaseSimple.value = []
  if (route.params.id > 0) {
    reportID = route.params.id
  }
  tableDdata.value = reportData.value.details
  await getTestCaseDetailFunc(reportID)
  testCaseSimple.value.push({label: '运行状态', name: reportData.value.success, key: 'success'})
  testCaseSimple.value.push({label: '开始时间', name: reportData.value.time.start_at, key: 'start_at'})
  testCaseSimple.value.push({label: '运行时长', name: reportData.value.time.duration, key: 'duration'})
  testCasesData.value = [
    {value: reportData.value.stat.testcases['success'], name: '成功'},
    {value: reportData.value.stat.testcases['fail'], name: '失败'}
  ]
  testStepsData.value = [
    {value: reportData.value.stat.teststeps['successes'], name: '成功'},
    {value: reportData.value.stat.teststeps['failures'], name: '失败'}
  ]
  testCaseSimple.value.push({label: '用例数', name: reportData.value.stat.testcases['total'], key: 'caseTotal'})
  testCaseSimple.value.push({label: '接口数', name: reportData.value.stat.teststeps['total'], key: 'stepTotal'})

}
initData()
watch(() => route.params.id, () => {
  if (route.params.id){
    initData()
  }
})


pieOption = {
  tooltip: {
    trigger: 'item'
  },
  legend: {
    top: '0%',
    left: 'center'
  },
  color: ['#91cc75', '#ee6666', 'yellow', 'blue', 'purple'],
  series: [
    {
      name: '用例运行情况',
      type: 'pie',
      radius: ['40%', '70%'],
      avoidLabelOverlap: false,
      label: {
        show: false,
        position: 'center'
      },
      emphasis: {
        label: {
          fontSize: '40',
          fontWeight: 'bold'
        }
      },
      labelLine: {
        show: false
      },
      data: testCasesData
    }
  ]
}

let testCaseChart = null
let testStepChart = null
onMounted(async () => {
  currentInstance = getCurrentInstance()
  pieOption.series[0].data = testCasesData.value
  pieOption.series[0].name = '用例运行情况'
  const testCaseDom = document.getElementById('testcases');
  testCaseChart = echarts.init(testCaseDom, null, {
    renderer: 'canvas',
    useDirtyRect: false
  });
  testCaseChart.setOption(pieOption);

  pieOption.series[0].data = testStepsData.value
  pieOption.series[0].name = '接口运行情况'
  const testStepDom = document.getElementById('testSteps');
  testStepChart = echarts.init(testStepDom, null, {
    renderer: 'canvas',
    useDirtyRect: false
  });
  testStepChart.setOption(pieOption);
})

watch(
    testStepsData, () => {
      pieOption.series[0].data = testCasesData.value
      pieOption.series[0].name = '用例运行情况'
      testCaseChart.setOption(pieOption);
      pieOption.series = [
        {
          data: testStepsData.value,
          name: '接口运行情况'
        }
      ]
      testStepChart.setOption(pieOption)
    })

const toggleExpand = (row) => {
  let table = currentInstance.ctx.$refs.reportDataId
  reportData.value.details.map((item) => {
    if (row.ID !== item.ID) {
      table.toggleRowExpansion(item, false)
    }
  })
  table.toggleRowExpansion(row)
  if (currentIndex.value === row.ID) {
    currentIndex.value = 0
  } else {
    currentIndex.value = row.ID
  }
}


</script>

<style lang="scss" scoped>
@import 'src/style/apiList';

.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }

  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}

#testcases, #testSteps {
  height: 270px;
  width: 270px;
}

.tableDetail {
  padding-bottom: 5px;
}


#caseDetail {
  width: 400px;
}

#requestTimeEl {
  height: 100px;
  margin-top: 20px;
  margin-bottom: 20px;
}

</style>
