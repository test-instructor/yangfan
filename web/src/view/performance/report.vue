<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="报告名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
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
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="mini" type="text" @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button size="mini" type="primary" @click="onDelete"
              >确定</el-button
            >
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="mini"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              >删除
            </el-button>
          </template>
        </el-popover>
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
        <el-table-column align="left" label="执行时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="报告名称"
          prop="name"
          width="200"
        />
        <el-table-column align="left" label="执行状态" width="120">
          <template #default="scope">
            <el-tag
              :key="runState(scope.row.state)[0]"
              :type="runState(scope.row.state)[1]"
              >{{ runState(scope.row.state)[0] }}</el-tag
            >
          </template>
        </el-table-column>

        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              v-if="runState(scope.row.state)[2]"
              type="text"
              icon="document"
              size="small"
              class="table-button"
              @click="reportDetailFunc(scope.row)"
              >详情</el-button
            >
            <el-button
              v-if="!runState(scope.row.state)[2]"
              type="text"
              icon="document"
              size="small"
              class="table-button"
              @click="reportDetailFunc(scope.row)"
              disabled
              >详情</el-button
            >
            <!--            <el-button type="text" size="small" class="table-button" @click="reportDetailFunc(scope.row)">详情</el-button>-->
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
  </div>
</template>

<script>
export default {
  name: "pReport",
};
</script>

<script setup>
import { ref } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { deleteApiConfigByIds } from "@/api/apiConfig";
import {
  deletePerformance,
  getReportList,
  deleteReport,
} from "@/api/performance";
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { useRouter } from "vue-router";
import { delApisCase } from "@/api/apiCase";

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

const router = useRouter();

const runState = (t) => {
  if (t === 1) {
    return ["\xa0准备中\xa0", "warning", false];
  }
  if (t === 2) {
    return ["\xa0运行中\xa0", "", true];
  }
  if (t === 3) {
    return ["\xa0运行中\xa0", "", true];
  }
  if (t === 4) {
    return ["\xa0停止中\xa0", "danger", true];
  }
  if (t === 5) {
    return ["\xa0已完成\xa0", "success", true];
  }
  return ["未知状态"];
};

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  // const res = await deleteApiConfigByIds({ ids });
  // if (res.code === 0) {
  //   ElMessage({
  //     type: "success",
  //     message: "删除成功",
  //   });
  //   if (tableData.value.length === ids.length && page.value > 1) {
  //     page.value--;
  //   }
  //   deleteVisible.value = false;
  //   await getTableData();
  // }
};

const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const res = await deleteReport({ ID: row.ID });
    if (res.code === 0) {
      await getTableData();
      ElMessage({
        type: "success",
        message: "用例删除成功",
      });
    }
  });
};

const reportDetailFunc = (row) => {
  router.push({
    name: "pReportDetail",
    params: {
      id: row.ID,
    },
  });
};

// 查询
const getTableData = async () => {
  const table = await getReportList({
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

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  if (searchInfo.value.default === "") {
    searchInfo.value.default = null;
  }
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

getTableData();
</script>

<style scoped></style>
