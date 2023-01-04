package logger

import (
	"core/config"
	"io"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(*options)

type options struct {
	Sync []zapcore.WriteSyncer
}

func WithAddSync(w io.Writer) Option {
	return func(o *options) {
		o.Sync = append(o.Sync, zapcore.AddSync(w))
	}
}

func defaultOptions(logs *config.Logs) (*options, zapcore.EncoderConfig) {
	lumberjack := &lumberjack.Logger{
		Filename:   logs.Filename, // 日志文件路径
		MaxSize:    100,           // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,             // 日志文件最多保存多少个备份
		MaxAge:     28,            // 文件最多保存多少天
		Compress:   true,          // 是否压缩
	}

	opt := &options{}
	opt.Sync = append(opt.Sync, zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjack))

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	}

	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	// 自定义文件：行号输出项
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,   // 自定义时间格式
		EncodeLevel:    customLevelEncoder,  // 小写编码器
		EncodeCaller:   customCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// level大写染色编码器
	//encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return opt, encoderConf
}

// InitLogger TODO 动态设置级别
func InitLogger(logs *config.Logs, opt ...Option) error {
	def, encoderConf := defaultOptions(logs)
	for _, o := range opt {
		o(def)
	}

	logger, err := zap.NewProductionConfig().Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConf),
			zapcore.NewMultiWriteSyncer(def.Sync...),
			zap.NewAtomicLevelAt(LogLevel(logs.Level)),
		)
	}))
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	return nil
}

func LogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
