<template>
  <div>
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        ref="searchForm"
        class="demo-form-inline"
      >
        <el-form-item>
          <el-input v-model="searchInfo.name" placeholder="名称" size="mini" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="searchInfo.type">
            <el-option
              v-for="item in typeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="仅失败通知" prop="fail">
          <el-switch
            v-model="searchInfo.fail"
            :active-value="true"
            :inactive-value="false"
          />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button size="mini" icon="refresh" @click="onReset"
            >重置</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
      </div>
      <el-table
        ref="multipleTable"
        row-key="ID"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
      >
        <el-table-column label="名称" prop="name" width="150" />
        <el-table-column label="通知类型" width="150">
          <template #default="scope">
            <span>{{ getMessageType(scope.row.type) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="Webhook" width="240">
          <template #default="scope">
            <el-tooltip
              :content="scope.row.webhook"
              raw-content
              placement="top-start"
              v-if="scope.row.webhook"
            >
              <span v-if="scope.row.webhook && scope.row.webhook.length <= 26">
                {{ scope.row.webhook }}
              </span>
              <span v-if="scope.row.webhook && scope.row.webhook.length > 26">
                {{ scope.row.webhook.substr(0, 26) + "..." }}
              </span>
            </el-tooltip>
            <span v-else-if="scope.row.webhook == null"> NA </span>
          </template>
        </el-table-column>
        <el-table-column label="签名" prop="signature" width="240">
          <template #default="scope">
            <el-tooltip
              :content="scope.row.signature"
              raw-content
              placement="top-start"
              v-if="scope.row.signature"
            >
              <span
                v-if="scope.row.signature && scope.row.signature.length <= 26"
              >
                {{ scope.row.signature }}
              </span>
              <span
                v-if="scope.row.signature && scope.row.signature.length > 26"
              >
                {{ scope.row.signature.substr(0, 26) + "..." }}
              </span>
            </el-tooltip>
            <span v-else-if="scope.row.signature == null"> NA </span>
          </template>
        </el-table-column>
        <el-table-column label="仅失败通知" prop="fail" width="150">
          <template #default="scope">
            <el-switch
              v-model="scope.row.fail"
              :active-value="true"
              :inactive-value="false"
              disabled
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组" width="160">
          <template #default="scope">
            <el-button
              class="table-button"
              icon="edit"
              size="small"
              type="text"
              @click="updateRow(scope.row)"
              >变更
            </el-button>
            <el-button
              icon="delete"
              size="mini"
              type="text"
              @click="deleteRow(scope.row)"
              >删除</el-button
            ></template
          >
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
      :title="dialogTitle"
      v-model="dialogFormVisible"
      width="50%"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-form :model="dialogForm" label-width="100px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="dialogForm.name" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="dialogForm.type">
            <el-option
              v-for="item in typeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Webhook" prop="webhook">
          <el-input v-model="dialogForm.webhook" />
        </el-form-item>
        <el-form-item label="签名" prop="signature">
          <el-input v-model="dialogForm.signature" />
        </el-form-item>
        <el-form-item label="仅失败通知" prop="fail">
          <el-switch
            v-model="dialogForm.fail"
            :active-value="true"
            :inactive-value="false"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm">提交</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from "vue";
import {
  getMessageList,
  createMessage,
  updateMessage,
  deleteMessage,
  findMessage,
} from "@/api/message";
import { ElMessage, ElMessageBox } from "element-plus";

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

const dialogForm = ref({
  name: "",
  type: "",
  webhook: "",
  signature: "",
  fail: false,
});

const typeList = ref([
  { label: "钉钉", value: "dingtalk" },
  { label: "飞书", value: "feishu" },
]);

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 弹窗控制标记
const dialogFormVisible = ref(false);
const dialogTitle = ref("");

const getTableData = async () => {
  const table = await getMessageList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  if (searchInfo.value.default === "") {
    searchInfo.value.default = null;
  }
  getTableData();
};

const resetForm = () => {
  dialogForm.value = {
    name: "",
    type: "",
    webhook: "",
    signature: "",
    fail: false,
  };
};

const init = () => {
  getTableData();
};
init();
const openDialog = () => {
  type.value = "create";
  dialogTitle.value = "新增消息";
  dialogFormVisible.value = true;
};

const onReset = () => {
  searchInfo.value = {};
};

const submitForm = async () => {
  console.log("dialogForm", dialogForm.value);
  if (dialogForm.value.name === "") {
    ElMessage.warning("名称不能为空");
    return;
  }
  if (dialogForm.value.type === "") {
    ElMessage.warning("类型不能为空");
    return;
  }
  if (dialogForm.value.webhook === "") {
    ElMessage.warning("Webhook不能为空");
    return;
  }
  let res = await createMessage(dialogForm.value);
  if (res.code === 0) {
    ElMessage.success("新增成功");
    dialogFormVisible.value = false;
    getTableData();
  }
};
const updateRow = (row) => {};

const deleteRow = (row) => {
  ElMessageBox.confirm(
    "删除消息类型，会导致已引用该消息内容的任务无法正常发送消息，确定要删除吗？",
    "提示",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "error",
    }
  ).then(() => {
    deleteMsgRow(row);
  });
};

const deleteMsgRow = async (row) => {
  let res = await deleteMessage({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "变量消息类型成功",
    });
    await getTableData();
  }
};

const getMessageType = (type) => {
  let res = typeList.value.find((item) => item.value === type);
  return res ? res.label : "";
};
</script>

<style scoped></style>
