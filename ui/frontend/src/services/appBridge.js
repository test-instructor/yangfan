import { handleNetworkError } from '@/utils/errorHandler'

const callApp = async (method, ...args) => {
  try {
    const fn = window?.go?.main?.App?.[method]
    if (!fn) {
      throw new Error(`Go 方法不可用: ${method}`)
    }
    
    const result = await fn(...args)
    
    // 统一处理返回结果
    if (result && typeof result === 'object' && result.error) {
      throw new Error(result.error)
    }
    
    return result
  } catch (error) {
    console.error(`调用 Go 方法失败 [${method}]:`, error)
    handleNetworkError(error)
    throw error
  }
}

export const getBaseURL = async () => {
  const res = await callApp('GetBaseURL')
  if (Array.isArray(res)) {
    return { baseURL: res[0] || '', ok: Boolean(res[1]) }
  }
  if (res && typeof res === 'object') {
    return { baseURL: res.baseURL || '', ok: Boolean(res.ok) }
  }
  return { baseURL: String(res || ''), ok: Boolean(res) }
}

export const setBaseURL = async (baseURL) => {
  return await callApp('SetBaseURL', baseURL)
}

export const checkBaseURLConnectivity = async (baseURL) => {
  const res = await callApp('CheckBaseURLConnectivity', baseURL)
  if (res && typeof res === 'object') {
    return { ok: Boolean(res.ok), baseURL: res.baseURL || '', ...res }
  }
  return { ok: Boolean(res), baseURL: '' }
}

export const clearAuth = async () => {
  return await callApp('ClearAuth')
}

export const hasToken = async () => {
  return await callApp('HasToken')
}

export const captcha = async () => {
  const res = await callApp('Captcha')
  return res || {}
}

export const login = async ({ username, password, captcha, captchaId }) => {
  return await callApp('Login', username, password, captcha, captchaId)
}

export const getUserInfo = async () => {
  return await callApp('GetUserInfo')
}

export const setUserAuthority = async ({ authorityId, projectId }) => {
  return await callApp('SetUserAuthority', authorityId, projectId)
}

export const getLogConfig = async () => {
  return await callApp('GetLogConfig')
}

export const setLogConfig = async ({ level, prefix, retention }) => {
  return await callApp('SetLogConfig', level, prefix, retention)
}

export const setSelfInfo = async (info) => {
  return await callApp('SetSelfInfo', info)
}

export const changePassword = async ({ password, newPassword }) => {
  return await callApp('ChangePassword', password, newPassword)
}

export const getAppConfig = async () => {
  return await callApp('GetAppConfig')
}

export const setAppConfig = async (config) => {
  const { environment, debugMode, theme, language, autoLogin, rememberMe } = config
  return await callApp('SetAppConfig', environment, debugMode, theme, language, autoLogin, rememberMe)
}

export const getSystemInfo = async () => {
  return await callApp('GetSystemInfo')
}

export const trackPerformance = async (metrics) => {
  return await callApp('TrackPerformance', metrics)
}

export const trackError = async (errorInfo) => {
  return await callApp('TrackError', errorInfo)
}
