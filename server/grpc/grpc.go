package grpc

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/test-instructor/yangfan/proto/tools"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"sync"
	"time"
	// 替换为你的 protobuf 包路径
)

//	type server struct {
//		clients sync.Map
//		mu      sync.Mutex
//		tools.UnimplementedToolsPackageServer
//	}

type GrpcServerInstallPackage struct {
	clients     map[string]tools.ToolsPackage_InstallPackageStreamingMessageServer
	clientsLock sync.RWMutex
	tools.UnimplementedToolsPackageServer
}

func (s *GrpcServerInstallPackage) InstallPackageStreamingMessage(_ *tools.InstallPackageReq, stream tools.ToolsPackage_InstallPackageStreamingMessageServer) error {
	// 接收连接请求中的用户名
	_, ok := metadata.FromIncomingContext(stream.Context())
	grpc.Method(stream.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	clientID, err := s.generateClientID()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to generate client ID: %v", err)
	}

	// 保存连接的客户端
	s.clientsLock.Lock()
	s.clients[clientID] = stream
	s.clientsLock.Unlock()

	for {
		select {
		case <-stream.Context().Done():
			// 客户端连接已断开，退出循环 global.GVA_LOG.Debug("客户端连接失败", zap.String("clientID", clientID))
			s.clientsLock.Lock()
			delete(s.clients, clientID)
			s.clientsLock.Unlock()
			return nil
		default:
			global.GVA_LOG.Debug("已连接的客户端", zap.String("clientID", clientID))
			time.Sleep(3 * time.Second)
		}
	}
}

func (s *GrpcServerInstallPackage) SendMessageToSavedClients(res *tools.InstallPackageRes) {
	s.clientsLock.RLock()
	defer s.clientsLock.RUnlock()
	global.GVA_LOG.Debug("准备给客户端推送消息", zap.String("res", res.String()))

	for clientID, client := range s.clients {
		err := client.Send(res)
		if err != nil {
			// 处理发送失败的情况
			global.GVA_LOG.Error("Failed to send message to client", zap.String("clientID", clientID), zap.Error(err))
		}
	}
}

func (s *GrpcServerInstallPackage) generateClientID() (string, error) {
	bytes := make([]byte, 16) // 生成 16 字节的随机字节
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	clientID := hex.EncodeToString(bytes) // 将随机字节转换为十六进制字符串表示
	return clientID, nil
}

var once sync.Once

var ServerInstallPackage *GrpcServerInstallPackage

func newGrpcServerInstallPackage() *GrpcServerInstallPackage {
	if ServerInstallPackage == nil {
		once.Do(func() {
			if ServerInstallPackage == nil {
				ServerInstallPackage = &GrpcServerInstallPackage{
					clients: make(map[string]tools.ToolsPackage_InstallPackageStreamingMessageServer),
				}
			}
		})
	}
	return ServerInstallPackage
}

func RunGrpcServer() {
	listenAddr := "0.0.0.0:" + global.GVA_CONFIG.GrpcServer.BackgroundGrpcPort // 服务监听的地址和端口
	global.GVA_LOG.Debug("准备监听", zap.String("listenAddr", listenAddr))
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		global.GVA_LOG.Panic("Failed to listen", zap.Error(err))
	}

	s := grpc.NewServer()
	svr := newGrpcServerInstallPackage()
	tools.RegisterToolsPackageServer(s, svr)

	global.GVA_LOG.Debug("Server listening on %s\n", zap.String("listenAddr", listenAddr))
	if err := s.Serve(lis); err != nil {
		global.GVA_LOG.Panic("Failed to listen", zap.Error(err))
	}
}
