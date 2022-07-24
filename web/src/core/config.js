/**
 * 网站配置文件
 */

const config = {
  appName: 'cheetah',
  // appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  appLogo: '/src/assets/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用 cheetah`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:v1.0.0`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信号：taylorter QQ群：873175584`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log('\n')
  }
}

export default config
