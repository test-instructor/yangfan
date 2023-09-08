<template>
  <div>
    <div class="gva-search-box">
      <el-form class="demo-form-inline">
        <el-form-item>
          <el-button size="mini" @click="onAddItem" type="primary">{{
            "添加环境"
          }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table
        highlight-current-row
        :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
        strpe
        :data="envTableData"
        @cell-mouse-enter="cellMouseEnter"
        @cell-mouse-leave="cellMouseLeave"
        max-height="500"
      >
        <el-table-column label="环境名称" width="300">
          <template #default="scope">
            <el-input
              v-model="scope.row.name"
              placeholder="请输入环境名称"
              :disabled="!showBtn(scope.row)"
            ></el-input>
          </template>
        </el-table-column>

        <el-table-column label="备注信息">
          <template #default="scope">
            <el-input
              v-model="scope.row.remarks"
              placeholder="请输入备注信息"
              :disabled="!showBtn(scope.row)"
            ></el-input>
          </template>
        </el-table-column>

        <el-table-column label="操作" mini-width="100">
          <template #default="scope">
            <el-button
              class="mt-4"
              @click="editEnv(scope.row)"
              v-if="!showBtn(scope.row)"
              >{{ "编辑" }}</el-button
            >
            <el-button
              class="mt-4"
              @click="cancelEnv(scope.row)"
              v-if="showCancelBtn(scope.row)"
              >{{ "取消" }}</el-button
            >
            <el-button
              class="mt-4"
              @click="saveEnv(scope.$index, scope.row)"
              v-if="showBtn(scope.row)"
              >{{ "保存" }}</el-button
            >
            <el-button
              class="mt-4"
              @click="deleteEnvs(scope.$index, scope.row)"
              >{{ "删除" }}</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  name: "Env",
};
</script>

<script setup>
import { createEnv, deleteEnv, getEnvList } from "@/api/env";
import { ref } from "vue";
import { ElMessage } from "element-plus";
const showTable = ref(false);
const envTableData = ref([]);
const currentRow = ref({});
const editEnvRow = ref(0);

const init = async () => {
  const tables = await getEnvList();
  envTableData.value = tables.data.list;
};

init();

const onAddItem = () => {
  envTableData.value.push({
    ID: 0,
    name: "",
    remarks: "",
  });
};
const shouAddBtn = ref(true);
const showBtn = (row) => {
  return editEnvRow.value === row.ID;
};

const showCancelBtn = (row) => {
  return row.ID && row.ID > 0 && editEnvRow.value === row.ID;
};

const editEnv = (row) => {
  editEnvRow.value = row.ID;
};

const cancelEnv = (row) => {
  editEnvRow.value = 0;
};

const saveEnv = async (index, row) => {
  if (!(row.name && row !== "")) {
    ElMessage({
      type: "error",
      message: "环境名称不能为空",
    });
    return;
  }
  let res = await createEnv(row);
  if (res.code === 0) {
    let message = "变量【" + row.name + "】创建成功";
    if (row.ID && row.ID > 0) {
      message = "变量【" + row.name + "】编辑成功";
    }
    ElMessage({
      type: "success",
      message: message,
    });
    envTableData.value = res.data.list;
  }
  editEnvRow.value = 0;
};

const deleteEnvs = async (index, row) => {
  if (row.ID && row.ID > 0) {
    let res = await deleteEnv({ ID: row.ID });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "标签【" + row.name + "】删除成功",
      });
      {
        envTableData.value.splice(index, 1);
      }
    } else {
      ElMessage({
        type: "error",
        message: "标签【" + row.name + "】删除失败",
      });
    }
  } else {
    envTableData.value.splice(index, 1);
  }
};

const cellMouseEnter = (row) => {
  currentRow.value = row;
};

const cellMouseLeave = (row) => {
  currentRow.value = {};
};
</script>

<style scoped></style>
