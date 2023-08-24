package yangfan

import (
	"time"

	system2 "github.com/test-instructor/yangfan/server/service/system"
	"go.uber.org/zap"

	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/system"
)

func RegisterApis() {
	entities := []system.SysApi{

		{ApiGroup: "base", Method: "POST", Path: "/base/login", Description: "用户登录(必选)"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "更改角色api权限"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "发送测试邮件"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "发送邮件示例"},

		{ApiGroup: "excel", Method: "GET", Path: "/excel/loadExcel", Description: "下载excel"},
		{ApiGroup: "excel", Method: "POST", Path: "/excel/exportExcel", Description: "导出excel"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/downloadTemplate", Description: "下载excel模板"},
		{ApiGroup: "excel", Method: "GET", Path: "/excel/downloadTemplate", Description: "下载excel模板"},
		{ApiGroup: "excel", Method: "POST", Path: "/excel/importExcel", Description: "导入excel"},

		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)"},

		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getColumn", Description: "获取所选table的所有字段"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/createPlug", Description: "自动创建插件包"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/installPlugin", Description: "安装插件"},
		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getDB", Description: "获取所有数据库"},
		{ApiGroup: "代码生成器", Method: "GET", Path: "/autoCode/getTables", Description: "获取数据库表"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/createTemp", Description: "自动化代码"},
		{ApiGroup: "代码生成器", Method: "POST", Path: "/autoCode/preview", Description: "预览自动化代码"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/rollback", Description: "回滚自动生成代码"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getSysHistory", Description: "查询回滚记录"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/delSysHistory", Description: "删除回滚记录"},
		{ApiGroup: "代码生成器历史", Method: "POST", Path: "/autoCode/getMeta", Description: "获取meta信息"},

		{ApiGroup: "分片上传", Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "寻找目标文件（秒传）"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "断点续传"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "断点续传完成"},
		{ApiGroup: "分片上传", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "上传完成移除文件"},

		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/delPackage", Description: "删除包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/createPackage", Description: "生成包(package)"},
		{ApiGroup: "包（pkg）生成器", Method: "POST", Path: "/autoCode/getPackage", Description: "获取所有包(package)"},

		{ApiGroup: "系统字典", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "删除字典"},
		{ApiGroup: "系统字典", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "更新字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "根据ID获取字典"},
		{ApiGroup: "系统字典", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "获取字典列表"},
		{ApiGroup: "系统字典", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "新增字典"},

		{ApiGroup: "系统字典详情", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "新增字典内容"},
		{ApiGroup: "系统字典详情", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "删除字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "根据ID获取字典内容"},
		{ApiGroup: "系统字典详情", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "获取字典内容列表"},
		{ApiGroup: "系统字典详情", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "更新字典内容"},

		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getSystemConfig", Description: "获取配置文件内容"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/setSystemConfig", Description: "设置配置文件内容"},
		{ApiGroup: "系统服务", Method: "POST", Path: "/system/getServerInfo", Description: "获取服务器信息"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setUserInfo", Description: "设置用户信息"},
		{ApiGroup: "系统用户", Method: "PUT", Path: "/user/setSelfInfo", Description: "设置自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "GET", Path: "/user/getUserInfo", Description: "获取自身信息(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserProjects", Description: "分配项目"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthorities", Description: "设置权限组"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/changePassword", Description: "修改密码（建议选择)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/setUserAuthority", Description: "修改用户角色(必选)"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/resetPassword", Description: "重置用户密码"},
		{ApiGroup: "系统用户", Method: "DELETE", Path: "/user/deleteUser", Description: "删除用户"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/admin_register", Description: "用户注册"},
		{ApiGroup: "系统用户", Method: "POST", Path: "/user/getUserList", Description: "获取用户列表"},

		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuAuthority", Description: "获取指定角色menu"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/addBaseMenu", Description: "新增菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenu", Description: "获取菜单树(必选)"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "删除菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/updateBaseMenu", Description: "更新菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuById", Description: "根据id获取菜单"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getMenuList", Description: "分页获取基础menu列表"},
		{ApiGroup: "菜单", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由"},

		{ApiGroup: "角色", Method: "PUT", Path: "/authority/updateAuthority", Description: "更新角色信息"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/getAuthorityList", Description: "获取角色列表"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/setDataAuthority", Description: "设置角色资源权限"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/copyAuthority", Description: "拷贝角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/createAuthority", Description: "创建角色"},
		{ApiGroup: "角色", Method: "POST", Path: "/authority/deleteAuthority", Description: "删除角色"},

		{ApiGroup: "客户", Method: "GET", Path: "/customer/customerList", Description: "获取客户列表"},
		{ApiGroup: "客户", Method: "PUT", Path: "/customer/customer", Description: "更新客户"},
		{ApiGroup: "客户", Method: "POST", Path: "/customer/customer", Description: "创建客户"},
		{ApiGroup: "客户", Method: "DELETE", Path: "/customer/customer", Description: "删除客户"},
		{ApiGroup: "客户", Method: "GET", Path: "/customer/customer", Description: "获取单一客户"},

		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "删除按钮"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "设置按钮权限"},
		{ApiGroup: "按钮权限", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "获取已有按钮权限"},

		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录"},
		{ApiGroup: "操作记录", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史"},
		{ApiGroup: "操作记录", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录"},
		{ApiGroup: "操作记录", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录"},

		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "文件名或者备注编辑"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "获取上传文件列表"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "文件上传示例"},
		{ApiGroup: "文件上传与下载", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "删除文件"},

		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "上传完成合并文件"},
		{ApiGroup: "断点续传(插件版)", Method: "POST", Path: "/simpleUploader/upload", Description: "插件版分片上传"},
		{ApiGroup: "断点续传(插件版)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "文件完整度验证"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "创建api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "删除Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "更新Api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "获取api列表"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "获取所有api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "获取api详细信息"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "批量删除api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/createApiMenu", Description: "新增接口菜单"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApiMenu", Description: "删除接口菜单"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApiMenuByIds", Description: "批量删除接口菜单"},
		{ApiGroup: "api", Method: "PUT", Path: "/api/updateApiMenu", Description: "更新接口菜单"},
		{ApiGroup: "api", Method: "GET", Path: "/api/findApiMenu", Description: "根据ID获取接口菜单"},
		{ApiGroup: "api", Method: "GET", Path: "/api/getApiMenuList", Description: "获取接口菜单列表"},

		{ApiGroup: "api模板", Method: "POST", Path: "/case/:project/createInterfaceTemplate", Description: "新增api 模版"},
		{ApiGroup: "api模板", Method: "DELETE", Path: "/case/:project/deleteInterfaceTemplate", Description: "删除api 模版"},
		{ApiGroup: "api模板", Method: "DELETE", Path: "/case/:project/deleteInterfaceTemplateByIds", Description: "批量删除api 模版"},
		{ApiGroup: "api模板", Method: "PUT", Path: "/case/:project/updateInterfaceTemplate", Description: "更新api 模版"},
		{ApiGroup: "api模板", Method: "GET", Path: "/case/:project/findInterfaceTemplate", Description: "根据ID获取api 模版"},
		{ApiGroup: "api模板", Method: "GET", Path: "/case/:project/getInterfaceTemplateList", Description: "获取api 模版列表"},
		{ApiGroup: "api模板", Method: "PUT", Path: "/case/:project/updateDebugTalk", Description: "更新DebugTalk文件"},
		{ApiGroup: "api模板", Method: "GET", Path: "/case/:project/getDebugTalk", Description: "获取DebugTalk文件(新)"},
		{ApiGroup: "api模板", Method: "POST", Path: "/case/:project/createDebugTalk", Description: "创建DebugTalk文件"},
		{ApiGroup: "api模板", Method: "POST", Path: "/case/:project/deleteDebugTalk", Description: "删除DebugTalk文件"},
		{ApiGroup: "api模板", Method: "POST", Path: "/case/:project/getDebugTalkList", Description: "获取DebugTalk列表"},
		{ApiGroup: "api模板", Method: "GET", Path: "/case/:project/getGrpc", Description: "获取grpc信息(新)"},
		{ApiGroup: "api模板", Method: "POST", Path: "/case/:project/createUserConfig", Description: "创建/更新用户配置"},
		{ApiGroup: "api模板", Method: "GET", Path: "/case/:project/getUserConfig", Description: "获取用户配置"},

		{ApiGroup: "py库管理", Method: "POST", Path: "/case/:project/pyPkg/installPyPkg", Description: "安装python第三方库"},
		{ApiGroup: "py库管理", Method: "POST", Path: "/case/:project/pyPkg/uninstallPyPkg", Description: "卸载ython第三方库"},
		{ApiGroup: "py库管理", Method: "POST", Path: "/case/:project/pyPkg/updatePyPkg", Description: "更新ython第三方库"},
		{ApiGroup: "py库管理", Method: "GET", Path: "/case/:project/pyPkg/getPkgVersionList", Description: "获取python第三方库版本信息(新)"},
		{ApiGroup: "py库管理", Method: "GET", Path: "/case/:project/pyPkg/pyPkgList", Description: "获取python第三方库列表"},

		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/addTaskTestCase", Description: "添加测试用例"},
		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/setTaskCase", Description: "定时任务设置测试用例"},
		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/createTimerTask", Description: "新增定时任务"},
		{ApiGroup: "定时任务", Method: "DELETE", Path: "/task/:project/deleteTimerTask", Description: "删除定时任务"},
		{ApiGroup: "定时任务", Method: "DELETE", Path: "/task/:project/deleteTimerTaskByIds", Description: "批量删除定时任务"},
		{ApiGroup: "定时任务", Method: "PUT", Path: "/task/:project/updateTimerTask", Description: "更新定时任务"},
		{ApiGroup: "定时任务", Method: "GET", Path: "/task/:project/findTimerTask", Description: "根据ID获取定时任务"},
		{ApiGroup: "定时任务", Method: "GET", Path: "/task/:project/getTimerTaskList", Description: "获取定时任务列表"},
		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/sortTaskCase", Description: "测试用例排序"},
		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/addTaskCase", Description: "任务添加测试用例"},
		{ApiGroup: "定时任务", Method: "DELETE", Path: "/task/:project/delTaskCase", Description: "定时任务删除测试用例"},
		{ApiGroup: "定时任务", Method: "GET", Path: "/task/:project/getTimerTaskTagList", Description: "获取TimerTaskTag列表"},
		{ApiGroup: "定时任务", Method: "POST", Path: "/task/:project/createTimerTaskTag", Description: "创建定时任务标签"},
		{ApiGroup: "定时任务", Method: "DELETE", Path: "/task/:project/deleteTimerTaskTag", Description: "删除定时任务标签"},
		{ApiGroup: "定时任务", Method: "GET", Path: "/task/:project/findTaskTestCase", Description: "定时任务用例"},

		{ApiGroup: "性能测试", Method: "PUT", Path: "/performance/:project/updatePerformance", Description: "更新性能测试任务"},
		{ApiGroup: "性能测试", Method: "POST", Path: "/performance/:project/addPerformanceCase", Description: "性能任务添加测试用例"},
		{ApiGroup: "性能测试", Method: "POST", Path: "/performance/:project/sortPerformanceCase", Description: "性能测试用例排序"},
		{ApiGroup: "性能测试", Method: "DELETE", Path: "/performance/:project/delPerformanceCase", Description: "性能测试删除测试用例"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/findPerformanceCase", Description: "性能测试通过id查找用例"},
		{ApiGroup: "性能测试", Method: "POST", Path: "/performance/:project/addOperation", Description: "添加事务、集合点"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/findPerformanceStep", Description: "查看性能测试步骤"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/getReportList", Description: "性能测试报告列表"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/findReport", Description: "性能测试报告详情"},
		{ApiGroup: "性能测试", Method: "DELETE", Path: "/performance/:project/deleteReport", Description: "删除性能测试报告"},
		{ApiGroup: "性能测试", Method: "POST", Path: "/performance/:project/createPerformance", Description: "创建性能测试任务"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/getPerformanceList", Description: "获取性能任务列表"},
		{ApiGroup: "性能测试", Method: "DELETE", Path: "/performance/:project/deletePerformance", Description: "删除性能任务"},
		{ApiGroup: "性能测试", Method: "GET", Path: "/performance/:project/findPerformance", Description: "通过id查找性能任务"},

		{ApiGroup: "接口分组树形菜单", Method: "GET", Path: "/case/:project/getApiMenuList", Description: "获取接口菜单列表"},
		{ApiGroup: "接口分组树形菜单", Method: "POST", Path: "/case/:project/createApiMenu", Description: "新增接口菜单"},
		{ApiGroup: "接口分组树形菜单", Method: "DELETE", Path: "/case/:project/deleteApiMenu", Description: "删除接口菜单"},
		{ApiGroup: "接口分组树形菜单", Method: "DELETE", Path: "/case/:project/deleteApiMenuByIds", Description: "批量删除接口菜单"},
		{ApiGroup: "接口分组树形菜单", Method: "PUT", Path: "/case/:project/updateApiMenu", Description: "更新接口菜单"},
		{ApiGroup: "接口分组树形菜单", Method: "GET", Path: "/case/:project/findApiMenu", Description: "根据ID获取接口菜单"},

		{ApiGroup: "测试步骤", Method: "GET", Path: "/case/:project/step/findTestCase", Description: "根据ID获取测试步骤"},
		{ApiGroup: "测试步骤", Method: "GET", Path: "/case/:project/step/getTestCaseList", Description: "获取测试步骤列表"},
		{ApiGroup: "测试步骤", Method: "POST", Path: "/case/:project/step/getTestCaseList", Description: "测试步骤列表排序"},
		{ApiGroup: "测试步骤", Method: "POST", Path: "/case/:project/step/addTestCase", Description: "测试步骤添加api"},
		{ApiGroup: "测试步骤", Method: "DELETE", Path: "/case/:project/step/delTestCase", Description: "删除测试步骤关联的api"},
		{ApiGroup: "测试步骤", Method: "POST", Path: "/case/:project/step/sortTestCase", Description: "测试步骤排序"},
		{ApiGroup: "测试步骤", Method: "POST", Path: "/case/:project/step/createTestCase", Description: "新增测试步骤"},
		{ApiGroup: "测试步骤", Method: "DELETE", Path: "/case/:project/step/deleteTestCase", Description: "删除测试步骤"},
		{ApiGroup: "测试步骤", Method: "DELETE", Path: "/case/:project/step/deleteTestCaseByIds", Description: "批量删除测试步骤"},
		{ApiGroup: "测试步骤", Method: "PUT", Path: "/case/:project/step/updateTestCase", Description: "更新测试步骤"},

		{ApiGroup: "测试报告", Method: "GET", Path: "/case/report/:project/getReportList", Description: "测试报告列表"},
		{ApiGroup: "测试报告", Method: "GET", Path: "/case/report/:project/getReportDetail", Description: "测试报告详情"},
		{ApiGroup: "测试报告", Method: "DELETE", Path: "/case/report/:project/delReport", Description: "删除测试报告"},
		{ApiGroup: "测试报告", Method: "GET", Path: "/case/report/:project/findReport", Description: "测试报告详情"},

		{ApiGroup: "测试用例", Method: "DELETE", Path: "/testcase/:project/delApisCase", Description: "定时任务删除测试步骤"},
		{ApiGroup: "测试用例", Method: "PUT", Path: "/testcase/:project/updateApiCase", Description: "更新测试用例"},
		{ApiGroup: "测试用例", Method: "GET", Path: "/testcase/:project/findApiCase", Description: "根据ID获取测试用例"},
		{ApiGroup: "测试用例", Method: "GET", Path: "/testcase/:project/getApiCaseList", Description: "获取测试用例列表"},
		{ApiGroup: "测试用例", Method: "POST", Path: "/testcase/:project/createApiCase", Description: "新增测试用例"},
		{ApiGroup: "测试用例", Method: "GET", Path: "/testcase/:project/findApiTestCase", Description: "测试用例用例"},
		{ApiGroup: "测试用例", Method: "POST", Path: "/testcase/:project/addApisCase", Description: "添加测试用例"},
		{ApiGroup: "测试用例", Method: "POST", Path: "/testcase/:project/setApisCase", Description: "测试用例设置测试步骤"},
		{ApiGroup: "测试用例", Method: "DELETE", Path: "/testcase/:project/deleteApiCase", Description: "删除测试用例"},
		{ApiGroup: "测试用例", Method: "DELETE", Path: "/testcase/:project/deleteApiCaseByIds", Description: "批量删除测试用例"},
		{ApiGroup: "测试用例", Method: "POST", Path: "/testcase/:project/sortApisCase", Description: "测试步骤排序"},
		{ApiGroup: "测试用例", Method: "POST", Path: "/testcase/:project/AddApiTestCase", Description: "任务添加测试步骤"},

		{ApiGroup: "测试配置", Method: "GET", Path: "/ac/:project/getApiConfigList", Description: "获取配置管理列表"},
		{ApiGroup: "测试配置", Method: "POST", Path: "/ac/:project/createApiConfig", Description: "新增配置管理"},
		{ApiGroup: "测试配置", Method: "DELETE", Path: "/ac/:project/deleteApiConfig", Description: "删除配置管理"},
		{ApiGroup: "测试配置", Method: "DELETE", Path: "/ac/:project/deleteApiConfigByIds", Description: "批量删除配置管理"},
		{ApiGroup: "测试配置", Method: "PUT", Path: "/ac/:project/updateApiConfig", Description: "更新配置管理"},
		{ApiGroup: "测试配置", Method: "GET", Path: "/ac/:project/findApiConfig", Description: "根据ID获取配置管理"},

		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runApi", Description: "调试接口"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runTimerTask", Description: "运行定时任务"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runApiCase", Description: "运行定时任务"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runBoomerDebug", Description: "性能测试调试运行"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runBoomer", Description: "运行性能测试"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/rebalance", Description: "调整性能测试参数"},
		{ApiGroup: "运行", Method: "GET", Path: "/case/run/:project/stop", Description: "停止性能测试"},
		{ApiGroup: "运行", Method: "POST", Path: "/case/run/:project/runTestCaseStep", Description: "运行测试用例"},

		{ApiGroup: "项目管理", Method: "DELETE", Path: "/project/deleteProject", Description: "删除项目管理"},
		{ApiGroup: "项目管理", Method: "DELETE", Path: "/project/deleteProjectByIds", Description: "批量删除项目管理"},
		{ApiGroup: "项目管理", Method: "PUT", Path: "/project/updateProject", Description: "更新项目管理"},
		{ApiGroup: "项目管理", Method: "GET", Path: "/project/findProject", Description: "根据ID获取项目管理"},
		{ApiGroup: "项目管理", Method: "GET", Path: "/project/getProjectList", Description: "获取项目管理列表"},
		{ApiGroup: "项目管理", Method: "POST", Path: "/project/createProject", Description: "新增项目管理"},
		{ApiGroup: "项目管理", Method: "POST", Path: "/project/setUserProjectAuth", Description: "设置用户项目权限"},
		{ApiGroup: "项目管理", Method: "DELETE", Path: "/project/deleteUserProjectAuth", Description: "删除项目成员"},
		{ApiGroup: "项目管理", Method: "GET", Path: "/project/getProjectUserList", Description: "获取项目用户列表"},
		{ApiGroup: "项目管理", Method: "POST", Path: "/project/setKey", Description: "设置项目密钥"},
		{ApiGroup: "项目管理", Method: "GET", Path: "/project/findKey", Description: "查询项目密钥"},

		{ApiGroup: "环境变量", Method: "POST", Path: "/env/:project/createEnv", Description: "新增环境"},
		{ApiGroup: "环境变量", Method: "PUT", Path: "/env/:project/updateEnv", Description: "修改环境"},
		{ApiGroup: "环境变量", Method: "DELETE", Path: "/env/:project/deleteEnv", Description: "删除环境"},
		{ApiGroup: "环境变量", Method: "GET", Path: "/env/:project/findEnv", Description: "通过id查找环境"},
		{ApiGroup: "环境变量", Method: "GET", Path: "/env/:project/getEnvList", Description: "查询环境列表"},

		{ApiGroup: "环境变量", Method: "POST", Path: "/env/:project/createEnvVariable", Description: "新增变量"},
		{ApiGroup: "环境变量", Method: "DELETE", Path: "/env/:project/deleteEnvVariable", Description: "删除变量"},
		{ApiGroup: "环境变量", Method: "GET", Path: "/env/:project/findEnvVariable", Description: "通过id查找变量"},
		{ApiGroup: "环境变量", Method: "GET", Path: "/env/:project/getEnvVariableList", Description: "查询变量列表"},

		{ApiGroup: "消息通知", Method: "POST", Path: "/message/:project/createMessage", Description: "新建Message"},
		{ApiGroup: "消息通知", Method: "DELETE", Path: "/message/:project/deleteMessage", Description: "删除Message"},
		{ApiGroup: "消息通知", Method: "PUT", Path: "/message/:project/updateMessage", Description: "更新Message"},
		{ApiGroup: "消息通知", Method: "GET", Path: "/message/:project/findMessage", Description: "根据ID获取Message"},
		{ApiGroup: "消息通知", Method: "GET", Path: "/message/:project/getMessageList", Description: "获取Message列表"},

		{ApiGroup: "CI", Method: "GET", Path: "/ci/runTag", Description: "运行tag"},
		{ApiGroup: "CI", Method: "GET", Path: "/ci/getReport", Description: "获取测试报告"},
	}

	var count int64

	for i := 0; i < len(entities); i++ {
		global.GVA_DB.Find(&[]system.SysApi{}, "path = ? and method = ?", entities[i].Path, entities[i].Method).Count(&count)
		if count > 0 {
			//global.GVA_LOG.Debug("插件已安装或存在同名路由" + entities[i].Path)
			continue
		}
		err := global.GVA_DB.Create(&entities[i]).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
	}
}

func RegisterMenus() {
	entities := []system.SysBaseMenu{
		{GVA_MODEL: global.GVA_MODEL{ID: 32, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "project", Name: "project", Component: "view/project/project.vue", Sort: 1, Meta: system.Meta{KeepAlive: true, Title: "项目管理", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 33, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "interfaces", Name: "interfaces", Component: "view/interface/index.vue", Sort: 4, Meta: system.Meta{KeepAlive: true, Title: "接口自动化", Icon: "box"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 34, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "interfacetemplate", Name: "interfacetemplate", Component: "view/interface/interfaceTemplate/interfaceTemplate.vue", Sort: 200, Meta: system.Meta{KeepAlive: true, Title: "接口管理", Icon: "coin"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 35, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "DebugReport", Name: "DebugReport", Component: "view/interface/interfaceReport/DebugReport.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "DebugReport", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 36, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "apiConfig", Name: "apiConfig", Component: "view/interface/interfaceTemplate/apiConfig.vue", Sort: 100, Meta: system.Meta{KeepAlive: true, Title: "配置管理", Icon: "expand"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 37, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "testCaseStep", Name: "testCaseStep", Component: "view/interface/testCaseStep/testCaseStep.vue", Sort: 300, Meta: system.Meta{KeepAlive: true, Title: "测试步骤", Icon: "suitcase"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 38, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "testCaseStepDetail/:id", Name: "testCaseStepDetail", Component: "view/interface/testCaseStep/testCaseStepDetail.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "步骤详情-${id}", Icon: "finished"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 39, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "apiTest", Name: "apiTest", Component: "view/interface/test.vue", Sort: 0, Meta: system.Meta{KeepAlive: true,Title: "测试", Icon: "aim"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 40, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "testCaseAdd", Name: "testCaseAdd", Component: "view/interface/testCase/testCaseAdd.vue", Sort: 0, Meta: system.Meta{KeepAlive: true,Title: "添加测试用例", Icon: "aim"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 41, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "apiTest2", Name: "apiTest2", Component: "view/interface/test2.vue", Sort: 0, Meta: system.Meta{KeepAlive: true,Title: "测试2", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 42, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "report", Name: "report", Component: "view/interface/Reports/report.vue", Sort: 600, Meta: system.Meta{KeepAlive: true, Title: "测试报告", Icon: "compass"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 43, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "reportDetail/:report_id", Name: "reportDetail", Component: "view/interface/Reports/reportDetail.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "测试报告详情-${report_id}", Icon: "document"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 44, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "testCase", Name: "testCase", Component: "view/interface/apiCase/apiCase.vue", Sort: 400, Meta: system.Meta{KeepAlive: true, Title: "测试用例", Icon: "briefcase"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 46, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "debugtalk", Name: "debugtalk", Component: "view/interface/debugtalk/debugtalk.vue", Sort: 700, Meta: system.Meta{KeepAlive: true, Title: "驱动函数", Icon: "reading"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 47, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "debugtalkGen", Name: "debugtalkGen", Component: "view/interface/debugtalk/debugtalkGen.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "debugtalkGen", Icon: "aim"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 48, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "taskAddCase", Name: "taskAddCase", Component: "view/interface/timerTask/taskAddCase.vue", Sort: 0, Meta: system.Meta{KeepAlive: true,Title: "任务添加测试用例", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 49, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "apisCaseDetail/:id", Name: "apisCaseDetail", Component: "view/interface/apiCase/apisCaseDetail.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "用例详情-${id}", Icon: "aim"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 50, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "envConfig", Name: "envConfig", Component: "view/interface/interfaceComponents/envConfig.vue", Sort: 1, Meta: system.Meta{KeepAlive: true,Title: "envConfig", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 51, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "taskCaseDetail/:id", Name: "taskCaseDetail", Component: "view/interface/timerTask/taskCaseDetail.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "任务详情-${id}", Icon: "aim"}},
		//{GVA_MODEL: global.GVA_MODEL{ID: 52, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "taskAddCase", Name: "taskAddCase", Component: "view/interface/timerTask/taskAddCase.vue", Sort: 0, Meta: system.Meta{KeepAlive: true,Title: "任务添加测试用例", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 53, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "timerTask", Name: "timerTask", Component: "view/interface/timerTask/timerTask.vue", Sort: 500, Meta: system.Meta{KeepAlive: true, Title: "定时任务", Icon: "timer"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 55, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "performance", Name: "performance", Component: "view/performance/index.vue", Sort: 5, Meta: system.Meta{KeepAlive: true, Title: "性能测试", Icon: "stopwatch"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 56, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "55", Path: "performanceTask", Name: "performanceTask", Component: "view/performance/task/index.vue", Sort: 1, Meta: system.Meta{KeepAlive: true, Title: "性能任务", Icon: "cpu"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 57, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "55", Path: "performanceDetail/:id", Name: "performanceDetail", Component: "view/performance/task/taskDetail.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "性能任务详情-${id}", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 58, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "55", Path: "pReport", Name: "pReport", Component: "view/performance/report.vue", Sort: 2, Meta: system.Meta{KeepAlive: true, Title: "性能测试报告", Icon: "compass"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 59, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "55", Path: "pReportDetail/:id", Name: "pReportDetail", Component: "view/performance/reportDetail.vue", Sort: 999, Meta: system.Meta{KeepAlive: true, Title: "性能测试报告详情-${id}", Icon: "document"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 60, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: true, ParentId: "33", Path: "jsonCompare", Name: "jsonCompare", Component: "view/interface/interfaceComponents/jsonCompare.vue", Sort: 99999, Meta: system.Meta{KeepAlive: true, Title: "json", Icon: "aim"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 61, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "env", Name: "env", Component: "view/interface/environment/environment.vue", Sort: 0, Meta: system.Meta{KeepAlive: true, Title: "环境变量", Icon: "grid"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 62, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "py_pkg", Name: "py_pkg", Component: "view/py_pkg/py_pkg.vue", Sort: 680, Meta: system.Meta{KeepAlive: true, Title: "py库管理", Icon: "office-building"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 80, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 0, Meta: system.Meta{KeepAlive: true, Title: "关于我们", Icon: "info-filled"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 81, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "project-auth", Name: "project-auth", Component: "view/project/authority.vue", Sort: 0, Meta: system.Meta{KeepAlive: true, Title: "项目权限", Icon: "avatar"}},
		{GVA_MODEL: global.GVA_MODEL{ID: 82, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}}, MenuLevel: 0, Hidden: false, ParentId: "33", Path: "message", Name: "message", Component: "view/interface/message/index.vue", Sort: 150, Meta: system.Meta{KeepAlive: true, Title: "消息通知", Icon: "message"}},
	}
	baseMenuService := system2.BaseMenuService{}
	for i := 0; i < len(entities); i++ {
		var entitie system.SysBaseMenu
		entitie = entities[i]
		err := global.GVA_DB.First(&entitie).Error
		if err == gorm.ErrRecordNotFound {
			err := system2.MenuServiceApp.AddBaseMenu(entities[i])
			if err != nil {
				global.GVA_LOG.Error("创建菜单失败"+entities[i].Name, zap.Error(err))
			}
			continue
		}
		if err != nil {
			global.GVA_LOG.Error("查找菜单出错：" + entities[i].Path)
			continue
		}
		err = baseMenuService.UpdateBaseMenu(entities[i])
		if err != nil {
			global.GVA_LOG.Error("更新菜单失败"+entities[i].Name, zap.Error(err))
		}
	}

}
