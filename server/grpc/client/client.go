package client

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/proto/run"
	"github.com/test-instructor/yangfan/proto/tools"
	"github.com/test-instructor/yangfan/server/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type Client struct {
	host              string
	MasterClient      master.MasterClient
	ToolsServerClient tools.ToolsServerClient
	RunClient         run.RunCaseClient
}

var clientMap sync.Map
var clientLock sync.Mutex
var apiClient *Client

func NewClientMap(host string) (*Client, error) {
	var c *Client
	var err error
	clientLock.Lock()
	defer clientLock.Unlock()
	if ct, ok := clientMap.Load(host); ok {
		return ct.(*Client), nil
	}
	c, err = newClient(host)
	if err != nil {
		return nil, err
	}
	clientMap.Store(host, c)
	return c, err
}

func NewClient(host string) (*Client, error) {
	var c *Client
	var err error
	global.GVA_LOG.Debug("[NewClient]host", zap.Any("host", host))

	c, err = newClient(host)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Reconnect() (*Client, error) {
	var c *Client
	var err error
	c, err = newClient(apiClient.host)
	global.GVA_LOG.Debug("[Reconnect]重新连接", zap.Any("apiClient.host", apiClient.host))
	if err != nil {
		global.GVA_LOG.Error("[Reconnect]重新连接失败", zap.Error(err))
		global.GVA_LOG.Error("[Reconnect]重新连接失败", zap.Any("apiClient.host", apiClient.host))
		return nil, err
	}
	apiClient = c
	return apiClient, nil
}

func newClient(host string) (*Client, error) {
	global.GVA_LOG.Debug("[newClient]host", zap.Any("host", host))
	retryMiddlewareConfig := []retry.CallOption{
		retry.WithCodes(codes.Unavailable),
		retry.WithBackoff(retry.BackoffExponential(100 * time.Millisecond)),
		retry.WithMax(5),
		retry.WithOnRetryCallback(func(_ context.Context, attempt uint, err error) {
			if attempt == 0 {
				return
			}
			global.GVA_LOG.Warn(fmt.Sprintf("retry %d times, err: %v\n", attempt, err))
		}),
	}
	bc := backoff.DefaultConfig
	c, err := grpc.Dial(
		host,
		grpc.WithAuthority(host),
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: bc, MinConnectTimeout: time.Minute}),
		grpc.WithChainUnaryInterceptor(retry.UnaryClientInterceptor(retryMiddlewareConfig...)),
		grpc.WithChainStreamInterceptor(retry.StreamClientInterceptor(retryMiddlewareConfig...)),
	)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("dial remote server fail %v", err))
		return nil, err
	}
	apiClient = &Client{
		host:              host,
		MasterClient:      master.NewMasterClient(c),
		ToolsServerClient: tools.NewToolsServerClient(c),
		RunClient:         run.NewRunCaseClient(c),
	}
	return apiClient, nil
}
