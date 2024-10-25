package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(development bool) *zap.SugaredLogger {
	var err error
	var logger *zap.Logger
	if development {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = config.Build()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	slogger := logger.Sugar()
	return slogger
}
