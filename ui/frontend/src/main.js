import { createApp } from 'vue'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
import './style.css'
import router from './router'
import { initTheme } from './utils/theme'
import { initPerformance } from './utils/performance'
import { errorHandler } from './utils/errorHandler'
import { devTools } from './utils/devTools'
import i18n from './i18n'

// 初始化主题
initTheme()

// 初始化性能监控
initPerformance()

// 初始化开发工具（仅在开发环境）
if (import.meta.env.DEV) {
  devTools.init()
}

// 全局错误处理已在 ErrorHandler 构造函数中设置

const app = createApp(App)

// 全局属性
app.config.globalProperties.$error = errorHandler
app.config.globalProperties.$dev = devTools

app.use(ArcoVue)
app.use(router)
app.use(i18n)
app.mount('#app')
