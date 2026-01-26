import { createRouter, createWebHashHistory } from 'vue-router'
import { getBaseURL, hasToken } from '../services/appBridge'
import SettingsPage from '../views/SettingsPage.vue'
import LoginPage from '../views/LoginPage.vue'
import HomePage from '../views/HomePage.vue'

let notifiedMissingBaseURL = false

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', redirect: '/home' },
    { path: '/settings', name: 'settings', component: SettingsPage },
    { path: '/login', name: 'login', component: LoginPage },
    { path: '/home', name: 'home', component: HomePage }
  ]
})

router.beforeEach(async (to) => {
  const { baseURL, ok } = await getBaseURL()

  if (!ok && to.name !== 'settings') {
    const query = notifiedMissingBaseURL ? {} : { missing: '1' }
    notifiedMissingBaseURL = true
    return { name: 'settings', query }
  }

  if (to.name === 'settings') return true

  const authed = await hasToken()
  if (!authed && to.name !== 'login') return { name: 'login' }
  if (authed && to.name === 'login') return { name: 'home' }

  void baseURL
  return true
})

export default router
