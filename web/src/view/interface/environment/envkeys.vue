<template>
  <div>
    <span>测试页面</span>
    {{ keys }}
  </div>
</template>

<script>
export default {
  name: "envkeys",
};
</script>

<script setup>
import { getEnvVariableList } from "@/api/env";
var keys = ref([]);
const getTableData = async () => {
  let table = await getEnvVariableList({
    page: 1,
    pageSize: 9999,
    show_key: true,
  });
  if (table.code === 0) {
  }
};
</script>

<style scoped></style>
