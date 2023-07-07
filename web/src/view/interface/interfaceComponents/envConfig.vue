<template>
  <el-form>
    <el-form-item label="配置">
      <el-select v-model="configId" filterable placeholder="配置">
        <el-option
          v-model="configId"
          v-for="item in configOptions"
          :key="item.ID"
          :label="item.name"
          :value="item.ID"
          @click.native="getConfigID(item)"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="环境">
      <el-select v-model="envId" filterable placeholder="环境">
        <el-option
          v-model="envId"
          v-for="item in envOptions"
          :key="item.ID"
          :label="item.name"
          :value="item.ID"
          @click.native="getEnvID(item)"
        />
      </el-select>
    </el-form-item>
    <el-form-item>
      <env-keys />
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: "envConfig",
};
</script>

<script setup>
import { ref, defineEmits } from "vue";

const configOptions = ref();
const envOptions = ref();
const configId = ref();
const envId = ref();
import { getApiConfigList } from "@/api/apiConfig";
import { getUserConfig } from "@/api/interfaceTemplate";
import { getEnvList } from "@/api/env";
const emit = defineEmits(["configId"]);
const userConfigDialog = ref(false);
const userConfig = ref({
  api_config_id: 0,
  api_env_id: 0,
});

const getConfigID = (item) => {
  emit("configId", item.ID);
};

const getEnvID = (item) => {
  emit("envId", item.ID);
};

const getConfigData = async () => {
  const table = await getApiConfigList({ page: 1, pageSize: 9999 });
  if (table.code === 0) {
    configOptions.value = table.data.list;
  }
  let env = await getEnvList();
  if (env.code === 0) {
    envOptions.value = env.data.list;
  }
  getConfig();
};
const userConfigs = ref({
  api_config_id: "",
  api_env_id: "",
  ID: "",
});
const getConfig = async () => {
  let res = await getUserConfig();
  if (res.code === 0 && res.data) {
    if (res.data.api_env_id > 0) {
      userConfigs.value.api_env_id = res.data.api_env_id;
      envId.value = res.data.api_env_id;
    }
    if (res.data.api_config_id > 0) {
      userConfigs.value.api_config_id = res.data.api_config_id;
      configId.value = res.data.api_config_id;
    }
    if (res.data.ID > 0) {
      userConfigs.value.ID = res.data.ID;
    }
  }
};

const setUserConfig = () => {
  userConfigDialog.value = true;
};

const closeDialogUserConfig = () => {
  userConfigDialog.value = false;
};

const init = () => {
  getConfigData();
};
init();
</script>

<style scoped></style>
