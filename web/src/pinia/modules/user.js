import { login, getUserInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useRouterStore } from './router'
import { useCookies } from '@vueuse/integrations/useCookies'
import { useStorage } from '@vueuse/core'

import { useAppStore } from '@/pinia'

export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {}
  })
  const token = useStorage('token', '')
  const xToken = useCookies('x-token')
  const currentToken = computed(() => token.value || xToken.value || '')

  const setUserInfo = (val) => {
    userInfo.value = val
    if (val.originSetting) {
      Object.keys(appStore.config).forEach((key) => {
        if (val.originSetting[key] !== undefined) {
          appStore.config[key] = val.originSetting[key]
        }
      })
    }
    if (val.projectList) {
      SetProjectInfo(val.projectList,val.projectId)
    }
  }

  // 设置用户项目信息、项目列表，同时把项目信息写入localStorage
  const SetProjectInfo = (projectList, projectId) => {
    if (!projectList || projectList.length === 0) {
      toast.error("请联系管理员添加项目权限")
      return
    }

    // 情况1：优先使用传入的projectId
    if (projectId) {
      const targetProject = projectList.find(project => project.id === projectId)
      if (targetProject) {
        setProject(targetProject)
        localStorage.setItem('currentProject', JSON.stringify(targetProject))
        localStorage.setItem('projectList', JSON.stringify(projectList))
        return
      } else {
        console.warn('Provided projectId not found in projectList:', projectId)
      }
    }

    // 情况2：检查并更新已保存的项目
    const savedProject = localStorage.getItem('currentProject')
    if (savedProject) {
      try {
        const parsedProject = JSON.parse(savedProject)
        // 查找projectList中对应的完整项目数据
        const currentProjectData = projectList.find(project => project.id === parsedProject.id)
        if (currentProjectData) {
          // 使用projectList中的最新数据更新
          setProject(currentProjectData)
          localStorage.setItem('currentProject', JSON.stringify(currentProjectData))
          localStorage.setItem('projectList', JSON.stringify(projectList))
          return
        } else {
        }
      } catch (error) {
        console.error('Error parsing saved project:', error)
      }
    }

    // 情况3：使用默认的第一个项目
    const defaultProject = projectList[0]
    setProject(defaultProject)
    localStorage.setItem('currentProject', JSON.stringify(defaultProject))
    localStorage.setItem('projectList', JSON.stringify(projectList))
  }

  const setProject = (project) => {
    userInfo.value.projectId = project.id
    localStorage.setItem('currentProject', JSON.stringify(project))
  }
  const setToken = (val) => {
    token.value = val
    xToken.value = val
  }

  const NeedInit = async () => {
    await ClearStorage()
    await router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async () => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }
  /* 登录*/
  const LoginIn = async (loginInfo) => {
    try {
      loadingInstance.value = ElLoading.service({
        fullscreen: true,
        text: '登录中，请稍候...'
      })

      const res = await login(loginInfo)

      if (res.code !== 0) {
        return false
      }
      // 登陆成功，设置用户信息和权限相关信息
      setUserInfo(res.data.user)
      setToken(res.data.token)
      // 初始化路由信息
      const routerStore = useRouterStore()
      await routerStore.SetAsyncRouter()
      const asyncRouters = routerStore.asyncRouters

      // 注册到路由表里
      asyncRouters.forEach((asyncRouter) => {
        router.addRoute(asyncRouter)
      })

      if(router.currentRoute.value.query.redirect) {
        await router.replace(router.currentRoute.value.query.redirect)
        return true
      }

      if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
        ElMessage.error('不存在可以登陆的首页，请联系管理员进行配置')
      } else {
        await router.replace({ name: userInfo.value.authority.defaultRouter })
      }

      const isWindows = /windows/i.test(navigator.userAgent)
      window.localStorage.setItem('osType', isWindows ? 'WIN' : 'MAC')

      // 全部操作均结束，关闭loading并返回
      return true
    } catch (error) {
      console.error('LoginIn error:', error)
      return false
    } finally {
      loadingInstance.value?.close()
    }
  }
  /* 登出*/
  const LoginOut = async () => {
    const res = await jsonInBlacklist()

    // 登出失败
    if (res.code !== 0) {
      return
    }

    await ClearStorage()

    // 把路由定向到登录页，无需等待直接reload
    router.push({ name: 'Login', replace: true })
    window.location.reload()
  }
  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = ''
    // 使用remove方法正确删除cookie
    xToken.remove()
    sessionStorage.clear()
    // 清理所有相关的localStorage项
    localStorage.removeItem('originSetting')
    localStorage.removeItem('token')
  }

  return {
    userInfo,
    token: currentToken,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setToken,
    loadingInstance,
    ClearStorage,
    setUserInfo,
    setProject
  }
})
