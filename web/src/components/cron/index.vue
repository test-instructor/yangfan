<template>
  <div class="cron-container">
    <el-popover v-model:visible="popoverVisible" trigger="click" width="700">
      <template #reference>
        <el-input
          :model-value="modelValue"
          placeholder="点击选择Cron表达式"
          readonly
        >
          <template #prefix>
            <el-icon><Calendar /></el-icon>
          </template>
        </el-input>
      </template>

      <div class="cron-content">
        <el-tabs type="border-card">
          <el-tab-pane label="秒">
            <el-radio-group v-model="second.type">
              <div class="item">
                <el-radio label="*">每秒</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-input-number v-model="second.range.start" :min="0" :max="59" size="small" /> -
                  <el-input-number v-model="second.range.end" :min="0" :max="59" size="small" /> 秒
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="/">
                  从
                  <el-input-number v-model="second.loop.start" :min="0" :max="59" size="small" /> 秒开始，每
                  <el-input-number v-model="second.loop.interval" :min="1" :max="59" size="small" /> 秒执行一次
                </el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="second.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="i in 60" :key="i-1" :label="i-1" :value="String(i-1)" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>

          <el-tab-pane label="分">
            <el-radio-group v-model="minute.type">
              <div class="item">
                <el-radio label="*">每分</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-input-number v-model="minute.range.start" :min="0" :max="59" size="small" /> -
                  <el-input-number v-model="minute.range.end" :min="0" :max="59" size="small" /> 分
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="/">
                  从
                  <el-input-number v-model="minute.loop.start" :min="0" :max="59" size="small" /> 分开始，每
                  <el-input-number v-model="minute.loop.interval" :min="1" :max="59" size="small" /> 分执行一次
                </el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="minute.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="i in 60" :key="i-1" :label="i-1" :value="String(i-1)" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>

          <el-tab-pane label="时">
            <el-radio-group v-model="hour.type">
              <div class="item">
                <el-radio label="*">每小时</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-input-number v-model="hour.range.start" :min="0" :max="23" size="small" /> -
                  <el-input-number v-model="hour.range.end" :min="0" :max="23" size="small" /> 时
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="/">
                  从
                  <el-input-number v-model="hour.loop.start" :min="0" :max="23" size="small" /> 时开始，每
                  <el-input-number v-model="hour.loop.interval" :min="1" :max="23" size="small" /> 小时执行一次
                </el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="hour.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="i in 24" :key="i-1" :label="i-1" :value="String(i-1)" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>

          <el-tab-pane label="日">
            <el-radio-group v-model="day.type">
              <div class="item">
                <el-radio label="*">每日</el-radio>
              </div>
              <div class="item">
                <el-radio label="?">不指定</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-input-number v-model="day.range.start" :min="1" :max="31" size="small" /> -
                  <el-input-number v-model="day.range.end" :min="1" :max="31" size="small" /> 日
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="/">
                  从
                  <el-input-number v-model="day.loop.start" :min="1" :max="31" size="small" /> 日开始，每
                  <el-input-number v-model="day.loop.interval" :min="1" :max="31" size="small" /> 日执行一次
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="W">
                  每月
                  <el-input-number v-model="day.work.day" :min="1" :max="31" size="small" /> 号最近的那个工作日
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="L">本月最后一天</el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="day.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="i in 31" :key="i" :label="i" :value="String(i)" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>

          <el-tab-pane label="月">
            <el-radio-group v-model="month.type">
              <div class="item">
                <el-radio label="*">每月</el-radio>
              </div>
              <div class="item">
                <el-radio label="?">不指定</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-input-number v-model="month.range.start" :min="1" :max="12" size="small" /> -
                  <el-input-number v-model="month.range.end" :min="1" :max="12" size="small" /> 月
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="/">
                  从
                  <el-input-number v-model="month.loop.start" :min="1" :max="12" size="small" /> 月开始，每
                  <el-input-number v-model="month.loop.interval" :min="1" :max="12" size="small" /> 月执行一次
                </el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="month.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="i in 12" :key="i" :label="i" :value="String(i)" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>

          <el-tab-pane label="周">
            <el-radio-group v-model="week.type">
              <div class="item">
                <el-radio label="*">每周</el-radio>
              </div>
              <div class="item">
                <el-radio label="?">不指定</el-radio>
              </div>
              <div class="item">
                <el-radio label="-">
                  周期从
                  <el-select v-model="week.range.start" placeholder="周几" size="small" style="width: 100px">
                    <el-option v-for="(v, k) in weekMap" :key="k" :label="v" :value="k" />
                  </el-select> -
                  <el-select v-model="week.range.end" placeholder="周几" size="small" style="width: 100px">
                    <el-option v-for="(v, k) in weekMap" :key="k" :label="v" :value="k" />
                  </el-select>
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="#">
                  第
                  <el-input-number v-model="week.nth.nth" :min="1" :max="5" size="small" /> 周的
                  <el-select v-model="week.nth.day" placeholder="周几" size="small" style="width: 100px">
                    <el-option v-for="(v, k) in weekMap" :key="k" :label="v" :value="k" />
                  </el-select>
                </el-radio>
              </div>
              <div class="item">
                <el-radio label="L">
                  本月最后一个
                  <el-select v-model="week.last" placeholder="周几" size="small" style="width: 100px">
                    <el-option v-for="(v, k) in weekMap" :key="k" :label="v" :value="k" />
                  </el-select>
                </el-radio>
              </div>
              <div class="item">
                <el-radio label=",">
                  指定
                  <el-select v-model="week.list" multiple placeholder="请选择" size="small" style="width: 200px">
                    <el-option v-for="(v, k) in weekMap" :key="k" :label="v" :value="k" />
                  </el-select>
                </el-radio>
              </div>
            </el-radio-group>
          </el-tab-pane>
        </el-tabs>
        <div class="footer-btn">
          <el-button size="small" @click="popoverVisible = false">取消</el-button>
          <el-button type="primary" size="small" @click="handleConfirm">确定</el-button>
        </div>
      </div>
    </el-popover>
  </div>
</template>

<script setup>
import { ref, reactive, watch, onMounted } from 'vue'
import { Calendar } from '@element-plus/icons-vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: '* * * * * ?'
  }
})

const emit = defineEmits(['update:modelValue'])

const popoverVisible = ref(false)

const weekMap = {
  '1': '周日',
  '2': '周一',
  '3': '周二',
  '4': '周三',
  '5': '周四',
  '6': '周五',
  '7': '周六'
}

const second = reactive({
  type: '*',
  range: { start: 0, end: 0 },
  loop: { start: 0, interval: 1 },
  list: []
})

const minute = reactive({
  type: '*',
  range: { start: 0, end: 0 },
  loop: { start: 0, interval: 1 },
  list: []
})

const hour = reactive({
  type: '*',
  range: { start: 0, end: 0 },
  loop: { start: 0, interval: 1 },
  list: []
})

const day = reactive({
  type: '*',
  range: { start: 1, end: 1 },
  loop: { start: 1, interval: 1 },
  work: { day: 1 },
  list: []
})

const month = reactive({
  type: '*',
  range: { start: 1, end: 1 },
  loop: { start: 1, interval: 1 },
  list: []
})

const week = reactive({
  type: '?',
  range: { start: '2', end: '2' },
  nth: { nth: 1, day: '2' },
  last: '2',
  list: []
})

// 解析Cron表达式
const parseCron = (val) => {
  if (!val) return
  const arr = val.split(' ')
  if (arr.length < 6) return

  // 秒
  parseItem(arr[0], second)
  // 分
  parseItem(arr[1], minute)
  // 时
  parseItem(arr[2], hour)
  // 日
  parseItem(arr[3], day)
  // 月
  parseItem(arr[4], month)
  // 周
  parseItem(arr[5], week)
}

const parseItem = (str, obj) => {
  if (str === '*') {
    obj.type = '*'
  } else if (str === '?') {
    obj.type = '?'
  } else if (str.includes('-')) {
    obj.type = '-'
    const [start, end] = str.split('-')
    obj.range.start = isNaN(start) ? start : Number(start)
    obj.range.end = isNaN(end) ? end : Number(end)
  } else if (str.includes('/')) {
    obj.type = '/'
    const [start, interval] = str.split('/')
    obj.loop.start = isNaN(start) ? start : Number(start)
    obj.loop.interval = Number(interval)
  } else if (str.includes('W')) {
    obj.type = 'W'
    obj.work.day = Number(str.replace('W', ''))
  } else if (str === 'L') {
    obj.type = 'L'
  } else if (str.includes('L') && obj === week) {
    obj.type = 'L'
    obj.last = str.replace('L', '')
  } else if (str.includes('#')) {
    obj.type = '#'
    const [day, nth] = str.split('#')
    obj.nth.day = day
    obj.nth.nth = Number(nth)
  } else {
    obj.type = ','
    obj.list = str.split(',')
  }
}

// 生成Cron表达式
const generateCron = () => {
  const s = generateItem(second)
  const m = generateItem(minute)
  const h = generateItem(hour)
  const d = generateItem(day)
  const M = generateItem(month)
  const w = generateItem(week)
  return `${s} ${m} ${h} ${d} ${M} ${w}`
}

const generateItem = (obj) => {
  switch (obj.type) {
    case '*':
      return '*'
    case '?':
      return '?'
    case '-':
      return `${obj.range.start}-${obj.range.end}`
    case '/':
      return `${obj.loop.start}/${obj.loop.interval}`
    case ',':
      return obj.list.join(',') || '*'
    case 'W':
      return `${obj.work.day}W`
    case 'L':
      return obj === week ? `${obj.last}L` : 'L'
    case '#':
      return `${obj.nth.day}#${obj.nth.nth}`
    default:
      return '*'
  }
}

const handleConfirm = () => {
  emit('update:modelValue', generateCron())
  popoverVisible.value = false
}

watch(() => popoverVisible.value, (val) => {
  if (val) {
    parseCron(props.modelValue)
  }
})

onMounted(() => {
  parseCron(props.modelValue)
})
</script>

<style scoped>
.cron-container {
  width: 100%;
}
.item {
  margin-bottom: 10px;
}
.footer-btn {
  text-align: right;
  margin-top: 10px;
}
</style>
