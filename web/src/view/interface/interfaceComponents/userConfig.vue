<template>
  <el-form :inline="true" style="margin-left: 20px">
    <el-form-item
      v-if="userConfig && userConfig.api_config_id"
      label="环境配置："
    >
      <el-tag class="mx-1" size="large" type="success">
        {{ userConfig.api_config.name }}
      </el-tag>
    </el-form-item>
    <el-form-item v-if="userConfig && userConfig.api_env_id" label="环境变量：">
      <el-tag class="mx-1" size="large" type="success">
        {{ userConfig.api_env.name }}
      </el-tag>
    </el-form-item>
    <el-form-item>
      <env-copy></env-copy>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: "userConfig",
};
</script>

<script setup>
import { ref } from "vue";
import { getUserConfig } from "@/api/interfaceTemplate";
import EnvCopy from "@/view/interface/interfaceComponents/envCopy.vue";

const props = defineProps({
  api_config_name: "",
  api_env_name: "",
});

const userConfig = ref({
  api_config: { name: "" },
  api_env: { name: "" },
  api_config_id: "",
  api_env_id: "",
});

const getUserConfigs = async () => {
  let res = await getUserConfig();
  if (res.code === 0 && res.data) {
    if (res.data.api_env_id > 0) {
      userConfig.value.api_env_id = res.data.api_env_id;
      userConfig.value.api_env.name = res.data.api_env.name;
    }
    if (res.data.api_config_id > 0) {
      userConfig.value.api_config_id = res.data.api_config_id;
      userConfig.value.api_config.name = res.data.api_config.name;
    }
  }
};
const init = () => {
  if (props.api_config_name && props.api_config_name !== "") {
    userConfig.value.api_config.name = props.api_config_name;
    userConfig.value.api_env.name = props.api_env_name;
  } else {
    getUserConfigs();
  }
};
init();
</script>

<style scoped></style>
