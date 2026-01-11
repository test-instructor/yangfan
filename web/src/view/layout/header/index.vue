<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div
    class="flex justify-between fixed top-0 left-0 right-0 z-10 h-16 bg-white text-slate-700 dark:text-slate-300 dark:bg-slate-900 shadow dark:shadow-gray-700 items-center px-2"
  >
    <div class="flex items-center cursor-pointer flex-1">
      <div
        class="flex items-center cursor-pointer"
        :class="isMobile ? '' : 'min-w-48'"
        @click="router.push({ path: '/' })"
      >
        <img
          alt
          class="h-12 bg-white rounded-full"
          :src="$GIN_VUE_ADMIN.appLogo"
        />
        <div
          v-if="!isMobile"
          class="inline-flex font-bold text-2xl ml-2"
          :class="
            (config.side_mode === 'head' ||
              config.side_mode === 'combination') &&
            'min-w-fit'
          "
        >
          {{ $GIN_VUE_ADMIN.appName }}
        </div>
      </div>

      <el-breadcrumb
        v-show="!isMobile"
        v-if="config.side_mode !== 'head' && config.side_mode !== 'combination'"
        class="ml-4"
      >
        <el-breadcrumb-item
          v-for="item in matched.slice(1, matched.length)"
          :key="item.path"
        >
          {{ fmtTitle(item.meta.title, route) }}
        </el-breadcrumb-item>
      </el-breadcrumb>
      <gva-aside
        v-if="config.side_mode === 'head' && !isMobile"
        class="flex-1"
      />
      <gva-aside
        v-if="config.side_mode === 'combination' && !isMobile"
        mode="head"
        class="flex-1"
      />
    </div>
    <div>
      <el-dropdown
        trigger="hover"
        @command="handleProjectChange"
        class="project-dropdown"
      >
          <span class="dropdown-trigger">
            当前项目：{{ currentProject?.name || '选择项目' }}
            <el-icon><arrow-down /></el-icon>
          </span>

        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item
              v-for="project in projectList"
              :key="project.id"
              :command="project"
              :class="{ 'is-active': currentProject?.id === project.id }"
            >
              <span> 切换为：{{ project.name }} </span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <div class="ml-2 flex items-center">
      <tools />


      <el-dropdown>
        <div class="flex justify-center items-center h-full w-full">
          <span
            class="cursor-pointer flex justify-center items-center text-black dark:text-gray-100"
          >
            <CustomPic />
            <span v-show="!isMobile" class="w-16">{{
              userStore.userInfo.nickName
            }}</span>
            <el-icon>
              <arrow-down />
            </el-icon>
          </span>
        </div>



        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <span class="font-bold">
                当前角色：{{ userStore.userInfo.authority.authorityName }}
              </span>
            </el-dropdown-item>
            <template v-if="userStore.userInfo.authorities">
              <el-dropdown-item
                v-for="item in userStore.userInfo.authorities.filter(
                  (i) => i.authorityId !== userStore.userInfo.authorityId
                )"
                :key="item.authorityId"
                @click="changeUserAuth(item.authorityId)"
              >
                <span> 切换为：{{ item.authorityName }} </span>
              </el-dropdown-item>
            </template>
            <el-dropdown-item icon="avatar" @click="toPerson">
              个人信息
            </el-dropdown-item>
            <el-dropdown-item icon="reading-lamp" @click="userStore.LoginOut">
              登 出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import tools from './tools.vue'
  import CustomPic from '@/components/customPic/index.vue'
  import { useUserStore } from '@/pinia/modules/user'
  import { useRoute, useRouter } from 'vue-router'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { computed } from 'vue'
  import { setUserAuthority } from '@/api/user'
  import { fmtTitle } from '@/utils/fmtRouterTitle'
  import gvaAside from '@/view/layout/aside/index.vue'
  const userStore = useUserStore()
  const router = useRouter()
  const route = useRoute()
  const appStore = useAppStore()
  const { device, config } = storeToRefs(appStore)
  const isMobile = computed(() => {
    return device.value === 'mobile'
  })
  const toPerson = () => {
    router.push({ name: 'person' })
  }
  const matched = computed(() => route.meta.matched)

  const changeUserAuth = async (id) => {
    const currentProject = JSON.parse(localStorage.getItem('currentProject'))
    const res = await setUserAuthority({
      authorityId: id,
      projectId: currentProject.id
    })
    if (res.code === 0) {
      window.sessionStorage.setItem('needCloseAll', 'true')
      window.sessionStorage.setItem('needToHome', 'true')
      window.location.reload()
    }
  }

  const currentProject = ref({})
  const projectList = ref([])

  const initProjectData = () => {
    // 初始化数据逻辑（同上）
    const storedProjectList = JSON.parse(localStorage.getItem('projectList') || '[]')

    if (storedProjectList.length === 0) {

    } else {
      projectList.value = storedProjectList
    }

    const storedCurrentProject = JSON.parse(localStorage.getItem('currentProject') || 'null')
    currentProject.value = storedCurrentProject || projectList.value[0]
  }

  const handleProjectChange = (project) => {
    // 检查是否选择了相同的项目
    if (currentProject.value?.id === project.id) {
      console.log('已经是当前项目')
      return
    }

    currentProject.value = project
    userStore.setProject(project)
    changeUserAuthProject(project.id)
  }
  const changeUserAuthProject = async (id) => {
    const res = await setUserAuthority({
      authorityId: userStore.userInfo.authority.authorityId,
      projectId: id
    })
    if (res.code === 0) {
      window.sessionStorage.setItem('needCloseAll', 'true')
      window.sessionStorage.setItem('needToHome', 'true')
      window.location.reload()
    }
  }
  initProjectData()
</script>

<style scoped lang="scss">




</style>
