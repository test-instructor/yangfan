package server

import (
	"net"
	"strings"

	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ns *grpc.Server

func StartRunServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		global.GVA_LOG.Panic("failed to listen", zap.Error(err))
	}
	ns = grpc.NewServer()
	run.RegisterRunCaseServer(ns, &runServer{})
	if strings.ToLower(global.GVA_CONFIG.Zap.Level) == "debug" {
		global.GVA_LOG.Debug("启用grpc反射服务")
		reflection.Register(ns)
	}
	if err := ns.Serve(lis); err != nil {
		global.GVA_LOG.Panic("failed to serve", zap.Error(err))
	}
}
