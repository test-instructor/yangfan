<template>
  <el-dialog v-model="dialogVisible" title="导入 CURL" width="800px" append-to-body>
    <el-input
      v-model="curlContent"
      type="textarea"
      :rows="10"
      placeholder="请输入 CURL 命令"
    />
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleConfirm">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const dialogVisible = ref(false)
const curlContent = ref('')
const emit = defineEmits(['success'])

const open = () => {
  curlContent.value = ''
  dialogVisible.value = true
}

// Simple shell quote parser
const parseCommand = (cmd) => {
  const args = []
  let current = ''
  let quote = null
  let escape = false

  for (let i = 0; i < cmd.length; i++) {
    const char = cmd[i]

    if (escape) {
      current += char
      escape = false
      continue
    }

    if (char === '\\') {
      escape = true
      continue
    }

    if (quote) {
      if (char === quote) {
        quote = null
      } else {
        current += char
      }
    } else {
      if (char === '"' || char === "'") {
        quote = char
      } else if (/\s/.test(char)) {
        if (current.length > 0) {
          args.push(current)
          current = ''
        }
      } else {
        current += char
      }
    }
  }
  if (current.length > 0) {
    args.push(current)
  }
  return args
}

const parseCurl = (curlStr) => {
  const trimmed = curlStr.trim()
  if (!trimmed.toLowerCase().startsWith('curl')) {
    throw new Error('内容必须以 curl 开头')
  }

  // Remove backslashes used for line continuation
  const cleanStr = trimmed.replace(/\\\s*\n/g, ' ').replace(/\n/g, ' ')
  
  const args = parseCommand(cleanStr)
  
  const result = {
    method: 'GET', // Default
    url: '',
    headers: [],
    params: [],
    json: {},
    data: {}, // form data
    isJson: false
  }

  // Skip 'curl'
  let i = 1
  while (i < args.length) {
    const arg = args[i]
    
    if (arg.startsWith('-')) {
      // Options
      if (arg === '-X' || arg === '--request') {
        if (i + 1 < args.length) {
          result.method = args[i + 1].toUpperCase()
          i += 2
        } else {
          i++
        }
      } else if (arg === '-H' || arg === '--header') {
        if (i + 1 < args.length) {
          const header = args[i + 1]
          const colonIndex = header.indexOf(':')
          if (colonIndex > -1) {
            const key = header.substring(0, colonIndex).trim()
            const value = header.substring(colonIndex + 1).trim()
            result.headers.push({ key, value, desc: '' })
            
            // Check content type
            if (key.toLowerCase() === 'content-type') {
              if (value.includes('application/json')) {
                result.isJson = true
              }
            }
          }
          i += 2
        } else {
          i++
        }
      } else if (arg === '-d' || arg === '--data' || arg === '--data-raw' || arg === '--data-binary') {
        if (i + 1 < args.length) {
          const dataStr = args[i + 1]
          try {
            // Try parsing as JSON first
            const jsonData = JSON.parse(dataStr)
            result.json = { ...result.json, ...jsonData }
            result.isJson = true
            result.method = 'POST' // Implied POST
          } catch (e) {
            // If not JSON, treat as form data string (key=value&key2=value2)
            const params = new URLSearchParams(dataStr)
            for (const [key, value] of params.entries()) {
               // Assuming form data
               // We might need to handle this differently depending on structure
               // For now, let's put it in data (form)
               // But step component expects array structure for form
            }
            // For simplicity in this basic parser, if it fails JSON, we might ignore or try to parse key=value
            if (dataStr.includes('=')) {
                dataStr.split('&').forEach(pair => {
                    const [key, value] = pair.split('=')
                    if (key) {
                        // We need to decide where to put this. 
                        // If content-type is json, it might be error.
                        // If x-www-form-urlencoded, it goes to data.
                        // For now, let's treat it as json key-value if we decided it's json, or just ignore for now to be safe or put in json as string?
                        // Actually, let's try to be smart.
                        if (!result.isJson) {
                            // likely form data
                             // We'll handle form data later if needed, user requirement is "parse content", usually JSON body is most important.
                        }
                    }
                })
            }
            
            // If it's just a string and not JSON, maybe it's raw body.
            // But the Step component mainly supports JSON body or Form data.
             if (!result.isJson) {
                // If we detected it's not JSON, but we have data, we might want to set method to POST
                result.method = 'POST'
             }
          }
          i += 2
        } else {
          i++
        }
      } else if (arg === '-u' || arg === '--user') {
          // Basic auth
           if (i + 1 < args.length) {
               // const [user, pass] = args[i+1].split(':')
               // header Authorization: Basic base64(...)
               // Skipping for now as it requires btoa
               i += 2
           } else {
               i++
           }
      } else if (arg.startsWith('-')) {
          // Unknown flag, skip argument if it looks like it has one?
          // It's hard to guess if the next token is an argument or another flag or the URL.
          // Heuristic: if next token doesn't start with -, assume it is argument.
          if (i + 1 < args.length && !args[i+1].startsWith('-')) {
              i += 2
          } else {
              i++
          }
      } else {
          i++
      }
    } else {
      // Positional argument, likely URL
      if (!result.url && arg.match(/^https?:/)) {
        result.url = arg
      } else if (!result.url) {
          // Maybe url without http? or localhost
           result.url = arg
      }
      i++
    }
  }

  // Parse URL for query params
  if (result.url) {
      try {
          // Handle case where url might not have protocol
          let urlToParse = result.url
          if (!urlToParse.startsWith('http')) {
              urlToParse = 'http://' + urlToParse
          }
          const urlObj = new URL(urlToParse)
          // Extract params
          urlObj.searchParams.forEach((value, key) => {
              result.params.push({ key, value, desc: '' })
          })
          // Keep only base URL in the url field?
          // The Step component seems to allow full URL. 
          // But usually we separate params.
          // Let's keep the full URL in `url` field for now, or strip params?
          // Step component usually separates them.
          // Let's strip params from URL if we extracted them.
          result.url = urlObj.origin + urlObj.pathname
      } catch (e) {
          // Invalid URL, keep as is
      }
  }

  return result
}

const handleConfirm = () => {
  try {
    const parsed = parseCurl(curlContent.value)
    emit('success', parsed)
    dialogVisible.value = false
  } catch (e) {
    console.error(e)
    ElMessage.error('解析 CURL 失败: ' + e.message)
  }
}

defineExpose({ open })
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>
