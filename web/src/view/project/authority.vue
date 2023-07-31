<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="项目成员">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <!--        <el-form-item label="管理员">-->
        <!--          <el-input v-model="searchInfo.admin" placeholder="搜索条件" />-->
        <!--        </el-form-item>-->
        <!--        <el-form-item label="创建人">-->
        <!--          <el-input v-model="searchInfo.creator" placeholder="搜索条件" />-->
        <!--        </el-form-item>-->
        <!--        <el-form-item label="项目描述">-->
        <!--          <el-input v-model="searchInfo.describe" placeholder="搜索条件" />-->
        <!--        </el-form-item>-->
        <el-form-item>
          <el-button icon="search" size="mini" type="primary" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" size="mini" @click="onReset"
            >重置</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="SysUserID"
        style="width: 100%"
        tooltip-effect="dark"
      >
        <el-table-column label="项目成员" width="300">
          <template #default="scope">
            <span>{{ scope.row.username }}</span>
          </template>
        </el-table-column>
        <el-table-column label="权限" width="300">
          <template #default="scope">
            <el-checkbox-group>
              <el-checkbox label="select">查询</el-checkbox>
              <el-checkbox label="save">新增/修改</el-checkbox>
              <el-checkbox label="delete">删除</el-checkbox>
            </el-checkbox-group>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { getProjectUserList, setUserProjectAuth } from "@/api/project";

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});
const checkboxGroup = ref([]);

const checkboxGroupFunc = (row) => {
  return checkboxGroup.value;
};

const checkboxValue = (row, key) => {
  return row[key];
};

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

const getTableData = async () => {
  let project = JSON.parse(window.localStorage.getItem("project")).ID;
  const table = await getProjectUserList({
    projectId: project,
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
  console.log("tableData.value", tableData.value);
  console.log("table.data.list", table.data.list);
};
getTableData();
</script>

<style scoped></style>
