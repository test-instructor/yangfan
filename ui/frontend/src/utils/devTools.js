// 开发调试工具
export class DevTools {
  constructor() {
    this.isEnabled = process.env.NODE_ENV === 'development'
    this.init()
  }
  
  init() {
    if (!this.isEnabled) return
    
    // 添加全局调试对象
    window.$yangfan = {
      version: '1.0.0',
      env: process.env.NODE_ENV,
      utils: {
        clearStorage: () => localStorage.clear(),
        getConfig: () => localStorage.getItem('yangfan-config'),
        showRoutes: () => console.table(this.getRoutes()),
        testApi: (method, ...args) => this.testApi(method, args)
      }
    }
    
    console.log('%c🚀 扬帆自动化测试平台 - 调试模式', 'color: #667eea; font-size: 16px; font-weight: bold;')
    console.log('%c使用 window.$yangfan 访问调试工具', 'color: #666; font-size: 12px;')
  }
  
  getRoutes() {
    // 获取当前路由信息
    const routes = []
    const router = window.$router
    if (router && router.getRoutes) {
      return router.getRoutes().map(route => ({
        path: route.path,
        name: route.name,
        component: route.component?.name || 'Anonymous'
      }))
    }
    return routes
  }
  
  async testApi(method, args = []) {
    try {
      console.log(`🧪 测试 API: ${method}`, args)
      const result = await window.go?.main?.App?.[method]?.(...args)
      console.log(`✅ ${method} 成功:`, result)
      return result
    } catch (error) {
      console.error(`❌ ${method} 失败:`, error)
      throw error
    }
  }
  
  // 性能测试工具
  benchmark(fn, name = 'Benchmark', iterations = 1000) {
    console.log(`🏃‍♂️ 开始性能测试: ${name}`)
    const start = performance.now()
    
    for (let i = 0; i < iterations; i++) {
      fn()
    }
    
    const end = performance.now()
    const total = end - start
    const average = total / iterations
    
    console.log(`📊 ${name} 测试结果:`)
    console.log(`   总耗时: ${total.toFixed(2)}ms`)
    console.log(`   平均耗时: ${average.toFixed(4)}ms`)
    console.log(`   迭代次数: ${iterations}`)
    
    return { total, average, iterations }
  }
  
  // 内存使用监控
  monitorMemory() {
    if (!performance.memory) {
      console.warn('当前浏览器不支持内存监控')
      return
    }
    
    const memory = performance.memory
    console.log('💾 内存使用情况:')
    console.log(`   已用内存: ${(memory.usedJSHeapSize / 1024 / 1024).toFixed(2)} MB`)
    console.log(`   总内存: ${(memory.totalJSHeapSize / 1024 / 1024).toFixed(2)} MB`)
    console.log(`   内存限制: ${(memory.jsHeapSizeLimit / 1024 / 1024).toFixed(2)} MB`)
    
    return memory
  }
  
  // API 调用日志
  logApiCall(method, args, result, error = null) {
    if (!this.isEnabled) return
    
    const timestamp = new Date().toISOString()
    const logEntry = {
      timestamp,
      method,
      args,
      result: error ? null : result,
      error: error ? error.message : null,
      duration: result?.duration || null
    }
    
    console.group(`📡 API 调用: ${method}`)
    console.log('时间:', timestamp)
    console.log('参数:', args)
    if (error) {
      console.error('错误:', error)
    } else {
      console.log('结果:', result)
    }
    console.groupEnd()
    
    // 存储到本地，便于调试
    const logs = JSON.parse(localStorage.getItem('yangfan-api-logs') || '[]')
    logs.push(logEntry)
    if (logs.length > 100) logs.shift() // 保持最近100条
    localStorage.setItem('yangfan-api-logs', JSON.stringify(logs))
  }
  
  // 显示 API 调用历史
  showApiLogs() {
    const logs = JSON.parse(localStorage.getItem('yangfan-api-logs') || '[]')
    if (logs.length === 0) {
      console.log('暂无 API 调用记录')
      return
    }
    
    console.table(logs.map(log => ({
      时间: new Date(log.timestamp).toLocaleTimeString(),
      方法: log.method,
      状态: log.error ? '❌ 失败' : '✅ 成功',
      错误: log.error || '-'
    })))
  }
}

// 创建开发工具实例
export const devTools = new DevTools()

// 便捷函数
export const benchmark = (fn, name, iterations) => devTools.benchmark(fn, name, iterations)
export const monitorMemory = () => devTools.monitorMemory()
export const showApiLogs = () => devTools.showApiLogs()

export default devTools