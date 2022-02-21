package logger

import (
	"context"
	"os"
	"strings"

	"github.com/butlerhq/butler/internal/environment"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	EncodingConsole string = "console"
	EncodingJSON    string = "json"
)

var (
	logTimeFormat = "2006-01-02T15:04:05.000+08:00"
	zapLogger     *zap.Logger
	// For mapping config logger to app logger levels
	loggerLevelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	DefaultLoggerConfig = LoggerConfig{
		DisableCaller:     true,
		DisableStacktrace: false,
		Encoding:          EncodingJSON,
		Level:             "debug",
	}
)

// init zap logger
func init() {
	c := zap.NewDevelopmentConfig()
	c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapLogger, _ = c.Build()
}

type LoggerConfig struct {
	DisableCaller     bool   `env:"DISABLE_CALLER"`
	DisableStacktrace bool   `env:"DISABLE_STACK_TRACE"`
	Encoding          string `env:"DISABLE_CALLER"`
	Level             string `env:"LEVEL"`
}

// Update Logger configuration
func UpdateAppLoggerWithConfig(env string, cfg *LoggerConfig) {
	logLevel := getLoggerLevel(cfg)
	logWriter := zapcore.AddSync(os.Stderr)
	var encoderCfg zapcore.EncoderConfig
	if environment.IsProductionEnv(env) {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}

	var encoder zapcore.Encoder
	if cfg.Encoding == EncodingConsole {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	if cfg.DisableCaller {
		encoderCfg.CallerKey = ""
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Return zapcore.Level for a given LoggerConfig
func getLoggerLevel(cfg *LoggerConfig) zapcore.Level {
	strLevel := strings.ToLower(cfg.Level)
	level, exist := loggerLevelMap[strLevel]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Return zapLogger instance
func GetLogger() *zap.Logger {
	return zapLogger
}

func WithGRPCTags(ctx context.Context) *zap.Logger {
	// Add grpc_ctxtags tags metadata until now.
	fields := ctxzap.TagsToFields(ctx)
	return zapLogger.With(fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Error(msg, fields...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Errorf(format, args...)
}

func Errorw(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Errorw(msg, args...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Warn(msg, fields...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Warnf(format, args...)
}

func Warnw(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Warnw(msg, args...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Info(msg, fields...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Infof(format, args...)
}

func Infow(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Infow(msg, args...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Debug(msg, fields...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Debugf(format, args...)
}

func Debugw(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Debugw(msg, args...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Fatal(msg, fields...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Fatalf(format, args...)
}

func Fatalw(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Fatalw(msg, args...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	WithGRPCTags(ctx).Fatal(msg, fields...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Panicf(format, args...)
}

func Panicw(ctx context.Context, msg string, args ...interface{}) {
	WithGRPCTags(ctx).Sugar().Panicw(msg, args...)
}
