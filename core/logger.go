package core

import (
	"GO1/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger() {
	// 日志轮转器（写入文件）
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   global.Config.Logger.FilenameOther,
		MaxSize:    global.Config.Logger.MaxSize,
		MaxBackups: global.Config.Logger.MaxBackups,
		MaxAge:     global.Config.Logger.MaxAge,
		Compress:   global.Config.Logger.Compress,
	})

	// 公共编码配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime:    zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// 控制台编码器（带颜色）
	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// 文件编码器（JSON，无颜色）
	fileEncoderConfig := encoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	global.Logger = logger.Sugar()
}

func AddTrace(traceID, requestID string) *zap.SugaredLogger {
	return global.Logger.With("trace_id", traceID, "request_id", requestID)
}
