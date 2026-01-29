package config

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Environment 环境类型
type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvTesting     Environment = "testing"
	EnvProduction  Environment = "production"
)

type Config struct {
	BaseURL      string      `json:"baseURL"`
	Token        string      `json:"token"`
	ExpiresAt    int64       `json:"expiresAt"`
	LogLevel     string      `json:"logLevel"`     // debug, info, warn, error, dpanic, panic, fatal
	LogPrefix    string      `json:"logPrefix"`    // e.g. [ https://github.com/test-instructor/yangfan/ui ]
	LogRetention int         `json:"logRetention"` // days
	Environment  Environment `json:"environment"`
	DebugMode    bool        `json:"debugMode"`
	Theme        string      `json:"theme"`
	Language     string      `json:"language"`
	AutoLogin    bool        `json:"autoLogin"`
	RememberMe   bool        `json:"rememberMe"`
}

type Store struct {
	mu   sync.RWMutex
	path string
	cfg  Config
}

func New(appName string) (*Store, error) {
	if strings.TrimSpace(appName) == "" {
		return nil, errors.New("appName 不能为空")
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	dir = filepath.Join(dir, appName)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	s := &Store{
		path: filepath.Join(dir, "config.json"),
	}
	_ = s.load()
	return s, nil
}

func (s *Store) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, err := os.ReadFile(s.path)
	if err != nil {
		// Default values if file not found
		s.cfg = Config{
			LogLevel:     "info",
			LogPrefix:    "[ https://github.com/test-instructor/yangfan/ui ]",
			LogRetention: 30,
			Environment:  EnvDevelopment,
			DebugMode:    false,
			Theme:        "light",
			Language:     "zh-CN",
			AutoLogin:    false,
			RememberMe:   false,
		}
		return nil // treat as no error, use defaults
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return err
	}
	cfg.BaseURL = normalizeBaseURL(cfg.BaseURL)
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}
	if cfg.LogPrefix == "" {
		cfg.LogPrefix = "[ https://github.com/test-instructor/yangfan/ui ]"
	}
	if cfg.LogRetention <= 0 {
		cfg.LogRetention = 30
	}
	if cfg.Environment == "" {
		cfg.Environment = EnvDevelopment
	}
	if cfg.Theme == "" {
		cfg.Theme = "light"
	}
	if cfg.Language == "" {
		cfg.Language = "zh-CN"
	}
	s.cfg = cfg
	return nil
}

func (s *Store) saveLocked() error {
	b, err := json.MarshalIndent(s.cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, b, 0o600)
}

func (s *Store) Get() Config {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg
}

func (s *Store) BaseURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg.BaseURL
}

func (s *Store) Token() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg.Token
}

func (s *Store) SetBaseURL(baseURL string) error {
	baseURL, err := ValidateBaseURL(baseURL)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cfg.BaseURL != baseURL {
		s.cfg.BaseURL = baseURL
		s.cfg.Token = ""
		s.cfg.ExpiresAt = 0
	}
	return s.saveLocked()
}

func (s *Store) ClearAuth() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.Token = ""
	s.cfg.ExpiresAt = 0
	return s.saveLocked()
}

func (s *Store) SetLogConfig(level string, prefix string, retention int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.LogLevel = level
	s.cfg.LogPrefix = prefix
	s.cfg.LogRetention = retention
	return s.saveLocked()
}

func (s *Store) SetToken(token string, expiresAt int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.Token = token
	s.cfg.ExpiresAt = expiresAt
	return s.saveLocked()
}

// Environment 获取环境
func (s *Store) Environment() Environment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg.Environment
}

// IsDebugMode 是否为调试模式
func (s *Store) IsDebugMode() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cfg.DebugMode
}

// SetEnvironment 设置环境
func (s *Store) SetEnvironment(env Environment) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.Environment = env
	return s.saveLocked()
}

// SetDebugMode 设置调试模式
func (s *Store) SetDebugMode(debug bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.DebugMode = debug
	return s.saveLocked()
}

// SetTheme 设置主题
func (s *Store) SetTheme(theme string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.Theme = theme
	return s.saveLocked()
}

// SetLanguage 设置语言
func (s *Store) SetLanguage(language string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.Language = language
	return s.saveLocked()
}

// SetAutoLogin 设置自动登录
func (s *Store) SetAutoLogin(autoLogin bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.AutoLogin = autoLogin
	return s.saveLocked()
}

// SetRememberMe 设置记住我
func (s *Store) SetRememberMe(remember bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cfg.RememberMe = remember
	return s.saveLocked()
}

func ValidateBaseURL(baseURL string) (string, error) {
	baseURL = normalizeBaseURL(baseURL)
	if baseURL == "" {
		return "", errors.New("baseURL 不能为空")
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" || u.Host == "" {
		return "", errors.New("baseURL 必须包含 scheme 与 host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("baseURL scheme 仅支持 http/https")
	}
	return normalizeBaseURL(u.String()), nil
}

func normalizeBaseURL(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimRight(v, "/")
	return v
}
