package logs

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func init() {
	initLogger()
	sugarLogger.Sync()
}

func initLogger() {
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.WriteSyncer(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func ConsoleDebug(format string, v ...interface{}) {
	sugarLogger.Debugf(format, v...)
}

func ConsoleInfo(format string, v ...interface{}) {
	sugarLogger.Infof(format, v...)
}

func ConsoleError(format string, v ...interface{}) {
	sugarLogger.Errorf(format, v...)
}

func ConsoleWarn(format string, v ...interface{}) {
	sugarLogger.Warnf(format, v...)
}

func ConsoleFatal(format string, v ...interface{}) {
	sugarLogger.Fatalf(format, v...)
}

func ConsolePanic(format string, v ...interface{}) {
	sugarLogger.Panicf(format, v...)
}
