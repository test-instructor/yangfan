<template>
  <div>
    <div class="dashboard-line-box" style="display: flex">
      <div id="caseDetail">
        <el-table
          border
          :data="testCaseSimple"
          :cell-style="{ textAlign: 'center' }"
          :show-header="false"
        >
          <el-table-column property="label" label="label" width="120" />
          <el-table-column property="name" label="label" width="278" />
        </el-table>
      </div>
      <div id="testcases"></div>
      <div id="testSteps"></div>
    </div>
    <div></div>
  </div>
</template>

<script setup>
name = "ReportDetail";
import { nextTick, onMounted, onUnmounted, ref } from "vue";
import "echarts/theme/macarons";
// import * as echarts from 'echarts';
import * as echarts from "echarts/core";
import {
  TooltipComponent,
  LegendComponent,
  TitleComponent,
} from "echarts/components";
import { PieChart } from "echarts/charts";
import { LabelLayout } from "echarts/features";
import { CanvasRenderer } from "echarts/renderers";

echarts.use([
  TooltipComponent,
  LegendComponent,
  PieChart,
  CanvasRenderer,
  LabelLayout,
  TitleComponent,
]);

let pieOption;
const reportData = ref({});
let testCasesData = [];
let testStepsData = [];
const testCaseSimple = ref([]);

const initData = () => {
  reportData.value = {
    success: true,
    stat: {
      testcases: { total: 10, success: 7, fail: 3 },
      teststeps: { total: 270, successes: 200, failures: 70 },
    },
    time: { start_at: "2022-06-01T18:31:07.489339+08:00", duration: 0.5088207 },
    platform: {
      httprunner_version: "v4.1.0-beta",
      go_version: "go1.18",
      platform: "windows-amd64",
    },
    details: [
      {
        name: "测试名称0",
        success: true,
        stat: { total: 3, successes: 3, failures: 0 },
        time: {
          start_at: "2022-06-01T18:31:07.8227227+08:00",
          duration: 0.0166467,
        },
        in_out: {
          config_vars: { base_url: "http://localhost:8081/" },
          export_vars: { "": null },
        },
        records: [
          {
            ID: 0,
            parntID: 6,
            name: "Header1",
            step_type: "request",
            success: true,
            elapsed_ms: 5,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 4,
              StartTransfer: 4,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 5,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 7,
            name: "Header2",
            step_type: "request",
            success: true,
            elapsed_ms: 3,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 3,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 5,
            name: "Header",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
            export_vars: { Extract: null },
          },
        ],
        root_dir: "",
      },
      {
        name: "测试名称1",
        success: true,
        stat: { total: 3, successes: 3, failures: 0 },
        time: {
          start_at: "2022-06-01T18:31:07.8598116+08:00",
          duration: 0.0139837,
        },
        in_out: {
          config_vars: { base_url: "http://localhost:8081/" },
          export_vars: { "": null },
        },
        records: [
          {
            ID: 0,
            parntID: 6,
            name: "Header1",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 4,
              StartTransfer: 4,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 7,
            name: "Header2",
            step_type: "request",
            success: true,
            elapsed_ms: 3,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 3,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 5,
            name: "Header",
            step_type: "request",
            success: true,
            elapsed_ms: 3,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 3,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
            export_vars: { Extract: null },
          },
        ],
        root_dir: "",
      },
      {
        name: "测试名称2",
        success: true,
        stat: { total: 3, successes: 3, failures: 0 },
        time: {
          start_at: "2022-06-01T18:31:07.8955395+08:00",
          duration: 0.0170075,
        },
        in_out: {
          config_vars: { base_url: "http://localhost:8081/" },
          export_vars: { "": null },
        },
        records: [
          {
            ID: 0,
            parntID: 6,
            name: "Header1",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 4,
              StartTransfer: 4,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 7,
            name: "Header2",
            step_type: "request",
            success: true,
            elapsed_ms: 5,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 4,
              StartTransfer: 4,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 5,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 5,
            name: "Header",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
            export_vars: { Extract: null },
          },
        ],
        root_dir: "",
      },
      {
        name: "测试名称3",
        success: true,
        stat: { total: 3, successes: 3, failures: 0 },
        time: {
          start_at: "2022-06-01T18:31:07.9340496+08:00",
          duration: 0.0168644,
        },
        in_out: {
          config_vars: { base_url: "http://localhost:8081/" },
          export_vars: { "": null },
        },
        records: [
          {
            ID: 0,
            parntID: 6,
            name: "Header1",
            step_type: "request",
            success: true,
            elapsed_ms: 5,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 5,
              StartTransfer: 5,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 5,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 7,
            name: "Header2",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 4,
              StartTransfer: 4,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 5,
            name: "Header",
            step_type: "request",
            success: true,
            elapsed_ms: 4,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 3,
              StartTransfer: 3,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 4,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
            export_vars: { Extract: null },
          },
        ],
        root_dir: "",
      },
      {
        name: "测试名称4",
        success: true,
        stat: { total: 3, successes: 3, failures: 0 },
        time: {
          start_at: "2022-06-01T18:31:07.9725276+08:00",
          duration: 0.0256321,
        },
        in_out: {
          config_vars: { base_url: "http://localhost:8081/" },
          export_vars: { "": null },
        },
        records: [
          {
            ID: 0,
            parntID: 6,
            name: "Header1",
            step_type: "request",
            success: true,
            elapsed_ms: 9,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 9,
              StartTransfer: 9,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 9,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 7,
            name: "Header2",
            step_type: "request",
            success: true,
            elapsed_ms: 7,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 7,
              StartTransfer: 7,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 7,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
          },
          {
            ID: 0,
            parntID: 5,
            name: "Header",
            step_type: "request",
            success: true,
            elapsed_ms: 6,
            httpstat: {
              Connect: 0,
              ContentTransfer: 0,
              DNSLookup: 0,
              NameLookup: 0,
              Pretransfer: 0,
              ServerProcessing: 5,
              StartTransfer: 5,
              TCPConnection: 0,
              TLSHandshake: 0,
              Total: 5,
            },
            data: {
              success: true,
              req_resps: {
                request: {
                  body: {},
                  data: { Form: "Form" },
                  headers: {
                    "Content-Type": "application/json; charset=utf-8",
                    Header: "Header",
                  },
                  method: "GET",
                  params: { Params: "Params" },
                  url: "Header",
                },
                response: {
                  body: '""',
                  cookies: {},
                  headers: {
                    "Access-Control-Allow-Origin": "*",
                    Connection: "keep-alive",
                    "Content-Length": "0",
                    Date: "Wed, 01 Jun 2022 10:31:07 GMT",
                    "Keep-Alive": "timeout=5",
                  },
                  proto: "HTTP/1.1",
                  status_code: 404,
                },
              },
            },
            content_size: 0,
            export_vars: { Extract: null },
          },
        ],
        root_dir: "",
      },
    ],
  };
  testCaseSimple.value.push({
    label: "运行状态",
    name: reportData.value.success,
    key: "success",
  });
  testCaseSimple.value.push({
    label: "开始时间",
    name: reportData.value.time.start_at,
    key: "start_at",
  });
  testCaseSimple.value.push({
    label: "运行时长",
    name: reportData.value.time.duration,
    key: "duration",
  });

  testCasesData = [
    { value: reportData.value.stat.testcases["success"], name: "成功" },
    { value: reportData.value.stat.testcases["fail"], name: "失败" },
  ];
  testStepsData = [
    { value: reportData.value.stat.teststeps["successes"], name: "成功" },
    { value: reportData.value.stat.teststeps["failures"], name: "失败" },
  ];
  testCaseSimple.value.push({
    label: "用例数",
    name: reportData.value.stat.testcases["total"],
    key: "caseTotal",
  });
  testCaseSimple.value.push({
    label: "接口数",
    name: reportData.value.stat.teststeps["total"],
    key: "stepTotal",
  });
};

initData();

pieOption = {
  tooltip: {
    trigger: "item",
  },
  legend: {
    top: "0%",
    left: "center",
  },
  color: ["#91cc75", "#ee6666", "yellow", "blue", "purple"],
  series: [
    {
      name: "用例运行情况",
      type: "pie",
      radius: ["40%", "70%"],
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
onMounted(async () => {
  pieOption.series[0].data = testCasesData;
  pieOption.series[0].name = "用例运行情况";
  const testCaseDom = document.getElementById("testcases");
  const testCaseChart = echarts.init(testCaseDom, null, {
    renderer: "canvas",
    useDirtyRect: false,
  });
  testCaseChart.setOption(pieOption);

  pieOption.series[0].data = testStepsData;
  pieOption.series[0].name = "接口运行情况";
  const testStepDom = document.getElementById("testSteps");
  const testStepChart = echarts.init(testStepDom, null, {
    renderer: "canvas",
    useDirtyRect: false,
  });
  testStepChart.setOption(pieOption);
});
</script>

<style lang="scss" scoped>
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
  height: 270px;
  width: 270px;
}
#caseDetail {
  width: 400px;
}
</style>
