package config

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局公开日志记录器
var Logger *zap.SugaredLogger

// 初始化日志记录器
func InitLogger() {
	logWriter := &lumberjack.Logger{
		Filename:   "./运行日志.log",
		MaxSize:    10,   // 日志文件空间限制(MB)
		MaxBackups: 3,    // 保留的最大备份文件数量
		MaxAge:     7,    // 保留日期, 过期后自动删除
		Compress:   true, // 旧的日志文件在滚动时是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05 MST"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(time.FixedZone("UTC+8", 8*3600)).Format("2006-01-02 15:04:05 MST"))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(logWriter),
		zap.InfoLevel,
	)
	logger := zap.New(core)
	Logger = logger.Sugar()
}
