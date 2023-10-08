(window.webpackJsonp=window.webpackJsonp||[]).push([[46],{467:function(a,t,r){"use strict";r.r(t);var s=r(2),n=Object(s.a)({},(function(){var a=this,t=a._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":a.$parent.slotKey}},[t("h1",{attrs:{id:"平台简介"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#平台简介"}},[a._v("#")]),a._v(" 平台简介")]),a._v(" "),t("p",[a._v("扬帆测试平台是一款基于 gin-vue-admin 为框架，以 HttpRunner v4 go 模块（以下简称hrp）为测试引擎搭建的自动化测试平台，致力于打造最易使用的开源测试平台。与大多数测试平台不同，扬帆测试平台采用了go语言进行开发，具有良好的性能和稳定性，同时在部署方式和复杂度方面也更加简单，减轻了用户的部署负担。")]),a._v(" "),t("p",[a._v("在设计理念上，扬帆测试平台注重实用性和易用性，平台界面简洁明了，用户可以通过简单的操作完成测试任务的创建、执行、查看和管理。平台提供了完整的测试流程支持，包括测试用例管理、测试计划管理、测试报告生成等，让测试工作更加规范和高效。")]),a._v(" "),t("p",[a._v("作为一款自动化测试平台，扬帆测试平台自然也支持接口自动化测试。平台已经实现了接口自动化测试中最关键的部分，包括测试用例的编写、执行、性能测试和结果分析等。同时，平台还提供了丰富的接口测试功能，包括参数化测试、前置后置处理、断言验证、函数驱动、hooks等，满足用户在接口测试中的不同需求。")]),a._v(" "),t("p",[a._v("除此之外，扬帆测试平台后续将支持多种测试类型，包括UI自动化测试、k8s部署、分布式压测、消息通知等，满足不同场景下的测试需求。平台提供了灵活的扩展机制，用户可以根据实际需求开发自己的测试插件，实现更多的测试类型和功能。")]),a._v(" "),t("p",[a._v("总的来说，扬帆测试平台是一款易用且功能丰富的自动化测试平台，适用于各类软件测试工作，为用户提供高效的测试支持，助力测试工作的顺利进行。")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/%E6%89%AC%E5%B8%86%E6%B5%8B%E8%AF%95%E5%B9%B3%E5%8F%B0%E6%9E%B6%E6%9E%84%E5%9B%BE.jpg",alt:"扬帆测试平台架构图"}})]),a._v(" "),t("h2",{attrs:{id:"部署方式"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#部署方式"}},[a._v("#")]),a._v(" 部署方式")]),a._v(" "),t("p",[a._v("您可以使用Docker和Kubernetes（K8s）。请查看"),t("a",{attrs:{href:"/documentation/deploy"}},[a._v("部署文档")]),a._v("以获取详细的部署说明。当前部署流程需要预先准备MySQL数据库。我们计划在后续版本中添加一键部署脚本，以简化整个部署过程。")]),a._v(" "),t("h1",{attrs:{id:"平台功能"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#平台功能"}},[a._v("#")]),a._v(" 平台功能")]),a._v(" "),t("h2",{attrs:{id:"基础功能"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#基础功能"}},[a._v("#")]),a._v(" 基础功能")]),a._v(" "),t("ol",[t("li",[a._v("API管理：通过鉴权和角色分配，确保只有具备权限的用户可以访问相应的API。")]),a._v(" "),t("li",[a._v("页面管理：通过鉴权和角色分配，确保只有具备权限的用户可以访问相应的页面。")]),a._v(" "),t("li",[a._v("用户管理：使用用户管理功能添加用户，并可设置他们的角色、项目等信息。当前版本暂不支持自助注册功能。")]),a._v(" "),t("li",[a._v("角色管理：为不同角色分配对应的访问权限。")]),a._v(" "),t("li",[a._v("项目管理：创建项目后会自动初始化函数驱动。可根据实际需求对项目进行划分，各项目数据相互独立，无法查看或引用其他项目的数据。")])]),a._v(" "),t("h2",{attrs:{id:"自动化功能"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#自动化功能"}},[a._v("#")]),a._v(" 自动化功能")]),a._v(" "),t("ol",[t("li",[a._v("环境变量：用于不同环境中相同变量的设置，所有模块都必须有环境变量，自行设置"),t("code",[a._v("开发环境")]),a._v("、"),t("code",[a._v("测试环境")]),a._v("、"),t("code",[a._v("预发布环境")]),a._v("等多个环境，相对固定的变量进行设置，如：域名、账号等")]),a._v(" "),t("li",[a._v("配置管理：公共数据配置，可以配置域名、请求头、变量和前置套件等")]),a._v(" "),t("li",[a._v("树形菜单：接口管理、测试套件、测试用例都包含了树形菜单，可以根据树形菜单对接口按功能模块、服务等进行划分，方便用例管理")]),a._v(" "),t("li",[a._v("接口管理：接口测试最基础模块，测试用例、测试套件、定时任务等都依赖与接口管理")]),a._v(" "),t("li",[a._v("测试套件：数据从接口管理的数据复制过来，数据相互独立，互不影响；运行配置只在调试时生效，测试用例、定时任务执行时无效")]),a._v(" "),t("li",[a._v("测试用例：引用测试套件，执行时以测试用例的配置为主；测试套件的修改，会导致测试用例运行报错、无法运行等")]),a._v(" "),t("li",[a._v("定时任务：引用多个定时任务，执行时各用例项目独立，没有依赖")]),a._v(" "),t("li",[a._v("性能任务：引用测试套件，增加性能测试相关特性（如：事务、集合点等）")]),a._v(" "),t("li",[a._v("测试报告：展示除压测任务的报告外的所有接口调试、运行报告")]),a._v(" "),t("li",[a._v("性能测试报告：展示性能测试报告")]),a._v(" "),t("li",[a._v("函数驱动：目前只能用python作为函数驱动，使用docker时会初始化python环境，安装所有依赖")]),a._v(" "),t("li",[a._v("py库管理：用于管理python第三方库，安装后会通过到master、work、run的所有节点，部署后会默认安装已有环境")])]),a._v(" "),t("h1",{attrs:{id:"页面预览"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#页面预览"}},[a._v("#")]),a._v(" 页面预览")]),a._v(" "),t("h3",{attrs:{id:"环境变量管理"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#环境变量管理"}},[a._v("#")]),a._v(" 环境变量管理")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230706103232835.png",alt:"环境变量"}})]),a._v(" "),t("h3",{attrs:{id:"配置管理"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#配置管理"}},[a._v("#")]),a._v(" 配置管理")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230706103301861.png",alt:"配置管理"}})]),a._v(" "),t("h3",{attrs:{id:"接口管理"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#接口管理"}},[a._v("#")]),a._v(" 接口管理")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230706103324932.png",alt:"接口管理"}})]),a._v(" "),t("h3",{attrs:{id:"测试步骤"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#测试步骤"}},[a._v("#")]),a._v(" 测试步骤")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710214256760.png",alt:"测试步骤"}})]),a._v(" "),t("h3",{attrs:{id:"测试用例"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#测试用例"}},[a._v("#")]),a._v(" 测试用例")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710214315331.png",alt:"测试用例"}})]),a._v(" "),t("h3",{attrs:{id:"定时任务"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#定时任务"}},[a._v("#")]),a._v(" 定时任务")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710214332669.png",alt:"定时任务"}})]),a._v(" "),t("h3",{attrs:{id:"定时任务-标签管理"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#定时任务-标签管理"}},[a._v("#")]),a._v(" 定时任务-标签管理")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710214416134.png",alt:"定时任务-标签管理"}})]),a._v(" "),t("h3",{attrs:{id:"测试报告列表"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#测试报告列表"}},[a._v("#")]),a._v(" 测试报告列表")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710215435068.png",alt:"测试报告"}})]),a._v(" "),t("h3",{attrs:{id:"测试报告详情"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#测试报告详情"}},[a._v("#")]),a._v(" 测试报告详情")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710215510007.png",alt:"测试报告详情"}})]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230710215523125.png",alt:"image-20230710215523125"}})]),a._v(" "),t("h3",{attrs:{id:"python-第三方库管理"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#python-第三方库管理"}},[a._v("#")]),a._v(" python 第三方库管理")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230711193741088.png",alt:"python包管理"}})]),a._v(" "),t("h3",{attrs:{id:"性能任务"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#性能任务"}},[a._v("#")]),a._v(" 性能任务")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230711163010339.png",alt:"性能任务"}})]),a._v(" "),t("h3",{attrs:{id:"性能任务详情"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#性能任务详情"}},[a._v("#")]),a._v(" 性能任务详情")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230711163030903.png",alt:"性能任务详情"}})]),a._v(" "),t("h3",{attrs:{id:"性能测试报告详情"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#性能测试报告详情"}},[a._v("#")]),a._v(" 性能测试报告详情")]),a._v(" "),t("p",[t("img",{attrs:{src:"https://qiniu.yangfan.gd.cn/markdown/image-20230711163104716.png",alt:"性能测试报告详情"}})])])}),[],!1,null,null,null);t.default=n.exports}}]);