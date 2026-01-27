package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/config"
	"yangfan-ui/internal/logger"
	"yangfan-ui/internal/platformapi"
	"yangfan-ui/internal/platformclient"
	"yangfan-ui/internal/service"

	"go.uber.org/zap"
)

type App struct {
	ctx    context.Context
	store  *config.Store
	auth   *auth.Manager
	client *platformclient.Client
	svc    *service.PlatformService
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
	a.client = platformclient.New(store)
	a.svc = service.NewPlatformService(a.client, a.auth)
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
	normalized, err := config.ValidateBaseURL(baseURL)
	if err != nil {
		logger.Warn("SetBaseURL validate failed", zap.String("baseURL_raw", baseURL), zap.Error(err))
		return err
	}
	logger.Info("SetBaseURL request", zap.String("baseURL_raw", baseURL), zap.String("baseURL_normalized", normalized))
	if err := a.store.SetBaseURL(normalized); err != nil {
		logger.Error("SetBaseURL save failed", zap.String("baseURL_normalized", normalized), zap.Error(err))
		return err
	}
	logger.Info("SetBaseURL saved", zap.String("baseURL_normalized", normalized))
	return nil
}

func (a *App) CheckBaseURLConnectivity(baseURL string) (map[string]any, error) {
	if a.store == nil || a.client == nil {
		return nil, fmt.Errorf("config store 未初始化")
	}
	normalized, err := config.ValidateBaseURL(baseURL)
	if err != nil {
		logger.Warn("CheckBaseURLConnectivity validate failed", zap.String("baseURL_raw", baseURL), zap.Error(err))
		return nil, err
	}

	ctx := a.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	logger.Info("CheckBaseURLConnectivity started", zap.String("baseURL_raw", baseURL), zap.String("baseURL_normalized", normalized))

	var lastStatus int
	var lastBody string
	attempts := []platformapi.Endpoint{platformapi.EndpointHealthAPI, platformapi.EndpointHealth}
	for _, endpoint := range attempts {
		spec, specErr := platformapi.GetSpec(endpoint)
		if specErr != nil {
			continue
		}
		status, _, body, reqErr := a.client.DoWithBaseURL(ctx, normalized, spec.Method, spec.Path, nil, nil)
		lastStatus = status
		bodyText := strings.TrimSpace(string(body))
		lastBody = bodyText
		logger.Info(
			"CheckBaseURLConnectivity response",
			zap.String("endpoint", string(endpoint)),
			zap.Int("status", status),
			zap.String("url", normalized+spec.Path),
			zap.String("body", truncateLogText(bodyText, 300)),
			zap.Error(reqErr),
		)
		if reqErr != nil {
			continue
		}
		if status < 200 || status >= 300 {
			continue
		}
		if strings.Contains(strings.ToLower(bodyText), "ok") {
			return map[string]any{
				"ok":               true,
				"baseURL":          normalized,
				"healthPathTried":  spec.Path,
				"healthStatusCode": status,
			}, nil
		}
	}

	logger.Warn(
		"CheckBaseURLConnectivity failed",
		zap.String("baseURL_normalized", normalized),
		zap.Int("lastStatus", lastStatus),
		zap.String("lastBody", truncateLogText(lastBody, 300)),
	)
	return nil, fmt.Errorf("域名连通性检查失败，请检查域名是否正确")
}

func truncateLogText(s string, max int) string {
	if max <= 0 {
		return ""
	}
	if len(s) <= max {
		return s
	}
	return s[:max]
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

func (a *App) SetSelfInfo(info map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.SetSelfInfo(a.ctx, info)
}

func (a *App) ChangePassword(password string, newPassword string) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.ChangePassword(a.ctx, password, newPassword)
}
