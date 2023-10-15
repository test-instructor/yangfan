//参考：
// https://github.com/vuejs/vuepress/blob/master/packages/docs/docs/.vuepress/config.js
// https://vuepress-theme-reco.recoluan.com/views/1.x/
module.exports = {
	title: '扬帆测试平台 官方文档',
	description: '扬帆测试平台，由个人开发的项目，已入选本年度码云最有价值开源项目。扬帆测试平台是一款高效、可靠的自动化测试平台，旨在帮助团队提升测试效率、降低测试成本。该平台包括用例管理、定时任务、执行记录等功能模块，支持多种类型的测试用例，目前支持API(http和grpc协议)、性能、CI调用等功能，并且可定制化，灵活满足不同场景的需求。 其中，支持批量执行、并发执行等高级功能。通过用例设置，可以设置用例的基本信息、运行配置、环境变量等，灵活控制用例的执行。',
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
				},{
					title: '操作手册',
					collapsable: true,
					children: [
						{title: '配置管理', path: '/documentation/config'},
						{title: '环境变量', path: '/documentation/env'},
						{title: '用例管理', path: '/documentation/case'},
						{title: '定时任务', path: '/documentation/task'},
						{title: '测试报告', path: '/documentation/report'},
						{title: '性能测试', path: '/documentation/performance'},
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
				},
				{
					title: '定制开发',
					collapsable: true,
					children: [
						{title: '编译&交叉编译', path: '/hrp/development/compile'},
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