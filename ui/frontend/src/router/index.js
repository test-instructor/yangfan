import { createRouter, createWebHashHistory } from 'vue-router'
import { hasToken } from '../services/appBridge'
import { mainRoutes } from './modules/main'
import { authRoutes } from './modules/auth'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [...mainRoutes, ...authRoutes]
})

router.beforeEach(async (to) => {
  const authed = await hasToken()
  if (!authed && to.name !== 'login') return { name: 'login' }
  if (authed && to.name === 'login') return { name: 'home' }

  return true
})

export default router
