import { createRouter, createWebHashHistory } from 'vue-router'
import { hasToken } from '../services/appBridge'
import SettingsPage from '../views/SettingsPage.vue'
import PersonPage from '../views/PersonPage.vue'
import LoginPage from '../views/LoginPage.vue'
import HomePage from '../views/HomePage.vue'
import DashboardPage from '../views/DashboardPage.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { 
      path: '/', 
      component: HomePage,
      children: [
        { path: '', redirect: 'home' },
        { path: 'home', name: 'home', component: DashboardPage },
        { path: 'settings', name: 'settings', component: SettingsPage },
        { path: 'person', name: 'person', component: PersonPage }
      ]
    },
    { path: '/login', name: 'login', component: LoginPage }
  ]
})

router.beforeEach(async (to) => {
  const authed = await hasToken()
  if (!authed && to.name !== 'login') return { name: 'login' }
  if (authed && to.name === 'login') return { name: 'home' }

  return true
})

export default router
