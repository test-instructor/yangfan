import { createRouter, createWebHashHistory } from 'vue-router'
import { hasToken } from '../services/appBridge'
import { mainRoutes } from './modules/main'
import { authRoutes } from './modules/auth'
import { ensureUINodeDynamicRoutes, getFirstLeafMenuPath } from './uiNodeDynamicRoutes'
import { uiNodeMenus } from './uiNodeState'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [...mainRoutes, ...authRoutes]
})

router.beforeEach(async (to) => {
  const authed = await hasToken()
  if (!authed && to.name !== 'login') return { name: 'login' }
  if (authed && to.name === 'login') return { name: 'routeInit' }

  if (authed) {
    const ok = await ensureUINodeDynamicRoutes(router)
    if (!ok) return { name: 'routeInit' }
    if (to.name === 'routeInit') {
      return { path: getFirstLeafMenuPath(uiNodeMenus.value), replace: true }
    }
  }

  return true
})

export default router
