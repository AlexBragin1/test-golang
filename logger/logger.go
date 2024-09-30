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

//var xl *xlogger.XLogger
//
//func initX() {
//	var err error
//
//	xl, err = xlogger.NewXLogger(
//		config.C.AppName,
//		config.C.LogstashHost,
//		config.C.LogstashPort,
//		config.C.LogstashProtocol,
//		5)
//	if err != nil {
//		panic(err)
//	}
//}

func init() {
	initZap()
	//initX()
}

func Say(msg string, args ...any) {
	sugar.Infow(msg, args...)
}

func Info(msg string, args ...any) {
	sugar.Infow(msg, args...)
	//if err := xl.Info(msg, args...); err != nil {
	//	log.Error("logger error " + err.Error())
	//}
}

func Debug(msg string, args ...any) {
	sugar.Debugw(msg, args...)
	//if err := xl.Debug(msg, args...); err != nil {
	//	log.Error("logger error " + err.Error())
	//}
}

func Warn(msg string, args ...any) {
	sugar.Warnw(msg, args...)
	//if err := xl.Warn(msg, args...); err != nil {
	//	log.Error("logger error " + err.Error())
	//}
}

func Error(msg string, args ...any) {
	sugar.Errorw(msg, args...)
	//if err := xl.Error(msg, args...); err != nil {
	//	log.Error("logger error " + err.Error())
	//}
}
