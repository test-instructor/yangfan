package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"yangfan-ui/internal/auth"
	"yangfan-ui/internal/config"
	"yangfan-ui/internal/devtools"
	"yangfan-ui/internal/logger"
	"yangfan-ui/internal/platformapi"
	"yangfan-ui/internal/platformclient"
	"yangfan-ui/internal/service"

	"go.uber.org/zap"
)

type App struct {
	ctx      context.Context
	store    *config.Store
	auth     *auth.Manager
	client   *platformclient.Client
	svc      *service.PlatformService
	devTools *devtools.DevTools
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
	a.devTools = devtools.New(logger.Log)
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

func (a *App) GetUINodeMenuTree() ([]map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetUINodeMenuTree(a.ctx)
}

func (a *App) GetAndroidDeviceOptionsList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAndroidDeviceOptionsList(a.ctx, query)
}

func (a *App) CreateAndroidDeviceOptions(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateAndroidDeviceOptions(a.ctx, payload)
}

func (a *App) UpdateAndroidDeviceOptions(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateAndroidDeviceOptions(a.ctx, payload)
}

func (a *App) DeleteAndroidDeviceOptions(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAndroidDeviceOptions(a.ctx, id)
}

func (a *App) GetRunConfigList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetRunConfigList(a.ctx, query)
}

func (a *App) CreateRunConfig(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateRunConfig(a.ctx, payload)
}

func (a *App) UpdateRunConfig(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateRunConfig(a.ctx, payload)
}

func (a *App) DeleteRunConfig(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteRunConfig(a.ctx, id)
}

func (a *App) GetAutoStepList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAutoStepList(a.ctx, query)
}

func (a *App) CreateAutoStep(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateAutoStep(a.ctx, payload)
}

func (a *App) UpdateAutoStep(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateAutoStep(a.ctx, payload)
}

func (a *App) DeleteAutoStep(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAutoStep(a.ctx, id)
}

func (a *App) GetAutoCaseStepList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAutoCaseStepList(a.ctx, query)
}

func (a *App) CreateAutoCaseStep(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateAutoCaseStep(a.ctx, payload)
}

func (a *App) UpdateAutoCaseStep(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateAutoCaseStep(a.ctx, payload)
}

func (a *App) DeleteAutoCaseStep(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAutoCaseStep(a.ctx, id)
}

func (a *App) GetAutoCaseList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAutoCaseList(a.ctx, query)
}

func (a *App) CreateAutoCase(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateAutoCase(a.ctx, payload)
}

func (a *App) UpdateAutoCase(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateAutoCase(a.ctx, payload)
}

func (a *App) DeleteAutoCase(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAutoCase(a.ctx, id)
}

func (a *App) GetTimerTaskList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetTimerTaskList(a.ctx, query)
}

func (a *App) CreateTimerTask(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.CreateTimerTask(a.ctx, payload)
}

func (a *App) UpdateTimerTask(payload map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.UpdateTimerTask(a.ctx, payload)
}

func (a *App) DeleteTimerTask(id uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteTimerTask(a.ctx, id)
}

func (a *App) GetAutoReportList(query map[string]any) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAutoReportList(a.ctx, query)
}

func (a *App) FindAutoReport(id uint) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.FindAutoReport(a.ctx, id)
}

func (a *App) FindAutoCaseStepApis(autoCaseStepId uint) ([]map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.FindAutoCaseStepApis(a.ctx, autoCaseStepId)
}

func (a *App) AddAutoCaseStepApi(autoCaseStepId uint, apiId uint, sort uint) (map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.AddAutoCaseStepApi(a.ctx, autoCaseStepId, apiId, sort)
}

func (a *App) DeleteAutoCaseStepApi(autoStepId uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAutoCaseStepApi(a.ctx, autoStepId)
}

func (a *App) SortAutoCaseStepApis(data []map[string]any) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.SortAutoCaseStepApis(a.ctx, data)
}

func (a *App) GetAutoCaseSteps(autoCaseId uint) ([]map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetAutoCaseSteps(a.ctx, autoCaseId)
}

func (a *App) AddAutoCaseStep(caseId uint, stepId uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.AddAutoCaseStep(a.ctx, caseId, stepId)
}

func (a *App) DeleteAutoCaseStepRef(refId uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteAutoCaseStepRef(a.ctx, refId)
}

func (a *App) SortAutoCaseSteps(caseId uint, data []map[string]any) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.SortAutoCaseSteps(a.ctx, caseId, data)
}

func (a *App) SetAutoCaseStepConfig(refId uint, isConfig bool, isStepConfig bool) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.SetAutoCaseStepConfig(a.ctx, refId, isConfig, isStepConfig)
}

func (a *App) GetTimerTaskCases(taskId uint) ([]map[string]any, error) {
	if a.svc == nil {
		return nil, fmt.Errorf("service 未初始化")
	}
	return a.svc.GetTimerTaskCases(a.ctx, taskId)
}

func (a *App) AddTimerTaskCase(taskId uint, caseId uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.AddTimerTaskCase(a.ctx, taskId, caseId)
}

func (a *App) DeleteTimerTaskCaseRef(refId uint) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.DeleteTimerTaskCaseRef(a.ctx, refId)
}

func (a *App) SortTimerTaskCases(taskId uint, data []map[string]any) error {
	if a.svc == nil {
		return fmt.Errorf("service 未初始化")
	}
	return a.svc.SortTimerTaskCases(a.ctx, taskId, data)
}

func (a *App) GetAppConfig() (map[string]any, error) {
	if a.store == nil {
		return nil, fmt.Errorf("config store 未初始化")
	}
	cfg := a.store.Get()
	return map[string]any{
		"environment": cfg.Environment,
		"debugMode":   cfg.DebugMode,
		"theme":       cfg.Theme,
		"language":    cfg.Language,
		"autoLogin":   cfg.AutoLogin,
		"rememberMe":  cfg.RememberMe,
	}, nil
}

func (a *App) SetAppConfig(environment string, debugMode bool, theme string, language string, autoLogin bool, rememberMe bool) error {
	if a.store == nil {
		return fmt.Errorf("config store 未初始化")
	}

	// 验证环境
	env := config.Environment(environment)
	if env != config.EnvDevelopment && env != config.EnvTesting && env != config.EnvProduction {
		return fmt.Errorf("无效的环境: %s", environment)
	}

	// 更新配置
	if err := a.store.SetEnvironment(env); err != nil {
		return err
	}
	if err := a.store.SetDebugMode(debugMode); err != nil {
		return err
	}
	if err := a.store.SetTheme(theme); err != nil {
		return err
	}
	if err := a.store.SetLanguage(language); err != nil {
		return err
	}
	if err := a.store.SetAutoLogin(autoLogin); err != nil {
		return err
	}
	if err := a.store.SetRememberMe(rememberMe); err != nil {
		return err
	}

	logger.Info("App configuration updated",
		zap.String("environment", environment),
		zap.Bool("debugMode", debugMode),
		zap.String("theme", theme),
		zap.String("language", language))
	return nil
}

func (a *App) GetSystemInfo() (map[string]any, error) {
	if a.devTools == nil {
		return map[string]any{
			"appName":      "扬帆自动化测试平台-UI自动化节点",
			"version":      "1.0.0",
			"buildTime":    "2024-01-29",
			"goVersion":    runtime.Version(),
			"wailsVersion": "2.11.0",
			"platform":     runtime.GOOS,
			"arch":         runtime.GOARCH,
		}, nil
	}
	return a.devTools.GetSystemInfo(), nil
}

func (a *App) TrackPerformance(metrics map[string]any) error {
	logger.Debug("Performance metrics received", zap.Any("metrics", metrics))
	// 可以在这里添加性能数据存储或分析逻辑
	return nil
}

func (a *App) TrackError(errorInfo map[string]any) error {
	logger.Error("Frontend error tracked", zap.Any("error", errorInfo))
	// 可以在这里添加错误日志存储或分析逻辑
	return nil
}
