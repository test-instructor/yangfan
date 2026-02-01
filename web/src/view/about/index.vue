<template>
  <div class="p-4">
    <div class="bg-white dark:bg-slate-800 rounded-xl p-6">
      <div class="flex items-center justify-between gap-4">
        <div>
          <div class="text-2xl font-bold text-gray-800 dark:text-gray-100">
            关于我们
          </div>
          <div class="mt-2 text-sm text-gray-500 dark:text-gray-400">
            扬帆（Yangfan）自动化测试平台 · 让回归更高效、资产更可复用、质量更可度量
          </div>
        </div>
      </div>

      <el-divider />

      <div class="space-y-6 text-gray-700 dark:text-gray-200">
        <section>
          <div class="text-lg font-semibold mb-3">项目介绍</div>
          <div class="leading-7 text-gray-600 dark:text-gray-300">
            扬帆是一套面向企业研发团队的接口自动化测试平台，围绕“用例资产沉淀、任务调度执行、测试数据治理、报告与通知闭环”搭建一体化能力。
            平台以 HttpRunner V5 为核心运行引擎，通过标准化的资产管理与可扩展的运行能力，把零散的脚本执行升级为可复用、可追踪、可治理的质量体系。
          </div>
        </section>

        <section>
          <div class="text-lg font-semibold mb-3">核心能力</div>
          <ul class="list-disc pl-5 space-y-2 text-gray-600 dark:text-gray-300">
            <li>接口与用例资产管理：接口编排、步骤复用、参数化/断言/数据传递、CURL 导入与单次执行查看。</li>
            <li>回归执行与调度：定时任务、按标签/配置触发回归、并发执行与执行资源控制。</li>
            <li>报告与通知闭环：报告查看与统计，支持飞书/钉钉/企业微信等通知方式。</li>
            <li>平台扩展与运行保障：运行节点管理、运行配置统一管理，支持 Python 函数插件与第三方依赖管理。</li>
            <li>测试数据治理：数据分类与沉淀，支撑测试数据复用与规范化管理。</li>
            <li>项目协作与权限：项目配置、成员与权限管理，满足多项目团队协作场景。</li>
          </ul>
        </section>

        <section>
          <div class="text-lg font-semibold mb-3">技术栈</div>
          <div class="flex flex-wrap gap-2">
            <el-tag effect="plain" type="success">Go</el-tag>
            <el-tag effect="plain" type="success">Gin</el-tag>
            <el-tag effect="plain" type="info">Vue 3</el-tag>
            <el-tag effect="plain" type="info">Vite</el-tag>
            <el-tag effect="plain" type="warning">Element Plus</el-tag>
            <el-tag effect="plain">HttpRunner V5</el-tag>
          </div>
        </section>

        <section>
          <div class="text-lg font-semibold mb-3">作者介绍</div>
          <div class="leading-7 text-gray-600 dark:text-gray-300">
            本项目由作者长期维护与迭代，聚焦自动化测试与研发效能平台化建设，目标是让接口自动化从“能跑”走向“可用、可复用、可运营”。
            如果你正在推进回归体系建设、需要多项目协作治理、或希望把测试资产纳入团队工程化流程，欢迎一起交流与共建。
          </div>
        </section>

        <section>
          <el-divider content-position="center">交流与支持</el-divider>
          <div class="text-sm text-gray-500 dark:text-gray-400 mb-4">
            扫码联系作者或加入交流群（二维码位于本文末尾）
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div
              class="border border-gray-200 dark:border-slate-700 rounded-lg p-4 flex flex-col items-center bg-gray-50 dark:bg-slate-900"
            >
              <div class="font-medium text-gray-700 dark:text-gray-200 mb-3">
                作者二维码
              </div>
              <el-image
                class="w-48 h-48 rounded"
                :src="authorQrUrl"
                fit="cover"
                :preview-src-list="[authorQrUrl]"
                preview-teleported
              />
            </div>

            <div
              class="border border-gray-200 dark:border-slate-700 rounded-lg p-4 flex flex-col items-center bg-gray-50 dark:bg-slate-900"
            >
              <div class="font-medium text-gray-700 dark:text-gray-200 mb-3">
                群二维码
              </div>
              <el-image
                class="w-48 h-48 rounded"
                :src="groupQrUrl"
                fit="cover"
                :preview-src-list="[groupQrUrl]"
                preview-teleported
              />
            </div>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
  const authorQrUrl = 'http://qiniu.yangfan.gd.cn/about/author.jpeg'
  const groupQrUrl = 'http://qiniu.yangfan.gd.cn/about/group.jpeg'
</script>
