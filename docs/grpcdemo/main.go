package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcdemo/pb"
)

type GrpcDemoServer struct {
	*pb.UnimplementedServerDemoServer
}

func (s *GrpcDemoServer) ProcessData(_ context.Context, req *pb.ProcessDataReq) (*pb.ProcessDataResp, error) {
	response := &pb.ProcessDataResp{
		Hostname: req.Hostname,
		Ip:       req.Ip, // 将请求数据直接作为返回数据
	}
	return response, nil
}

func (s *GrpcDemoServer) Get(_ context.Context, _ *pb.GetReq) (*pb.GetResp, error) {
	hostname, _ := os.Hostname()
	response := &pb.GetResp{
		Hostname: hostname,
	}
	return response, nil
}

func main() {
	// 创建gRPC服务器
	server := grpc.NewServer()

	// 注册您的gRPC服务
	stringConverterServer := &GrpcDemoServer{}
	pb.RegisterServerDemoServer(server, stringConverterServer)

	// 启用gRPC反射服务
	reflection.Register(server)

	// 监听服务器端口
	listener, err := net.Listen("tcp", ":7967")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 启动服务器
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
