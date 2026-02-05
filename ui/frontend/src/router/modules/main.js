export const mainRoutes = [
  {
    path: '/',
    name: 'index',
    component: () => import('../../views/layout/index.vue'),
    children: [
      {
        path: 'route-init',
        name: 'routeInit',
        component: () => import('../../views/routeInit/index.vue'),
        meta: { title: 'common.loading' }
      }
    ]
  }
]
