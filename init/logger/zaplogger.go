package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var logger *zap.Logger

func InitZapLogger() {
	logConfiguration := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ = logConfiguration.Build()
}

func Info(message string, tags ...zap.Field) {
	logger.Info(message, tags...)
	logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Error(message, tags...)
	logger.Sync()
}

func Fatal(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Fatal(message, tags...)
	logger.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("log_output")))

	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("log_level"))) {
	case "info":
		return zapcore.InfoLevel
	case "fatal":
		return zapcore.FatalLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
