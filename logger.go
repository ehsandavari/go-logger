package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

//go:generate mockgen -destination=./mocks/logger.go -package=mocks github.com/ehsandavari/go-logger ILogger

type ILogger interface {
	Debug(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warn(ctx context.Context, message string)
	Error(ctx context.Context, message string)
	DPanic(ctx context.Context, message string)
	Panic(ctx context.Context, message string)
	Fatal(ctx context.Context, message string)
	IField
	Sync() error
}

type sLogger struct {
	sConfig *sConfig
	sLogger *zap.Logger
	fields  []zap.Field
	cores   []zapcore.Core
}

func NewLogger(isDevelopment bool, disableStacktrace bool, level string, serviceId int, serviceName string, serviceNamespace string, serviceInstanceId string, serviceVersion string, serviceMode string, serviceCommitId string, options ...Option) ILogger {
	logger := &sLogger{
		sConfig: &sConfig{
			isDevelopment:     isDevelopment,
			disableStacktrace: disableStacktrace,
			level:             level,
			serviceId:         serviceId,
			serviceName:       serviceName,
			serviceNamespace:  serviceNamespace,
			serviceInstanceId: serviceInstanceId,
			serviceVersion:    serviceVersion,
			serviceMode:       serviceMode,
			serviceCommitId:   serviceCommitId,
		},
	}
	for _, option := range options {
		option.apply(logger)
	}
	logger.init()
	return logger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dPanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (r *sLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[r.sConfig.level]
	if !exist {
		log.Fatalln("logger level is not valid")
	}
	return level
}

func (r *sLogger) config() zap.Config {
	loggerConfig := zap.NewProductionConfig()
	if r.sConfig.isDevelopment {
		loggerConfig = zap.NewDevelopmentConfig()
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		loggerConfig.EncoderConfig.ConsoleSeparator = " || "
	}

	loggerConfig.Level = zap.NewAtomicLevelAt(r.getLoggerLevel())
	loggerConfig.DisableStacktrace = r.sConfig.disableStacktrace

	loggerConfig.EncoderConfig.NameKey = "[ServiceName]"
	loggerConfig.EncoderConfig.TimeKey = "[Time]"
	loggerConfig.EncoderConfig.LevelKey = "[Level]"
	loggerConfig.EncoderConfig.CallerKey = "[Caller]"
	loggerConfig.EncoderConfig.FunctionKey = "[Function]"
	loggerConfig.EncoderConfig.MessageKey = "[Message]"
	loggerConfig.EncoderConfig.StacktraceKey = "[Stacktrace]"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return loggerConfig
}

func (r *sLogger) init() {
	if r.sConfig.UseStdout {
		r.cores = append(r.cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(r.config().EncoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(r.getLoggerLevel()),
		))
	}
	r.sLogger = zap.New(
		zapcore.NewTee(r.cores...),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zap.Int("[ServiceId]", r.sConfig.serviceId),
			zap.String("[ServiceNamespace]", r.sConfig.serviceNamespace),
			zap.String("[ServiceInstanceId]", r.sConfig.serviceInstanceId),
			zap.String("[ServiceVersion]", r.sConfig.serviceVersion),
			zap.String("[ServiceMode]", r.sConfig.serviceMode),
			zap.String("[ServiceCommitId]", r.sConfig.serviceCommitId),
		),
	)
	r.named(r.sConfig.serviceName)
}

func (r *sLogger) named(name string) {
	r.sLogger = r.sLogger.Named(name)
}

func (r *sLogger) setRequestId(ctx context.Context) *sLogger {
	value, ok := ctx.Value(requestId).(string)
	if ok {
		r.WithString(requestId, value)
	}
	return r
}

func (r *sLogger) setTraceId(ctx context.Context) *sLogger {
	value, ok := ctx.Value(traceId).(string)
	if ok {
		r.WithString(traceId, value)
	}
	return r
}

func (r *sLogger) logger(ctx context.Context) *sLogger {
	return r.setRequestId(ctx).setTraceId(ctx)
}

func (r *sLogger) Debug(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Debug(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Info(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Info(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Warn(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Warn(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Error(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Error(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) DPanic(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).DPanic(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Panic(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Panic(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Fatal(ctx context.Context, message string) {
	r.logger(ctx).sLogger.With(zap.Namespace("[Meta]")).Fatal(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Sync() error {
	return r.sLogger.Sync()
}
