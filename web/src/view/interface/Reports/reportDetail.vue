<template>
  <div style="display: flex">
    <div class="dashboard-line-box">
      <div id="caseDetail">
        <el-table
          border
          :data="testCaseSimple"
          :cell-style="{ textAlign: 'center' }"
          :show-header="false"
        >
          <el-table-column property="label" label="label" width="100" />

          <el-table-column width="230">
            <template #default="scope">
              <el-table
                border
                :show-header="false"
                v-if="scope.row.str === 'case'"
                :data="scope.row.name"
              >
                <template #default="scope">
                  <el-table-column title="label" prop="label">
                  </el-table-column>
                  <el-table-column title="name" align="center" width="100">
                    <template #default="scope">
                      <el-tag v-if="scope.row.str === 'fail'" type="warning">{{
                        scope.row.name
                      }}</el-tag>
                      <el-tag v-if="scope.row.str === 'success'" type="success">
                        {{ scope.row.name }}
                      </el-tag>
                      <el-tag v-if="scope.row.str === 'error'" type="danger">
                        {{ scope.row.name }}
                      </el-tag>
                      <el-tag v-if="scope.row.str === 'skip'" type="info">
                        {{ scope.row.name }}
                      </el-tag>
                      <el-tag v-if="scope.row.str === 'total'">
                        {{ scope.row.name }}
                      </el-tag>
                    </template>
                  </el-table-column>
                </template>
              </el-table>
              <el-tag
                v-if="scope.row.label === '运行状态'"
                :type="scope.row.name ? 'success' : 'danger'"
                effect="dark"
                >{{ scope.row.name ? "成功" : "失败" }}</el-tag
              >
              {{ scope.row.str === "str" ? scope.row.name : "" }}
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>
    <div style="width: 1010px; margin-left: 20px">
      <el-table ref="reportDataId" :data="reportData.details" height="780px">
        <el-table-column label="运行状态" width="80">
          <template #default="scope">
            <el-tag
              :type="scope.row.success ? 'success' : 'danger'"
              effect="dark"
              >{{ scope.row.success ? "成功" : "失败" }}</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column property="name" label="用例名称" width="350">
          <template #default="scope">
            <el-tag type="danger" v-if="setupCaseShow(scope.row)">{{
              "前置步骤"
            }}</el-tag>
            {{ scope.row.name }}
          </template>
        </el-table-column>
        <el-table-column
          label="运行时间"
          :formatter="runTime"
          width="165"
        ></el-table-column>
        <el-table-column label="运行时长" width="115">
          <template #default="scope">
            {{ formatTime(scope.row.time.duration) }}
          </template>
        </el-table-column>
        <el-table-column
          property="stat.successes"
          label="成功"
          width="60"
        ></el-table-column>
        <el-table-column
          property="stat.failures"
          label="失败"
          width="60"
        ></el-table-column>
        <el-table-column
          property="stat.error"
          label="错误"
          width="60"
        ></el-table-column>
        <el-table-column
          property="stat.skip"
          label="跳过"
          width="60"
        ></el-table-column>
        <el-table-column width="79">
          <template #default="scope">
            <el-button type="text" @click="toggleExpand(scope.row)">
              <span>{{ scope.row.ID === currentIndex ? "收起" : "展开" }}</span>
            </el-button>
          </template>
        </el-table-column>
        <el-table-column type="expand" width="1">
          <template #default="scope">
            <el-table
              :show-header="false"
              id="apiTableData"
              :data="scope.row.records"
              :default-expand-all="true"
              style="padding-left: 15px"
            >
              <el-table-column width="950" label="name">
                <template #default="scope">
                  <div class="block" :class="`block_patch`">
                    <span
                      class="block-method block_method_color"
                      :class="`block_method_patch`"
                    >
                      {{ "STEP" }}
                    </span>
                    <div class="block"></div>
                    <span class="block-method block_url">{{
                      scope.row.name
                    }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column
                type="expand"
                :default-expand-all="true"
                :resizable="false"
              >
                <template #default="scope">
                  <el-table
                    ref="apiTableData"
                    id="apiTableData"
                    :data="scope.row.data"
                    :show-header="false"
                    v-if="shouStep(scope.row.data)"
                  >
                    <el-table-column width="70">
                      <template #default="scope">
                        <el-tag
                          :type="shouStatus(scope.row)[0]"
                          effect="dark"
                          >{{ shouStatus(scope.row)[1] }}</el-tag
                        >
                      </template>
                    </el-table-column>
                    <el-table-column min-width="600" align="center">
                      <template #default="scope">
                        <div
                          class="block"
                          :class="`block_${dataMethod(scope.row)[0]}`"
                        >
                          <span
                            class="block-method block_method_color"
                            :class="`block_method_${dataMethod(scope.row)[0]}`"
                          >
                            {{ dataMethod(scope.row)[1] }}
                          </span>
                          <span class="block-method block_url">{{
                            scope.row.data
                              ? scope.row.data.req_resps.request.url
                              : ""
                          }}</span>
                          <span class="block-summary-description">{{
                            scope.row.name
                          }}</span>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column min-width="40">
                      <template #default="scope">
                        <el-button type="text" @click="openDrawer(scope.row)">
                          <span> 详情 </span>
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </template>
              </el-table-column>
            </el-table>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
      <div id="testcases"></div>
      <div id="testSteps"></div>
    </div>

    <el-drawer
      v-if="drawer"
      v-model="drawer"
      :with-header="false"
      size="70%"
      title="请求详情"
      :tabindex="-1"
    >
      <div v-if="requestTimeShow" id="requestTimeEl"></div>

      <div style="margin: 20px; overflow: auto; padding-right: 10px">
        <div class="tableDetail">
          <el-form-item label="重试次数:" v-if="isRetry"
            ><el-tag>{{ retry }}</el-tag></el-form-item
          >
          <div>
            <el-button type="info" @click="requestFunc">
              {{ requestTable ? "收起" : "展开" }} Request 详情
            </el-button>

            <el-button type="info" @click="reportDetailFunc('request')">
              request json 数据
            </el-button>
          </div>
          <br />
          <el-table
            border
            :data="activeRow.requestData"
            v-show="requestTable"
            class="tableDetail"
            :row-style="{ height: '20px' }"
          >
            <el-table-column width="120" align="center" prop="key" label="key">
            </el-table-column>
            <!--            <el-table-column width="80" align="center" prop="key" label="操作">-->
            <!--              <template v-slot="scope">-->
            <!--                <el-button type="text" @click="copy_text(scope.row)"-->
            <!--                  >复制</el-button-->
            <!--                >-->
            <!--              </template>-->
            <!--            </el-table-column>-->
            <el-table-column align="center" label="value">
              <template #default="scope">
                <span
                  v-if="
                    !scope.row.isTable ||
                    JSON.stringify(scope.row.value) === '{}'
                  "
                  >{{ scope.row.value }}</span
                >
                <tableKeyValue
                  :tableData="scope.row.value"
                  v-if="scope.row.isTable"
                ></tableKeyValue>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <div class="tableDetail" v-if="responseShow">
          <div>
            <el-button type="info" @click="responseFunc">
              {{ responseTable ? "收起" : "展开" }} Response 详情
            </el-button>
            <el-button type="info" @click="reportDetailFunc('response')">
              response json 数据
            </el-button>
          </div>
          <br />
          <el-table
            border
            :data="activeRow.responseData"
            v-show="responseTable"
            :row-style="{ height: '20px' }"
          >
            <el-table-column width="120" align="center" prop="key" label="key">
            </el-table-column>
            <!--            <el-table-column width="80" align="center" prop="key" label="操作">-->
            <!--              <template v-slot="scope">-->
            <!--                <el-button type="text" @click="copy_text(scope.row)"-->
            <!--                  >复制</el-button-->
            <!--                >-->
            <!--              </template>-->
            <!--            </el-table-column>-->
            <el-table-column align="center" label="value">
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

        <div v-if="activeRow.validators" class="tableDetail">
          <div>
            <el-button type="info" @click="validatorsFunc">
              {{ validatorsTable ? "收起" : "展开" }} 断言结果
            </el-button>
          </div>
          <br />

          <el-table
            border
            :data="activeRow.validators"
            v-show="validatorsTable"
            :row-style="{ height: '20px' }"
          >
            <el-table-column align="center" prop="check" label="状态">
              <template #default="scope">
                <el-tag
                  :type="
                    scope.row.check_result === 'pass' ? 'success' : 'danger'
                  "
                  effect="dark"
                >
                  {{ scope.row.check_result === "pass" ? "成功" : "失败" }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column align="center" prop="check" label="断言字段">
            </el-table-column>
            <el-table-column align="center" prop="assert" label="断言类型">
            </el-table-column>
            <el-table-column align="center" prop="expect" label="期望结果">
            </el-table-column>
            <el-table-column align="center" prop="check_value" label="实际结果">
            </el-table-column>
            <el-table-column align="center" prop="msg" label="描述">
            </el-table-column>
          </el-table>
        </div>

        <div
          v-if="activeRow.exportVars && activeRow.exportVars.length > 0"
          class="tableDetail"
        >
          <div>
            <el-button type="info" @click="exportFunc">
              {{ exportTable ? "收起" : "展开" }} 提取参数详情
            </el-button>
          </div>
          <br />
          <el-table
            border
            :data="activeRow.exportVars"
            v-show="exportTable"
            :row-style="{ height: '20px' }"
          >
            <el-table-column width="148" align="center" prop="key" label="key">
            </el-table-column>
            <el-table-column align="center" label="value">
              <template #default="scope">
                <span>{{ scope.row.value }}</span>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </el-drawer>

    <el-dialog
      :title="dialogTitle"
      :before-close="dialogClose"
      @close="dialogClose"
      v-model="dialogFormDetail"
    >
      <jsons
        v-if="dialogFormDetail"
        :heights="600"
        :jsons="dialogData ? dialogData : ''"
      >
      </jsons>
    </el-dialog>
  </div>
</template>

<script setup>
import tableKeyValue from "@/view/interface/Reports/tableKeyValue.vue";

import { findReport, getReportDetail } from "@/api/report";

name = "ReportDetail";
import { onBeforeMount, onMounted, onUpdated, reactive, ref, watch } from "vue";
import "echarts/theme/macarons";
// import * as echarts from 'echarts';
import * as echarts from "echarts/core";
import {
  TooltipComponent,
  LegendComponent,
  GridComponent,
  TitleComponent,
} from "echarts/components";
import { PieChart } from "echarts/charts";
import { LabelLayout } from "echarts/features";
import { BarChart } from "echarts/charts";
import { CanvasRenderer } from "echarts/renderers";
import { useRoute } from "vue-router";
import { getCurrentInstance } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { formatDate } from "@/utils/format";
import jsons from "@/view/interface/interfaceComponents/Jsons.vue";

import { Discount } from "@element-plus/icons-vue";
import { defineComponent } from "vue";
import message from "@element-plus/icons-vue/dist/es/message.mjs";
import Jsons from "@/view/interface/interfaceComponents/Jsons.vue";

const route = useRoute();
echarts.use([
  TooltipComponent,
  LegendComponent,
  PieChart,
  CanvasRenderer,
  BarChart,
  LabelLayout,
  GridComponent,
  TitleComponent,
]);

let pieOption;
const reportData = ref({});
const reportDataDetail = new Map();
const testCasesData = ref([]);
const testStepsData = ref([]);
const testCaseSimple = ref([]);
let reportID = 1;
const validatorsTable = ref(true);
const responseTable = ref(true);
const requestTable = ref(true);
const exportTable = ref(true);
const dialogFormVisible = ref(false);
const dialogFormDetail = ref(false);
const dialogTitle = ref("");
const dialogData = ref({});
const dialogDataRequest = ref({});
const dialogDataResponse = ref({});

const validatorsFunc = () => {
  validatorsTable.value = !validatorsTable.value;
};
const responseFunc = () => {
  responseTable.value = !responseTable.value;
};
const requestFunc = () => {
  requestTable.value = !requestTable.value;
};

const reportDetailFunc = (dataType) => {
  if (dataType === "request") {
    dialogTitle.value = "request json 数据";
    dialogData.value = dialogDataRequest.value;
  }
  if (dataType === "response") {
    dialogTitle.value = "response json 数据";
    dialogData.value = dialogDataResponse.value;
  }
  dialogFormDetail.value = true;
};
const exportFunc = () => {
  dialogFormVisible.value = true;
};
let copy_report_data = "";
const copy_text = (row) => {
  copy_report_data = JSON.stringify(row);
  navigator.clipboard
    .writeText(copy_report_data)
    .then(() => {
      ElMessage.success("复制成功");
    })
    .catch((error) => {
      ElMessage.error("复制失败");
    });
};

let currentInstance;
const currentIndex = ref(0);

const drawer = ref(false);
const activeRow = ref({});
const retry = ref(0);
const isRetry = ref(false);
let requestTimeOption;
requestTimeOption = {
  tooltip: {
    trigger: "axis",
    position: function (point, params, dom, rect, size) {
      let point0 = point[0] + 10;
      if (point[0] > 300) {
        point0 = point[0] - 175;
      }
      return [point0, "10%"];
    },
    axisPointer: {
      type: "shadow",
    },
  },
  legend: {},
  grid: {
    left: "3%",
    bottom: "2%",
    containLabel: true,
  },
  xAxis: {
    type: "value",
  },
  yAxis: {
    type: "category",
    data: ["响应时间"],
  },
};
const tableDatas = ref();
const requestTimeShow = ref(false);
const responseShow = ref(false);
const tableKeyToValue = (data) => {
  let tableData = [];
  for (let k in data) {
    let tableJson = { key: k, value: data[k] };
    tableData.push(tableJson);
  }
  tableDatas.value = tableData;
  return tableData;
};

const dataMethod = (row) => {
  if (!row.data) {
    return ["delete", "执行错误"];
  } else {
    if (row.data.req_resps.response.proto === "gRPC") {
      return ["put", "gRPC"];
    } else {
      let method = row.data.req_resps.request.method.toLowerCase();
      return [method, method];
    }
  }
};

const formatTime = (duration) => {
  let hours = Math.floor(duration / 3600);
  let minutes = Math.floor((duration % 3600) / 60);
  let seconds = Math.floor(duration % 60);

  let formattedTime = "";

  if (hours > 0) {
    formattedTime += hours + "小时";
  }

  if (minutes > 0 || (hours > 0 && seconds > 0)) {
    formattedTime += minutes + "分钟";
  }

  if (seconds <= 0 && formattedTime === "") {
    formattedTime = "1秒";
  } else if (seconds > 0) {
    formattedTime += seconds + "秒";
  }

  return formattedTime;
};

const setupCaseShow = (row) => {
  return !!(
    reportData.value.setup_case && row.ID === reportData.value.details[0].ID
  );
};

const shouStep = (data) => {
  return data.length > 0;
};

const openDrawer = async (row) => {
  let row_id = row.ID;
  row = reportDataDetail.get(row_id);
  if (!row) {
    let res = await getReportDetail({ ID: row_id });
    if (res.code === 0) {
      row = res.data.data;
    }
  }
  if (row.attachments === "") {
    requestTimeShow.value = false;
    responseShow.value = false;
    if (row.data.req_resps.response != null) {
      responseShow.value = true;
      requestTimeShow.value = row.data.req_resps.response.proto !== "gRPC";
    }
    console.log("=============responseShow.value", responseShow.value);
    console.log("=============requestTimeShow.value", requestTimeShow.value);
    drawer.value = true;
    validatorsTable.value = true;
    responseTable.value = true;
    requestTable.value = true;
    let requestData = [];
    let responseData = [];
    {
      dialogDataRequest.value = row.data.req_resps.request;
      dialogDataResponse.value = row.data.req_resps.response;
      requestData.push({ key: "url", value: row.data.req_resps.request.url });
      if (row.data.req_resps.request.method) {
        requestData.push({
          key: "method",
          value: row.data.req_resps.request.method,
        });
      }
      requestData.push({
        key: "headers",
        value: row.data.req_resps.request.headers,
        isTable: true,
      });
      if (row.data.req_resps.request.body || requestTimeShow) {
        requestData.push({
          key: "body",
          value: row.data.req_resps.request.body,
          isTable: true,
        });
      }
      if (row.data.req_resps.request.data) {
        requestData.push({
          key: "data",
          value: row.data.req_resps.request.data,
          isTable: true,
        });
      }
      if (row.data.req_resps.request.params) {
        requestData.push({
          key: "params",
          value: row.data.req_resps.request.params,
          isTable: true,
        });
      }
      if (responseShow.value) {
        responseData.push({
          key: "status_code",
          value: row.data.req_resps.response.status_code,
        });
        if (row.data.req_resps.response.err) {
          responseData.push({
            key: "err",
            value: row.data.req_resps.response.err,
          });
        }
        responseData.push({
          key: "body",
          value: row.data.req_resps.response.body,
        });
        if (row.data.req_resps.response.cookies) {
          responseData.push({
            key: "cookies",
            value: row.data.req_resps.response.cookies,
            isTable: true,
          });
        }
        responseData.push({
          key: "headers",
          value: row.data.req_resps.response.headers,
          isTable: true,
        });
      }
    }
    let export_vars = tableKeyToValue(row.export_vars);
    activeRow.value = {
      requestData: requestData,
      responseData: responseData,
      validators: row.data.validators,
      exportVars: export_vars,
    };
    console.log("row.data.req_resps", row.data.req_resps);
    retry.value = row.retry;
    isRetry.value = row.retry > 0;
    console.log("retry", retry.value, isRetry.value, row.data.retry);
    console.log("row.data", row.data);
    let series = [];
    if (requestTimeShow.value) {
      series = [
        {
          name: "DNS 解析",
          type: "bar",
          stack: "total",
          label: {
            show: true,
          },
          emphasis: {
            focus: "series",
          },
          data: [row.httpstat.DNSLookup],
        },
        {
          name: "TCP 连接",
          type: "bar",
          stack: "total",
          label: {
            show: true,
          },
          emphasis: {
            focus: "series",
          },
          data: [row.httpstat.TCPConnection],
        },
        {
          name: "TLS 握手",
          type: "bar",
          stack: "total",
          label: {
            show: true,
          },
          emphasis: {
            focus: "series",
          },
          data: [row.httpstat.TLSHandshake],
        },
        {
          name: "服务端处理",
          type: "bar",
          stack: "total",
          label: {
            show: true,
          },
          emphasis: {
            focus: "series",
          },
          data: [row.httpstat.ServerProcessing],
        },
        {
          name: "数据传输",
          type: "bar",
          stack: "total",
          label: {
            show: true,
          },
          emphasis: {
            focus: "series",
          },
          data: [row.httpstat.ContentTransfer],
        },
      ];
      requestTimeOption.series = series;
      setTimeout(() => {
        const requestTimeChartDom = document.getElementById("requestTimeEl");
        const requestTimeChart = echarts.init(requestTimeChartDom, null, {
          renderer: "canvas",
          useDirtyRect: false,
        });
        requestTimeChart.setOption(requestTimeOption);
      }, 50);
    }
  } else {
    ElMessageBox.alert(
      "当前用例执行错误，错误详情：" + row.attachments,
      "用例执行出错",
      {
        type: "error",
      }
    );
  }
};

const getTestCaseDetailFunc = async (testCaseID) => {
  const res = await findReport({ ID: testCaseID });
  let reset = true;
  if (res.code === 0) {
    let reapicase = JSON.parse(JSON.stringify(res.data.reapicase));
    res.data.reapicase.details.forEach((item, index, arr) => {
      item.records.forEach((items, indexs, arrs) => {
        let stepName =
          res.data.reapicase.details[index].records[indexs].name + " - ";
        res.data.reapicase.details[index].records[indexs].data.forEach(
          (item2, index2) => {
            let casename =
              res.data.reapicase.details[index].records[indexs].data[index2]
                .name;
            res.data.reapicase.details[index].records[indexs].data[
              index2
            ].name = casename.substring(stepName.length);
          }
        );
      });
      reportData.value = reapicase;
      return false;
    });
  }
  if (reset) {
    reportData.value = res.data.reapicase;
  }
  getTestCaseDetailData();
};

const shouStatus = (row) => {
  console.log("row.skip", row);
  if (row.success) {
    if (row.skip) {
      return ["info", "跳过"];
    }
    return ["success", "成功"];
  }
  if (row.attachments && row.attachments != "") {
    return ["danger", "错误"];
  }
  return ["warning", "失败"];
};

const getTestCaseDetailData = async () => {
  for (var i = 0; i < reportData.value.details.length; i++) {
    for (var j = 0; j < reportData.value.details[i].records.length; j++) {
      for (
        var k = 0;
        k < reportData.value.details[i].records[j].data.length;
        k++
      ) {
        let data = reportData.value.details[i].records[j].data[k];
        let res = await getReportDetail({ ID: data.ID });
        if (res.code === 0) {
          reportDataDetail.set(data.ID, res.data.data);
        }
      }
    }
  }
};

const runTime = (row, column) => {
  let dt = new Date(row.time.start_at);
  return (
    dt.getFullYear() +
    "-" +
    (dt.getMonth() + 1) +
    "-" +
    dt.getDate() +
    " " +
    dt.getHours() +
    ":" +
    dt.getMinutes() +
    ":" +
    dt.getSeconds()
  );
};

const tableDdata = ref([]);
const initData = async () => {
  testStepsData.value = [];
  testStepsData.value = [];
  testCaseSimple.value = [];
  if (route.params.report_id > 0) {
    reportID = route.params.report_id;
  }
  tableDdata.value = reportData.value.details;
  await getTestCaseDetailFunc(reportID);
  testCaseSimple.value.push({
    label: "报告名称",
    name: reportData.value.name,
    key: "name",
    str: "str",
  });
  testCaseSimple.value.push({
    label: "运行节点",
    name: reportData.value.hostname,
    key: "hostname",
    str: "str",
  });
  testCaseSimple.value.push({
    label: "运行状态",
    name: reportData.value.success,
    key: "success",
  });
  testCaseSimple.value.push({
    label: "开始时间",
    name: formatDate(reportData.value.time.start_at),
    key: "start_at",
    str: "str",
  });
  testCaseSimple.value.push({
    label: "运行时长",
    name: formatTime(reportData.value.time.duration),
    key: "duration",
    str: "str",
  });
  testCasesData.value = [
    { value: reportData.value.stat.testcases["success"], name: "成功" },
    { value: reportData.value.stat.testcases["fail"], name: "失败" },
  ];
  testStepsData.value = [
    { value: reportData.value.stat.teststeps["successes"], name: "成功" },
    { value: reportData.value.stat.teststeps["error"], name: "错误" },
    { value: reportData.value.stat.teststeps["failures"], name: "失败" },
    { value: reportData.value.stat.teststeps["skip"], name: "跳过" },
  ];
  const tesecase = ref([]);
  const apicase = ref([]);

  // testCaseSimple.value.push({label: '执行用例数', name: reportData.value.stat.testcases['total'], key: 'caseTotal', str:"total"})
  // testCaseSimple.value.push({label: '成功用例数', name: reportData.value.stat.testcases['success'], key: 'caseTotal', str:"success"})
  // testCaseSimple.value.push({label: '失败用例数', name: reportData.value.stat.testcases['fail'], key: 'caseTotal', str:"fail"})
  apicase.value.push({
    label: "执行接口数",
    name: reportData.value.stat.teststeps["total"],
    key: "stepTotal",
    str: "total",
  });
  apicase.value.push({
    label: "成功接口数",
    name: reportData.value.stat.teststeps["successes"],
    key: "stepTotal",
    str: "success",
  });
  apicase.value.push({
    label: "失败接口数",
    name: reportData.value.stat.teststeps["failures"],
    key: "stepTotal",
    str: "fail",
  });
  apicase.value.push({
    label: "错误接口数",
    name: reportData.value.stat.teststeps["error"],
    key: "stepTotal",
    str: "error",
  });
  apicase.value.push({
    label: "错误接口数",
    name: reportData.value.stat.teststeps["skip"],
    key: "stepTotal",
    str: "skip",
  });

  tesecase.value.push({
    label: "执行用例数",
    name: reportData.value.stat.testcases["total"],
    key: "caseTotal",
    str: "total",
  });
  tesecase.value.push({
    label: "成功用例数",
    name: reportData.value.stat.testcases["success"],
    key: "caseTotal",
    str: "success",
  });
  tesecase.value.push({
    label: "失败用例数",
    name: reportData.value.stat.testcases["fail"],
    key: "caseTotal",
    str: "fail",
  });

  testCaseSimple.value.push({
    label: "用例状态",
    name: tesecase,
    key: "stepTotal",
    str: "case",
  });
  testCaseSimple.value.push({
    label: "接口状态",
    name: apicase,
    key: "stepTotal",
    str: "case",
  });
};

const columns = reactive([
  {
    title: "label",
    dataIndex: "label",
  },
  {
    title: "name",
    dataIndex: "name",
  },
]);

initData();
watch(
  () => route.params.report_id,
  () => {
    if (route.params.report_id) {
      initData();
    }
  }
);

pieOption = {
  title: {
    text: "用例运行情况",
    left: "8%",
    top: "0%",
  },
  tooltip: {
    trigger: "item",
  },
  legend: {
    top: "8%",
    right: "8%",
  },
  color: ["#67C23A", "#F56C6C", "#E6A23C", "#909399", "purple"],
  series: [
    {
      name: "用例运行情况",
      type: "pie",
      radius: ["50%", "75%"],
      avoidLabelOverlap: false,
      label: {
        show: false,
        position: "center",
      },
      emphasis: {
        label: {
          fontSize: "40",
          fontWeight: "bold",
        },
      },
      labelLine: {
        show: false,
      },
      data: testCasesData,
    },
  ],
};

let testCaseChart = null;
let testStepChart = null;
currentInstance = getCurrentInstance();
onMounted(async () => {
  pieOption.series[0].data = testCasesData.value;
  pieOption.series[0].name = "用例";
  const testCaseDom = document.getElementById("testcases");
  testCaseChart = echarts.init(testCaseDom, null, {
    renderer: "canvas",
    useDirtyRect: false,
  });
  testCaseChart.setOption(pieOption);
  pieOption.series[0].data = testStepsData.value;
  pieOption.series[0].name = "接口运行情况";
  pieOption.title.text = "用例运行情况";
  const testStepDom = document.getElementById("testSteps");
  testStepChart = echarts.init(testStepDom, null, {
    renderer: "canvas",
    useDirtyRect: false,
  });
  testStepChart.setOption(pieOption);
});

watch(testStepsData, () => {
  pieOption.series[0].data = testCasesData.value;
  pieOption.series[0].name = "用例运行情况";
  testCaseChart.setOption(pieOption);
  pieOption.title.text = "接口接口运行情况";
  pieOption.series = [
    {
      data: testStepsData.value,
      name: "接口运行情况",
    },
  ];
  console.log("testStepsData.value", testStepsData.value);
  testStepChart.setOption(pieOption);
});

const toggleExpand = (row) => {
  let table = currentInstance.refs.reportDataId;
  reportData.value.details.map((item) => {
    if (row.ID !== item.ID) {
      table.toggleRowExpansion(item, false);
    }
  });
  table.toggleRowExpansion(row);
  if (currentIndex.value === row.ID) {
    currentIndex.value = 0;
  } else {
    currentIndex.value = row.ID;
  }
};

const dialogClose = () => {
  dialogFormDetail.value = false;
  dialogData.value = {};
};
</script>

<style lang="scss" scoped>
@import "src/style/apiList";

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

#testcases,
#testSteps {
  height: 400px;
  width: 300px;
}

.tableDetail {
  padding-bottom: 5px;
}

#caseDetail {
  width: 330px;
}

#requestTimeEl {
  height: 10%;
  margin-top: 40px;
  margin-bottom: 40px;
}

.el-table__body-wrapper {
  &::-webkit-scrollbar {
    // 整个滚动条
    width: 0; // 纵向滚动条的宽度
    background: rgba(213, 215, 220, 0.3);
    border: none;
  }
  &::-webkit-scrollbar-track {
    // 滚动条轨道
    border: none;
  }
}

.el-table th.gutter {
  display: none;
  width: 0;
}
.el-table colgroup col[name="gutter"] {
  display: none;
  width: 0;
}

.el-table__body {
  width: 100% !important;
}
</style>
