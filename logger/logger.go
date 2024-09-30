package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var sugar *zap.SugaredLogger

func initZap() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	sugar = log.Sugar()
}

func init() {
	initZap()
}

func Say(msg string, args ...any) {
	sugar.Infow(msg, args...)
}

func Info(msg string, args ...any) {
	sugar.Infow(msg, args...)
}

func Debug(msg string, args ...any) {
	sugar.Debugw(msg, args...)
}

func Warn(msg string, args ...any) {
	sugar.Warnw(msg, args...)
}

func Error(msg string, args ...any) {
	sugar.Errorw(msg, args...)
}
