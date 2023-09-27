import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/init",
    name: "Init",
    component: () => import("@/view/init/index.vue"),
  },
  {
    path: "/response/:project/:report_id/:ci",
    name: "CIResponse",
    component: () => import("@/view/interface/Reports/reportDetail.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/view/login/index.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
