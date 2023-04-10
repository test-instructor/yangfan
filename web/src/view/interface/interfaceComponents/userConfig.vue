<template>
  <el-form :inline="true"  style="margin-left:20px;" >
    <el-form-item v-if="userConfig && userConfig.api_config_id" label="环境配置：">
      <el-tag
          class="mx-1"
          size="large"
          type="success"
      >
        {{userConfig.api_config.name}}
      </el-tag>
    </el-form-item>
    <el-form-item v-if="userConfig && userConfig.api_env_id" label="环境变量：">
      <el-tag
          class="mx-1"
          size="large"
          type="success"
      >
        {{userConfig.api_env.name}}
      </el-tag>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: "userConfig"
}
</script>

<script setup>

import {ref} from "vue";
import {getUserConfig} from "@/api/interfaceTemplate";

const props = defineProps({
    api_config_name:"",
    api_env_name:""
})

const userConfig = ref({
    api_config:{name:""},
    api_env:{name:""}
})
const getUserConfigs = async () => {
  let res = await getUserConfig()
  if (res.code === 0 && res.data) {
    userConfig.value = res.data
  }
}
const init = () => {
    console.log("=============1", props )
    console.log("=============2", props.api_config_name )
    console.log("=============3", props.api_env_name )
  if (props.api_config_name && props.api_config_name !== ""){
      userConfig.value.api_config.name = props.api_config_name
      userConfig.value.api_env.name = props.api_env_name
  }else {
      getUserConfigs()
  }
}
init()
</script>


<style scoped>

</style>