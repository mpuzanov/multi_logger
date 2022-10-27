package glogger

import (
	"context"
	"multi_logger/pkg/logging"
	"multi_logger/pkg/logstd"
	"multi_logger/pkg/logzap"
	"strings"
)

var loggerCtxKey struct{}

// Logger represent common interface for logging function
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// BuildLogger создаем выбранный логгер
func BuildLogger(logType, Level string) Logger {
	var logger Logger
	switch strings.ToUpper(logType) {
	case "LOGRUS":
		logger = logging.New(Level, false) // logrus
	case "STD":
		logger = logstd.New(Level) // log
	default:
		logType = "ZAP"
		logger = logzap.New(Level) // zap
	}
	logger.Debugf("Используем logger: %s", logType)
	return logger
}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(loggerCtxKey).(Logger); ok {
		return l
	}
	l := BuildLogger("", "")
	return l
}
