/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'

const greenText = (text) => `\x1b[32m${text}\x1b[0m`

export const config = {
  appName: '扬帆自动化测试平台',
  // appLogo: 'logo.png',
  showViteLogo: true,
  KeepAliveTabs: true,
  logs: []
}

export const viteLogo = (env) => {}

export default config
