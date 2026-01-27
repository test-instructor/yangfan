package service

import (
	"context"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/platformapi"
	"yangfan-ui/internal/platformclient"
)

type PlatformService struct {
	client *platformclient.Client
	auth   *auth.Manager
}

type loginData struct {
	User      map[string]any `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type userInfoData struct {
	UserInfo map[string]any `json:"userInfo"`
}

func NewPlatformService(client *platformclient.Client, authManager *auth.Manager) *PlatformService {
	return &PlatformService{
		client: client,
		auth:   authManager,
	}
}

func (s *PlatformService) Captcha(ctx context.Context) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointCaptcha, nil, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return map[string]any{}, nil
	}
	return data, nil
}

func (s *PlatformService) Login(ctx context.Context, username string, password string, captcha string, captchaId string, node string) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[loginData](ctx, s.client, platformapi.EndpointLogin, map[string]any{
		"username":  username,
		"password":  password,
		"captcha":   captcha,
		"captchaId": captchaId,
		"node":      node,
	}, nil)
	if err != nil {
		return nil, err
	}
	if err := s.auth.Set(data.Token, data.ExpiresAt); err != nil {
		return nil, err
	}
	return data.User, nil
}

func (s *PlatformService) GetUserInfo(ctx context.Context) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[userInfoData](ctx, s.client, platformapi.EndpointUserInfo, nil, nil)
	if err != nil {
		return nil, err
	}
	return data.UserInfo, nil
}

func (s *PlatformService) SetUserAuthority(ctx context.Context, authorityId uint, projectId uint) (map[string]any, error) {
	headers, _, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointSetAuth, map[string]any{
		"authorityId": authorityId,
		"projectId":   projectId,
	}, nil)
	if err == nil {
		_ = s.auth.ApplyNewTokenFromHeaders(headers)
	}
	if err != nil {
		return nil, err
	}
	return s.GetUserInfo(ctx)
}

func (s *PlatformService) SetSelfInfo(ctx context.Context, info map[string]any) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointSetSelf, info, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PlatformService) ChangePassword(ctx context.Context, password string, newPassword string) (map[string]any, error) {
	_, data, err := platformapi.CallAndDecodeData[map[string]any](ctx, s.client, platformapi.EndpointChangePwd, map[string]any{
		"password":    password,
		"newPassword": newPassword,
	}, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}
