<template>
  <div style="background-color: #ffffff">
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="任务名称：">
          <span>{{ reportName }}</span>
        </el-form-item>
        <el-form-item label="任务状态：">
          <el-tag :key="stateName" :type="stateStyle">{{ stateName }}</el-tag>
        </el-form-item>
        <el-form-item label="操作：">
          <el-button
            @click="updateDetail"
            type="primary"
            :disabled="boomerButton"
            >手动刷新</el-button
          >
          <el-button
            @click="resetBoomer"
            type="success"
            :disabled="boomerButton"
            >调整运行参数</el-button
          >
          <el-button @click="stopBoomer" type="danger" :disabled="boomerButton"
            >停止运行</el-button
          >
        </el-form-item>

        <el-form-item>
          <el-select
            :disabled="boomerButton"
            v-model="timerValue"
            @change="updateTimerData(timerValue)"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in timerOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </div>
    <div>
      <el-tabs
        v-model="activeName"
        type="border-card"
        class="demo-tabs"
        @tab-click="handleClick"
      >
        <el-tab-pane label="性能指标" name="grafana">
          <iframe class="grafana-iframe" :src="grafanaUrl"></iframe>
        </el-tab-pane>
        <el-tab-pane label="节点性能指标" name="node">
          <iframe class="grafana-iframe" :src="grafanaStatsUrl"></iframe>
        </el-tab-pane>
      </el-tabs>
    </div>
    <el-dialog
      v-model="dialogRunner"
      :before-close="closeRunner"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      width="400px"
      title="调整运行参数"
    >
      <el-form :model="runnerConfig" label-position="right" label-width="160px">
        <el-form-item label="并发用户数：">
          <el-input-number
            v-model="runnerConfig.spawnCount"
            :min="1"
            step-strictly
          />
        </el-form-item>
        <el-form-item label="初始每秒增加用户数：">
          <el-input-number
            v-model="runnerConfig.spawnRate"
            :min="1"
            step-strictly
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeRunner">取 消</el-button>
          <el-button size="small" type="primary" @click="updateUser"
            >调整运行参数</el-button
          >
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "pReportDetail",
};
</script>

<script setup>
import { findReport } from "@/api/performance";
import { useRoute } from "vue-router";
import { onMounted, ref, watch } from "vue";
import { Discount } from "@element-plus/icons-vue";

import { runBoomer, stopBoom, rebalance } from "@/api/runTestCase";
import { ElMessage } from "element-plus";

const activeName = ref("grafana");
let performance_id = 0;
const handleClick = (Event) => {
  updateDetail();
};
const boomerButton = ref(false);
const route = useRoute();
const state = ref(0);
const respData = ref({});
const errData = ref([]);
const errDataKey = ref({});
const reportName = ref("");
const PCT95 = ref(0);
let reportID = 1;
const dialogRunner = ref(false);
const grafanaUrl = ref("");
const grafanaStatsUrl = ref("");
const grafana_host = ref("");
const grafana_dashboard = ref("");
const grafana_dashboard_name = ref("");
const grafana_dashboard_stats = ref("");
const grafana_dashboard_stats_name = ref("");
const CreatedAt = ref("");
const UpdatedAt = ref("");
let detailId = 0;
const tableData = ref([]);

let runnerConfig = {
  spawnCount: 1,
  spawnRate: 1,
};
const getTestCaseDetailFunc = async (testCaseID) => {
  const res = await findReport({ ID: testCaseID });
  if (res.code === 0) {
    performance_id = res.data.reapicase.performance_id;
    respData.value = res.data.reapicase;
    state.value = res.data.reapicase.state;
    runState(res.data.reapicase.state);
    reportName.value = res.data.reapicase.name;
    CreatedAt.value = res.data.reapicase.CreatedAt;
    UpdatedAt.value = res.data.reapicase.UpdatedAt;
    grafana_host.value = res.data.grafana_host;
    grafana_dashboard.value = res.data.grafana_dashboard;
    grafana_dashboard_name.value = res.data.grafana_dashboard_name;
    grafana_dashboard_stats.value = res.data.grafana_dashboard_stats;
    grafana_dashboard_stats_name.value = res.data.grafana_dashboard_stats_name;
    getGrafanaUrl();
  }
};

const removeData = () => {
  respData.value = {};
  state.value = 0;
  respData.value = {};
  errData.value = [];
  errDataKey.value = {};
  reportName.value = "";
  PCT95.value = 0;
};

const getGrafanaUrl = () => {
  let url = "";
  if (
    grafana_host.value.startsWith("http://") ||
    grafana_host.value.startsWith("https://")
  ) {
    url = grafana_host.value;
  } else {
    url = "http://" + grafana_host.value;
  }
  if (grafana_host.value.endsWith("/")) {
    url = url + "d/";
  } else {
    url = url + "/d/";
  }
  let grafana_url =
    url + grafana_dashboard.value + "/" + grafana_dashboard_name.value;
  let grafana_stats_url =
    url +
    grafana_dashboard_stats.value +
    "/" +
    grafana_dashboard_stats_name.value;
  let params = {};
  params["orgId"] = 1;
  params["var-report"] = reportName.value + "_id_" + reportID;
  params["kiosk"] = "tv";
  {
    if (state.value < 5) {
      params["refresh"] = "30s";
      params["from"] = "now-5m";

      let timeIntervals = ["15m", "30m", "1h", "3h", "6h", "12h", "24h"];
      for (let i = 0; i < timeIntervals.length; i++) {
        let interval = timeIntervals[i];
        if (isTimeExpired(interval, CreatedAt.value)) {
          params["from"] = "now-" + interval;
        }
      }
    } else {
      let startDate = new Date(CreatedAt.value);
      startDate.setSeconds(0);

      let date = new Date(UpdatedAt.value);
      let currentMinute = date.getMinutes();
      let nextMinute = Math.ceil(currentMinute / 5) * 5;
      date.setMinutes(nextMinute);
      date.setSeconds(0);
      date.setMilliseconds(0);
      params["from"] = startDate.getTime();
      params["to"] = date.getTime();
    }
  }
  grafanaUrl.value =
    grafana_url +
    "?" +
    Object.keys(params)
      .map(function (key) {
        return encodeURIComponent(key) + "=" + encodeURIComponent(params[key]);
      })
      .join("&");

  grafanaStatsUrl.value =
    grafana_stats_url +
    "?" +
    Object.keys(params)
      .map(function (key) {
        return encodeURIComponent(key) + "=" + encodeURIComponent(params[key]);
      })
      .join("&");
};

const isTimeExpired = (timeInterval, CreatedAt) => {
  let currentDate = new Date();
  let date = new Date(CreatedAt);
  let interval;
  switch (timeInterval) {
    case "15m":
      interval = 5 * 60 * 1000;
      break;
    case "30m":
      interval = 15 * 60 * 1000;
      break;
    case "1h":
      interval = 30 * 60 * 1000;
      break;
    case "3h":
      interval = 60 * 60 * 1000;
      break;
    case "6h":
      interval = 3 * 60 * 60 * 1000;
      break;
    case "12h":
      interval = 6 * 60 * 60 * 1000;
      break;
    case "24h":
      interval = 12 * 60 * 60 * 1000;
      break;
    default:
      // 如果传入的时间间隔无效，可以抛出错误或返回默认值
      throw new Error("无效的时间间隔");
  }

  let timeDiff = currentDate - date;
  return timeDiff > interval;
};

const updateDetail = async () => {
  const res = await findReport({ ID: reportID, DetailID: detailId });
  if (res.code === 0) {
    respData.value = res.data.reapicase;
    state.value = res.data.reapicase.state;
    runState(res.data.reapicase.state);
  }
};

const resetBoomer = async () => {
  dialogRunner.value = true;
};

const closeRunner = () => {
  runnerConfig = {
    spawnCount: 1,
    spawnRate: 1,
  };
  dialogRunner.value = false;
};

const updateUser = async () => {
  let data = {
    caseID: performance_id,
    operation: {
      spawnCount: runnerConfig.spawnCount,
      spawnRate: runnerConfig.spawnRate,
    },
  };
  const res = await rebalance(data);
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "运行成功",
    });
  }
  closeRunner();
};

const stopBoomer = async () => {
  let params = {
    caseID: performance_id,
    reportID: reportID,
  };
  const res = await stopBoom(params);
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "操作成功",
    });
  }
};

const stateName = ref("");
const stateStyle = ref("");
const runState = (t) => {
  if (t === 1) {
    stateName.value = "\xa0准备中\xa0";
    stateStyle.value = "warning";
    return;
  }
  if (t === 2) {
    stateName.value = "\xa0运行中\xa0";
    stateStyle.value = "";
    return;
  }
  if (t === 3) {
    stateName.value = "\xa0运行中\xa0";
    stateStyle.value = "";
    return;
  }
  if (t === 4) {
    stateName.value = "\xa0停止中\xa0";
    stateStyle.value = "danger";
    return;
  }
  if (t === 5) {
    stateName.value = "\xa0已完成\xa0";
    stateStyle.value = "success";
    boomerButton.value = true;
    return;
  }
  stateName.value = "未知状态";
};

const initData = async () => {
  if (route.params.id > 0) {
    reportID = route.params.id;
  }
  await getTestCaseDetailFunc(reportID);
};
initData();
watch(
  () => route.params.id,
  () => {
    if (route.params.id) {
      removeData();
      initData();
    }
  }
);

// 定时刷新
let timerData = setInterval(function () {}, 10000000000);
const timerOptions = [
  {
    value: 5,
    label: "5s",
  },
  {
    value: 15,
    label: "15s",
  },
  {
    value: 30,
    label: "30s",
  },
  {
    value: 30,
    label: "60s",
  },
  {
    value: 0,
    label: "关闭自动刷新",
  },
];
const timerValue = ref("手动刷新");
const updateTimerData = (timer) => {
  clearInterval(timerData);
  if (timer === 0 || state.value === 5) {
    return;
  }
  timerData = setInterval(function () {
    if (state.value === 5) {
      clearInterval(timerData);
    }
    updateDetail();
  }, timer * 1000);
};
</script>

<style scoped>
h1 {
  text-align: center;
}

p {
  margin-top: 25px;
  margin-bottom: 5px;
  text-align: center;
  font-size: 16px;
}

.grafana-iframe {
  width: 100%;
  height: 100vh; /* 或者使用height: 100%; 来让iframe撑满父容器的高度 */
  border: none; /* 可选：如果不需要边框的话 */
}
</style>
