package interfacecase

import (
	"context"
	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/server/core/client"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"
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

func (r *RunCaseService) RunMasterBoomer(runCase request.RunCaseReq, runType interfacecase.RunType) (report *interfacecase.ApiReport, err error) {
	c, err := client.NewClient(global.GVA_CONFIG.GrpcServer.Master)
	if err != nil {
		return nil, err
	}
	_, err = c.MasterClient.Start(context.Background(), &master.StartReq{
		Profile: &master.Profile{
			SpawnCount: runCase.Operation.SpawnCount,
			SpawnRate:  runCase.Operation.SpawnRate,
			ID:         uint64(runCase.CaseID),
		},
	})
	if err != nil {
		return nil, err
	}
	return
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
