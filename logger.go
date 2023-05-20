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
	value, ok := ctx.Value(RequestID).(string)
	if ok {
		r.WithString(RequestID, value)
	}
	return r
}

func (r *sLogger) setTraceId(ctx context.Context) *sLogger {
	value, ok := ctx.Value(TraceID).(string)
	if ok {
		r.WithString(TraceID, value)
	}
	return r
}

func (r *sLogger) logger(ctx context.Context) *sLogger {
	return r.setRequestId(ctx).setTraceId(ctx)
}

func (r *sLogger) Debug(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Debug(message)
}

func (r *sLogger) Info(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Info(message)
}

func (r *sLogger) Warn(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Warn(message)
}

func (r *sLogger) Error(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Error(message)
}

func (r *sLogger) DPanic(ctx context.Context, message string) {
	r.logger(ctx).sLogger.DPanic(message)
}

func (r *sLogger) Panic(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Panic(message)
}

func (r *sLogger) Fatal(ctx context.Context, message string) {
	r.logger(ctx).sLogger.Fatal(message)
}

func (r *sLogger) Sync() error {
	return r.sLogger.Sync()
}
