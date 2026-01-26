package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Log *zap.Logger
	mu  sync.Mutex
	// Current config state to detect changes
	currentLevel     string
	currentRetention int
	currentPrefix    string
)

type Config struct {
	Level     string
	Prefix    string
	Retention int
}

// Setup initializes the global logger
func Setup(cfg Config) {
	mu.Lock()
	defer mu.Unlock()

	// Update current state
	currentLevel = cfg.Level
	currentRetention = cfg.Retention
	currentPrefix = cfg.Prefix

	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// Define encoders
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Console encoder (with color)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// Determine min level based on config
	var minLevel zapcore.Level
	switch strings.ToLower(cfg.Level) {
	case "debug":
		minLevel = zapcore.DebugLevel
	case "info":
		minLevel = zapcore.InfoLevel
	case "warn":
		minLevel = zapcore.WarnLevel
	case "error":
		minLevel = zapcore.ErrorLevel
	case "fatal":
		minLevel = zapcore.FatalLevel
	default:
		minLevel = zapcore.InfoLevel
	}

	// Level enablers - strict matching for separate files, but respecting minLevel
	// Actually, usually you want info.log to contain info AND above, or just info?
	// The requirement "separate recording" usually means separate files.
	// But commonly debug.log has everything, info.log has info+, etc.
	// Let's stick to the previous pattern: specific files for specific levels,
	// BUT we should respect the global configured level.
	// If configured level is INFO, Debug logs should NOT be written anywhere.

	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && lvl >= minLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && lvl >= minLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && lvl >= minLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && lvl >= minLevel
	})

	// Console enabler
	consoleLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= minLevel
	})

	// Log writers with date in filename: logs/2024-01-27-debug.log
	dateStr := time.Now().Format("2006-01-02")

	debugWriter := getLogWriter(filepath.Join(logDir, fmt.Sprintf("%s-debug.log", dateStr)), cfg.Retention)
	infoWriter := getLogWriter(filepath.Join(logDir, fmt.Sprintf("%s-info.log", dateStr)), cfg.Retention)
	warnWriter := getLogWriter(filepath.Join(logDir, fmt.Sprintf("%s-warn.log", dateStr)), cfg.Retention)
	errorWriter := getLogWriter(filepath.Join(logDir, fmt.Sprintf("%s-error.log", dateStr)), cfg.Retention)

	// Cores
	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), consoleLevel),
		zapcore.NewCore(fileEncoder, debugWriter, debugLevel),
		zapcore.NewCore(fileEncoder, infoWriter, infoLevel),
		zapcore.NewCore(fileEncoder, warnWriter, warnLevel),
		zapcore.NewCore(fileEncoder, errorWriter, errorLevel),
	}

	core := zapcore.NewTee(cores...)

	// Create logger with initial fields (prefix)
	options := []zap.Option{zap.AddCaller()}
	if cfg.Prefix != "" {
		options = append(options, zap.Fields(zap.String("app", cfg.Prefix)))
	}

	Log = zap.New(core, options...)

	// Replace global logger
	zap.ReplaceGlobals(Log)
}

func getLogWriter(filename string, retentionDays int) zapcore.WriteSyncer {
	if retentionDays <= 0 {
		retentionDays = 30
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     retentionDays, // days
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Helper functions for easy access
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}
