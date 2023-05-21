<template>
  <div style="padding: 2px">
    <el-table border v-show="isTable" :data="tableDatas">
      <el-table-column width="180" align="center" prop="key" label="key">
      </el-table-column>
      <el-table-column width="80" align="center" prop="key" label="操作">
        <template v-slot="scope">
          <el-button type="text" @click="copy(scope.row)">复制</el-button>
        </template>
      </el-table-column>
      <el-table-column align="center" label="value">
        <template #default="scope">
          <span>{{ scope.row.value }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "tableKeyValue",
};
</script>

<script setup>
import { ref } from "vue";

const tableDatas = ref();
const isTable = ref(false);

const props = defineProps({
  tableData: ref(),
});

const tableKeyToValue = async (data) => {
  let tableData = [];
  for (let k in data) {
    let tableJson = { key: k, value: data[k] };
    tableData.push(tableJson);
  }
  tableDatas.value = tableData;
  if (tableDatas.value.length > 0) {
    isTable.value = true;
  }
};

const initData = () => {
  tableKeyToValue(props.tableData);
};

const copy = (row) => {
  let last = JSON.stringify(row);
  navigator.clipboard.writeText(last);
};

initData();
</script>

<style scoped></style>
