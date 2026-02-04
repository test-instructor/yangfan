// 性能监控工具
export const performanceMonitor = {
  // 记录页面加载时间
  trackPageLoad() {
    if (!window.performance) return
    
    window.addEventListener('load', () => {
      setTimeout(() => {
        const perfData = window.performance.timing
        const pageLoadTime = perfData.loadEventEnd - perfData.navigationStart
        const connectTime = perfData.responseEnd - perfData.requestStart
        const renderTime = perfData.domComplete - perfData.domLoading
        
        console.log('Performance Metrics:', {
          pageLoadTime,
          connectTime,
          renderTime,
          domReady: perfData.domContentLoadedEventEnd - perfData.navigationStart
        })
        
        // 可以发送到后端进行监控
        if (window.go?.main?.App?.TrackPerformance) {
          window.go.main.App.TrackPerformance({
            pageLoadTime,
            connectTime,
            renderTime
          })
        }
      }, 0)
    })
  },
  
  // 监控资源加载性能
  trackResourceLoad() {
    if (!window.performance || !window.performance.getEntriesByType) return
    
    const resources = window.performance.getEntriesByType('resource')
    const slowResources = resources.filter(resource => 
      resource.duration > 1000 // 超过1秒的资源
    )
    
    if (slowResources.length > 0) {
      console.warn('Slow Resources:', slowResources.map(r => ({
        name: r.name,
        duration: r.duration,
        size: r.transferSize
      })))
    }
  }
}

// 初始化性能监控
export function initPerformance() {
  performanceMonitor.trackPageLoad()
  
  // 延迟检查资源加载
  setTimeout(() => {
    performanceMonitor.trackResourceLoad()
  }, 3000)
}