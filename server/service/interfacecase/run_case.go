package interfacecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"
	"go.uber.org/zap"
)

type RunCaseService struct {
	c *client.Client
}

// RunTestCase TestCase排序

func (r *RunCaseService) newClient() error {
	c, err := client.NewClient(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.RunServer, global.GVA_CONFIG.YangFan.RunServerGrpcPort))
	if err != nil {
		return err
	}
	r.c = c
	return err
}

func (r *RunCaseService) getRunCase(runCaseReq request.RunCaseReq) (req *run.RunCaseReq) {
	req = new(run.RunCaseReq)
	req.ApiID = uint32(runCaseReq.ApiID)
	req.ConfigID = uint32(runCaseReq.ConfigID)
	req.CaseID = uint32(runCaseReq.CaseID)
	req.RunType = run.RunType(runCaseReq.RunType)
	req.TaskID = uint32(runCaseReq.TaskID)
	req.TagID = uint32(runCaseReq.TagID)
	req.ProjectID = uint32(runCaseReq.ProjectID)
	req.TaskID = uint32(runCaseReq.TaskID)
	req.Env = uint32(runCaseReq.Env)
	return
}

func (r *RunCaseService) setRunCaseMsg(req *run.RunCaseReq, msg *interfacecase.ApiMessage) {
	var Type run.NotifierType
	switch msg.Type {
	case interfacecase.MessageTypeFeishu:
		Type = run.NotifierType_Feishu
	case interfacecase.MessageTypeDingtalk:
		Type = run.NotifierType_Dingtalk
	case interfacecase.MessageTypeWechat:
		Type = run.NotifierType_Wechat
	default:
		Type = run.NotifierType_Default
	}
	req.Msg = &run.Msg{
		Id:        uint64(msg.ID),
		Name:      msg.Name,
		Type:      Type,
		TypeName:  msg.TypeName,
		Webhook:   msg.WebHook,
		Signature: msg.Signature,
		Fail:      msg.Fail,
	}
}

func (r *RunCaseService) RunTestCaseStep(runCase request.RunCaseReq) (reports *interfacecase.ApiReport, err error) {
	err = r.newClient()
	if err != nil {
		return nil, err
	}
	step, err := r.c.RunClient.RunStep(context.Background(), r.getRunCase(runCase))
	if err != nil {
		return nil, err
	}
	reports = new(interfacecase.ApiReport)
	reports.ID = uint(step.ReportID)
	return
}

func (r *RunCaseService) RunApiCase(runCase request.RunCaseReq) (report *interfacecase.ApiReport, err error) {
	err = r.newClient()
	if err != nil {
		return nil, err
	}
	step, err := r.c.RunClient.RunCase(context.Background(), r.getRunCase(runCase))
	if err != nil {
		return nil, err
	}
	report = new(interfacecase.ApiReport)
	report.ID = uint(step.ReportID)
	return
}

func (r *RunCaseService) RunBoomerDebug(runCase request.RunCaseReq) (report *interfacecase.ApiReport, err error) {
	err = r.newClient()
	if err != nil {
		return nil, err
	}
	step, err := r.c.RunClient.RunBoomerDebug(context.Background(), r.getRunCase(runCase))
	if err != nil {
		return nil, err
	}
	report = new(interfacecase.ApiReport)
	report.ID = uint(step.ReportID)
	return
}

func (r *RunCaseService) RunBoomer(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunBoomer(runCase, runType)
	return
}

func (r *RunCaseService) RunMasterBoomer(runCase request.RunCaseReq, _ interfacecase.RunType) (*interfacecase.ApiReport, error) {
	global.GVA_LOG.Debug("RunMasterBoomer", zap.Any("master host", fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt)))

	c, err := client.NewClient(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))
	if err != nil {
		global.GVA_LOG.Error("RunMasterBoomer", zap.Any("err host", errors.New(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))))
		return nil, err
	}
	_, err = c.MasterClient.Start(context.Background(), &master.StartReq{
		Profile: &master.Profile{
			SpawnCount:               runCase.Operation.SpawnCount,
			SpawnRate:                runCase.Operation.SpawnRate,
			ID:                       uint64(runCase.CaseID),
			PrometheusPushgatewayURL: global.GVA_CONFIG.YangFan.PrometheusPushgatewayURL,
		},
	})
	if err == nil {
		var report interfacecase.ApiReport
		global.GVA_DB.Model(&interfacecase.ApiReport{}).Order("id desc").First(&report)
		if err == nil {
			global.GVA_LOG.Error("RunMasterBoomer2", zap.Any("err host", errors.New(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))))
			return &report, nil
		}
	}
	return nil, err
}

func (r *RunCaseService) Rebalance(runCase request.RunCaseReq) (*interfacecase.ApiReport, error) {
	c, err := client.NewClient(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))
	if err != nil {
		return nil, err
	}
	_, err = c.MasterClient.Rebalance(context.Background(), &master.RebalanceReq{
		SpawnCount: runCase.Operation.SpawnCount,
		SpawnRate:  runCase.Operation.SpawnRate,
	})
	if err == nil {
		var report interfacecase.ApiReport
		global.GVA_DB.Model(&interfacecase.ApiReport{}).Order("id desc").First(&report)
		if err == nil {
			return &report, nil
		}
	}
	return nil, nil
}

func (r *RunCaseService) Stop(runCase request.RunCaseReq) (err error) {
	defer func() {
		if err == nil || err.Error() == "rpc error: code = Unknown desc = already stopped" {
			err = global.GVA_DB.Model(&interfacecase.PerformanceReport{}).Where("id = ?", runCase.ReportID).Update("state", interfacecase.StateStopped).Error
			if err != nil {
				global.GVA_LOG.Error("修改性能测试报告状态失败", zap.Error(err))
			}
		}
	}()
	c, err := client.NewClient(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))
	if err != nil {
		return err
	}
	_, err = c.MasterClient.Stop(context.Background(), &master.StopReq{})
	return err
}

func (r *RunCaseService) RunTimerTask(runCase request.RunCaseReq) {
	err := r.newClient()
	if err != nil {
		return
	}
	if runCase.TaskID > 0 {
		req := r.getRunCase(runCase)
		var task interfacecase.ApiTimerTask
		err = global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).Preload("ApiMessage").Where("id = ?", runCase.TaskID).First(&task).Error
		if err == nil && task.ApiMessage != nil {
			r.setRunCaseMsg(req, task.ApiMessage)
		}
		_, err = r.c.RunClient.RunTimerTask(context.Background(), req)
		if err != nil {
			return
		}
		return
	}
	if runCase.TagID > 0 {
		req := r.getRunCase(runCase)
		if runCase.ApiMessageID > 0 {
			var msg interfacecase.ApiMessage
			err = global.GVA_DB.Model(&interfacecase.ApiMessage{}).Where("id = ?", runCase.ApiMessageID).First(&msg).Error
			if err == nil && msg.ID > 0 {
				r.setRunCaseMsg(req, &msg)
			}
		}
		_, err := r.c.RunClient.RunTimerTag(context.Background(), req)
		if err != nil {
			return
		}
		return
	}
	return
}

func (r *RunCaseService) RunTimerTaskBack(taskID uint) {

	req := request.RunCaseReq{
		TagID:   taskID,
		RunType: uint(interfacecase.RunTypeRunTimer),
	}
	r.RunTimerTask(req)

}

func RunTimerTaskBack(taskID uint) func() {
	return func() {
		var r RunCaseService
		r.RunTimerTaskBack(taskID)
	}
}

func (r *RunCaseService) RunApi(runCase request.RunCaseReq) (report *interfacecase.ApiReport, err error) {
	err = r.newClient()
	if err != nil {
		return nil, err
	}
	step, err := r.c.RunClient.RunApi(context.Background(), r.getRunCase(runCase))
	if err != nil {
		return nil, err
	}
	report = new(interfacecase.ApiReport)
	report.ID = uint(step.ReportID)
	return report, nil
}
