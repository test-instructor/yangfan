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
  console.log(last);

  try {
    navigator.clipboard.writeText(last);
  } catch (error) {
    let element = createElement(last);
    element.select();
    element.setSelectionRange(0, element.value.length);
    document.execCommand("copy");
    element.remove();
    alert("已复制到剪切板");
  }
};

const createElement = (text) => {
  let isRTL = document.documentElement.getAttribute("dir") === "rtl";
  let element = document.createElement("textarea");
  // 防止在ios中产生缩放效果
  element.style.fontSize = "12pt";
  // 重置盒模型
  element.style.border = "0";
  element.style.padding = "0";
  element.style.margin = "0";
  // 将元素移到屏幕外
  element.style.position = "absolute";
  element.style[isRTL ? "right" : "left"] = "-9999px";
  // 移动元素到页面底部
  let yPosition = window.pageYOffset || document.documentElement.scrollTop;
  element.style.top = `${yPosition}px`;
  //设置元素只读
  element.setAttribute("readonly", "");
  element.value = text;
  document.body.appendChild(element);
  return element;
};

initData();
</script>

<style scoped></style>
