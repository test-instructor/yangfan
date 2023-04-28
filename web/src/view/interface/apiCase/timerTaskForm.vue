<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="任务名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="时间配置:">
          <el-input v-model="formData.runTime" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下次执行时间:">
          <el-date-picker
            v-model="formData.nextRunTime"
            clearable
            placeholder="选择日期"
            type="date"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="状态:">
          <el-switch
            v-model="formData.status"
            active-color="#13ce66"
            active-text="是"
            clearable
            inactive-color="#ff4949"
            inactive-text="否"
          ></el-switch>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input
            v-model="formData.describe"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="运行次数:">
          <el-input
            v-model.number="formData.runNumber"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="运行配置:">
          <el-input
            v-model="formData.runConfig"
            clearable
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "ApiCase",
};
</script>

<script setup>
import { createApiCase, updateApiCase, findApiCase } from "@/api/apiCase";

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref } from "vue";

const route = useRoute();
const router = useRouter();
const type = ref("");
const formData = ref({
  name: "",
  runTime: "",
  nextRunTime: new Date(),
  status: false,
  describe: "",
  runNumber: 0,
  runConfig: "",
});

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findApiCase({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.retask;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
};

init();
// 保存按钮
const save = async () => {
  let res;
  switch (type.value) {
    case "create":
      res = await createApiCase(formData.value);
      break;
    case "update":
      res = await updateApiCase(formData.value);
      break;
    default:
      res = await createApiCase(formData.value);
      break;
  }
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "创建/更改成功",
    });
  }
};

// 返回按钮
const back = () => {
  router.go(-1);
};
</script>

<style></style>
