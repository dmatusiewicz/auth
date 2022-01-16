package server

import "go.uber.org/zap"

type ilogger interface {
	Debug(string, ...zap.Field)
	Panic(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Info(string, ...zap.Field)
}
