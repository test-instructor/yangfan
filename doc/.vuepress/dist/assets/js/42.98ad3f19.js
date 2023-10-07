(window.webpackJsonp=window.webpackJsonp||[]).push([[42],{467:function(s,t,a){"use strict";a.r(t);var n=a(2),e=Object(n.a)({},(function(){var s=this,t=s._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":s.$parent.slotKey}},[t("h2",{attrs:{id:"前端环境"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#前端环境"}},[s._v("#")]),s._v(" 前端环境")]),s._v(" "),t("ol",[t("li",[s._v("前往https://nodejs.org/zh-cn/下载当前版本node")]),s._v(" "),t("li",[s._v("命令行运行 node -v 若控制台输出版本号则前端环境搭建成功")]),s._v(" "),t("li",[s._v("node 版本需大于 16.4")]),s._v(" "),t("li",[s._v("开发工具推荐vscode https://code.visualstudio.com/")]),s._v(" "),t("li",[s._v("安装依赖"),t("div",{staticClass:"language-shell line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" web\n"),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("npm")]),s._v(" "),t("span",{pre:!0,attrs:{class:"token function"}},[s._v("install")]),s._v("\n")])]),s._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[s._v("1")]),t("br"),t("span",{staticClass:"line-number"},[s._v("2")]),t("br")])])])]),s._v(" "),t("h2",{attrs:{id:"后端环境"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#后端环境"}},[s._v("#")]),s._v(" 后端环境")]),s._v(" "),t("ol",[t("li",[s._v("下载golang安装 版本号需>=1.19\n"),t("ul",[t("li",[s._v("国际: https://golang.org/dl/")]),s._v(" "),t("li",[s._v("国内: https://golang.google.cn/dl/")])])]),s._v(" "),t("li",[s._v("命令行运行 go 若控制台输出各类提示命令 则安装成功 输入 go version 确认版本大于 1.18")]),s._v(" "),t("li",[s._v("开发工具推荐 Goland")]),s._v(" "),t("li",[s._v("安装依赖"),t("div",{staticClass:"language-shell line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" server\ngo mod tidy\n")])]),s._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[s._v("1")]),t("br"),t("span",{staticClass:"line-number"},[s._v("2")]),t("br")])])])]),s._v(" "),t("h3",{attrs:{id:"服务初始化"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#服务初始化"}},[s._v("#")]),s._v(" 服务初始化")]),s._v(" "),t("ol",[t("li",[s._v("服务列表\n"),t("ul",[t("li",[s._v("server 后端服务")]),s._v(" "),t("li",[s._v("run 用例运行服务")]),s._v(" "),t("li",[s._v("timer 定时任务服务")]),s._v(" "),t("li",[s._v("master 性能测试master服务")]),s._v(" "),t("li",[s._v("work 性能测试worker服务")]),s._v(" "),t("li",[s._v("web 前端服务")])])]),s._v(" "),t("li",[s._v("安装依赖"),t("div",{staticClass:"language-shell line-numbers-mode"},[t("pre",{pre:!0,attrs:{class:"language-shell"}},[t("code",[t("span",{pre:!0,attrs:{class:"token builtin class-name"}},[s._v("cd")]),s._v(" server "),t("span",{pre:!0,attrs:{class:"token comment"}},[s._v("# 其他服务同理，每个服务都需要安装")]),s._v("\ngo mod tidy\n")])]),s._v(" "),t("div",{staticClass:"line-numbers-wrapper"},[t("span",{staticClass:"line-number"},[s._v("1")]),t("br"),t("span",{staticClass:"line-number"},[s._v("2")]),t("br")])])])]),s._v(" "),t("h3",{attrs:{id:"本地调试"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#本地调试"}},[s._v("#")]),s._v(" 本地调试")]),s._v(" "),t("h4",{attrs:{id:"软件包路径"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#软件包路径"}},[s._v("#")]),s._v(" 软件包路径")]),s._v(" "),t("ul",[t("li",[s._v("server ：github.com/test-instructor/yangfan/server")]),s._v(" "),t("li",[s._v("run ：github.com/test-instructor/yangfan/run")]),s._v(" "),t("li",[s._v("timer ：github.com/test-instructor/yangfan/timer")]),s._v(" "),t("li",[s._v("master ：github.com/test-instructor/yangfan/master")]),s._v(" "),t("li",[s._v("work：github.com/test-instructor/yangfan/work")])]),s._v(" "),t("h4",{attrs:{id:"添加配置"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#添加配置"}},[s._v("#")]),s._v(" 添加配置")]),s._v(" "),t("ol",[t("li",[s._v("进入编辑配置页面")]),s._v(" "),t("li",[s._v("添加“Go 构建”的配置")]),s._v(" "),t("li",[s._v("运行种类选择软件包")]),s._v(" "),t("li",[s._v("软件包路径选择对应的服务")]),s._v(" "),t("li",[s._v("名称默认为软件包路径，为了更加简洁可以直接改成对应服务的名称")])]),s._v(" "),t("p",[t("img",{attrs:{src:"http://qiniu.yangfan.gd.cn//markdown/image-20230928162540358.png",alt:"image-20230928162540358"}})]),s._v(" "),t("p",[t("img",{attrs:{src:"http://qiniu.yangfan.gd.cn//markdown/image-20230928162759825.png",alt:"image-20230928162759825"}})]),s._v(" "),t("p",[t("img",{attrs:{src:"http://qiniu.yangfan.gd.cn//markdown/image-20230928162928134.png",alt:"image-20230928162928134"}})])])}),[],!1,null,null,null);t.default=e.exports}}]);