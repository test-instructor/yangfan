package main

import (
	"context"
	"fmt"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/config"
	"yangfan-ui/internal/logger"
	"yangfan-ui/internal/platformclient"
	"yangfan-ui/internal/service"

	"go.uber.org/zap"
)

type App struct {
	ctx   context.Context
	store *config.Store
	auth  *auth.Manager
	svc   *service.PlatformService
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	logger.Info("App startup initiated")

	store, err := config.New("yangfan-ui")
	if err != nil {
		logger.Error("Failed to initialize config store", zap.Error(err))
		return
	}
	a.store = store

	// Reconfigure logger with stored settings
	cfg := store.Get()
	logger.Setup(logger.Config{
		Level:     cfg.LogLevel,
		Prefix:    cfg.LogPrefix,
		Retention: cfg.LogRetention,
	})

	a.auth = auth.New(store)
	a.svc = service.NewPlatformService(platformclient.New(store), a.auth)
	logger.Info("App services initialized", zap.String("BaseURL", store.BaseURL()))
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RunTestCase(path string) string {
	return fmt.Sprintf("Preparing to run test case at: %s. HRP Integration Active.", path)
}

func (a *App) GetBaseURL() (string, bool) {
	if a.store == nil {
		return "", false
	}
	u := a.store.BaseURL()
	return u, u != ""
}

func (a *App) SetBaseURL(baseURL string) error {
	if a.store == nil {
		return fmt.Errorf("config store 未初始化")
	}
	return a.store.SetBaseURL(baseURL)
}

func (a *App) HasToken() bool {
	if a.store == nil {
		return false
	}
	return a.store.Token() != ""
}

func (a *App) ClearAuth() error {
	if a.auth == nil {
		return fmt.Errorf("auth manager 未初始化")
	}
	return a.auth.Clear()
}

func (a *App) Captcha() (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	res, err := a.svc.Captcha(a.ctx)
	if res == nil && err == nil {
		return map[string]any{}, nil
	}
	return res, err
}

func (a *App) Login(username string, password string, captcha string, captchaId string) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.Login(a.ctx, username, password, captcha, captchaId, "ui-node")
}

func (a *App) GetUserInfo() (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetUserInfo(a.ctx)
}

func (a *App) GetLogConfig() (map[string]any, error) {
	if a.store == nil {
		return nil, fmt.Errorf("config store 未初始化")
	}
	cfg := a.store.Get()
	return map[string]any{
		"logLevel":     cfg.LogLevel,
		"logPrefix":    cfg.LogPrefix,
		"logRetention": cfg.LogRetention,
	}, nil
}

func (a *App) SetLogConfig(level string, prefix string, retention int) error {
	if a.store == nil {
		return fmt.Errorf("config store 未初始化")
	}
	if err := a.store.SetLogConfig(level, prefix, retention); err != nil {
		return err
	}
	// Reload logger immediately
	logger.Setup(logger.Config{
		Level:     level,
		Prefix:    prefix,
		Retention: retention,
	})
	logger.Info("Logger configuration updated", zap.String("level", level), zap.Int("retention", retention))
	return nil
}

func (a *App) SetUserAuthority(authorityId uint, projectId uint) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.SetUserAuthority(a.ctx, authorityId, projectId)
}
