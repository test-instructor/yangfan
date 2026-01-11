<template>
  <el-table
    highlight-current-row
    :cell-style="{ paddingTop: '4px', paddingBottom: '4px' }"
    stripe
    :data="formData"
    :height="height"
    style="width: 100%"
    @cell-mouse-enter="cellMouseEnter"
    @cell-mouse-leave="cellMouseLeave"
  >
    <!-- Request params -->
    <el-table-column label="请求Key" width="220">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.key"
          placeholder="Key"
          @input="fromDatas"
        ></el-input>
      </template>
    </el-table-column>

    <!-- 类型选择 -->
    <el-table-column label="类型" width="120">
      <template #default="scope">
        <el-select v-model="scope.row.type" @change="fromDatas">
          <el-option
            v-for="item in dataTypeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </template>
    </el-table-column>

    <!-- 请求值 -->
    <el-table-column label="请求Value" width="320">
      <template #default="scope">
        <el-input
          v-model="scope.row.value"
          placeholder="Value"
          @input="fromDatas"
        ></el-input>
      </template>
    </el-table-column>

    <!-- 描述 -->
    <el-table-column label="描述" width="300">
      <template #default="scope">
        <el-input
          clearable
          v-model="scope.row.desc"
          placeholder="参数简要描述"
          @input="fromDatas"
        ></el-input>
      </template>
    </el-table-column>

    <!-- 操作按钮 -->
    <el-table-column>
      <template #default="scope">
        <el-row>
          <el-button
            size="small"
            type="info"
            @click="handleEdit"
            icon="add"
          />
          <el-button
            size="small"
            type="info"
            @click="handleCopy(scope.$index, scope.row)"
            icon="copy"
          />
          <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
            icon="delete"
          />
        </el-row>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup>
  import { ref, computed, onMounted, onUnmounted, watch } from "vue";
  import { ElNotification } from "element-plus";
  import { getDict } from '@/utils/dictionary.js'

  // 定义props
  const props = defineProps({
    forms: {
      type: Array,
      default: () => []
    },
    heights: {
      type: Number,
      default: 400
    }
  });

  // 定义emits
  const emit = defineEmits(['requestFormData', 'request']);

  // 响应式数据
  const formData = ref([
    {
      key: "",
      value: "",
      type: "String", // 默认字符串类型
      desc: "",
    }
  ]);

  const currentRow = ref("");
  const contentStyleObj = ref({
    height: "",
    width: ""
  });
  const fileList = ref([]);
  const currentIndex = ref(0);
  const timeStamp = ref("");

  // 数据类型选项
  const dataTypeOptions = ref([])

  const getDataTypeOptions = async () => {
    try {
      const res = await getDict('fieldType')
      if (res && Array.isArray(res)) {
        // 简化数据映射，避免冗余变量
        dataTypeOptions.value = res.map(item => ({
          label: item.label,
          value: item.value
        }))
      }
    } catch (error) {
      console.error('获取变量类型选项失败：', error)
      // 可在此处添加错误提示（如ElMessage）
    }
  }

  // 计算属性
  const height = computed(() => props.heights - 70);

  // 方法定义
  const getHeight = () => {
    contentStyleObj.value.height = height.value + "px";
    contentStyleObj.value.width = "98%";
  };

  const fromDatas = () => {
    const emitdata = formData.value.filter(item =>
      item.desc !== "" || item.key !== "" || item.value !== ""
    );
    emit("requestFormData", emitdata);
  };

  const cellMouseEnter = (row) => {
    currentRow.value = row;
  };

  const cellMouseLeave = () => {
    currentRow.value = "";
  };

  const handleEdit = () => {
    formData.value.push({
      key: "",
      value: "",
      type: "String",
      desc: "",
    });
  };

  const handleCopy = (index, row) => {
    formData.value.splice(index + 1, 0, {
      key: row.key,
      value: row.value,
      type: row.type,
      desc: row.desc,
    });
    fromDatas();
  };

  const handleDelete = (index, row) => {
    emit("request", row.ID);
    formData.value.splice(index, 1);
    fromDatas();
    if (formData.value.length === 0) {
      handleEdit();
    }
  };

  const parseJson = () => {
    let json = {};
    // 注意：原代码中jsonData未定义，这里保持原逻辑
    if (this.jsonData !== "") {
      try {
        json = JSON.parse(this.jsonData);
      } catch (err) {
        ElNotification.error({
          title: "json错误",
          message: "不是标准的json数据格式",
          duration: 2000,
        });
      }
    }
    return json;
  };


  // 生命周期
  onMounted(() => {
    getDataTypeOptions()
    window.addEventListener("resize", getHeight);
    getHeight();

    // 监听初始表单数据
    if (props.forms && props.forms.length) {
      formData.value = [...props.forms];
      fromDatas();
    }
  });

  onUnmounted(() => {
    window.removeEventListener("resize", getHeight);
  });

  // 监听表单数据变化
  watch(
    () => props.forms,
    (newVal) => {
      if (newVal && newVal.length) {
        formData.value = [...newVal];
        fromDatas();
      }
    },
    { deep: true }
  );
  watch(
    () => formData.value,
    (newVal) => {
      // console.log("formData.value======", newVal);
      // emit("formDataJson", newVal);
    },
    { deep: true }
  );
</script>

<style scoped></style>
