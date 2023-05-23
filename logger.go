package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

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
}

func NewLogger(isDevelopment bool, level string, serviceId int, serviceName string, serviceNamespace string, serviceInstanceId string, serviceVersion string, serviceMode string, serviceCommitId string) ILogger {
	logger := &sLogger{
		sConfig: &sConfig{
			isDevelopment:     isDevelopment,
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
	logger.config(logger.getLoggerLevel())
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

func (r *sLogger) config(logLevel zapcore.Level) {
	loggerConfig := zap.NewProductionConfig()
	if r.sConfig.isDevelopment {
		loggerConfig = zap.NewDevelopmentConfig()
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		loggerConfig.EncoderConfig.ConsoleSeparator = " || "
	}

	loggerConfig.EncoderConfig.NameKey = "[ServiceName]"
	loggerConfig.EncoderConfig.TimeKey = "[Time]"
	loggerConfig.EncoderConfig.LevelKey = "[Level]"
	loggerConfig.EncoderConfig.CallerKey = "[Caller]"
	loggerConfig.EncoderConfig.FunctionKey = "[Function]"
	loggerConfig.EncoderConfig.MessageKey = "[Message]"
	loggerConfig.EncoderConfig.StacktraceKey = "[Stacktrace]"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	loggerConfig.Level = zap.NewAtomicLevelAt(logLevel)
	logger, err := loggerConfig.Build(
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
	if err != nil {
		log.Fatalln("error in build logger : ", err)
	}
	r.sLogger = logger
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
