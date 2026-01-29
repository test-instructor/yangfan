package devtools

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
)

// DevTools 开发调试工具
type DevTools struct {
	logger *zap.Logger
}

// New 创建开发工具实例
func New(logger *zap.Logger) *DevTools {
	return &DevTools{
		logger: logger,
	}
}

// GetSystemInfo 获取系统信息
func (d *DevTools) GetSystemInfo() map[string]any {
	return map[string]any{
		"appName":      "扬帆自动化测试平台-UI自动化节点",
		"version":      "1.0.0",
		"buildTime":    "2024-01-29",
		"goVersion":    runtime.Version(),
		"platform":     runtime.GOOS,
		"arch":         runtime.GOARCH,
		"numCPU":       runtime.NumCPU(),
		"numGoroutine": runtime.NumGoroutine(),
		"memory":       d.getMemoryInfo(),
	}
}

// getMemoryInfo 获取内存信息
func (d *DevTools) getMemoryInfo() map[string]any {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return map[string]any{
		"alloc":         m.Alloc / 1024 / 1024,      // MB
		"totalAlloc":    m.TotalAlloc / 1024 / 1024, // MB
		"sys":           m.Sys / 1024 / 1024,        // MB
		"numGC":         m.NumGC,
		"lastGC":        time.Unix(0, int64(m.LastGC)).Format("2006-01-02 15:04:05"),
		"gcCPUFraction": m.GCCPUFraction,
	}
}

// GetRuntimeInfo 获取运行时信息
func (d *DevTools) GetRuntimeInfo() map[string]any {
	return map[string]any{
		"goroutines": runtime.NumGoroutine(),
		"threads":    runtime.GOMAXPROCS(0),
		"cpuCount":   runtime.NumCPU(),
		"cgocalls":   runtime.NumCgoCall(),
		"memory":     d.getMemoryInfo(),
		"gcStats":    d.getGCStats(),
	}
}

// getGCStats 获取GC统计
func (d *DevTools) getGCStats() map[string]any {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return map[string]any{
		"numGC":         m.NumGC,
		"pauseTotalNs":  m.PauseTotalNs,
		"pauseNs":       m.PauseNs[(m.NumGC+255)%256], // 最近一次的GC暂停时间
		"gcCPUFraction": m.GCCPUFraction,
	}
}

// TrackPerformance 跟踪性能指标
func (d *DevTools) TrackPerformance(metrics map[string]any) error {
	metricsJSON, _ := json.Marshal(metrics)
	d.logger.Debug("Performance metrics received", zap.String("metrics", string(metricsJSON)))

	// 可以在这里添加性能数据存储或分析逻辑
	return nil
}

// TrackError 跟踪前端错误
func (d *DevTools) TrackError(errorInfo map[string]any) error {
	errorJSON, _ := json.Marshal(errorInfo)
	d.logger.Error("Frontend error tracked", zap.String("error", string(errorJSON)))

	// 可以在这里添加错误日志存储或分析逻辑
	return nil
}

// GetConfigDiff 获取配置差异
func (d *DevTools) GetConfigDiff(current, previous map[string]any) map[string]any {
	diff := make(map[string]any)

	for key, currentValue := range current {
		if prevValue, exists := previous[key]; exists {
			if fmt.Sprintf("%v", currentValue) != fmt.Sprintf("%v", prevValue) {
				diff[key] = map[string]any{
					"old": prevValue,
					"new": currentValue,
				}
			}
		} else {
			diff[key] = map[string]any{
				"old": nil,
				"new": currentValue,
			}
		}
	}

	for key, prevValue := range previous {
		if _, exists := current[key]; !exists {
			diff[key] = map[string]any{
				"old": prevValue,
				"new": nil,
			}
		}
	}

	return diff
}

// ValidateConfig 验证配置
func (d *DevTools) ValidateConfig(config map[string]any) []string {
	var errors []string

	// 基础验证逻辑
	if baseURL, ok := config["baseURL"].(string); ok && baseURL != "" {
		if len(baseURL) < 10 || (baseURL[:7] != "http://" && baseURL[:8] != "https://") {
			errors = append(errors, "baseURL 格式不正确")
		}
	}

	if logLevel, ok := config["logLevel"].(string); ok {
		validLevels := []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}
		valid := false
		for _, level := range validLevels {
			if logLevel == level {
				valid = true
				break
			}
		}
		if !valid {
			errors = append(errors, fmt.Sprintf("logLevel 无效: %s", logLevel))
		}
	}

	return errors
}

// Benchmark 性能基准测试
func (d *DevTools) Benchmark(name string, fn func(), iterations int) map[string]any {
	if iterations <= 0 {
		iterations = 1000
	}

	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	duration := time.Since(start)

	result := map[string]any{
		"name":       name,
		"iterations": iterations,
		"total":      duration.String(),
		"average":    fmt.Sprintf("%.4fms", float64(duration.Nanoseconds())/float64(iterations)/1000000),
		"opsPerSec":  fmt.Sprintf("%.0f", float64(iterations)/duration.Seconds()),
	}

	d.logger.Info("Benchmark completed",
		zap.String("name", name),
		zap.Int("iterations", iterations),
		zap.Duration("duration", duration))

	return result
}
