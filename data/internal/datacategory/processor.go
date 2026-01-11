package datacategory

import (
	"fmt"
	"strconv"
	"time"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// DataProcessor 统一数据处理器
type DataProcessor struct {
	ctx *ExecuteContext
}

// NewDataProcessor 创建数据处理器
func NewDataProcessor(ctx *ExecuteContext) *DataProcessor {
	return &DataProcessor{ctx: ctx}
}

// UpdateDataStatus 批量更新数据状态（只更新 status 和 cleaned_at，不更新 value）
func (p *DataProcessor) UpdateDataStatus(dataList []datawarehouse.DataCategoryData) error {
	if len(dataList) == 0 {
		return nil
	}

	now := time.Now()

	// 提取所有 ID
	ids := make([]uint, len(dataList))
	for i, data := range dataList {
		ids[i] = data.ID
	}

	// 批量更新状态
	if err := global.GVA_DB.Model(&datawarehouse.DataCategoryData{}).
		Where("id IN ?", ids).
		Updates(map[string]interface{}{
			"status":     DataStatusCleaned,
			"cleaned_at": now,
		}).Error; err != nil {
		return fmt.Errorf("批量更新状态失败: %w", err)
	}

	global.GVA_LOG.Info("批量更新数据状态完成",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
		zap.Int("count", len(dataList)),
	)

	return nil
}

// DeleteDataBatch 批量删除数据
func (p *DataProcessor) DeleteDataBatch(dataList []datawarehouse.DataCategoryData) (int, error) {
	if len(dataList) == 0 {
		return 0, nil
	}

	ids := make([]uint, len(dataList))
	for i, data := range dataList {
		ids[i] = data.ID
	}

	if err := global.GVA_DB.Delete(&datawarehouse.DataCategoryData{}, "id IN ?", ids).Error; err != nil {
		return 0, fmt.Errorf("批量删除数据失败: %w", err)
	}

	global.GVA_LOG.Info("批量删除数据完成",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
		zap.Int("count", len(ids)),
	)

	return len(ids), nil
}

// InsertDataBatch 批量插入新创建的数据
func (p *DataProcessor) InsertDataBatch(newData []map[string]interface{}) (int, error) {
	if len(newData) == 0 {
		return 0, nil
	}

	var insertData []datawarehouse.DataCategoryData
	for _, item := range newData {
		insertData = append(insertData, datawarehouse.DataCategoryData{
			DataCategoryId: p.ctx.CategoryID,
			Type:           p.ctx.Type,
			EnvId:          p.ctx.EnvID,
			ProjectId:      p.ctx.ProjectID,
			Value:          datatypes.JSONMap(item),
			Status:         DataStatusAvailable,
		})
	}

	if err := global.GVA_DB.Create(&insertData).Error; err != nil {
		return 0, fmt.Errorf("批量插入数据失败: %w", err)
	}

	global.GVA_LOG.Info("批量插入数据完成",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
		zap.Int("count", len(insertData)),
	)

	return len(insertData), nil
}

// SyncAvailableCount 同步更新 AvailableCount
// 根据数据库中实际可用的数据数量更新 (Status=0)
func (p *DataProcessor) SyncAvailableCount() error {
	envIDStr := strconv.FormatUint(uint64(p.ctx.EnvID), 10)

	// 统计数据库中 Status=0 (Available) 的可用数据数量
	var count int64
	if err := global.GVA_DB.Model(&datawarehouse.DataCategoryData{}).
		Where("data_category_id = ? AND env_id = ? AND status = ?",
			p.ctx.CategoryID, p.ctx.EnvID, DataStatusAvailable).
		Count(&count).Error; err != nil {
		return fmt.Errorf("统计可用数据失败: %w", err)
	}

	// 获取当前的 DataCategoryManagement
	var category datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.First(&category, p.ctx.CategoryID).Error; err != nil {
		return fmt.Errorf("获取数据分类失败: %w", err)
	}

	// 更新 AvailableCount
	if category.AvailableCount == nil {
		category.AvailableCount = make(datatypes.JSONMap)
	}
	category.AvailableCount[envIDStr] = int(count)

	if err := global.GVA_DB.Model(&datawarehouse.DataCategoryManagement{}).
		Where("id = ?", p.ctx.CategoryID).
		Update("available_count", category.AvailableCount).Error; err != nil {
		return fmt.Errorf("更新 AvailableCount 失败: %w", err)
	}

	global.GVA_LOG.Info("同步 AvailableCount 完成",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
		zap.Int64("availableCount", count),
	)

	return nil
}

// GetUsedData 获取需要清洗的数据 (Status=1 已占用)
func (p *DataProcessor) GetUsedData() ([]datawarehouse.DataCategoryData, error) {
	var dataList []datawarehouse.DataCategoryData
	if err := global.GVA_DB.Where(
		"data_category_id = ? AND env_id = ? AND status = ?",
		p.ctx.CategoryID, p.ctx.EnvID, DataStatusUsed,
	).Find(&dataList).Error; err != nil {
		return nil, fmt.Errorf("查询已占用数据失败: %w", err)
	}
	return dataList, nil
}

// CalculateNeedCreateCount 计算需要创建的数据数量
// 返回 needCreate, totalCount, availableCount, error
// availableCount 通过实时查询数据库获取，而不是使用缓存值
func (p *DataProcessor) CalculateNeedCreateCount() (int, int, int, error) {
	envIDStr := strconv.FormatUint(uint64(p.ctx.EnvID), 10)

	// 获取 DataCategoryManagement
	var category datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.First(&category, p.ctx.CategoryID).Error; err != nil {
		return 0, 0, 0, fmt.Errorf("获取数据分类失败: %w", err)
	}

	// 获取 Count 中对应环境的总数量
	totalCount := 0
	var rawVal interface{}
	var found bool

	if val, ok := category.Count[envIDStr]; ok {
		totalCount = toInt(val)
		rawVal = val
		found = true
	} else if p.ctx.EnvName != "" {
		if val, ok2 := category.Count[p.ctx.EnvName]; ok2 {
			totalCount = toInt(val)
			rawVal = val
			found = true
		}
	}

	global.GVA_LOG.Info("计算NeedCreate调试信息",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
		zap.String("envName", p.ctx.EnvName),
		zap.String("envKey", envIDStr),
		zap.Bool("foundInCount", found),
		zap.Any("rawVal", rawVal),
		zap.String("rawValType", fmt.Sprintf("%T", rawVal)),
		zap.Int("totalCount", totalCount),
	)

	// 实时查询数据库中 Status=0 (Available) 的可用数据数量
	var availableCount int64
	if err := global.GVA_DB.Model(&datawarehouse.DataCategoryData{}).
		Where("data_category_id = ? AND env_id = ? AND status = ?",
			p.ctx.CategoryID, p.ctx.EnvID, DataStatusAvailable).
		Count(&availableCount).Error; err != nil {
		return 0, 0, 0, fmt.Errorf("统计可用数据失败: %w", err)
	}

	// 计算需要创建的数量
	needCreate := totalCount - int(availableCount)

	return needCreate, totalCount, int(availableCount), nil
}

// PrepareOldData 将 DataCategoryData 列表转换为 map 切片（用于传递给 Python）
func (p *DataProcessor) PrepareOldData(dataList []datawarehouse.DataCategoryData) []map[string]interface{} {
	var oldData []map[string]interface{}
	for _, data := range dataList {
		oldData = append(oldData, map[string]interface{}(data.Value))
	}
	return oldData
}

// UpdateLastData 更新 LastData
func (p *DataProcessor) UpdateLastData(lastData map[string]interface{}) error {
	if lastData == nil {
		return nil
	}

	envIDStr := strconv.FormatUint(uint64(p.ctx.EnvID), 10)

	// 获取当前的 DataCategoryManagement
	var category datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.First(&category, p.ctx.CategoryID).Error; err != nil {
		return fmt.Errorf("获取数据分类失败: %w", err)
	}

	// 更新 LastData
	if category.LastData == nil {
		category.LastData = make(datatypes.JSONMap)
	}
	category.LastData[envIDStr] = lastData

	if err := global.GVA_DB.Model(&datawarehouse.DataCategoryManagement{}).
		Where("id = ?", p.ctx.CategoryID).
		Update("last_data", category.LastData).Error; err != nil {
		return fmt.Errorf("更新 LastData 失败: %w", err)
	}

	global.GVA_LOG.Info("更新 LastData 完成",
		zap.Uint("categoryId", p.ctx.CategoryID),
		zap.Uint("envId", p.ctx.EnvID),
	)

	return nil
}
