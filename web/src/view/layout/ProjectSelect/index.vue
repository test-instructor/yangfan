<template>
  <el-select v-model="value" placeholder="请选择" @change="selectProject(value)">
    <el-option v-for="item in options" :key="item.ID" :label="item.name" :value="item.ID"></el-option>
  </el-select>
</template>

<script>
import {defineComponent} from 'vue'
import {useUserStore} from '@/pinia/modules/user'

const userStore = useUserStore()
var projects = userStore.userInfo.projects
var project = JSON.parse(window.localStorage.getItem('project'))
export default defineComponent({
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
