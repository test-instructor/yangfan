package platformapi

import (
	"context"
	"errors"
	"net/http"

	"yangfan-ui/internal/platformclient"
)

type Endpoint string

const (
	EndpointHealthAPI Endpoint = "health_api"
	EndpointHealth    Endpoint = "health"
	EndpointCaptcha   Endpoint = "captcha"
	EndpointLogin     Endpoint = "login"
	EndpointUserInfo  Endpoint = "user_info"
	EndpointSetAuth   Endpoint = "set_user_authority"
	EndpointSetSelf   Endpoint = "set_self_info"
	EndpointChangePwd Endpoint = "change_password"
)

type Spec struct {
	Method string
	Path   string
}

var specs = map[Endpoint]Spec{
	EndpointHealthAPI: {Method: http.MethodGet, Path: "/api/health/yangfan/server"},
	EndpointHealth:    {Method: http.MethodGet, Path: "/health"},
	EndpointCaptcha:   {Method: http.MethodPost, Path: "/api/base/captcha"},
	EndpointLogin:     {Method: http.MethodPost, Path: "/api/base/login"},
	EndpointUserInfo:  {Method: http.MethodGet, Path: "/api/user/getUserInfo"},
	EndpointSetAuth:   {Method: http.MethodPost, Path: "/api/user/setUserAuthority"},
	EndpointSetSelf:   {Method: http.MethodPut, Path: "/api/user/setSelfInfo"},
	EndpointChangePwd: {Method: http.MethodPost, Path: "/api/user/changePassword"},
}

func GetSpec(e Endpoint) (Spec, error) {
	s, ok := specs[e]
	if !ok {
		return Spec{}, errors.New("未知的接口标识")
	}
	if s.Method == "" || s.Path == "" {
		return Spec{}, errors.New("接口配置不完整")
	}
	return s, nil
}

func Call(ctx context.Context, client *platformclient.Client, endpoint Endpoint, reqBody any, extraHeaders map[string]string) (http.Header, []byte, error) {
	if client == nil {
		return nil, nil, errors.New("client 为空")
	}
	s, err := GetSpec(endpoint)
	if err != nil {
		return nil, nil, err
	}
	return client.Do(ctx, s.Method, s.Path, reqBody, extraHeaders)
}

func CallAndDecodeData[T any](ctx context.Context, client *platformclient.Client, endpoint Endpoint, reqBody any, extraHeaders map[string]string) (http.Header, T, error) {
	headers, body, err := Call(ctx, client, endpoint, reqBody, extraHeaders)
	if err != nil {
		var zero T
		return headers, zero, err
	}
	r, err := platformclient.DecodeAPIResponse[T](body)
	if err != nil {
		var zero T
		return headers, zero, err
	}
	if r.Code != 0 {
		var zero T
		return headers, zero, errors.New(r.Msg)
	}
	return headers, r.Data, nil
}
