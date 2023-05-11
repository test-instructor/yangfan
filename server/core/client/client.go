package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/test-instructor/yangfan/proto/master"
	"github.com/test-instructor/yangfan/server/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"sync"
	"time"
)

type Client struct {
	master.MasterClient
}

var clientMap sync.Map
var clientLock sync.Mutex

func GetGRpcCredentialOption() grpc.DialOption {
	return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12}))
}

func NewClient(host string) (*Client, error) {
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

func newClient(host string) (*Client, error) {
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
		GetGRpcCredentialOption(),
		grpc.WithAuthority(host),
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: bc, MinConnectTimeout: time.Minute}),
		grpc.WithChainUnaryInterceptor(retry.UnaryClientInterceptor(retryMiddlewareConfig...)),
		grpc.WithChainStreamInterceptor(retry.StreamClientInterceptor(retryMiddlewareConfig...)),
	)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("dial remote server fail %v", err))
		return nil, err
	}

	return &Client{
		MasterClient: master.NewMasterClient(c),
	}, nil
}
