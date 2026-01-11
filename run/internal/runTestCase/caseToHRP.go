package runTestCase

// 此文件包含已废弃的转换函数
// 新的转换逻辑已移至 converter.go
//
// 重构说明:
// - 原有的 yangfanTestCaseToHrpCase 通过 JSON 序列化/反序列化转换数据
// - 新的实现直接构建 httprunner 类型，避免多次序列化
// - 使用 converter.go 中的函数:
//   - convertConfigToTConfig: 转换配置
//   - convertAutoStepToIStep: 转换单个 API 步骤
//   - convertAutoCaseStepToIStep: 转换步骤集合
//   - wrapStepsInVirtualTestCase: 虚拟用例包装
