export const mainRoutes = [
  {
    path: '/',
    component: () => import('../../views/layout/index.vue'),
    children: [
      { path: '', redirect: 'home' },
      {
        path: 'home',
        name: 'home',
        component: () => import('../../views/home/index.vue'),
        meta: { title: 'Dashboard' }
      },
      {
        path: 'settings',
        name: 'settings',
        component: () => import('../../views/settings/index.vue'),
        meta: { title: 'Settings' }
      },
      {
        path: 'person',
        name: 'person',
        component: () => import('../../views/person/index.vue'),
        meta: { title: 'User Info' }
      }
    ]
  }
]

