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
          <el-input
            v-model="searchInfo.name"
            placeholder="名称"
            clearable
            size="mini"
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
    </div>
    <el-dialog
      :title="dialogTitle"
      :visible.sync="dialogFormVisible"
      width="50%"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
    >
      <el-form
        :model="dialogForm"
        ref="dialogForm"
        :rules="rules"
        label-width="240px"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="dialogForm.name" clearable />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="dialogForm.type" clearable>
            <el-option
              v-for="item in typeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Webhook" prop="webhook">
          <el-input v-model="dialogForm.webhook" clearable />
        </el-form-item>
        <el-form-item label="签名" prop="signature">
          <el-input v-model="dialogForm.signature" clearable />
        </el-form-item>
        <el-form-item label="失败通知" prop="fail">
          <el-switch
            v-model="dialogForm.fail"
            active-value="1"
            inactive-value="0"
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

const openDialog = () => {
  type.value = "create";
  dialogTitle.value = "新增消息";
  dialogFormVisible.value = true;
};
</script>

<style scoped></style>
