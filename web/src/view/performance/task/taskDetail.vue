<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" class="demo-form-inline">
        <el-form-item>
          <div class="block" :class="`block_head`" style="height: 30px">
            <span
              class="block-method block_method_color"
              :class="`block_method_head`"
            >
              {{ "CASE" }}
            </span>
            <div class="block"></div>
            <span class="block-method block_url">{{ taskName }}</span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addApiCaseFunc" round
            >添加测试用例</el-button
          >
          <el-button type="primary" @click="addTransaction" round
            >添加事务</el-button
          >
          <el-button type="primary" @click="addRendezvous" round
            >添加集合</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        :show-header="false"
        :data="tableData"
        row-key="ID"
        :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
      >
        <el-table-column width="50" type="index"> </el-table-column>
        <el-table-column min-width="600" align="center">
          <template #default="scope">
            <div class="block" :class="`block_${StepType(scope.row)[1]}`">
              <span
                class="block-method block_method_color"
                :class="`block_method_${StepType(scope.row)[1]}`"
                style="width: 60px"
              >
                {{ StepType(scope.row)[0] }}
              </span>
              <span class="block-method block_url">
                {{ scope.row.ApiCaseStep.name }}
                <a-tag
                  :key="StepType(scope.row)[2]"
                  :color="StepType(scope.row)[2]"
                  v-if="showTag(scope.row.ApiCaseStep.api_step_type)"
                  >{{
                    scope.row.ApiCaseStep.api_step_type == "TransactionStart"
                      ? "START"
                      : "END"
                  }}</a-tag
                >
              </span>
            </div>
          </template>
        </el-table-column>

        <el-table-column align="right" label="按钮组" min-width="150px">
          <template #default="scope">
            <el-button
              type="text"
              icon="document"
              size="mini"
              @click="detailRow(scope.row)"
              >详情</el-button
            >
            <el-button
              type="text"
              icon="delete"
              size="mini"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
      v-model="interfaceTempleFormVisible"
      :before-close="closeDialog"
      :visible.sync="interfaceTempleFormVisible"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="dialogTitle"
      width="1380px"
      top="30px"
    >
      <InterfaceTempleForm
        @close="closeDialog"
        v-if="interfaceTempleFormVisible"
        :heights="heightDiv"
        :eventType="type"
        :apiType="apiTypes"
        :formData="formDatas"
        ref="menuRole"
      >
      </InterfaceTempleForm>
    </el-dialog>

    <el-dialog
      v-model="taskCaseVisible"
      :before-close="closeDialogAddCase"
      :visible.sync="taskCaseVisible"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      title="添加测试套件"
      width="1250px"
      top="30px"
    >
      <ApisCaseAdd
        @close="closeDialogAddCase"
        v-if="taskCaseVisible"
        @caseID="addTeseCase"
        :types="performanceTypes"
        ref="menuRole"
      >
      </ApisCaseAdd>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="addTaskCases"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="taskTransaction"
      :before-close="closeTransaction"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="type === 'create' ? '添加事务' : '编辑事务'"
      width="500px"
    >
      <el-form
        :model="transactionData"
        label-position="right"
        label-width="80px"
      >
        <el-form-item label="事务类型:">
          <el-radio-group v-model="transactionData.type" class="ml-4">
            <el-radio label="start" size="large">开始事务</el-radio>
            <el-radio label="end" size="large">结束事务</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="任务名称:">
          <el-input
            v-model="transactionData.name"
            v-if="transactionData.type === 'start'"
            clearable
            placeholder="请输入"
          />
          <el-select
            v-model="transactionData.name"
            v-if="transactionData.type === 'end'"
            class="m-2"
            placeholder="请选择事务"
          >
            <el-option
              v-for="item in TransactionStart"
              :key="item.key"
              :label="item.key"
              :value="item.key"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeTransaction">取 消</el-button>
          <el-button size="small" type="primary" @click="addTransactions"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="taskRendezvous"
      :before-close="closeRendezvous"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :title="type === 'create' ? '添加集合点' : '编辑集合点'"
      width="800px"
    >
      <el-form
        :model="rendezvousData"
        label-position="right"
        label-width="120px"
      >
        <el-form-item label="集合点名称:">
          <el-input
            v-model="rendezvousData.name"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="集合点释放条件:">
          <el-radio-group v-model="rendezvousData.type" class="ml-4">
            <el-radio label="number" size="large">指定虚拟用户数</el-radio>
            <el-radio label="all" size="large">全部用户</el-radio>
            <el-radio label="percent" size="large">虚拟用户百分比</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item v-if="rendezvousData.type === 'number'" label="用户数量:">
          <el-input-number
            v-model="rendezvousData.number"
            :min="1"
            controls-position="right"
          />
        </el-form-item>
        <el-form-item
          v-if="rendezvousData.type === 'percent'"
          label="用户百分比:"
        >
          <el-input-number
            v-model="rendezvousData.percent"
            :precision="2"
            :min="0.01"
            :step="0.01"
            :max="1"
            controls-position="right"
          >
          </el-input-number>
        </el-form-item>

        <el-form-item label="超时时间:">
          <el-input-number
            v-model="rendezvousData.timeout"
            :min="1"
            :max="100"
            controls-position="right"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeRendezvous">取 消</el-button>
          <el-button size="small" type="primary" @click="addRendezvousClick"
            >确 定</el-button
          >
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { useRoute, useRouter } from "vue-router";
import { reactive, ref, onMounted, watch } from "vue";

import InterfaceTempleForm from "@/view/interface/interfaceTemplate/interfaceTemplateForm.vue";
import ApisCaseAdd from "@/view/interface/apiCase/apiCaseAddStep.vue";
import { findInterfaceTemplate } from "@/api/interfaceTemplate";
import { ElMessage, ElMessageBox } from "element-plus";
import Sortable from "sortablejs";
import {
  addOperation,
  addPerformanceCase,
  delPerformanceCase,
  findPerformance,
  findPerformanceCase,
  findPerformanceStep,
  sortPerformanceCase,
} from "@/api/performance";
import { runBoomerDebug } from "@/api/runTestCase";

const route = useRoute();
const task_id = ref();
const tableData = ref([]);
const apiTypes = 2;
const interfaceTempleFormVisible = ref(false);
const taskCaseVisible = ref(false);
const taskTransaction = ref(false);
const taskRendezvous = ref(false);
const dialogTitle = ref(false);
const type = ref("");
const heightDiv = ref();
const taskName = ref();
let caseID = [];
let sortIdList = "";
heightDiv.value = window.screen.height - 480;
const formDatas = reactive({});
const sortData = ref([]);
let performanceTypes = 2;
const transactionData = ref({
  name: "",
  type: "",
});
const router = useRouter();
const rendezvousData = ref({
  name: "",
  type: "number",
  percent: 0.01,
  number: 1,
  timeout: 0,
});

const showTag = (type) => {
  if (type === "TransactionStart" || type === "TransactionEnd") {
    return true;
  } else {
    return false;
  }
};

const init = () => {
  if (route.params.id > 0) {
    task_id.value = Number(route.params.id);
  } else {
    task_id.value = 6;
  }
  if (task_id.value) {
    getTaskCaseDetailFunc(task_id.value);
  }
};
const TransactionStart = ref([]);
const getTaskCaseDetailFunc = async (task_id) => {
  const res = await findPerformanceCase({ ID: task_id });
  if (res.code === 0) {
    tableData.value = res.data.reapicase;
    taskName.value = res.data.name;
    TransactionStart.value = [];
    tableData.value.forEach((item, index) => {
      if (item.ApiCaseStep.api_step_type === "TransactionStart") {
        let step = { key: item.ApiCaseStep.name };
        TransactionStart.value.push(step);
      }
      if (item.ApiCaseStep.api_step_type === "TransactionEnd") {
        TransactionStart.value.forEach(function (item2, index, arr) {
          if (item2.key === item.ApiCaseStep.name) {
            arr.splice(index, 1);
          }
        });
      }
    });
  }
};
init();

const StepType = (row) => {
  let step_type = [];

  if (row.ApiCaseStep.api_step_type === "Transaction") {
    step_type.push("事务");
    step_type.push("patch");
    return step_type;
  }
  if (row.ApiCaseStep.api_step_type === "Rendezvous") {
    step_type.push("集合点");
    step_type.push("delete");
    return step_type;
  }
  if (row.ApiCaseStep.api_step_type === "TransactionStart") {
    step_type.push("事务");
    step_type.push("options");
    step_type.push("#168cff");
    return step_type;
  }
  if (row.ApiCaseStep.api_step_type === "TransactionEnd") {
    step_type.push("事务");
    step_type.push("head");
    step_type.push("#ff5722");
    return step_type;
  }
  step_type.push("STEP");
  step_type.push("patch");
  return step_type;
};

const addApiCaseFunc = async () => {
  taskCaseVisible.value = true;
};

const addTransaction = async () => {
  taskTransaction.value = true;
  type.value = "create";
  transactionData.value.type = "start";
};

const transaction = ref([]);

const addRendezvous = async () => {
  taskRendezvous.value = true;
  type.value = "create";
};

const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const res = await delPerformanceCase({ ID: row.ID });
    if (res.code === 0) {
      await getTaskCaseDetailFunc(task_id.value);
      ElMessage({
        type: "success",
        message: "用例删除成功",
      });
    }
  });
};

const detailRow = async (row) => {
  if (row.ApiCaseStep.api_step_type) {
    const res = await findPerformanceStep({ ID: row.ApiCaseStep.ID });
    if (res.code === 0) {
    }
  } else {
    await router.push({
      name: "testCaseStepDetail",
      params: {
        id: row.ApiCaseStepId,
      },
    });
  }
};
onMounted(() => {
  rowDrop();
});

//行拖拽
const rowDrop = async () => {
  sortData.value.ID = Number(task_id.value);
  const tbody = document.querySelector(".el-table__body-wrapper tbody");
  let taskCases = [];

  Sortable.create(tbody, {
    animation: 1000,
    async onEnd({ newIndex, oldIndex }) {
      const currRow = tableData.value.splice(oldIndex, 1)[0];
      tableData.value.splice(newIndex, 0, currRow);
      tableData.value.forEach((item, index, arr) => {
        let tStep = { ID: item.ID, sort: index + 1 };
        taskCases.push(tStep);
      });
      sortData.value = taskCases;
      const res = await sortPerformanceCase(taskCases);
      if (res.code === 0) {
        taskCases = [];
        await getTaskCaseDetailFunc(task_id.value);
        ElMessage({
          type: "success",
          message: "已完成排序",
        });
      }
    },
  });
};

// 关闭弹窗
const closeDialog = () => {
  taskCaseVisible.value = false;
  getTaskCaseDetailFunc(task_id.value);
};

const closeTransaction = () => {
  taskTransaction.value = false;
  transactionData.value.name = "";
  getTaskCaseDetailFunc(task_id.value);
};

const closeRendezvous = () => {
  taskRendezvous.value = false;
  rendezvousData.value.name = "";
  rendezvousData.value.type = "number";
  rendezvousData.value.percent = 0.01;
  rendezvousData.value.number = 1;
  rendezvousData.value.timeout = 0;
  getTaskCaseDetailFunc(task_id.value);
};

const addTaskCases = async () => {
  const res = await addPerformanceCase({
    task_id: task_id.value,
    case_id: caseID,
  });
  if (res.code === 0) {
    closeDialogAddCase();
    ElMessage({
      type: "success",
      message: "添加用例成功",
    });
  }
};

const addTransactions = async () => {
  if (transactionData.value.type === "") {
    ElMessage({
      type: "error",
      message: "请选择事务类型",
    });
    return;
  }
  if (transactionData.value.name === "") {
    ElMessage({
      type: "error",
      message: "事务名称不能为空",
    });
    return;
  }
  const step = ref({
    api_step: {
      name: transactionData.value.name,
      transaction: {
        name: transactionData.value.name,
        type: transactionData.value.type,
      },
    },
    pid: task_id.value,
  });
  const res = await addOperation(step.value);
  if (res.code === 0) {
    closeDialogAddCase();
    ElMessage({
      type: "success",
      message: "添加事务成功",
    });
    transactionData.value.name = "";
    closeTransaction();
  }
};

const addRendezvousClick = async () => {
  if (rendezvousData.value.name === "") {
    ElMessage({
      type: "error",
      message: "集合点名称不能为空",
    });
    return;
  }
  let number = 0;
  let percent = 1.0;
  if (rendezvousData.value.type === "percent") {
    number = 0;
    percent = rendezvousData.value.percent;
  }
  if (rendezvousData.value.type === "all") {
    number = 0;
    percent = 1.0;
  }
  if (rendezvousData.value.type === "number") {
    percent = 0;
    number = rendezvousData.value.number;
  }
  const step = ref({
    api_step: {
      name: rendezvousData.value.name,
      rendezvous: {
        name: rendezvousData.value.name,
        type: rendezvousData.value.type,
        timeout: rendezvousData.value.timeout,
        number: number,
        percent: percent,
      },
    },
    pid: task_id.value,
  });
  const res = await addOperation(step.value);
  if (res.code === 0) {
    closeDialogAddCase();
    ElMessage({
      type: "success",
      message: "添加事务成功",
    });
    transactionData.value.name = "";
    closeRendezvous();
  }
};

const closeDialogAddCase = () => {
  taskCaseVisible.value = false;
  getTaskCaseDetailFunc(task_id.value);
};

const addTeseCase = (caseIDs) => {
  caseID = caseIDs;
};
</script>

<style lang="scss" scoped>
@import "src/style/apiList";
</style>
