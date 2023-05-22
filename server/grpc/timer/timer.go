package timer

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/test-instructor/yangfan/proto/tools"
	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/grpc/client"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	"github.com/test-instructor/yangfan/server/service/interfacecase/runTestCase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type SetTimerClient struct {
	client *client.Client
}

// 方法：接收消息
func (s *SetTimerClient) receiveMessages() {
	stream, err := s.client.TimerTaskClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
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

func (s *SetTimerClient) createClientConn(target string) (*grpc.ClientConn, error) {
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

func (s *SetTimerClient) RunClient() {
	var err error

	go func() {
		var stream tools.TimerTask_SetTaskStreamingMessageClient
		//stream, err = s.client.ToolsPackageClient.InstallPackageStreamingMessage(context.Background(), &tools.InstallPackageReq{})
		stream, err = s.client.TimerTaskClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
		if err != nil {
			global.GVA_LOG.Error("[RunClient]创建流失败", zap.Error(err))
			for {
				time.Sleep(3 * time.Second)
				s.client, err = client.Reconnect()
				if err != nil {
					global.GVA_LOG.Error("[RunClient]重新连接失败", zap.Error(err))
					continue
				}
				stream, err = s.client.TimerTaskClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
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
					s.client, err = client.Reconnect()
					if err != nil {
						global.GVA_LOG.Error("[RunClient]重新连接失败", zap.Error(err))
						continue
					}
					stream, err = s.client.TimerTaskClient.SetTaskStreamingMessage(context.Background(), &tools.SetTaskReq{})
					if err != nil {
						global.GVA_LOG.Error("[RunClient]流式接口连接失败", zap.Error(err))
						continue
					}
					break
				}
				continue
			}
			global.GVA_LOG.Debug("[RunClient]接收到消息", zap.Any("res", res))
			// TODO: 设置定时任务逻辑
			go s.setTimerTask(res)
		}
	}()

	// 等待程序退出
	<-make(chan struct{})
}

func (s *SetTimerClient) setTimerTask(res *tools.SetTaskRes) {
	if res == nil {
		return
	}
	var task interfacecase.ApiTimerTask
	err := global.GVA_DB.Model(&interfacecase.ApiTimerTask{}).Where("id = ?", res.ID).First(&task).Error
	if err != nil {
		global.GVA_LOG.Error("[setTimerTask]查询定时任务失败", zap.Error(err))
		return
	}
	if res.TimerStatus == tools.TimerStatusOperate_DELETE {
		global.GVA_LOG.Debug("[setTimerTask]删除定时任务", zap.Any("task", task))
		s.deleteTimer(task)
		return
	}
	if res.TimerStatus == tools.TimerStatusOperate_RESET {
		global.GVA_LOG.Debug("[setTimerTask]重置定时任务", zap.Any("task", task))
		s.deleteTimer(task)
		s.addTimer(task)
		return
	}
	if res.TimerStatus == tools.TimerStatusOperate_ADD {
		global.GVA_LOG.Debug("[setTimerTask]添加定时任务", zap.Any("task", task))
		s.addTimer(task)
		return
	}
}

func (s *SetTimerClient) deleteTimer(task interfacecase.ApiTimerTask) {
	global.GVA_LOG.Debug("[deleteTimer]删除定时任务", zap.Any("task", task))
	global.GVA_Timer.Remove(strconv.Itoa(int(task.ID)), task.EntryID)
	task.EntryID = 0
	global.GVA_DB.Save(&task)
	return
}

func (s *SetTimerClient) addTimer(task interfacecase.ApiTimerTask) {
	if *task.Status {
		global.GVA_LOG.Debug("[addTimer]添加定时任务", zap.Any("task", task))
		id, err := global.GVA_Timer.AddTaskByFunc(strconv.Itoa(int(task.ID)), task.RunTime, runTestCase.RunTimerTaskBack(task.ID), cron.WithSeconds())
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
