package interfacecase

import (
	"context"
	"fmt"
	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"
	"go.uber.org/zap"
)

type RunCaseService struct {
}

// RunTestCase TestCase排序

func (r *RunCaseService) RunTestCaseStep(runCase request.RunCaseReq, runType interfacecase.RunType) (reports *interfacecase.ApiReport, err error) {
	reports, err = runTestCase.RunStep(runCase, runType)
	return
}

func (r *RunCaseService) RunApiCase(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunCase(runCase, runType)
	if err != nil {
		return
	}
	return
}

func (r *RunCaseService) RunBoomerDebug(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunBoomerDebug(runCase, runType)
	return
}

func (r *RunCaseService) RunBoomer(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	report, err = runTestCase.RunBoomer(runCase, runType)
	return
}

func (r *RunCaseService) RunMasterBoomer(runCase request.RunCaseReq, runType interfacecase.RunType) (*interfacecase.ApiReport, error) {
	c, err := client.NewClient(fmt.Sprintf("%s:%s", global.GVA_CONFIG.YangFan.Master, global.GVA_CONFIG.YangFan.MasterBoomerProt))
	if err != nil {
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

func (r *RunCaseService) RunTimerTask(runCase request.RunCaseReq, runType interfacecase.RunType) {
	if runCase.TaskID > 0 {
		_, err := runTestCase.RunTimerTask(runCase, runType)
		if err != nil {
			return
		}
		return
	}
	if runCase.TagID > 0 {
		_, err := runTestCase.RunTimerTag(runCase, runType)
		if err != nil {
			return
		}
		return
	}
	return
}

func (r *RunCaseService) RunApi(runCase request.RunCaseReq) (reports *interfacecase.ApiReport, err error) {
	report, err := runTestCase.RunApi(runCase, interfacecase.RunType(runCase.RunType))
	if err != nil {
		return nil, err
	}
	return report, nil
}
