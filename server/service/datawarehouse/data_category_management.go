package datawarehouse

import (
	"context"
	"errors"
	"fmt"

	"github.com/test-instructor/yangfan/server/v2/global"
	"github.com/test-instructor/yangfan/server/v2/model/datawarehouse"
	datawarehouseReq "github.com/test-instructor/yangfan/server/v2/model/datawarehouse/request"
	"github.com/test-instructor/yangfan/server/v2/model/platform"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DataCategoryManagementService struct{}

const PythonCodeTypeDCM = 2 // 数据分类代码类型

// generateDCMUniqueKey 生成数据分类的 UniqueKey
func generateDCMUniqueKey(dcmId uint, envId int64) string {
	return fmt.Sprintf("dcm_%d_%d", dcmId, envId)
}

// CreateDataCategoryManagement 创建数据分类记录
func (dcmService *DataCategoryManagementService) CreateDataCategoryManagement(ctx context.Context, req datawarehouseReq.DataCategoryManagementSave) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 创建数据分类记录
		dcm := datawarehouse.DataCategoryManagement{
			Name:              req.Name,
			Type:              req.Type,
			Count:             map[string]interface{}{},
			MaxCreatePerRun:   req.MaxCreatePerRun,
			CreateCallType:    req.CreateCallType,
			CreateTestStepId:  req.CreateTestStepId,
			CleanCallType:     req.CleanCallType,
			CleanTestStepId:   req.CleanTestStepId,
			DirectDelete:      req.DirectDelete,
			ProjectId:         req.ProjectId,
			CreateRunConfigId: req.CreateRunConfigId,
			CleanRunConfigId:  req.CleanRunConfigId,
		}

		// 设置各环境 Count（允许为空时默认0，不在此处自动填充）
		if len(req.Count) > 0 {
			for k, v := range req.Count {
				dcm.Count[fmt.Sprintf("%d", k)] = v
			}
		}

		if err := tx.Create(&dcm).Error; err != nil {
			return err
		}

		// 为每个环境创建 PythonCode 记录
		for envIdStr, info := range req.PythonCodes {
			if info.Code == "" {
				continue
			}
			var envId int64
			fmt.Sscanf(envIdStr, "%d", &envId)

			pc := platform.PythonCode{
				Type:      PythonCodeTypeDCM,
				UniqueKey: generateDCMUniqueKey(dcm.ID, envId),
				ProjectId: req.ProjectId,
				Code:      info.Code,
			}
			if err := tx.Create(&pc).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// DeleteDataCategoryManagement 删除数据分类记录
func (dcmService *DataCategoryManagementService) DeleteDataCategoryManagement(ctx context.Context, ID string, projectId int64) error {
	var dcm datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.Where("id = ?", ID).First(&dcm).Error; err != nil {
		return err
	}
	if dcm.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 删除关联的 PythonCode
		uniqueKeyPrefix := fmt.Sprintf("dcm_%s_%%", ID)
		tx.Where("unique_key LIKE ?", uniqueKeyPrefix).Delete(&platform.PythonCode{})

		// 删除数据分类
		return tx.Delete(&datawarehouse.DataCategoryManagement{}, "id = ?", ID).Error
	})
}

// DeleteDataCategoryManagementByIds 批量删除数据分类记录
func (dcmService *DataCategoryManagementService) DeleteDataCategoryManagementByIds(ctx context.Context, IDs []string) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range IDs {
			uniqueKeyPrefix := fmt.Sprintf("dcm_%s_%%", id)
			tx.Where("unique_key LIKE ?", uniqueKeyPrefix).Delete(&platform.PythonCode{})
		}
		return tx.Delete(&[]datawarehouse.DataCategoryManagement{}, "id in ?", IDs).Error
	})
}

// UpdateDataCategoryManagement 更新数据分类记录
func (dcmService *DataCategoryManagementService) UpdateDataCategoryManagement(ctx context.Context, req datawarehouseReq.DataCategoryManagementSave, projectId int64, userId uint) error {
	var oldDCM datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.Where("id = ?", req.ID).First(&oldDCM).Error; err != nil {
		return err
	}
	if oldDCM.ProjectId != projectId {
		return errors.New("没有该项目的操作权限")
	}

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 更新数据分类基本信息
		updates := datawarehouse.DataCategoryManagement{
			Name:              req.Name,
			Type:              req.Type,
			MaxCreatePerRun:   req.MaxCreatePerRun,
			CreateCallType:    req.CreateCallType,
			CreateTestStepId:  req.CreateTestStepId,
			CleanCallType:     req.CleanCallType,
			CleanTestStepId:   req.CleanTestStepId,
			DirectDelete:      req.DirectDelete,
			CreateRunConfigId: req.CreateRunConfigId,
			CleanRunConfigId:  req.CleanRunConfigId,
		}

		// 单独处理 Count（各环境总数量）
		if req.Count != nil {
			countMap := datatypes.JSONMap{}
			for k, v := range req.Count {
				countMap[fmt.Sprintf("%d", k)] = v
			}
			if err := tx.Model(&datawarehouse.DataCategoryManagement{}).Where("id = ?", req.ID).Update("count", countMap).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&datawarehouse.DataCategoryManagement{}).Where("id = ?", req.ID).Updates(&updates).Error; err != nil {
			return err
		}

		// 为每个环境创建新的 PythonCode 记录（保留历史）
		for envIdStr, info := range req.PythonCodes {
			if info.Code == "" {
				continue
			}
			var envId int64
			fmt.Sscanf(envIdStr, "%d", &envId)

			uniqueKey := generateDCMUniqueKey(req.ID, envId)

			// 检查代码是否变化，无变化则跳过
			var latestPC platform.PythonCode
			err := tx.Where("unique_key = ?", uniqueKey).Order("created_at DESC").First(&latestPC).Error
			if err == nil && latestPC.Code == info.Code {
				continue
			}

			// 创建新记录（全量保留历史）
			pc := platform.PythonCode{
				Type:      PythonCodeTypeDCM,
				UniqueKey: uniqueKey,
				ProjectId: projectId,
				Code:      info.Code,
				UpdateBy:  userId,
			}
			if err := tx.Create(&pc).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetDataCategoryManagement 根据ID获取数据分类记录（含代码内容）
func (dcmService *DataCategoryManagementService) GetDataCategoryManagement(ctx context.Context, ID string, projectId int64) (dcm datawarehouse.DataCategoryManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dcm).Error
	if err != nil {
		return
	}

	// 获取项目下所有环境
	global.GVA_DB.Where("project_id = ?", projectId).Find(&dcm.EnvList)

	// 查询每个环境的最新代码及更新时间
	dcm.PythonCodes = make(map[string]datawarehouse.PythonCodeInfo)
	for _, env := range dcm.EnvList {
		uniqueKey := generateDCMUniqueKey(dcm.ID, int64(env.ID))
		var pc platform.PythonCode
		if err := global.GVA_DB.Where("unique_key = ?", uniqueKey).Order("created_at DESC").First(&pc).Error; err == nil {
			pcCopy := pc // 避免引用同一个 UpdatedAt 指针
			dcm.PythonCodes[fmt.Sprintf("%d", env.ID)] = datawarehouse.PythonCodeInfo{
				Code:     pcCopy.Code,
				UpdateAt: &pcCopy.UpdatedAt,
			}
		}
	}
	return
}

// GetDataCategoryManagementInfoList 分页获取数据分类记录
func (dcmService *DataCategoryManagementService) GetDataCategoryManagementInfoList(ctx context.Context, info datawarehouseReq.DataCategoryManagementSearch) (list []datawarehouse.DataCategoryManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&datawarehouse.DataCategoryManagement{})

	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db.Order("id DESC")
	db.Where("project_id = ?", info.ProjectId)

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&list).Error
	return
}

// GetDataCategoryManagementPublic 不需要鉴权的接口
func (dcmService *DataCategoryManagementService) GetDataCategoryManagementPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetPythonCodeForExecution 根据数据分类ID和环境ID获取执行代码
func (dcmService *DataCategoryManagementService) GetPythonCodeForExecution(dcmId uint, envId int64) (string, error) {
	uniqueKey := generateDCMUniqueKey(dcmId, envId)
	var pc platform.PythonCode
	err := global.GVA_DB.Where("unique_key = ?", uniqueKey).Order("created_at DESC").First(&pc).Error
	if err != nil {
		return "", err
	}
	return pc.Code, nil
}

// CreateDataCategoryData 创建数据池记录（自动继承 DataCategoryManagement 的 Type）
func (dcmService *DataCategoryManagementService) CreateDataCategoryData(ctx context.Context, data *datawarehouse.DataCategoryData) error {
	// 获取对应的 DataCategoryManagement 以获取 Type
	var dcm datawarehouse.DataCategoryManagement
	if err := global.GVA_DB.Where("id = ?", data.DataCategoryId).First(&dcm).Error; err != nil {
		return fmt.Errorf("获取数据分类失败: %w", err)
	}

	// 将 DataCategoryManagement 的 Type 传递给 DataCategoryData
	data.Type = dcm.Type
	data.ProjectId = dcm.ProjectId

	return global.GVA_DB.Create(data).Error
}

// GetDataCategoryTypeList 获取数据分类类型列表及最新数据值
func (dcmService *DataCategoryManagementService) GetDataCategoryTypeList(ctx context.Context, projectId int64) ([]map[string]interface{}, error) {
	var dcms []datawarehouse.DataCategoryManagement
	// 获取该项目下的所有数据分类
	if err := global.GVA_DB.Where("project_id = ?", projectId).Find(&dcms).Error; err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0, len(dcms))

	// 遍历获取每个分类的最新数据
	for _, dcm := range dcms {
		if dcm.Type == nil {
			continue
		}
		item := map[string]interface{}{
			"type":  *dcm.Type,
			"name":  *dcm.Name,
			"value": nil,
		}

		// 获取该分类下最新的一条数据
		var latestData datawarehouse.DataCategoryData
		// 使用 data_category_id 查询更准确
		err := global.GVA_DB.Where("data_category_id = ?", dcm.ID).Order("id desc").First(&latestData).Error
		if err == nil {
			item["value"] = latestData.Value
		}

		result = append(result, item)
	}
	return result, nil
}
