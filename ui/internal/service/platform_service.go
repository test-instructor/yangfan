package service

import (
	"context"
	"errors"
	"net/http"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/platformclient"
)

type PlatformService struct {
	client *platformclient.Client
	auth   *auth.Manager
}

func NewPlatformService(client *platformclient.Client, authManager *auth.Manager) *PlatformService {
	return &PlatformService{
		client: client,
		auth:   authManager,
	}
}

func (s *PlatformService) Captcha(ctx context.Context) (map[string]any, error) {
	// 尝试 /api/base/captcha (兼容 curl 示例中的路径)
	// 如果你的后端配置了 system.router-prefix: "/api"，则此处应为 /api/base/captcha
	// 如果后端是默认 ""，但通过 nginx 代理了 /api，则也需要带 /api
	// 这里根据用户 curl，明确是 /api/base/captcha
	headers, body, err := s.client.Do(ctx, http.MethodPost, "/api/base/captcha", nil, nil)
	_ = headers
	if err != nil {
		return nil, err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	if r.Data == nil {
		return map[string]any{}, nil
	}
	return r.Data, nil
}

func (s *PlatformService) Login(ctx context.Context, username string, password string, captcha string, captchaId string, node string) (map[string]any, error) {
	headers, body, err := s.client.Do(ctx, http.MethodPost, "/api/base/login", map[string]any{
		"username":  username,
		"password":  password,
		"captcha":   captcha,
		"captchaId": captchaId,
		"node":      node,
	}, nil)
	_ = headers
	if err != nil {
		return nil, err
	}
	type loginData struct {
		User      map[string]any `json:"user"`
		Token     string         `json:"token"`
		ExpiresAt int64          `json:"expiresAt"`
	}
	r, err := platformclient.DecodeAPIResponse[loginData](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	if err := s.auth.Set(r.Data.Token, r.Data.ExpiresAt); err != nil {
		return nil, err
	}
	return r.Data.User, nil
}

func (s *PlatformService) GetUserInfo(ctx context.Context) (map[string]any, error) {
	headers, body, err := s.client.Do(ctx, http.MethodGet, "/api/user/getUserInfo", nil, nil)
	_ = headers
	if err != nil {
		return nil, err
	}
	type userInfoData struct {
		UserInfo map[string]any `json:"userInfo"`
	}
	r, err := platformclient.DecodeAPIResponse[userInfoData](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	return r.Data.UserInfo, nil
}

func (s *PlatformService) SetUserAuthority(ctx context.Context, authorityId uint, projectId uint) (map[string]any, error) {
	headers, body, err := s.client.Do(ctx, http.MethodPost, "/api/user/setUserAuthority", map[string]any{
		"authorityId": authorityId,
		"projectId":   projectId,
	}, nil)
	if err == nil {
		_ = s.auth.ApplyNewTokenFromHeaders(headers)
	}
	if err != nil {
		return nil, err
	}
	r, err := platformclient.DecodeAPIResponse[map[string]any](body)
	if err != nil {
		return nil, err
	}
	if r.Code != 0 {
		return nil, errors.New(r.Msg)
	}
	return s.GetUserInfo(ctx)
}
