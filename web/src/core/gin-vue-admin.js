/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 yangfan
       当前版本:v1.2.0
       加群方式:微信号：test-instructor QQ群：873175584
    `)
  }
}
