package server

import (
	"context"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/run/runTestCase"
	"github.com/test-instructor/yangfan/server/model/common/request"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
)

type runServer struct {
	run.UnimplementedRunCaseServer
}

func (r runServer) getRunCase(req *run.RunCaseReq) (runCaseReq request.RunCaseReq) {
	runCaseReq.ApiID = uint(req.ApiID)
	runCaseReq.ConfigID = uint(req.ConfigID)
	runCaseReq.CaseID = uint(req.CaseID)
	runCaseReq.RunType = uint(req.RunType)
	runCaseReq.TaskID = uint(req.TaskID)
	runCaseReq.TagID = uint(req.TagID)
	runCaseReq.ProjectID = uint(req.ProjectID)
	runCaseReq.TagID = uint(req.TagID)
	runCaseReq.Env = uint(req.Env)
	return
}

func (r runServer) RunApi(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewRunApi(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}

func (r runServer) RunStep(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewRunCase(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}

func (r runServer) RunCase(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewRunCase(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}

func (r runServer) RunBoomerDebug(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewBoomerDebug(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}

func (r runServer) RunTimerTask(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewRunTask(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}

func (r runServer) RunTimerTag(ctx context.Context, req *run.RunCaseReq) (resp *run.RunCaseResponse, err error) {
	runCaseReq := r.getRunCase(req)
	api := runTestCase.NewRunTag(runCaseReq, interfacecase.RunType(req.RunType))
	report, err := runTestCase.RunTestCase(api)
	resp = new(run.RunCaseResponse)
	if err == nil {
		resp.ReportID = uint32(report.ID)
	}
	return
}
