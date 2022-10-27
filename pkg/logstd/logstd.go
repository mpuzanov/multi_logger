package logstd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// Logger ...
type Logger struct {
	logLevel Level
	logger   *log.Logger
}

// New ...
func New(Level string) *Logger {

	var fileLog string
	ex, _ := os.Executable()
	workDir := filepath.Dir(ex) // путь к программе
	fileLog = fileNameWithoutExtension(filepath.Base(ex)) + ".log"
	fileLog = filepath.Join(workDir, fileLog)
	file := &lumberjack.Logger{
		Filename:   fileLog, // Имя файла
		MaxSize:    10,      // Размер в МБ до ротации файла
		MaxBackups: 5,       // Максимальное количество файлов, сохраненных до перезаписи
		MaxAge:     30,      // Максимальное количество дней для хранения файлов
		Compress:   true,    // Следует ли сжимать файлы логов с помощью gzip
	}
	w := io.MultiWriter(os.Stdout, file)

	l := log.New(w, "", log.Ldate|log.Ltime) //log.LstdFlags  |log.Lshortfile

	levelDefault, _ := ParseLevel(Level)

	return &Logger{logLevel: levelDefault, logger: l}
}

// ParseLevel takes a string level and returns log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}
	var l Level = DebugLevel
	return l, fmt.Errorf("not a valid Level: %q", lvl)
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, Logger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(Logger{}).(*Logger); ok {
		return l
	}
	return &Logger{logLevel: DebugLevel, logger: log.Default()}
}

// =============================================
func (l *Logger) Error(args ...interface{}) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Printf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.SetPrefix("[FATAL] ")
	l.logger.Fatal(args...)
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.SetPrefix("[FATAL] ")
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	if l.logLevel >= InfoLevel {
		l.logger.SetPrefix("[INFO] ")
		l.logger.Println(args...)
	}

}
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logLevel <= InfoLevel {
		l.logger.SetPrefix("[INFO] ")
		l.logger.Printf(format, args...)
	}
}

func (l *Logger) Warn(args ...interface{}) {
	if l.logLevel <= WarnLevel {
		l.logger.SetPrefix("[WARN] ")
		l.logger.Println(args...)
	}
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logLevel <= WarnLevel {
		l.logger.SetPrefix("[WARN] ")
		l.logger.Printf(format, args...)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(args...)
	}
}
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Printf(format, args...)
	}
}
