<template>
  <el-card class="box-card mt-4" shadow="hover">
    <template #header>
      <span class="text-lg font-bold">运行统计</span>
    </template>
    <div class="flex flex-col md:flex-row justify-around items-center h-auto md:h-64 gap-8 md:gap-0 py-4 md:py-0">
      <!-- 用例统计 -->
      <div class="flex flex-col items-center w-full md:w-1/3">
        <div class="font-bold mb-2">用例统计</div>
        <div ref="caseChartRef" style="width: 100%; height: 180px;"></div>
        <div class="text-sm mt-2">
          总数: {{ detail.stat?.testcases?.total || 0 }} | 
          成功: {{ detail.stat?.testcases?.success || 0 }} | 
          失败: {{ detail.stat?.testcases?.fail || 0 }}
        </div>
      </div>
      <!-- 步骤统计 -->
      <div class="flex flex-col items-center w-full md:w-1/3">
        <div class="font-bold mb-2">步骤统计</div>
        <div ref="stepChartRef" style="width: 100%; height: 180px;"></div>
        <div class="text-sm mt-2">
          总数: {{ detail.stat?.teststeps?.total || 0 }} | 
          成功: {{ detail.stat?.teststeps?.successes || 0 }} | 
          失败: {{ detail.stat?.teststeps?.failures || 0 }}
        </div>
      </div>
      <!-- 接口统计 -->
      <div class="flex flex-col items-center w-full md:w-1/3">
        <div class="font-bold mb-2">接口统计</div>
        <div ref="apiChartRef" style="width: 100%; height: 180px;"></div>
        <div class="text-sm mt-2">
          总数: {{ apiStats.total }} | 
          成功: {{ apiStats.success }} | 
          失败: {{ apiStats.fail }}
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup>
import { ref, computed, onMounted, watch, nextTick, onBeforeUnmount } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  detail: {
    type: Object,
    required: true,
    default: () => ({})
  }
})

const caseChartRef = ref(null)
const stepChartRef = ref(null)
const apiChartRef = ref(null)
let charts = []

const apiStats = computed(() => {
  const api = props.detail.stat?.teststepapi || {}
  return {
    total: api.total || 0,
    success: api.success || 0,
    fail: api.fail || 0
  }
})

const renderChart = (el, name, success, fail) => {
  if (!el) return
  let chart = echarts.getInstanceByDom(el)
  if (!chart) {
    chart = echarts.init(el)
    charts.push(chart)
  }
  
  const total = success + fail
  const successRate = total > 0 ? Math.round((success / total) * 100) + '%' : '0%'

  const option = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    toolbox: {
      show: true,
      feature: {
        saveAsImage: { show: true, title: '保存图片' }
      }
    },
    legend: {
      top: '5%',
      left: 'center'
    },
    series: [
      {
        name: name,
        type: 'pie',
        radius: ['50%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 20,
            fontWeight: 'bold',
            formatter: '{b}\n{d}%'
          },
          scale: true,
          scaleSize: 10
        },
        labelLine: {
          show: false
        },
        data: [
          { value: success, name: '成功', itemStyle: { color: '#67C23A' } },
          { value: fail, name: '失败', itemStyle: { color: '#F56C6C' } }
        ]
      }
    ],
    graphic: {
      type: 'text',
      left: 'center',
      top: 'center',
      style: {
        text: successRate,
        textAlign: 'center',
        fill: '#333',
        fontSize: 24,
        fontWeight: 'bold'
      }
    }
  }
  chart.setOption(option)
}

const initCharts = () => {
  nextTick(() => {
    renderChart(caseChartRef.value, '用例', 
      props.detail.stat?.testcases?.success || 0, 
      props.detail.stat?.testcases?.fail || 0
    )
    renderChart(stepChartRef.value, '步骤',
      props.detail.stat?.teststeps?.successes || 0,
      props.detail.stat?.teststeps?.failures || 0
    )
    renderChart(apiChartRef.value, '接口',
      apiStats.value.success,
      apiStats.value.fail
    )
  })
}

watch(() => props.detail, () => {
  initCharts()
}, { deep: true })

onMounted(() => {
  initCharts()
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  charts.forEach(chart => chart.dispose())
})

const handleResize = () => {
  charts.forEach(chart => chart.resize())
}
</script>
