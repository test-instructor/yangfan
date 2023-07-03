package interfacecase

import (
	"github.com/pkg/errors"
	"github.com/test-instructor/grpc-plugin/plugin"
	"gorm.io/gorm"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type ApiParams struct {
	Requests string `json:"request" form:"request"`
	Validate string `json:"validate" form:"validate"`
}

type InterfaceTemplateService struct {
}

// CreateInterfaceTemplate 创建InterfaceTemplate记录

func (i *InterfaceTemplateService) CreateInterfaceTemplate(apicase interfacecase.ApiStep) (id *uint, err error) {
	apicase.ValidateNumber = uint(len(apicase.Validate))
	err = global.GVA_DB.Create(&apicase).Error
	if err != nil {
		return nil, err
	}
	return &apicase.ID, nil
}

// DeleteInterfaceTemplate 删除InterfaceTemplate记录

func (i *InterfaceTemplateService) DeleteInterfaceTemplate(apicase interfacecase.ApiStep) (err error) {
	err = global.GVA_DB.Delete(&apicase).Error
	return err
}

// DeleteInterfaceTemplateByIds 批量删除InterfaceTemplate记录

func (i *InterfaceTemplateService) DeleteInterfaceTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]interfacecase.ApiStep{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInterfaceTemplate 更新InterfaceTemplate记录

func (i *InterfaceTemplateService) UpdateInterfaceTemplate(apicase interfacecase.ApiStep) (id uint, err error) {
	var oId interfacecase.Operator
	global.GVA_DB.Model(interfacecase.ApiStep{}).Where("id = ?", apicase.ID).First(&oId)
	apicase.CreatedBy = oId.CreatedBy
	apicase.ValidateNumber = uint(len(apicase.Validate))
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Where(&interfacecase.ApiStep{
			GVA_MODEL: global.GVA_MODEL{ID: apicase.ID},
		}).Save(&apicase).Error
		if err != nil {
			return err
		}
		if apicase.Request != nil {
			err = tx.Where(&interfacecase.ApiRequest{GVA_MODEL: global.GVA_MODEL{ID: apicase.Request.ID}}).
				Save(&apicase.Request).Error
		}
		if apicase.Grpc != nil {
			err = tx.Where(&interfacecase.ApiGrpc{GVA_MODEL: global.GVA_MODEL{ID: apicase.Grpc.ID}}).
				Save(&apicase.Grpc).Error
		}

		return err
	})

	return apicase.ID, err
}

// GetInterfaceTemplate 根据id获取InterfaceTemplate记录

func (i *InterfaceTemplateService) GetInterfaceTemplate(id uint) (err error, apiCase interfacecase.ApiStep) {
	db := global.GVA_DB.
		Preload("Request").
		Preload("Grpc").
		Preload("Validate").
		Model(&interfacecase.ApiStep{})
	db.Where("id = ?", id).Find(&apiCase)
	return
}

// GetInterfaceTemplateInfoList 分页获取InterfaceTemplate记录

func (i *InterfaceTemplateService) GetInterfaceTemplateInfoList(info interfacecaseReq.InterfaceTemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.
		Model(&interfacecase.ApiStep{}).
		Preload("Project").Joins("Project").
		Preload("ApiMenu").Joins("ApiMenu").
		Where("Project.ID = ?", info.ProjectID)
	if info.ApiMenuID > 0 {
		db.Preload("ApiMenu").Joins("ApiMenu").Where("ApiMenu.ID = ?", info.ApiMenuID)
	}
	var apicases []*interfacecase.ApiStep

	//查询对应的类型
	db.Where("api_type = ?", info.ApiType)

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("api_steps.name LIKE ?", "%"+info.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Joins("Request").Joins("Grpc").
		Limit(limit).Offset(offset).Find(&apicases).Error
	resetApiCaseStep(apicases...)
	return err, apicases, total
}

func (i *InterfaceTemplateService) UpdateDebugTalk(debugTalk interfacecase.ApiDebugTalk) (err error) {
	err = global.GVA_DB.Save(&debugTalk).Error
	return err
}

func (i *InterfaceTemplateService) GetDebugTalk(debugTalk interfacecase.ApiDebugTalk) (err error, debugTalkFirst interfacecase.ApiDebugTalk) {

	db := global.GVA_DB.
		Model(&interfacecase.ApiDebugTalk{}).
		Preload("Project").Joins("Project").Where("Project.ID = ?", debugTalk.Project.ID)
	//查询对应的类型
	db.Where("file_type = ?", debugTalk.FileType).Order("id desc")
	err = db.First(&debugTalkFirst).Error
	if err != nil {
		defaultDB := global.GVA_DB.Model(&interfacecase.ApiDebugTalk{}).
			Preload("Project").Joins("Project").Where("Project.ID = ?", 2)
		defaultDB.Where("file_type = ?", debugTalk.FileType)
		err = defaultDB.First(&debugTalkFirst).Error
	}
	return err, debugTalkFirst
}

func (i *InterfaceTemplateService) GetGrpc(gRPC interfacecaseReq.GrpcFunc) (err error, data interface{}) {
	var gData = make(map[string]interface{})
	var g plugin.Grpc
	if gRPC.Host == nil {
		err = errors.New("请输入对应的服务信息")
	}
	g.Host = *gRPC.Host
	ig := plugin.NewInvokeGrpc(&g)
	err = ig.GetResource()
	if err != nil {
		return err, gData
	}
	if gRPC.Ref != nil && *gRPC.Ref {
		err = ig.Reset()
		return
	}
	svc, err := ig.GetSvs()
	if err != nil {
		return err, gData
	}
	gData["servers"] = svc
	if gRPC.Server == nil || *gRPC.Server == "" {
		return err, gData
	}
	methods, err := ig.GetMethod(*gRPC.Server)
	if err != nil {
		return err, gData
	}
	gData["methods"] = methods
	if gRPC.Method == nil || *gRPC.Method == "" {
		return err, gData
	}
	results, err := ig.GetReq(*gRPC.Server, *gRPC.Method)
	if err != nil {
		return err, gData
	}
	req := make(map[string]interface{})
	req["type"] = results.RequestType
	req["stream"] = results.RequestStream
	req["message"] = results.MessageTypes
	req["enum"] = results.EnumTypes
	req["body"] = results.Body
	gData["request"] = req

	return err, gData
}

func (i *InterfaceTemplateService) CreateUserConfig(userConfig interfacecase.ApiUserConfig) (err error) {
	userConfigOld := interfacecase.ApiUserConfig{}
	err = global.GVA_DB.Where("project_id = ? and user_id = ?", userConfig.ProjectID, userConfig.UserID).First(&userConfigOld).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GVA_LOG.Warn("用户配置不存在，创建新的配置")
		} else {
			return err
		}
	}
	userConfigOld.ApiConfigID = userConfig.ApiConfigID
	userConfigOld.UserID = userConfig.UserID
	userConfigOld.ApiEnvID = userConfig.ApiEnvID
	userConfigOld.ProjectID = userConfig.ProjectID
	err = global.GVA_DB.Where("id = ?", userConfig.ID).Save(&userConfigOld).Error
	return err
}

func (i *InterfaceTemplateService) GetUserConfig(projectID uint, userID uint) (userConfig *interfacecase.ApiUserConfig, err error) {
	err = global.GVA_DB.Model(interfacecase.ApiUserConfig{}).Preload("ApiConfig").Preload("ApiEnv").
		Scopes(projectDB(projectID)).
		Where("user_id = ?", userID).First(&userConfig).Error
	return
}
