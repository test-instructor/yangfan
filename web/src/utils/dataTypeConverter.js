const convertData = (dataList) => {
  const result = {
    success: true,
    data: {},
    errors: {}
  }

  dataList.forEach(item => {
    const { type, key, value } = item
    try {
      let convertedValue

      switch (type) {
        case 'String':
          convertedValue = String(value)
          break
        case 'Integer':
          const intVal = parseInt(value, 10)
          if (isNaN(intVal) || String(intVal) !== String(value).trim()) {
            throw new Error('不是有效的整数')
          }
          convertedValue = intVal
          break
        case 'Float':
          const floatVal = parseFloat(value)
          if (isNaN(floatVal)) {
            throw new Error('不是有效的浮点数')
          }
          convertedValue = floatVal
          break
        case 'Boolean':
          const lowerVal = String(value).toLowerCase()
          if (lowerVal === 'true') {
            convertedValue = true
          } else if (lowerVal === 'false') {
            convertedValue = false
          } else {
            throw new Error('不是有效的布尔值（true/false）')
          }
          break
        case 'List':
          try {
            convertedValue = JSON.parse(value)
            if (!Array.isArray(convertedValue)) {
              throw new Error('不是有效的列表')
            }
          } catch (e) {
            throw new Error('列表格式错误：' + e.message)
          }
          break
        case 'Dict':
          try {
            convertedValue = JSON.parse(value)
            if (typeof convertedValue !== 'object' || convertedValue === null || Array.isArray(convertedValue)) {
              throw new Error('不是有效的字典')
            }
          } catch (e) {
            throw new Error('字典格式错误：' + e.message)
          }
          break
        case 'None':
          if (value !== '' && value !== 'null' && value !== 'None') {
            throw new Error('None类型值必须为空、null或None')
          }
          convertedValue = null
          break
        default:
          throw new Error('不支持的数据类型')
      }

      result.data[key] = convertedValue
    } catch (error) {
      result.success = false
      result.errors = {
        key: key,
        error: error.message
      }
    }
  })

  return result
}

function assertionData(datalist) {
  const result = {
    success: true,
    data: [],
    errors: {}
  }
  const datas = []
  datalist.forEach((item, index) => {
    try {
      let convertedValue
      const expectValue = item.expect
      let data = { ...item }
      switch (item.type) {
        case 'String':
          convertedValue = String(expectValue)
          datas.push(data)
          break
        case 'Integer':
          // 检查是否为有效整数
          const intValue = parseInt(expectValue, 10)
          if (isNaN(intValue) || String(intValue) !== expectValue.trim()) {
            throw new Error(`无法将 "${expectValue}" 转换为整数`)
          }
          convertedValue = intValue
          data.expect = convertedValue
          datas.push(data)
          break
        case 'Float':
          // 检查是否为有效浮点数
          const floatValue = parseFloat(expectValue)
          if (isNaN(floatValue)) {
            throw new Error(`无法将 "${expectValue}" 转换为浮点数`)
          }
          convertedValue = floatValue
          data.expect = convertedValue
          datas.push(data)
          break
        case 'Boolean':
          // 严格匹配 'true' 或 'false'（不区分大小写）
          const lowerExpect = expectValue.toLowerCase()
          if (lowerExpect === 'true') {
            convertedValue = true
          } else if (lowerExpect === 'false') {
            convertedValue = false
          } else {
            throw new Error(`无法将 "${expectValue}" 转换为布尔值，必须是 "true" 或 "false"`)
          }
          onvertedValue = floatValue
          data.expect = convertedValue
          datas.push(data)
          break
        case 'List':
          try {
            // 尝试解析为数组（支持JSON格式的数组字符串）
            convertedValue = JSON.parse(expectValue)
            if (!Array.isArray(convertedValue)) {
              throw new Error()
            }
          } catch {
            throw new Error(`无法将 "${expectValue}" 转换为数组`)
          }
          onvertedValue = floatValue
          data.expect = convertedValue
          datas.push(data)
          break
        case 'Dict':
          try {
            // 尝试解析为对象（支持JSON格式的对象字符串）
            convertedValue = JSON.parse(expectValue)
            if (typeof convertedValue !== 'object' || convertedValue === null || Array.isArray(convertedValue)) {
              throw new Error()
            }
          } catch {
            throw new Error(`无法将 "${expectValue}" 转换为对象`)
          }
          onvertedValue = floatValue
          data.expect = convertedValue
          datas.push(data)
          break
        case 'None':
          // 仅当值为空字符串时转换为null
          if (expectValue !== '') {
            throw new Error(`None类型必须对应空字符串，实际值为 "${expectValue}"`)
          }
          convertedValue = null
          break
        default:
          throw new Error(`不支持的类型: ${item.type}`)
      }
    } catch (error) {
      // 转换失败，记录错误
      console.log('item.check', item)
      console.log('error.message', error.message)
      result.success = false
      result.errors = {
        key: item.check,
        error: error.message
      }
    }
  })
  result.data = datas
  console.log('result', result)
  return result
}

const processData = (data) => {
  const result = {
    success: true,
    data: {},
    errors: []
  }
  const keySet = new Set()

  for (const item of data) {
    const { key, value } = item
    // 检查key是否已存在
    if (keySet.has(key)) {
      result.success = false
      result.errors.push(key)
    } else {
      keySet.add(key)
      result.data[key] = value
    }
  }

  return result
}

export default {
  convertData,
  processData,
  assertionData
}