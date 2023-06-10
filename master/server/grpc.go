package server

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"os"
	"time"
)

type masterServer struct {
	master.UnimplementedMasterServer
	*MasterBoom
}

var okResp = &master.Resp{Code: 0, Message: "success"}
var errResp = &master.Resp{Code: 1, Message: "error"}

func (b masterServer) Start(ctx context.Context, request *master.StartReq) (resp *master.StartResp, err error) {

	req := hrp.StartRequestPlatformBody{
		Profile: *boomer.NewProfile(),
	}
	req.Profile.SpawnCount = request.Profile.SpawnCount
	req.Profile.SpawnRate = request.Profile.SpawnRate
	req.Profile.RunTime = request.Profile.RunTime
	req.Profile.MaxRPS = request.Profile.MaxRPS
	req.Profile.LoopCount = request.Profile.LoopCount
	req.Profile.RequestIncreaseRate = request.Profile.RequestIncreaseRate
	req.Profile.MemoryProfile = request.Profile.MemoryProfile
	req.Profile.MemoryProfileDuration = time.Duration(request.Profile.MemoryProfileDuration)
	req.Profile.CPUProfile = request.Profile.CPUProfile
	req.Profile.CPUProfileDuration = time.Duration(request.Profile.CPUProfileDuration)
	req.Profile.PrometheusPushgatewayURL = request.Profile.PrometheusPushgatewayURL
	req.Profile.DisableCompression = request.Profile.DisableCompression
	req.Profile.DisableKeepalive = request.Profile.DisableKeepalive
	req.ID = uint(request.Profile.ID)

	if req.ID < 1 {
		err = errors.New("missing testcases ID")
		return
	}

	b.SetTestCasesID(req.ID)
	err = b.StartPlatform(&req.Profile)

	if err != nil {
		global.GVA_LOG.Error("start platform error", zap.Error(err))
		resp = new(master.StartResp)
		resp.Resp = errResp
	} else {
		resp = new(master.StartResp)
		resp.Resp = okResp
		global.IgnoreInstall = true
	}
	return
}

func (b masterServer) Rebalance(ctx context.Context, req *master.RebalanceReq) (resp *master.RebalanceResp, err error) {
	defer func() {
		resp = new(master.RebalanceResp)
		if err != nil {
			resp.Resp = errResp
		} else {
			resp.Resp = okResp
		}
	}()

	hrpReq := hrp.RebalanceRequestBody{
		Profile: *b.Boomer.GetProfile(),
	}
	hrpReq.Profile.SpawnCount = req.SpawnCount
	hrpReq.Profile.SpawnRate = req.SpawnRate
	err = b.Boomer.ReBalance(&hrpReq.Profile)
	return
}

func (b masterServer) Work(ctx context.Context, req *master.WorkReq) (*master.WorkResp, error) {
	workInfo := b.GetWorkersInfo()
	workResp := new(master.WorkResp)
	var works []*master.Work
	for k, _ := range workInfo {
		work := &master.Work{
			Id:                workInfo[k].ID,
			Ip:                workInfo[k].IP,
			Os:                workInfo[k].OS,
			Arch:              workInfo[k].Arch,
			State:             workInfo[k].State,
			Heartbeat:         workInfo[k].Heartbeat,
			UserCount:         workInfo[k].UserCount,
			WorkerCpuUsage:    workInfo[k].WorkerMemoryUsage,
			CpuUsage:          workInfo[k].CPUUsage,
			CpuWarningEmitted: workInfo[k].CPUWarningEmitted,
			WorkerMemoryUsage: workInfo[k].WorkerCPUUsage,
			MemoryUsage:       workInfo[k].MemoryUsage,
		}
		works = append(works, work)
	}
	workResp.Work = works
	return workResp, nil
}

func (b masterServer) Master(ctx context.Context, req *master.MasterReq) (*master.MasterResp, error) {
	masterInfo := b.GetMasterInfo()
	var masterResp = new(master.MasterResp)
	masterResp.State = masterInfo["state"].(int32)
	masterResp.Workers = int32(masterInfo["workers"].(int))
	masterResp.TargetUsers = masterInfo["target_users"].(int64)
	masterResp.CurrentUsers = int32(masterInfo["current_users"].(int))
	return masterResp, nil
}

func (b masterServer) Stop(ctx context.Context, req *master.StopReq) (resp *master.StopResp, err error) {
	defer func() {
		resp = new(master.StopResp)
		if err != nil {
			resp.Resp = errResp
		} else {
			go func() {
				time.Sleep(30 * time.Second)
				os.Exit(0)
			}()
			resp.Resp = okResp
		}
	}()

	err = b.Boomer.Stop()

	return
}

func (b masterServer) Quit(ctx context.Context, req *master.QuitReq) (*master.QuitResp, error) {
	//TODO implement me
	panic("implement me")
}
