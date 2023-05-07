package server

import (
	"context"
	"errors"
	"github.com/test-instructor/yangfan/hrp"
	"github.com/test-instructor/yangfan/hrp/pkg/boomer"
	"github.com/test-instructor/yangfan/proto/pb"
	"os"
	"time"
)

type masterServer struct {
	pb.UnimplementedBoomerSerServer
	*MasterBoom
}

var okResp = &pb.Resp{Code: 0, Message: "success"}
var errResp = &pb.Resp{Code: 1, Message: "error"}

func (b masterServer) Start(ctx context.Context, request *pb.StartReq) (resp *pb.StartResp, err error) {
	defer func() {
		resp = new(pb.StartResp)
		if err != nil {
			resp.Resp = errResp
		} else {
			resp.Resp = okResp
		}
	}()
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
	return
}

func (b masterServer) Rebalance(ctx context.Context, req *pb.RebalanceReq) (resp *pb.RebalanceResp, err error) {
	defer func() {
		resp = new(pb.RebalanceResp)
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

func (b masterServer) Work(ctx context.Context, req *pb.WorkReq) (*pb.WorkResp, error) {
	workInfo := b.GetWorkersInfo()
	workResp := new(pb.WorkResp)
	var works []*pb.Work
	for k, _ := range workInfo {
		work := &pb.Work{
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

func (b masterServer) Master(ctx context.Context, req *pb.MasterReq) (*pb.MasterResp, error) {
	masterInfo := b.GetMasterInfo()
	var masterResp = new(pb.MasterResp)
	masterResp.State = masterInfo["state"].(int32)
	masterResp.Workers = int32(masterInfo["workers"].(int))
	masterResp.TargetUsers = masterInfo["target_users"].(int64)
	masterResp.CurrentUsers = int32(masterInfo["current_users"].(int))
	return masterResp, nil
}

func (b masterServer) Stop(ctx context.Context, req *pb.StopReq) (resp *pb.StopResp, err error) {
	defer func() {
		resp = new(pb.StopResp)
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

func (b masterServer) Quit(ctx context.Context, req *pb.QuitReq) (*pb.QuitResp, error) {
	//TODO implement me
	panic("implement me")
}
