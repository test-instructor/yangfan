// 错误处理工具
export class ErrorHandler {
  constructor() {
    this.setupGlobalErrorHandler()
  }
  
  setupGlobalErrorHandler() {
    // 全局错误捕获
    window.addEventListener('error', (event) => {
      this.handleError({
        type: 'javascript',
        message: event.message,
        filename: event.filename,
        lineno: event.lineno,
        colno: event.colno,
        error: event.error
      })
    })
    
    // Promise 错误捕获
    window.addEventListener('unhandledrejection', (event) => {
      this.handleError({
        type: 'promise',
        message: event.reason?.message || 'Unhandled Promise Rejection',
        error: event.reason
      })
    })
  }
  
  handleError(errorInfo) {
    console.error('Global Error:', errorInfo)
    
    // 发送到后端
    if (window.go?.main?.App?.TrackError) {
      window.go.main.App.TrackError(errorInfo)
    }
    
    // 用户友好的错误提示
    this.showErrorNotification(errorInfo)
  }
  
  showErrorNotification(errorInfo) {
    // 可以集成 Arco Design 的消息提示
    if (window.$message) {
      window.$message.error({
        content: `发生错误: ${errorInfo.message}`,
        duration: 5000,
        closable: true
      })
    }
  }
  
  // 业务错误处理
  handleBusinessError(error, context = '') {
    const errorMessage = this.getErrorMessage(error)
    
    console.error(`Business Error [${context}]:`, error)
    
    // 显示用户友好的错误信息
    if (window.$message) {
      window.$message.error({
        content: errorMessage,
        duration: 4000,
        closable: true
      })
    }
    
    return errorMessage
  }
  
  getErrorMessage(error) {
    if (typeof error === 'string') {
      return error
    }
    
    if (error?.response?.data?.message) {
      return error.response.data.message
    }
    
    if (error?.message) {
      return error.message
    }
    
    return '操作失败，请稍后重试'
  }
  
  // 网络错误处理
  handleNetworkError(error) {
    if (!navigator.onLine) {
      return this.handleBusinessError('网络连接已断开，请检查网络连接')
    }
    
    if (error?.code === 'ECONNABORTED') {
      return this.handleBusinessError('请求超时，请稍后重试')
    }
    
    if (error?.response?.status === 401) {
      return this.handleBusinessError('登录已过期，请重新登录')
    }
    
    if (error?.response?.status >= 500) {
      return this.handleBusinessError('服务器错误，请稍后重试')
    }
    
    return this.handleBusinessError(error)
  }
}

// 创建全局错误处理器实例
export const errorHandler = new ErrorHandler()

// 便捷的错误处理函数
export function handleError(error, context = '') {
  return errorHandler.handleBusinessError(error, context)
}

export function handleNetworkError(error) {
  return errorHandler.handleNetworkError(error)
}