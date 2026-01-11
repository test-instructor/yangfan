package hrp

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/test-instructor/yangfan/httprunner/internal/config"
	"go.uber.org/zap"
)

func InitLogger(logLevel string, logJSON bool, logFile bool) {
	// Error Logging with Stacktrace
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// set log timestamp precise to milliseconds with Beijing timezone (UTC+8)
	beijingLoc, _ := time.LoadLocation("Asia/Shanghai")
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(beijingLoc)
	}
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z0700"

	var msg string
	var consoleWriter io.Writer
	consoleWriter = &zapBridgeWriter{}
	if logFile {
		msg = "log via zap with file output"
	} else {
		msg = "log via zap"
	}

	// parse console log level
	consoleLevel := parseLogLevel(logLevel)

	log.Logger = zerolog.New(consoleWriter).With().Caller().Timestamp().Logger().Level(consoleLevel)
	log.Info().Msg(msg)
	if !logFile {
		return
	}

	logFilePath := config.GetConfig().LogFilePath()
	_, err := os.Stat(logFilePath)
	if err != nil {
		log.Error().Err(err).Str("logFilePath", logFilePath).Msg(msg)
	} else {
		log.Info().Str("logFilePath", logFilePath).Msg(msg)
	}
}

func GetResultsPath() string {
	return config.GetConfig().ResultsPath()
}

// parseLogLevel converts string log level to zerolog.Level
func parseLogLevel(logLevel string) zerolog.Level {
	level := strings.ToUpper(logLevel)
	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "PANIC":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// leveledMultiWriter is a custom writer that applies different log levels to different outputs
type leveledMultiWriter struct {
	consoleWriter io.Writer
	consoleLevel  zerolog.Level
	fileWriter    io.Writer
	fileLevel     zerolog.Level
}

func (w *leveledMultiWriter) Write(p []byte) (n int, err error) {
	// Parse the log level from the JSON log entry
	logLevel := extractLogLevel(p)

	var writeErrors []error

	// Write to console if log level meets console threshold
	if logLevel >= w.consoleLevel {
		if _, err := w.consoleWriter.Write(p); err != nil {
			writeErrors = append(writeErrors, err)
		}
	}

	// Write to file if log level meets file threshold (always debug, so always write)
	if logLevel >= w.fileLevel {
		if _, err := w.fileWriter.Write(p); err != nil {
			writeErrors = append(writeErrors, err)
		}
	}

	// Return the length of the original message and any write errors
	if len(writeErrors) > 0 {
		return len(p), writeErrors[0]
	}
	return len(p), nil
}

// extractLogLevel extracts the log level from a JSON log entry
func extractLogLevel(p []byte) zerolog.Level {
	// Simple parsing to extract level from JSON
	logStr := string(p)
	if strings.Contains(logStr, `"level":"debug"`) {
		return zerolog.DebugLevel
	} else if strings.Contains(logStr, `"level":"info"`) {
		return zerolog.InfoLevel
	} else if strings.Contains(logStr, `"level":"warn"`) {
		return zerolog.WarnLevel
	} else if strings.Contains(logStr, `"level":"error"`) {
		return zerolog.ErrorLevel
	} else if strings.Contains(logStr, `"level":"fatal"`) {
		return zerolog.FatalLevel
	} else if strings.Contains(logStr, `"level":"panic"`) {
		return zerolog.PanicLevel
	}
	return zerolog.InfoLevel // default
}

type zapBridgeWriter struct{}

func (w *zapBridgeWriter) Write(p []byte) (n int, err error) {
	lvl := extractLogLevel(p)
	msg := extractMessage(p)
	caller := extractCaller(p)
	logger := zap.L().WithOptions(zap.WithCaller(false))
	formatted := msg
	if caller != "" {
		formatted = caller + " " + msg
	}
	switch lvl {
	case zerolog.DebugLevel:
		logger.Debug(formatted)
	case zerolog.InfoLevel:
		logger.Info(formatted)
	case zerolog.WarnLevel:
		logger.Warn(formatted)
	case zerolog.ErrorLevel:
		logger.Error(formatted)
	case zerolog.FatalLevel:
		logger.Fatal(formatted)
	case zerolog.PanicLevel:
		logger.Panic(formatted)
	default:
		logger.Info(formatted)
	}
	return len(p), nil
}

func extractMessage(p []byte) string {
	s := string(p)
	idx := strings.Index(s, "\"message\":\"")
	if idx == -1 {
		return s
	}
	rest := s[idx+11:]
	end := strings.Index(rest, "\"")
	if end == -1 {
		return s
	}
	return rest[:end]
}

func extractCaller(p []byte) string {
	s := string(p)
	idx := strings.Index(s, "\"caller\":\"")
	if idx == -1 {
		return ""
	}
	rest := s[idx+10:]
	end := strings.Index(rest, "\"")
	if end == -1 {
		return ""
	}
	return rest[:end]
}
