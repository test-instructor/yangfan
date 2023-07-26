<template>
  <div class="env">
    <el-table :data="envTableData">
      <el-table-column prop="key" label="变量key" width="240"></el-table-column>
      <el-table-column prop="name" label="变量名" width="160"></el-table-column>
      <el-table-column align="left" label="按钮组" width="160">
        <template #default="scope">
          <el-button
            class="table-button"
            icon="copy-document"
            size="small"
            type="text"
            @click="copyVariable(scope.row)"
            >复制变量
          </el-button>
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
</template>

<script setup>
import { getEnvVariableList } from "@/api/env";
import { ref } from "vue";
import { ElMessage } from "element-plus";
const envTableData = ref([]);
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const searchInfo = ref({});
const init = async () => {
  const table = await getEnvVariableList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    envTableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};
init();

const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

const copyVariable = (row) => {
  console.log(row.key);
  navigator.clipboard
    .writeText(row.key)
    .then(() => {
      ElMessage.success("复制成功");
    })
    .catch((err) => {
      ElMessage.error("复制失败");
    });
};
</script>

<style scoped></style>
