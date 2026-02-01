//参考：
// https://github.com/vuejs/vuepress/blob/master/packages/docs/docs/.vuepress/config.js
// https://vuepress-theme-reco.recoluan.com/views/1.x/
module.exports = {
	title: '扬帆测试平台 官方文档',
	description: '扬帆测试平台是一款面向团队的自动化测试平台，聚焦“测试资产沉淀、任务调度执行、运行与扩展、报告与通知、测试数据管理”等能力，帮助提升回归效率与质量可控性。',
	// base:'/docs/',
	markdown: {
		lineNumbers: true
	},
	theme: 'vuepress-theme-reco',
	themeConfig: {
		//腾讯 404 公益配置
		noFoundPageByTencent: false,

		mode: 'light',
		modePicker: false,

		subSidebar: 'auto',
		docsBranch: 'master',
		docsDir: 'doc',
		editLinks: true,
		editLinkText: '编辑此页面',

		codeTheme: 'tomorrow',



		lastUpdated: '更新时间',

		nav: [
			{text: '首页', link: '/'},
			{text: '使用文档', link: '/documentation/'},
			{text: 'hrp二次开发', link: '/hrp/'},
			{text: '提问', link: 'https://gitee.com/test-instructor/yangfan/issues'},
			{
				text: '源码下载', items: [
					{text: 'Gitee', link: 'https://gitee.com/test-instructor/yangfan'},
					{text: 'Github', link: 'https://github.com/test-instructor/yangfan'}
				]
			}
		],

		sidebar: {
			'/documentation/':[
				{
					title: '快速入门',
					collapsable: true,
					children: [
						{title: '测试平台简介', path: '/documentation/'},
						{title: '快速开始', path: '/documentation/start'},
						{title: '部署服务', path: '/documentation/deploy'},
						{title: '开发调试', path: '/documentation/debug'},
					],
				},
				{
					title: '项目管理',
					collapsable: true,
					children: [
						{title: '项目配置', path: '/documentation/pm/pj'},
						{title: '项目成员与权限', path: '/documentation/pm/upa'},
						{title: '报告通知', path: '/documentation/pm/reportNotify'},
					],
				},
				{
					title: '配置管理',
					collapsable: true,
					children: [
						{title: '运行配置', path: '/documentation/platform/rc'},
						{title: '环境变量管理', path: '/documentation/platform/envDetail'},
						{
							title: '函数插件',
							collapsable: true,
							children: [
								{title: 'Python 函数', path: '/documentation/platform/FunctionPlugin/pc'},
								{title: '调试信息', path: '/documentation/platform/FunctionPlugin/pcd'},
								{title: 'Python 第三方库', path: '/documentation/platform/FunctionPlugin/pp'},
							],
						},
						{title: '运行节点', path: '/documentation/platform/rn'},
					],
				},
				{
					title: '接口自动化',
					collapsable: true,
					children: [
						{title: '接口管理', path: '/documentation/APIAutomation/as'},
						{title: '测试步骤', path: '/documentation/APIAutomation/acs'},
						{title: '测试用例', path: '/documentation/APIAutomation/ac'},
						{title: '定时任务', path: '/documentation/APIAutomation/tk'},
						{title: '自动报告', path: '/documentation/APIAutomation/ar'},
					],
				},
				{
					title: '数据仓库',
					collapsable: true,
					children: [
						{title: '数据分类管理', path: '/documentation/dataWarehouse/dcm'},
					],
				},
				{
					title: '其他',
					collapsable: true,
					children: [
						{title: '代码测试3', path: '/documentation/pc3'},
					],
				}
			],
			"/hrp/":[
				{
					title: '概述',
					path: '/hrp/',
				},
				{
					title: '源码解析',
					collapsable: true,
					children: [
						{title: '目录结构', path: '/hrp/code/directory'},
						{title: '流程解析', path: '/hrp/code/flow_path'},
						{title: '配置管理', path: '/hrp/code/config'},
						{title: '用例类型', path: '/hrp/code/case_type'},
						{title: '用例读取', path: '/hrp/code/case_read'},
						{title: '用例执行流程', path: '/hrp/code/case_run_flow'},
						{title: '用例执行对象', path: '/hrp/code/case_obj'},
						{title: '用例运行', path: '/hrp/code/case_run'},
					],
				}
			]
		},
		sidebarDepth: 1
	},

	head: [
		['link', {rel: 'icon', href: '/logo.png'}],
	]
}
