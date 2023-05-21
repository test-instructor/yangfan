<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="管理员:">
          <el-input v-model="formData.admin" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:">
          <el-input v-model="formData.creator" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="项目描述:">
          <el-input
            v-model="formData.describe"
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
  name: "Project",
};
</script>

<script setup>
import { createProject, updateProject, findProject } from "@/api/project";

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
  admin: "",
  creator: "",
  describe: "",
});

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findProject({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.reproject;
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
      res = await createProject(formData.value);
      break;
    case "update":
      res = await updateProject(formData.value);
      break;
    default:
      res = await createProject(formData.value);
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
