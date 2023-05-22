package pkg

import (
	"context"
	"github.com/test-instructor/yangfan/server/core/pkg"
	"github.com/test-instructor/yangfan/server/grpc/client"

	"github.com/test-instructor/yangfan/proto/tools"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/model/interfacecase/request"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type ToolsClient struct {
	client *client.Client
}

// 方法：接收消息
func (p *ToolsClient) receivePkgMessages() {
	stream, err := p.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
	if err != nil {
		global.GVA_LOG.Error("[receiveMessages]创建流失败", zap.Error(err))
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			global.GVA_LOG.Error("[receiveMessages]接收消息失败", zap.Error(err))
		}

		global.GVA_LOG.Info("[receiveMessages]接收消息成功", zap.Any("msg", msg))
	}
}

func (p *ToolsClient) createClientConn(target string) (*grpc.ClientConn, error) {
	// 创建连接选项
	opts := []grpc.DialOption{
		grpc.WithInsecure(), // 使用不安全的连接（仅供示例，请根据实际情况配置安全连接）
		grpc.WithBlock(),    // 阻塞连接，直到连接成功或失败
	}

	// 创建上下文和取消函数
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 连接服务端
	conn, err := grpc.DialContext(ctx, target, opts...)
	return conn, err
}

func (p *ToolsClient) RunClient() {
	var err error

	go func() {
		var stream tools.ToolsServer_InstallPackageStreamingMessageClient
		stream, err = p.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
		if err != nil {
			global.GVA_LOG.Error("[RunClient]创建流失败", zap.Error(err))
			for {
				time.Sleep(3 * time.Second)
				p.client, err = client.Reconnect()
				if err != nil {
					global.GVA_LOG.Error("[RunClient]重新连接失败", zap.Error(err))
					continue
				}
				stream, err = p.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
				if err != nil {
					global.GVA_LOG.Error("[RunClient]流式接口连接失败", zap.Error(err))
					continue
				}
				break
			}
		}

		for {
			res, err := stream.Recv()
			if err != nil {
				global.GVA_LOG.Error("[RunClient]接收消息失败", zap.Error(err))
				// 处理连接断开的情况
				// 尝试重连并继续接收消息
				for {
					// 等待一段时间后尝试重连
					time.Sleep(5 * time.Second)
					//conn, err := p.createClientConn(target)
					p.client, err = client.Reconnect()
					if err != nil {
						global.GVA_LOG.Error("[RunClient]重新连接失败", zap.Error(err))
						continue
					}
					stream, err = p.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
					if err != nil {
						global.GVA_LOG.Error("[RunClient]流式接口连接失败", zap.Error(err))
						continue
					}
					break
				}
				continue
			}
			global.GVA_LOG.Debug("[RunClient]接收到消息", zap.Any("res", res))
			p.installPythonPackage(res)
		}
	}()

	// 等待程序退出
	<-make(chan struct{})
}

func (p *ToolsClient) installPythonPackage(res *tools.InstallPackageRes) {
	var isUninstall bool
	var pyPkg interfacecase.HrpPyPkg
	pyPkg.Name = res.Name
	pyPkg.Version = res.Version
	if res.Operate == tools.Operate_REMOVE {
		isUninstall = true
	}
	pyPkg.IsUninstall = &isUninstall
	if err := pkg.PyPkgInstallServiceV2(request.HrpPyPkgRequest{HrpPyPkg: pyPkg}); err != nil {
		global.GVA_LOG.Error("[InitPythonPackage]安装 python 第三方库失败", zap.Any("pyPkg", pyPkg), zap.Error(err))
		return
	}
	global.GVA_LOG.Debug("[InitPythonPackage]安装 python 第三方库成功", zap.Any("pyPkg", pyPkg))
}

func NewRunInstallPkg(client *client.Client) *ToolsClient {
	return &ToolsClient{client: client}
}
