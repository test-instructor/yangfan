package datacategory

import (
	"fmt"
	"strconv"

	"github.com/test-instructor/yangfan/data/internal/python"
	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	"go.uber.org/zap"
)

// Executor 数据执行器
type Executor struct {
	ctx       *ExecuteContext
	runner    *python.Runner
	processor *DataProcessor
}

// NewExecutor 创建执行器
func NewExecutor(ctx *ExecuteContext) *Executor {
	return &Executor{
		ctx:       ctx,
		processor: NewDataProcessor(ctx),
	}
}

// SetRunner 设置 Python Runner（由外部管理生命周期）
func (e *Executor) SetRunner(runner *python.Runner) {
	e.runner = runner
}

// CleanData 执行清洗数据
func (e *Executor) CleanData() (*ProcessResult, error) {
	result := &ProcessResult{
		CategoryID:   e.ctx.CategoryID,
		CategoryName: e.ctx.CategoryName,
		EnvID:        e.ctx.EnvID,
		Action:       "clean",
		Success:      false,
	}

	if e.ctx.CleanCallType == nil {
		result.Message = "未配置清洗类型"
		return result, nil
	}

	cleanType := *e.ctx.CleanCallType

	// 使用 processor 获取需要清洗的数据
	dataList, err := e.processor.GetUsedData()
	if err != nil {
		result.Message = fmt.Sprintf("查询数据失败: %v", err)
		return result, err
	}

	if len(dataList) == 0 {
		result.Success = true
		result.Message = "没有需要清洗的数据"
		return result, nil
	}

	global.GVA_LOG.Info("开始清洗数据",
		zap.Uint("categoryId", e.ctx.CategoryID),
		zap.Uint("envId", e.ctx.EnvID),
		zap.Int64("cleanType", cleanType),
		zap.Int("dataCount", len(dataList)),
	)

	switch cleanType {
	case CallTypeTestStep:
		// 类型1: 测试步骤 - 保留入口，打印当前类型
		global.GVA_LOG.Info("清洗类型为测试步骤，当前仅打印日志",
			zap.Uint("categoryId", e.ctx.CategoryID),
			zap.String("categoryName", e.ctx.CategoryName),
			zap.Int("dataCount", len(dataList)),
		)
		result.Success = true
		result.Message = "测试步骤类型，已打印日志"
		result.DataCount = len(dataList)

	case CallTypePython:
		// 类型2: Python - 调用update_data
		count, err := e.cleanWithPython(dataList)
		if err != nil {
			result.Message = fmt.Sprintf("Python清洗失败: %v", err)
			return result, err
		}
		result.Success = true
		result.Message = "Python清洗完成"
		result.DataCount = count

	case CallTypeDelete:
		// 类型3: 直接删除
		count, err := e.processor.DeleteDataBatch(dataList)
		if err != nil {
			result.Message = fmt.Sprintf("删除数据失败: %v", err)
			return result, err
		}
		// 同步 AvailableCount
		if syncErr := e.processor.SyncAvailableCount(); syncErr != nil {
			global.GVA_LOG.Warn("删除后同步 AvailableCount 失败", zap.Error(syncErr))
		}
		result.Success = true
		result.Message = "数据已删除"
		result.DataCount = count

	default:
		result.Message = fmt.Sprintf("未知的清洗类型: %d", cleanType)
	}

	return result, nil
}

// cleanWithPython 使用Python脚本清洗数据
func (e *Executor) cleanWithPython(dataList []datawarehouse.DataCategoryData) (int, error) {
	if e.runner == nil {
		return 0, fmt.Errorf("Python Runner 未设置")
	}

	// 准备旧数据
	oldData := e.processor.PrepareOldData(dataList)

	// 调用Python update_data
	newData, err := e.runner.UpdateData(len(dataList), oldData)
	if err != nil {
		return 0, err
	}

	// 更新原数据的状态和清理时间（不更新value）
	if err := e.processor.UpdateDataStatus(dataList); err != nil {
		return 0, err
	}

	// 将Python返回的新数据入库
	if len(newData) > 0 {
		if _, err := e.processor.InsertDataBatch(newData); err != nil {
			return 0, fmt.Errorf("插入清洗后的新数据失败: %w", err)
		}
	}

	// 同步更新 AvailableCount
	if err := e.processor.SyncAvailableCount(); err != nil {
		global.GVA_LOG.Warn("清洗后同步 AvailableCount 失败", zap.Error(err))
	}

	return len(dataList), nil
}

// CreateData 执行创建数据
func (e *Executor) CreateData() (*ProcessResult, error) {
	result := &ProcessResult{
		CategoryID:   e.ctx.CategoryID,
		CategoryName: e.ctx.CategoryName,
		EnvID:        e.ctx.EnvID,
		Action:       "create",
		Success:      false,
	}

	if e.ctx.CreateCallType == nil {
		result.Message = "未配置创建类型"
		return result, nil
	}

	createType := *e.ctx.CreateCallType

	global.GVA_LOG.Info("开始创建数据",
		zap.Uint("categoryId", e.ctx.CategoryID),
		zap.Uint("envId", e.ctx.EnvID),
		zap.Int64("createType", createType),
	)

	switch createType {
	case CallTypeTestStep:
		// 类型1: 测试步骤 - 保留入口，打印当前类型
		global.GVA_LOG.Info("创建类型为测试步骤，当前仅打印日志",
			zap.Uint("categoryId", e.ctx.CategoryID),
			zap.String("categoryName", e.ctx.CategoryName),
		)
		result.Success = true
		result.Message = "测试步骤类型，已打印日志"

	case CallTypePython:
		// 类型2: Python - 调用create_data，内部处理数据入库和LastData更新
		count, err := e.createWithPython()
		if err != nil {
			result.Message = fmt.Sprintf("Python创建失败: %v", err)
			return result, err
		}
		result.Success = true
		result.Message = "Python创建完成"
		result.DataCount = count

	default:
		result.Message = fmt.Sprintf("未知的创建类型: %d", createType)
	}

	return result, nil
}

// createWithPython 使用Python脚本创建数据
func (e *Executor) createWithPython() (int, error) {
	if e.runner == nil {
		return 0, fmt.Errorf("Python Runner 未设置")
	}

	// 使用 processor 计算需要创建的数量
	needCreate, totalCount, availableCount, err := e.processor.CalculateNeedCreateCount()
	if err != nil {
		return 0, err
	}

	if needCreate <= 0 {
		global.GVA_LOG.Info("无需创建数据",
			zap.Uint("categoryId", e.ctx.CategoryID),
			zap.Uint("envId", e.ctx.EnvID),
			zap.Int("totalCount", totalCount),
			zap.Int("availableCount", availableCount),
		)
		return 0, nil
	}

	// 获取 DataCategoryManagement 以拿到 LastData
	var category datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.First(&category, e.ctx.CategoryID).Error; err != nil {
		return 0, fmt.Errorf("获取数据分类失败: %w", err)
	}
	envIDStr := strconv.FormatUint(uint64(e.ctx.EnvID), 10)

	// 获取当前环境的 LastData
	var currentLastData map[string]interface{}
	if val, ok := category.LastData[envIDStr]; ok {
		if m, ok := val.(map[string]interface{}); ok {
			currentLastData = m
		}
	}

	global.GVA_LOG.Info("计算需要创建的数据数量",
		zap.Uint("categoryId", e.ctx.CategoryID),
		zap.Uint("envId", e.ctx.EnvID),
		zap.Int("totalCount", totalCount),
		zap.Int("availableCount", availableCount),
		zap.Int("needCreate", needCreate),
	)

	// 调用Python create_data
	newData, newLastData, err := e.runner.CreateData(needCreate, currentLastData)
	if err != nil {
		return 0, err
	}

	if len(newData) == 0 {
		global.GVA_LOG.Info("Python返回空数据列表")
		return 0, nil
	}

	// 使用 processor 批量插入数据
	count, err := e.processor.InsertDataBatch(newData)
	if err != nil {
		return 0, err
	}

	// 更新 LastData
	if newLastData != nil {
		if err := e.processor.UpdateLastData(newLastData); err != nil {
			global.GVA_LOG.Warn("更新 LastData 失败", zap.Error(err))
		}
	}

	// 同步更新 AvailableCount
	if err := e.processor.SyncAvailableCount(); err != nil {
		global.GVA_LOG.Warn("创建后同步 AvailableCount 失败", zap.Error(err))
	}

	return count, nil
}

// toInt 将 interface{} 转换为 int
// 兼容 JSON 反序列化后可能出现的各种数字类型
func toInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		// 先按整数解析
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
		// 再按浮点解析
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return int(f)
		}
	default:
		// 兜底：使用 fmt 转成字符串再解析，避免直接返回 0
		s := fmt.Sprintf("%v", v)
		if i, err := strconv.Atoi(s); err == nil {
			return i
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return int(f)
		}
		global.GVA_LOG.Warn("toInt 转换失败，返回0",
			zap.Any("rawValue", val),
			zap.String("type", fmt.Sprintf("%T", val)),
		)
	}
	return 0
}
