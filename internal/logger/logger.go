package logger

import "go.uber.org/zap"

func New(env string) *zap.SugaredLogger {
	var logger *zap.Logger
	if env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	return logger.Sugar()
}