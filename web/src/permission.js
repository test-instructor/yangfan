import { useUserStore } from "@/pinia/modules/user";
import { useRouterStore } from "@/pinia/modules/router";
import getPageTitle from "@/utils/page";
import router from "@/router";

let asyncRouterFlag = 0;

const whiteList = ["Login", "Init", "CIResponse"];

const getRouter = async (userStore) => {
  const routerStore = useRouterStore();
  await routerStore.SetAsyncRouter();
  await userStore.GetUserInfo();
  const asyncRouters = routerStore.asyncRouters;
  asyncRouters.forEach((asyncRouter) => {
    router.addRoute(asyncRouter);
  });
};

async function handleKeepAlive(to) {
  if (to.matched.some((item) => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1];
        if (element.name === "layout") {
          to.matched.splice(i, 1);
          await handleKeepAlive(to);
        }
        // 如果没有按需加载完成则等待加载
        if (typeof element.components.default === "function") {
          await element.components.default();
          await handleKeepAlive(to);
        }
      }
    }
  }
}

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  to.meta.matched = [...to.matched];
  handleKeepAlive(to);
  const token = userStore.token;
  // 在白名单中的判断情况
  document.title = getPageTitle(to.meta.title, to);
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        asyncRouterFlag++;
        await getRouter(userStore);
      }
      // token 可以解析但是却是不存在的用户 id 或角色 id 会导致无限调用
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        next({ name: userStore.userInfo.authority.defaultRouter });
      } else {
        // 强制退出账号
        userStore.ClearStorage();
        next({
          name: "Login",
          query: {
            redirect: document.location.hash,
          },
        });
        return;
      }
    } else {
      next();
    }
  } else {
    // 不在白名单中并且已经登录的时候
    if (token) {
      // 添加flag防止多次获取动态路由和栈溢出
      if (!asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        asyncRouterFlag++;
        await getRouter(userStore);
        if (userStore.token) {
          next({ ...to, replace: true });
        } else {
          next({
            name: "Login",
            query: { redirect: to.href },
          });
        }
      } else {
        if (to.matched.length) {
          next();
        } else {
          next({ path: "/layout/404" });
        }
      }
    }
    // 不在白名单中并且未登录的时候
    if (!token) {
      next({
        name: "Login",
        query: {
          redirect: document.location.hash,
        },
      });
    }
  }
});
