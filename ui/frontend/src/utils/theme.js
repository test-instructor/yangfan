const THEME_STORAGE_KEY = 'yangfan:theme'

export const ThemeMode = Object.freeze({
  light: 'light',
  dark: 'dark'
})

export function getStoredTheme() {
  try {
    const value = localStorage.getItem(THEME_STORAGE_KEY)
    if (value === ThemeMode.dark || value === ThemeMode.light) return value
  } catch (e) {
    return ThemeMode.light
  }
  return ThemeMode.light
}

export function applyTheme(mode) {
  const resolved = mode === ThemeMode.dark ? ThemeMode.dark : ThemeMode.light
  const root = document.documentElement
  const body = document.body

  if (resolved === ThemeMode.dark) {
    root.setAttribute('arco-theme', 'dark')
    body.setAttribute('arco-theme', 'dark')
  } else {
    root.removeAttribute('arco-theme')
    body.removeAttribute('arco-theme')
  }
}

export function setStoredTheme(mode) {
  const resolved = mode === ThemeMode.dark ? ThemeMode.dark : ThemeMode.light
  try {
    localStorage.setItem(THEME_STORAGE_KEY, resolved)
  } catch (e) {
    return
  }
}

export function setTheme(mode) {
  setStoredTheme(mode)
  applyTheme(mode)
}

export function initTheme() {
  applyTheme(getStoredTheme())
}

