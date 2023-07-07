<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    strpe
    :data="tagTableData"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
    max-height="500"
  >
    <el-table-column label="标签名称" width="300">
      <template #default="scope">
        <el-input
          v-model="scope.row.name"
          placeholder="请输入标签名称"
          :disabled="!showBtn(scope.row)"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="备注信息">
      <template #default="scope">
        <el-input
          v-model="scope.row.remarks"
          placeholder="请输入标签备注"
          :disabled="!showBtn(scope.row)"
        ></el-input>
      </template>
    </el-table-column>

    <el-table-column label="操作" mini-width="100">
      <template #default="scope">
        <el-button
          class="mt-4"
          @click="editTag(scope.row)"
          v-if="!showBtn(scope.row)"
          >{{ "编辑" }}</el-button
        >
        <el-button
          class="mt-4"
          @click="cancelTag(scope.row)"
          v-if="showCancelBtn(scope.row)"
          >{{ "取消" }}</el-button
        >
        <el-button
          class="mt-4"
          @click="saveTag(scope.$index, scope.row)"
          v-if="showBtn(scope.row)"
          >{{ "保存" }}</el-button
        >
        <el-button class="mt-4" @click="deleteTag(scope.$index, scope.row)">{{
          "删除"
        }}</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-button class="mt-4" style="width: 100%" @click="onAddItem">{{
    "添加标签"
  }}</el-button>
</template>

<script>
export default {
  name: "TimerTaskTag",
};
</script>

<script setup>
import {
  deleteTimerTaskTag,
  createTimerTaskTag,
  getTimerTaskTagList,
} from "@/api/timerTask";
import { ref } from "vue";
import { ElMessage } from "element-plus";
const showTable = ref(false);
const tagTableData = ref([]);
const currentRow = ref({});
const editTagRow = ref(0);

const init = async () => {
  const tables = await getTimerTaskTagList();
  tagTableData.value = tables.data.list;
};

init();

const onAddItem = () => {
  tagTableData.value.push({
    ID: 0,
    name: "",
    remarks: "",
  });
};

const showBtn = (row) => {
  return editTagRow.value === row.ID;
};

const showCancelBtn = (row) => {
  return row.ID && row.ID > 0 && editTagRow.value === row.ID;
};

const editTag = (row) => {
  editTagRow.value = row.ID;
};

const cancelTag = (row) => {
  editTagRow.value = 0;
};

const saveTag = async (index, row) => {
  if (!(row.name && row !== "")) {
    ElMessage({
      type: "error",
      message: "标签名称不能为空",
    });
    return;
  }
  let res = await createTimerTaskTag(row);
  if (res.code === 0) {
    let message = "标签【" + row.name + "】创建成功";
    if (row.ID && row.ID > 0) {
      message = "标签【" + row.name + "】编辑成功";
    }
    ElMessage({
      type: "success",
      message: message,
    });
    tagTableData.value = res.data;
  }
  editTagRow.value = 0;
};

const deleteTag = async (index, row) => {
  if (row.ID && row.ID > 0) {
    let res = await deleteTimerTaskTag({ ID: row.ID });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "标签【" + row.name + "】删除成功",
      });
      {
        tagTableData.value.splice(index, 1);
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
