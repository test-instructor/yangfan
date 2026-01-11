package python

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/lingcetech/funplugin"
	"github.com/test-instructor/yangfan/httprunner/hrp"
	"github.com/test-instructor/yangfan/httprunner/python"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"go.uber.org/zap"
)

// Runner Python代码执行器
type Runner struct {
	ProjectID  int64
	EnvID      uint
	CategoryID uint // 数据分类 ID，用于生成 unique_key

	// 生命周期管理
	initialized bool
	tempDir     string
	plugin      funplugin.IPlugin
	parser      *hrp.Parser
	envVars     map[string]string // 缓存环境变量
}

// NewRunner 创建新的Python执行器
// categoryID: 数据分类 ID (DataCategoryManagement.ID)
func NewRunner(projectID int64, envID uint, categoryID uint) *Runner {
	return &Runner{
		ProjectID:  projectID,
		EnvID:      envID,
		CategoryID: categoryID,
	}
}

// Init 初始化 Runner，创建临时目录、加载 Python 代码和环境变量
func (r *Runner) Init() error {
	if r.initialized {
		return nil
	}

	// 获取环境变量
	envVars, err := r.getEnvVariables()
	if err != nil {
		return fmt.Errorf("获取环境变量失败: %w", err)
	}
	r.envVars = envVars

	// 获取 Python 代码
	pc, err := r.getPythonCode()
	if err != nil {
		return fmt.Errorf("获取Python代码失败: %w", err)
	}

	// 创建临时目录
	timestamp := time.Now().Format("20060102150405")
	secondDir := fmt.Sprintf("datacategory_%d_%d_%s", pc.ID, r.EnvID, timestamp)
	relativePath := filepath.Join("./debugcode", secondDir)
	path, err := filepath.Abs(relativePath)
	if err != nil {
		return fmt.Errorf("获取绝对路径失败: %w", err)
	}
	r.tempDir = path

	// 创建目录
	if err = os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 写入 Python 文件
	filePath := filepath.Join(path, "debugtalk.py")
	if err = os.WriteFile(filePath, []byte(pc.Code), 0644); err != nil {
		os.RemoveAll(path)
		return fmt.Errorf("写入文件失败: %w", err)
	}

	// 初始化插件
	plugin, err := python.InitPlugin(path, global.PythonVENV, true)
	if err != nil {
		os.RemoveAll(path)
		return fmt.Errorf("初始化Python插件失败: %w", err)
	}
	r.plugin = plugin

	// 创建 Parser
	r.parser = hrp.NewParser()
	r.parser.Plugin = plugin

	r.initialized = true
	global.GVA_LOG.Info("Python Runner 初始化成功",
		zap.Int64("projectId", r.ProjectID),
		zap.Uint("envId", r.EnvID),
	)

	return nil
}

// Close 关闭 Runner，清理资源
func (r *Runner) Close() {
	if !r.initialized {
		return
	}

	if r.plugin != nil {
		r.plugin.Quit()
		r.plugin = nil
	}

	if r.tempDir != "" {
		if err := os.RemoveAll(r.tempDir); err != nil {
			global.GVA_LOG.Warn("清理临时目录失败", zap.String("path", r.tempDir), zap.Error(err))
		}
		r.tempDir = ""
	}

	r.parser = nil
	r.envVars = nil
	r.initialized = false

	global.GVA_LOG.Info("Python Runner 已关闭",
		zap.Int64("projectId", r.ProjectID),
		zap.Uint("envId", r.EnvID),
	)
}

// IsInitialized 检查 Runner 是否已初始化
func (r *Runner) IsInitialized() bool {
	return r.initialized
}

// GetEnvVars 获取缓存的环境变量
func (r *Runner) GetEnvVars() map[string]string {
	return r.envVars
}

// getPythonCode 从数据库获取对应环境的Python代码
// Type=2 为数据分类代码
// UniqueKey 格式为: dcm_{categoryId}_{envId}
func (r *Runner) getPythonCode() (*platform.PythonCode, error) {
	var pc platform.PythonCode
	// 生成 unique_key，格式为 dcm_{categoryId}_{envId}
	uniqueKey := fmt.Sprintf("dcm_%d_%d", r.CategoryID, r.EnvID)

	db := global.GVA_DB.Model(&platform.PythonCode{})
	db.Select("id, created_at, updated_at, deleted_at, type, project_id, unique_key, update_by, code")
	db.Where("unique_key = ?", uniqueKey)
	db.Order("created_at DESC") // 获取最新的代码记录

	if err := db.First(&pc).Error; err != nil {
		return nil, fmt.Errorf("获取数据分类(%d)环境(%d)的Python代码失败: %w", r.CategoryID, r.EnvID, err)
	}
	return &pc, nil
}

// getEnvVariables 从 EnvDetail 获取当前环境的所有变量
// 返回 map[string]string，key 是变量名(EnvDetail.Key)，value 是对应 env_id 的值
func (r *Runner) getEnvVariables() (map[string]string, error) {
	envVars := make(map[string]string)

	// 查询项目下所有的环境变量
	var envDetails []platform.EnvDetail
	if err := global.GVA_DB.Where("project_id = ?", r.ProjectID).Find(&envDetails).Error; err != nil {
		return nil, fmt.Errorf("获取环境变量失败: %w", err)
	}

	// 将 env_id 转换为字符串，用于从 JSONMap 中提取值
	envIDStr := strconv.FormatUint(uint64(r.EnvID), 10)

	for _, detail := range envDetails {
		// JSONMap 可以直接作为 map[string]interface{} 使用
		if val, ok := detail.Value[envIDStr]; ok {
			// 将值转换为字符串
			switch v := val.(type) {
			case string:
				envVars[detail.Key] = v
			case float64:
				envVars[detail.Key] = strconv.FormatFloat(v, 'f', -1, 64)
			case bool:
				envVars[detail.Key] = strconv.FormatBool(v)
			case nil:
				envVars[detail.Key] = ""
			default:
				envVars[detail.Key] = fmt.Sprintf("%v", v)
			}
		} else {
			// 值不存在时设置为空串
			envVars[detail.Key] = ""
		}
	}

	return envVars, nil
}

// Execute 执行Python函数（使用已初始化的 Plugin）
// functionExpr: 函数调用表达式，如 "${create_data($ENV, $number)}"
// parameters: 变量映射，如 {"ENV": envVars, "number": 10}
func (r *Runner) Execute(functionExpr string, parameters map[string]interface{}) (interface{}, error) {
	if !r.initialized {
		return nil, fmt.Errorf("Runner 未初始化，请先调用 Init()")
	}

	// 调用 Parse，传入函数表达式和变量映射
	result, err := r.parser.Parse(functionExpr, parameters)
	if err != nil {
		return nil, fmt.Errorf("执行Python函数失败: %w", err)
	}

	global.GVA_LOG.Info("Python函数执行成功",
		zap.String("functionExpr", functionExpr),
		zap.Int64("projectId", r.ProjectID),
		zap.Uint("envId", r.EnvID),
	)

	return result, nil
}

// CreateData 调用create_data函数
// number: 创建数量
// lastData: 上一次生成的数据
// 返回: (新数据列表, 新的lastData, error)
func (r *Runner) CreateData(number int, lastData map[string]interface{}) ([]map[string]interface{}, map[string]interface{}, error) {
	// 使用缓存的环境变量
	envVars := r.envVars
	if envVars == nil {
		return nil, nil, fmt.Errorf("Runner 未初始化，环境变量为空")
	}

	// 兜底：如果 lastData 为 nil，初始化为空 map
	if lastData == nil {
		lastData = make(map[string]interface{})
	}

	// 构建函数表达式和变量映射
	functionExpr := "${create_data($ENV, $last_data, $number)}"
	parameters := map[string]interface{}{
		"ENV":       envVars,
		"last_data": lastData,
		"number":    number,
	}

	result, err := r.Execute(functionExpr, parameters)
	if err != nil {
		return nil, nil, err
	}

	// 解析返回结果 {"list": [], "last_data": {}}
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		// 兼容旧代码返回列表的情况（如果 Python 代码没更新）
		if list, err := convertToMapSlice(result); err == nil {
			return list, lastData, nil
		}
		return nil, nil, fmt.Errorf("Python函数返回类型错误，期望 dict, 实际: %T", result)
	}

	// 提取 list
	var listData []map[string]interface{}
	if listVal, ok := resultMap["list"]; ok {
		listData, err = convertToMapSlice(listVal)
		if err != nil {
			return nil, nil, fmt.Errorf("解析 list 失败: %w", err)
		}
	}

	// 提取 last_data
	var newLastData map[string]interface{}
	if lastVal, ok := resultMap["last_data"]; ok {
		if m, ok := lastVal.(map[string]interface{}); ok {
			newLastData = m
		}
	}
	// 如果 Python 没返回 last_data，沿用旧的（或者返回 nil）
	if newLastData == nil {
		newLastData = lastData
	}

	return listData, newLastData, nil
}

// UpdateData 调用update_data函数
// number: 数量
// dataOld: 旧数据
func (r *Runner) UpdateData(number int, dataOld []map[string]interface{}) ([]map[string]interface{}, error) {
	// 使用缓存的环境变量
	envVars := r.envVars
	if envVars == nil {
		return nil, fmt.Errorf("Runner 未初始化，环境变量为空")
	}

	// 构建函数表达式和变量映射
	functionExpr := "${update_data($ENV, $number, $data_old)}"
	parameters := map[string]interface{}{
		"ENV":      envVars,
		"number":   number,
		"data_old": dataOld,
	}

	result, err := r.Execute(functionExpr, parameters)
	if err != nil {
		return nil, err
	}

	return convertToMapSlice(result)
}

// convertToMapSlice 转换结果为 []map[string]interface{}
func convertToMapSlice(result interface{}) ([]map[string]interface{}, error) {
	if result == nil {
		return nil, nil
	}

	switch v := result.(type) {
	case []map[string]interface{}:
		return v, nil
	case []interface{}:
		var mapSlice []map[string]interface{}
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				mapSlice = append(mapSlice, m)
			}
		}
		return mapSlice, nil
	default:
		return nil, fmt.Errorf("无法转换结果类型: %T", result)
	}
}
