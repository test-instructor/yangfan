//参考：
// https://github.com/vuejs/vuepress/blob/master/packages/docs/docs/.vuepress/config.js
// https://vuepress-theme-reco.recoluan.com/views/1.x/
module.exports = {
	title: '扬帆测试平台 官方文档',
	description: '扬帆测试平台，由个人开发的项目，已入选本年度码云最有价值开源项目',
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
				}
			]
		},
		sidebarDepth: 1
	},

	head: [
		['link', {rel: 'icon', href: '/logo.png'}],
		['script', {}, `
            var _hmt = _hmt || [];
            (function() {
              var hm = document.createElement("script");
              hm.src = "https://hm.baidu.com/hm.js?13a39a1b1e7fb17e8f806d1fb6207796";
              var s = document.getElementsByTagName("script")[0]; 
              s.parentNode.insertBefore(hm, s);
            })();
        `],
	]
}