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
            <el-checkbox-group
              v-model="scope.row.status"
              @change="handleChecked(scope.row)"
            >
              <el-checkbox label="select" disabled>查询</el-checkbox>
              <el-checkbox label="save">新增/修改</el-checkbox>
              <el-checkbox label="delete">删除</el-checkbox>
            </el-checkbox-group>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteUser(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import {
  deleteUserProjectAuth,
  getProjectUserList,
  setUserProjectAuth,
} from "@/api/project";
import { ElMessage, ElMessageBox } from "element-plus";

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
    tableData.value.forEach((item) => {
      item.status = [];
      item.status.push("select");
      if (item.save) {
        item.status.push("save");
      }
      if (item.delete) {
        item.status.push("delete");
      }
    });
  }
  console.log("tableData.value", tableData.value);
  console.log("table.data.list", table.data.list);
};
getTableData();

const handleChecked = async (row) => {
  let params = {
    SysUserID: row.SysUserID,
    ProjectID: row.ProjectID,
    select: true,
    delete: false,
    save: false,
  };
  row.status.forEach((item) => {
    params[item] = true;
  });
  console.log("params", params);
  const res = await setUserProjectAuth(params);
  if (res.code === 0) {
    ElMessage.success("设置成功");
    getTableData();
  }
};

const deleteUser = (row) => {
  ElMessageBox.confirm("此操作将永久删除该该项目成员, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(async () => {
      let deleteData = {
        sys_user_id: row.SysUserID,
        project_id: row.ProjectID,
      };
      console.log("row", row);
      console.log("deleteData", deleteData);
      const res = await deleteUserProjectAuth(deleteData);
      if (res.code === 0) {
        ElMessage({
          type: "success",
          message: "删除成功!",
        });
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--;
        }
        await getTableData();
      }
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "已取消删除",
      });
    });
};
</script>

<style scoped></style>
