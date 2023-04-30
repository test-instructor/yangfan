package boomer

import (
	"github.com/test-instructor/yangfan/proto/pb"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var ns *grpc.Server

func StartGrpc(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		global.GVA_LOG.Panic("failed to listen", zap.Error(err))
	}
	ns = grpc.NewServer()
	pb.RegisterBoomerSerServer(ns, &masterServer{MasterBoom: NewMasterBoom()})
	reflection.Register(ns)
	if err := ns.Serve(lis); err != nil {
		global.GVA_LOG.Panic("failed to serve", zap.Error(err))
	}
}
