package logger

import (
	"GO1/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func InitLogger() {
	writeSyncer := getLogWriter()
	writeSyncerError := getLogWriterError()
	encoder := getEncoder()
	core1 := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	core2 := zapcore.NewCore(encoder, writeSyncerError, zap.ErrorLevel)

	core := zapcore.NewTee(core1, core2)
	logger := zap.New(core, zap.AddCaller())
	global.Logger = logger.Sugar()
}

func getEncoder() zapcore.Encoder { // 编码形式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeCaller:  zapcore.FullCallerEncoder,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
	}
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer { // 日志分割的配置信息
	lumberJackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   global.Config.Logger.Filename,
		MaxSize:    global.Config.Logger.MaxSize,
		MaxBackups: global.Config.Logger.MaxBackups,
		MaxAge:     global.Config.Logger.MaxAge,
		Compress:   global.Config.Logger.Compress,
	})
	ws := io.MultiWriter(lumberJackLogger, os.Stdout)
	return zapcore.AddSync(ws)
}

func getLogWriterError() zapcore.WriteSyncer {
	lumberJackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   global.Config.Logger.FilenameError,
		MaxSize:    global.Config.Logger.MaxSize,
		MaxBackups: global.Config.Logger.MaxBackups,
		MaxAge:     global.Config.Logger.MaxAge,
		Compress:   global.Config.Logger.Compress,
	})
	ws := io.MultiWriter(lumberJackLogger)
	return zapcore.AddSync(ws)
}
