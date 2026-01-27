import { createApp } from 'vue'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
import './style.css'
import router from './router'
import { initTheme } from './utils/theme'

initTheme()

const app = createApp(App)
app.use(ArcoVue)
app.use(router)
app.mount('#app')
