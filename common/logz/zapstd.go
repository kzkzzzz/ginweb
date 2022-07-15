package logz

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	levelMap = map[string]zapcore.Level{
		"DEBUG": zapcore.DebugLevel,
		"INFO":  zapcore.InfoLevel,
		"WARN":  zapcore.WarnLevel,
		"ERROR": zapcore.ErrorLevel,
		"FATAL": zapcore.FatalLevel,
	}
	// DefaultConfig 默认配置
	DefaultConfig = &Config{
		Color: true,
		Level: "DEBUG",
	}
	zapStdLog = &zapStd{}
)

type (
	zapStd struct {
		log *zap.SugaredLogger
	}

	Config struct {
		Color bool
		Level string
	}
)

func NewLogger(conf *Config, option ...zap.Option) *zapStd {
	if conf == nil {
		conf = DefaultConfig
	}
	var level zapcore.Level
	if v, ok := levelMap[conf.Level]; ok {
		level = v
	} else {
		level = zapcore.DebugLevel
	}

	cfg := zap.NewProductionEncoderConfig()
	if conf.Color {
		// 颜色
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	cfg.ConsoleSeparator = "  |  "

	// 指定日志时间格式
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	//cfg.EncodeCaller = zapcore.FullCallerEncoder

	// 使用控制台输出
	encoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
	l := zap.New(zapcore.NewTee(core), option...)

	zapStdLog = &zapStd{
		log: l.Sugar(),
	}
	return zapStdLog
}

func With(args ...interface{}) *zap.SugaredLogger {
	return zapStdLog.log.With(args...)
}

func Debug(args ...interface{}) {
	zapStdLog.log.Debug(args...)
}

func Info(args ...interface{}) {
	zapStdLog.log.Info(args...)
}

func Warn(args ...interface{}) {
	zapStdLog.log.Warn(args...)
}

func Error(args ...interface{}) {
	zapStdLog.log.Error(args...)
}

func Fatal(args ...interface{}) {
	zapStdLog.log.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	zapStdLog.log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	zapStdLog.log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	zapStdLog.log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	zapStdLog.log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	zapStdLog.log.Fatalf(format, args...)
}

func Println(args ...interface{}) {
	zapStdLog.log.Info(args...)
}

func Printf(format string, args ...interface{}) {
	zapStdLog.log.Infof(format, args...)
}

func (z *zapStd) Flush() {
	z.log.Sync()
}
