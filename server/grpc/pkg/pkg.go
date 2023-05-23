package pkg

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/test-instructor/yangfan/server/core/pkg"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"
	"strconv"

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
func (t *ToolsClient) receivePkgMessages() {
	stream, err := t.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
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

func (t *ToolsClient) createClientConn(target string) (*grpc.ClientConn, error) {
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

func (t *ToolsClient) RunInstallPkg() {
	var err error

	go func() {
		var stream tools.ToolsServer_InstallPackageStreamingMessageClient
		stream, err = t.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
		if err != nil {
			global.GVA_LOG.Error("[RunInstallPkg]创建流失败", zap.Error(err))
			for {
				time.Sleep(3 * time.Second)
				t.client, err = client.Reconnect()
				if err != nil {
					global.GVA_LOG.Error("[RunInstallPkg]重新连接失败", zap.Error(err))
					continue
				}
				stream, err = t.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
				if err != nil {
					global.GVA_LOG.Error("[RunInstallPkg]流式接口连接失败", zap.Error(err))
					continue
				}
				break
			}
		}

		for {
			res, err := stream.Recv()
			if err != nil {
				global.GVA_LOG.Error("[RunInstallPkg]接收消息失败", zap.Error(err))
				// 处理连接断开的情况
				// 尝试重连并继续接收消息
				for {
					// 等待一段时间后尝试重连
					time.Sleep(5 * time.Second)
					//conn, err := p.createClientConn(target)
					t.client, err = client.Reconnect()
					if err != nil {
						global.GVA_LOG.Error("[RunInstallPkg]重新连接失败", zap.Error(err))
						continue
					}
					stream, err = t.client.ToolsServerClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
					if err != nil {
						global.GVA_LOG.Error("[RunInstallPkg]流式接口连接失败", zap.Error(err))
						continue
					}
					break
				}
				continue
			}
			global.GVA_LOG.Debug("[RunInstallPkg]接收到消息", zap.Any("res", res))
			t.installPythonPackage(res)
		}
	}()

	// 等待程序退出
	<-make(chan struct{})
}

func (t *ToolsClient) installPythonPackage(res *tools.InstallPackageRes) {
	var isUninstall bool
	var pyPkg interfacecase.HrpPyPkg
	pyPkg.Name = res.GetName()
	pyPkg.Version = res.GetVersion()
	if res.GetOperate() == tools.Operate_REMOVE {
		isUninstall = true
	}
	pyPkg.IsUninstall = &isUninstall
	if err := pkg.PyPkgInstallServiceV2(request.HrpPyPkgRequest{HrpPyPkg: pyPkg}); err != nil {
		global.GVA_LOG.Error("[InitPythonPackage]安装 python 第三方库失败", zap.Any("pyPkg", pyPkg), zap.Error(err))
		return
	}
	global.GVA_LOG.Debug("[InitPythonPackage]安装 python 第三方库成功", zap.Any("pyPkg", pyPkg))
}

func (t *ToolsClient) RunSetTimerTask() {
	var err error

	go func() {
		var stream tools.ToolsServer_SetTaskStreamingMessageClient
		stream, err = t.client.ToolsServerClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
		if err != nil {
			global.GVA_LOG.Error("[RunSetTimerTask]创建流失败", zap.Error(err))
			for {
				time.Sleep(3 * time.Second)
				t.client, err = client.Reconnect()
				if err != nil {
					global.GVA_LOG.Error("[RunSetTimerTask]重新连接失败", zap.Error(err))
					continue
				}
				stream, err = t.client.ToolsServerClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
				if err != nil {
					global.GVA_LOG.Error("[RunSetTimerTask]流式接口连接失败", zap.Error(err))
					continue
				}
				break
			}
		}

		for {
			res, err := stream.Recv()
			if err != nil {
				global.GVA_LOG.Error("[RunSetTimerTask]接收消息失败", zap.Error(err))
				// 处理连接断开的情况
				// 尝试重连并继续接收消息
				for {
					// 等待一段时间后尝试重连
					time.Sleep(5 * time.Second)
					//conn, err := p.createClientConn(target)
					t.client, err = client.Reconnect()
					if err != nil {
						global.GVA_LOG.Error("[RunSetTimerTask]重新连接失败", zap.Error(err))
						continue
					}
					stream, err = t.client.ToolsServerClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
					if err != nil {
						global.GVA_LOG.Error("[RunSetTimerTask]流式接口连接失败", zap.Error(err))
						continue
					}
					break
				}
				continue
			}
			global.GVA_LOG.Debug("[RunSetTimerTask]接收到消息", zap.Any("res", res))
			t.setTimerTask(res)
		}
	}()

	// 等待程序退出
	<-make(chan struct{})
}

func (t *ToolsClient) setTimerTask(res *tools.SetTaskRes) {
	if res == nil {
		return
	}
	var task interfacecase.ApiTimerTask

	if res.GetTimerStatus() == tools.TimerStatusOperate_DELETE {
		global.GVA_LOG.Debug("[setTimerTask]删除定时任务", zap.Any("task", res))
		t.deleteTimer(res)
		return
	}

	err := global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).Where("id = ?", res.GetID()).First(&task).Error
	if err != nil {
		global.GVA_LOG.Error("[setTimerTask]查询定时任务失败", zap.Error(err))
		return
	}
	if res.GetTimerStatus() == tools.TimerStatusOperate_RESET {
		global.GVA_LOG.Debug("[setTimerTask]重置定时任务", zap.Any("task", task))
		t.resetTimer(task)
		return
	}
	if res.GetTimerStatus() == tools.TimerStatusOperate_ADD {
		global.GVA_LOG.Debug("[setTimerTask]添加定时任务", zap.Any("task", task))
		t.addTimer(task)
		return
	}
}

func (t *ToolsClient) deleteTimer(res *tools.SetTaskRes) {
	global.GVA_LOG.Debug("[deleteTimer]删除定时任务", zap.Any("task", res))
	global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(res.GetID())))))
	global.GVA_Timer.Remove(strconv.Itoa(int(res.GetID())), int(res.GetEntryID()))
	global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(res.GetID())))))
	return
}

func (t *ToolsClient) resetTimer(task interfacecase.ApiTimerTask) {
	global.GVA_LOG.Debug("[deleteTimer]删除定时任务", zap.Any("task", task))
	global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(task.ID)))))
	global.GVA_Timer.Remove(strconv.Itoa(int(task.ID)), task.EntryID)
	global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(task.ID)))))
	task.EntryID = 0
	global.GVA_DB.Save(&task)
	if *task.Status {
		global.GVA_LOG.Debug("[addTimer]添加定时任务", zap.Any("task", task))
		id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(task.ID)), task.RunTime, runTestCase.RunTimerTaskBack(task.ID), cron.WithSeconds())
		global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(task.ID)))))
		if err != nil {
			return
		}
		task.EntryID = int(id)
		err = global.GVA_DB.Save(&task).Error
		if err != nil {
			return
		}
	}
	return
}

func (t *ToolsClient) addTimer(task interfacecase.ApiTimerTask) {
	if *task.Status {
		global.GVA_LOG.Debug("[addTimer]添加定时任务", zap.Any("task", task))
		id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(task.ID)), task.RunTime, runTestCase.RunTimerTaskBack(task.ID), cron.WithSeconds())
		global.GVA_LOG.Debug(fmt.Sprintln(global.GVA_Timer.FindCron(strconv.Itoa(int(task.ID)))))
		if err != nil {
			return
		}
		task.EntryID = int(id)
		err = global.GVA_DB.Save(&task).Error
		if err != nil {
			return
		}
	}
	return
}

func NewRunInstallPkg(client *client.Client) *ToolsClient {
	return &ToolsClient{client: client}
}
