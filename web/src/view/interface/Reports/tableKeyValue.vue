<template>
  <div>
    <el-table
        border
        v-show="isTable"
        :data="tableDatas"
        style="padding: 5px"
        :row-style="{height:'10px'}"
    >
      <el-table-column
          width="148"
          align="center"
          prop="key"
          label="key"
      >
      </el-table-column>
      <el-table-column
          align="center"
          prop="value"
          label="value"
      >
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "tableKeyValue"
}
</script>

<script setup>

import {ref} from "vue";

const tableDatas = ref()
const isTable = ref(false)

const props = defineProps({
  tableData: ref(),
})

const tableKeyToValue = async (data) => {
  let tableData = []
  for (let k in data) {
    let tableJson = {key: k, value: data[k]}
    tableData.push(tableJson)
  }
  tableDatas.value = tableData
  if (tableDatas.value.length > 0) {
    isTable.value = true
  }
}

const initData = () => {
  tableKeyToValue(props.tableData)
}
initData()

</script>

<style scoped>

</style>
