<template>
  <el-form>
    <el-form-item label="环境配置">
      <el-select
          v-model="configId"
          filterable
          placeholder="环境配置"
      >
        <el-option
            v-model="configId"
            v-for="item in configOptions"
            :key="item.ID"
            :label="item.name"
            :value="item.ID"
            @click.native="getConfigID(item.ID)"
        />
      </el-select>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: "envConfig"
}
</script>

<script setup>
import {ref, defineEmits} from "vue";

const configOptions = ref()
const configId = ref()
import {getApiConfigList} from "@/api/apiConfig";
const emit = defineEmits(["configId"]);

const getConfigID = (id) => {
  configId.value = id
  emit("configId", configId.value)
}

const getConfigData = async () => {
  const table = await getApiConfigList({page: 1, pageSize: 9999})
  if (table.code === 0) {
    configOptions.value = table.data.list
    configOptions.value.forEach(item => {
      if (item.default){
        configId.value = item.ID
        emit("configId", configId.value)
      }
    })
  }
}

const init = () => {
  getConfigData()
}
init()

</script>

<style scoped>

</style>
