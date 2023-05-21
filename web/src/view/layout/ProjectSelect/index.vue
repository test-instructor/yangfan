<template>
  <el-dropdown>
    <div class="dp-flex justify-content-center align-items height-full width-full">
      <span class="header-avatar" style="cursor: pointer">
        <CustomPic />
        <span style="margin-left: 5px">项目：{{ label }}</span>
        <el-icon>
          <arrow-down />
        </el-icon>
      </span>
    </div>
    <template #dropdown>
      <el-dropdown-menu class="dropdown-group">
        <el-dropdown-item>
          <span style="font-weight: 600;">
            当前项目：{{ label }}
          </span>
        </el-dropdown-item>
        <template v-if="options">
          <el-dropdown-item v-for="item in options.filter(i=>i.label!==label)" :key="item.label" @click="selectProject(item.ID)"  >
            <span>
              切换为：{{ item.name }}
            </span>
          </el-dropdown-item>
        </template>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script>
import {defineComponent} from 'vue'
import {useUserStore} from '@/pinia/modules/user'

const userStore = useUserStore()
var projects = userStore.userInfo.projects
var project = JSON.parse(window.localStorage.getItem('project'))
export default defineComponent({
  name: 'SelectProject',
  methods: {
    selectProject(value) {
      this.value = value
      this.$forceUpdate()
      for (var i = 0; i < this.options.length; i++) {
        if (this.options[i].ID === this.value) {
          window.localStorage.setItem('project', JSON.stringify(this.options[i]))
          location.reload()
        }
      }
    }
  },

  setup() {
    return {
      options: projects,
      key: project.ID,
      value: project.name,
      label: project.name
    }
  }
})
</script>
<script setup>

</script>
<style lang="scss" scoped></style>
