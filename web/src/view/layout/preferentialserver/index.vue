<template>
  <el-dropdown>
    <span class="el-dropdown-link">
      特惠服务器
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item v-for="item in server" @click="openUrl(item.url)">{{item.name}}</el-dropdown-item>

      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script>
import {defineComponent} from 'vue'
import {useUserStore} from '@/pinia/modules/user'

export default defineComponent({
  name: 'PreferentialServer',
})

</script>
<script setup>
import {ref} from "vue";

const server = ref([
  {
    name:'腾讯云特惠服务器',
    url:'https://cloud.tencent.com/act/cps/redirect?redirect=36284&cps_key=2b81c9aacd37f7933c45bf18432afbe4',
  },{
    name:'阿里云特惠服务器',
    url:'https://www.aliyun.com/minisite/goods?userCode=f4a9e0fo',
  }
])
const openUrl = (url) => {
  window.open(url, "_search");
}
</script>
<style lang="scss" scoped></style>
