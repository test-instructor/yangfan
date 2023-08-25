package interfacecase

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type ApiCIService struct{}

var runCaseService = RunCaseService{}

func (ci *ApiCIService) RunTag(tagReq interfacecaseReq.CIRun) (error, interface{}) {
	if tagReq.TagID == 0 || tagReq.ProjectID == 0 || tagReq.EnvID == 0 {
		return errors.New("参数设置错误"), nil
	}
	key := uuid.NewV4()
	fmt.Print(key.String())
	arc := interfacecase.ApiReportCI{
		TagID:       tagReq.TagID,
		EnvID:       tagReq.EnvID,
		ProjectID:   tagReq.ProjectID,
		Key:         key.String(),
		Other:       tagReq.Other,
		CallbackUrl: tagReq.CallbackUrl,
	}
	err := global.GVA_DB.Create(&arc).Error
	if err != nil {
		return err, nil
	}
	req := request.RunCaseReq{
		TagID:        tagReq.TagID,
		Env:          tagReq.EnvID,
		ProjectID:    tagReq.ProjectID,
		ReportCIID:   arc.ID,
		ApiMessageID: tagReq.MessageID,
		RunType:      8,
	}
	go runCaseService.RunTimerTask(req)
	data := make(map[string]interface{})
	data["report"] = arc.ID
	data["key"] = arc.Key
	return nil, data

}

func (ci *ApiCIService) GetRepost(tagReq interfacecaseReq.CIRun) (error, interface{}) {
	if tagReq.ReportID == 0 || tagReq.Key == "" {
		return errors.New("参数设置错误"), nil
	}
	var arc interfacecase.ApiReportCI
	err := global.GVA_DB.Preload("Report").First(&arc, "id = ? and `key` = ?", tagReq.ReportID, tagReq.Key).Error
	if err != nil {
		return err, nil
	}
	if arc.Report == nil {
		return errors.New("测试执行中，请稍后查询"), nil
	}
	return nil, arc.Report

}
