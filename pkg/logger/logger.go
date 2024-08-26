package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger(profile string) {
	SetProfileLog(profile)
}

func SetProfileLog(profile string) {
	var level slog.Level

	switch profile {
	case "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	Logger = slog.New(handler)
}

func Debug(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Debug(msg, args...)
	}
}

func Info(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Info(msg, args...)
	}
}

func Warn(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Warn(msg, args...)
	}
}

func Error(msg string, args ...interface{}) {
	if Logger != nil {
		Logger.Error(msg, args...)
	}
}
