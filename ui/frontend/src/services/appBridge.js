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
    handleNetworkError(error, { method })
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

export const getUINodeMenuTree = async () => {
  const res = await callApp('GetUINodeMenuTree')
  return Array.isArray(res) ? res : (res || [])
}

export const setUserAuthority = async ({ authorityId, projectId }) => {
  return await callApp('SetUserAuthority', authorityId, projectId)
}

export const getAndroidDeviceOptionsList = async (query) => {
  const res = await callApp('GetAndroidDeviceOptionsList', query || {})
  return res || {}
}

export const createAndroidDeviceOptions = async (payload) => {
  const res = await callApp('CreateAndroidDeviceOptions', payload || {})
  return res || {}
}

export const updateAndroidDeviceOptions = async (payload) => {
  const res = await callApp('UpdateAndroidDeviceOptions', payload || {})
  return res || {}
}

export const deleteAndroidDeviceOptions = async (id) => {
  return await callApp('DeleteAndroidDeviceOptions', id)
}

export const getRunConfigList = async (query) => {
  const res = await callApp('GetRunConfigList', query || {})
  return res || {}
}

export const createRunConfig = async (payload) => {
  const res = await callApp('CreateRunConfig', payload || {})
  return res || {}
}

export const updateRunConfig = async (payload) => {
  const res = await callApp('UpdateRunConfig', payload || {})
  return res || {}
}

export const deleteRunConfig = async (id) => {
  return await callApp('DeleteRunConfig', id)
}

export const getAutoStepList = async (query) => {
  const res = await callApp('GetAutoStepList', query || {})
  return res || {}
}

export const createAutoStep = async (payload) => {
  const res = await callApp('CreateAutoStep', payload || {})
  return res || {}
}

export const updateAutoStep = async (payload) => {
  const res = await callApp('UpdateAutoStep', payload || {})
  return res || {}
}

export const deleteAutoStep = async (id) => {
  return await callApp('DeleteAutoStep', id)
}

export const getAutoCaseStepList = async (query) => {
  const res = await callApp('GetAutoCaseStepList', query || {})
  return res || {}
}

export const createAutoCaseStep = async (payload) => {
  const res = await callApp('CreateAutoCaseStep', payload || {})
  return res || {}
}

export const updateAutoCaseStep = async (payload) => {
  const res = await callApp('UpdateAutoCaseStep', payload || {})
  return res || {}
}

export const deleteAutoCaseStep = async (id) => {
  return await callApp('DeleteAutoCaseStep', id)
}

export const getAutoCaseList = async (query) => {
  const res = await callApp('GetAutoCaseList', query || {})
  return res || {}
}

export const createAutoCase = async (payload) => {
  const res = await callApp('CreateAutoCase', payload || {})
  return res || {}
}

export const updateAutoCase = async (payload) => {
  const res = await callApp('UpdateAutoCase', payload || {})
  return res || {}
}

export const deleteAutoCase = async (id) => {
  return await callApp('DeleteAutoCase', id)
}

export const getTimerTaskList = async (query) => {
  const res = await callApp('GetTimerTaskList', query || {})
  return res || {}
}

export const createTimerTask = async (payload) => {
  const res = await callApp('CreateTimerTask', payload || {})
  return res || {}
}

export const updateTimerTask = async (payload) => {
  const res = await callApp('UpdateTimerTask', payload || {})
  return res || {}
}

export const deleteTimerTask = async (id) => {
  return await callApp('DeleteTimerTask', id)
}

export const getAutoReportList = async (query) => {
  const res = await callApp('GetAutoReportList', query || {})
  return res || {}
}

export const findAutoReport = async (id) => {
  const res = await callApp('FindAutoReport', id)
  return res || {}
}

export const findAutoCaseStepApis = async (autoCaseStepId) => {
  const res = await callApp('FindAutoCaseStepApis', autoCaseStepId)
  return Array.isArray(res) ? res : (res || [])
}

export const addAutoCaseStepApi = async (autoCaseStepId, apiId, sort) => {
  const res = await callApp('AddAutoCaseStepApi', autoCaseStepId, apiId, sort)
  return res || {}
}

export const deleteAutoCaseStepApi = async (autoStepId) => {
  return await callApp('DeleteAutoCaseStepApi', autoStepId)
}

export const sortAutoCaseStepApis = async (data) => {
  return await callApp('SortAutoCaseStepApis', Array.isArray(data) ? data : [])
}

export const getAutoCaseSteps = async (autoCaseId) => {
  const res = await callApp('GetAutoCaseSteps', autoCaseId)
  return Array.isArray(res) ? res : (res || [])
}

export const addAutoCaseStep = async (caseId, stepId) => {
  return await callApp('AddAutoCaseStep', caseId, stepId)
}

export const deleteAutoCaseStepRef = async (refId) => {
  return await callApp('DeleteAutoCaseStepRef', refId)
}

export const sortAutoCaseSteps = async (caseId, data) => {
  return await callApp('SortAutoCaseSteps', caseId, Array.isArray(data) ? data : [])
}

export const setAutoCaseStepConfig = async (refId, isConfig, isStepConfig) => {
  return await callApp('SetAutoCaseStepConfig', refId, Boolean(isConfig), Boolean(isStepConfig))
}

export const getTimerTaskCases = async (taskId) => {
  const res = await callApp('GetTimerTaskCases', taskId)
  return Array.isArray(res) ? res : (res || [])
}

export const addTimerTaskCase = async (taskId, caseId) => {
  return await callApp('AddTimerTaskCase', taskId, caseId)
}

export const deleteTimerTaskCaseRef = async (refId) => {
  return await callApp('DeleteTimerTaskCaseRef', refId)
}

export const sortTimerTaskCases = async (taskId, data) => {
  return await callApp('SortTimerTaskCases', taskId, Array.isArray(data) ? data : [])
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
