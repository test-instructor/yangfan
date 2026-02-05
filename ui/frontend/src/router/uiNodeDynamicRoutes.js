import { RouterView } from 'vue-router'
import { uiNodeDynamicRouteNames, uiNodeMenus, uiNodeRouteInitError, uiNodeRoutesInitialized } from './uiNodeState'
import { getUINodeMenuTree } from '../services/appBridge'
import NotFound from '../views/error/NotFound.vue'
import RouteComponentMissing from '../views/error/RouteComponentMissing.vue'

const viewModules = import.meta.glob('../views/**/*.vue')

const normalizeComponentKey = (key) => {
  if (!key) return ''
  return String(key).trim().replace(/^\/+/, '').replace(/\.vue$/i, '')
}

const resolveComponent = (componentKey) => {
  const normalized = normalizeComponentKey(componentKey)
  if (!normalized || normalized === 'RouterView') return RouterView
  const modulePath = `../views/${normalized}.vue`
  const loader = viewModules[modulePath]
  if (!loader) {
    return RouteComponentMissing
  }
  return loader
}

const buildDefaultQuery = (parameters) => {
  const q = {}
  if (!Array.isArray(parameters)) return q
  for (const p of parameters) {
    if (!p || p.type !== 'query' || !p.key) continue
    q[p.key] = p.value ?? ''
  }
  return q
}

const buildDefaultParams = (parameters) => {
  const params = {}
  if (!Array.isArray(parameters)) return params
  for (const p of parameters) {
    if (!p || p.type !== 'params' || !p.key) continue
    params[p.key] = p.value ?? ''
  }
  return params
}

const menuToRoute = (menu) => {
  const componentKey = menu?.component || ''
  const resolved = resolveComponent(componentKey)
  const defaultQuery = buildDefaultQuery(menu?.parameters)
  const defaultParams = buildDefaultParams(menu?.parameters)

  const r = {
    path: menu?.path || '',
    name: menu?.name || '',
    component: resolved,
    meta: {
      activeName: menu?.meta?.activeName || '',
      title: menu?.meta?.title || '',
      icon: menu?.meta?.icon || '',
      hidden: Boolean(menu?.hidden),
      keepAlive: Boolean(menu?.meta?.keepAlive),
      defaultMenu: Boolean(menu?.meta?.defaultMenu),
      closeTab: Boolean(menu?.meta?.closeTab),
      transitionType: menu?.meta?.transitionType || '',
      defaultQuery,
      defaultParams,
      componentKey: normalizeComponentKey(componentKey),
      componentResolved: resolved !== RouteComponentMissing
    }
  }

  const children = Array.isArray(menu?.children) ? menu.children : []
  if (children.length) {
    r.children = children.map(menuToRoute)
  }

  if (resolved === RouteComponentMissing) {
    r.meta.missingComponentKey = normalizeComponentKey(componentKey)
  }

  return r
}

export const resetUINodeDynamicRoutes = (router) => {
  for (const name of uiNodeDynamicRouteNames.value) {
    if (router.hasRoute(name)) {
      router.removeRoute(name)
    }
  }
  uiNodeDynamicRouteNames.value = []
  uiNodeRoutesInitialized.value = false
  uiNodeRouteInitError.value = null
  uiNodeMenus.value = []
}

export const getFirstLeafMenuPath = (menus) => {
  const list = Array.isArray(menus) ? menus : []
  const stack = [...list]
  while (stack.length) {
    const m = stack.shift()
    if (!m || m.hidden) continue
    const children = Array.isArray(m.children) ? m.children : []
    if (children.length) {
      stack.unshift(...children)
      continue
    }
    if (m.path) return m.path
  }
  return '/home'
}

export const ensureUINodeDynamicRoutes = async (router) => {
  if (uiNodeRoutesInitialized.value) return true
  uiNodeRouteInitError.value = null

  try {
    const menus = await getUINodeMenuTree()
    uiNodeMenus.value = Array.isArray(menus) ? menus : []
    const routes = uiNodeMenus.value.map(menuToRoute)
    for (const r of routes) {
      if (!r?.name) continue
      router.addRoute('index', r)
      uiNodeDynamicRouteNames.value.push(r.name)
    }
    if (!router.hasRoute('notFound')) {
      router.addRoute({ path: '/:pathMatch(.*)*', name: 'notFound', component: NotFound })
    }
    uiNodeRoutesInitialized.value = true
    return true
  } catch (e) {
    uiNodeRouteInitError.value = e
    uiNodeRoutesInitialized.value = false
    return false
  }
}
