<template>
  <div class="flex flex-col md:flex-row justify-around items-center h-auto md:h-64 gap-4">
    <!-- 用例统计 -->
    <div class="flex flex-col items-center w-full md:w-1/3">
       <div class="font-bold mb-2">用例统计</div>
       <div ref="caseChartRef" style="width: 100%; height: 180px;"></div>
       <div class="text-sm mt-2">
         总数: {{ testcases?.total || 0 }} | 
         成功: {{ testcases?.success || 0 }} | 
         失败: {{ testcases?.fail || 0 }}
       </div>
    </div>
    <!-- 步骤统计 -->
    <div class="flex flex-col items-center w-full md:w-1/3">
        <div class="font-bold mb-2">步骤统计</div>
        <div ref="stepChartRef" style="width: 100%; height: 180px;"></div>
        <div class="text-sm mt-2">
          总数: {{ teststeps?.total || 0 }} | 
          成功: {{ teststeps?.successes || 0 }} | 
          失败: {{ teststeps?.failures || 0 }}
        </div>
    </div>
    <!-- 接口统计 -->
    <div class="flex flex-col items-center w-full md:w-1/3">
        <div class="font-bold mb-2">接口统计</div>
        <div ref="apiChartRef" style="width: 100%; height: 180px;"></div>
        <div class="text-sm mt-2">
           总数: {{ apiStats?.total || 0 }} | 
           成功: {{ apiStats?.success || 0 }} | 
           失败: {{ apiStats?.fail || 0 }}
        </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  testcases: {
    type: Object,
    default: () => ({})
  },
  teststeps: {
    type: Object,
    default: () => ({})
  },
  apiStats: {
    type: Object,
    default: () => ({})
  }
})

const caseChartRef = ref(null)
const stepChartRef = ref(null)
const apiChartRef = ref(null)

let charts = []

const initChart = (el, name, success, fail) => {
    if (!el) return
    let chart = echarts.getInstanceByDom(el)
    if (!chart) {
        chart = echarts.init(el)
        charts.push(chart)
    }
    const option = {
        tooltip: {
            trigger: 'item'
        },
        legend: {
            top: '5%',
            left: 'center'
        },
        series: [
            {
                name: name,
                type: 'pie',
                radius: ['40%', '70%'],
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
                        fontWeight: 'bold'
                    }
                },
                labelLine: {
                    show: false
                },
                data: [
                    { value: success, name: '成功', itemStyle: { color: '#67C23A' } },
                    { value: fail, name: '失败', itemStyle: { color: '#F56C6C' } }
                ]
            }
        ]
    }
    chart.setOption(option)
}

const renderAllCharts = () => {
    initChart(caseChartRef.value, '用例', 
        props.testcases?.success || 0, 
        props.testcases?.fail || 0
    )
    initChart(stepChartRef.value, '步骤',
        props.teststeps?.successes || 0,
        props.teststeps?.failures || 0
    )
    initChart(apiChartRef.value, '接口',
        props.apiStats?.success || 0,
        props.apiStats?.fail || 0
    )
}

watch(() => props, () => {
    nextTick(() => {
        renderAllCharts()
    })
}, { deep: true })

const handleResize = () => {
    charts.forEach(c => c.resize())
}

onMounted(() => {
    nextTick(() => {
        renderAllCharts()
    })
    window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize)
    charts.forEach(c => c.dispose())
    charts = []
})
</script>

<style scoped>
</style>
