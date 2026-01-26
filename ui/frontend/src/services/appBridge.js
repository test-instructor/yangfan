const callApp = (method, ...args) => {
  const fn = window?.go?.main?.App?.[method]
  if (!fn) {
    throw new Error(`Go 方法不可用: ${method}`)
  }
  return fn(...args)
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
