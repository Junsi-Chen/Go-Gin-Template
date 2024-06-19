package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
	"unsafe"
)

type ErrorType string

var (
	Logger *zap.Logger
)

func Init(name string) {
	logInit(name)
}

// Lumber 文件切割配置
type Lumber struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M  128
	MaxBackups int    // 日志文件最多保存多少个备份 30
	MaxAge     int    // 文件最多保存多少天 7
	Compress   bool   // 是否压缩
}

// GetDeferLumber 返回默认配置
// @ param：filename文件名
// @ return: lumberjack.Logger
func (l *Lumber) GetDeferLumber(filename string) lumberjack.Logger {
	today := time.Now().Format("2006-01-02")
	filename = fmt.Sprintf("logs/%s/%s", today, filename)
	return lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
}

// init 初始化日志
func logInit(name string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "line",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	var cores []zapcore.Core
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})

	debug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})

	dPanic := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DPanicLevel
	})

	//使用默认日志切割
	var lumber = &Lumber{}
	infoWriter := lumber.GetDeferLumber(name + "_info.log")
	warnWriter := lumber.GetDeferLumber(name + "_warn.log")
	errorWriter := lumber.GetDeferLumber(name + "_error.log")
	debugWriter := lumber.GetDeferLumber(name + "_debug.log")
	panicWriter := lumber.GetDeferLumber(name + "_panic.log")
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&infoWriter)), infoLevel))
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&warnWriter)), warnLevel))
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&errorWriter)), errorLevel))
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&debugWriter)), debug))
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&panicWriter)), dPanic))

	// 最后创建具体的Logger
	core := zapcore.NewTee(cores...)

	caller := zap.AddCaller()
	development := zap.Development()
	Logger = zap.New(core, caller, development, zap.AddStacktrace(zapcore.DPanicLevel))
	//logger.development设置为FALSE时，DPanic就不会panic
	//go.uber.org/zap/logger.go:279
	*(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(Logger)) + uintptr(16))) = false
}
