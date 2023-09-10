package logger

import (
	"github.com/ehsandavari/go-context-plus"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
)

//go:generate mockgen -destination=./mocks/logger.go -package=mocks github.com/ehsandavari/go-logger ILogger

type ILogger interface {
	Debug(ctx *contextplus.Context, message string)
	Info(ctx *contextplus.Context, message string)
	Warn(ctx *contextplus.Context, message string)
	Error(ctx *contextplus.Context, message string)
	DPanic(ctx *contextplus.Context, message string)
	Panic(ctx *contextplus.Context, message string)
	Fatal(ctx *contextplus.Context, message string)
	IField
	GormLogger() gormLogger.Interface
	Sync() error
}

type sLogger struct {
	sConfig    *sConfig
	zapLogger  *zap.Logger
	fields     []zap.Field
	cores      []zapcore.Core
	gormLogger gormLogger.Interface
}

func NewLogger(isDevelopment bool, disableStacktrace bool, disableStdout bool, level string, serviceId int, serviceName string, serviceNamespace string, serviceInstanceId string, serviceVersion string, serviceMode string, serviceCommitId string, options ...Option) ILogger {
	logger := &sLogger{
		sConfig: &sConfig{
			isDevelopment:     isDevelopment,
			disableStacktrace: disableStacktrace,
			disableStdout:     disableStdout,
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

func (r *sLogger) getLevel() zapcore.Level {
	level, exist := loggerLevelMap[r.sConfig.level]
	if !exist {
		log.Fatalln("log level is not valid")
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

	loggerConfig.Level = zap.NewAtomicLevelAt(r.getLevel())
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
	if !r.sConfig.disableStdout {
		r.cores = append(r.cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(r.config().EncoderConfig),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(r.getLevel()),
		))
	}
	r.zapLogger = zap.New(
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
	r.zapLogger = r.zapLogger.Named(name)
}

func (r *sLogger) setRequestId(ctx *contextplus.Context) *sLogger {
	requestId := ctx.RequestId()
	if len(requestId) != 0 {
		r.WithString("requestId", requestId)
	}
	return r
}

func (r *sLogger) setTraceId(ctx *contextplus.Context) *sLogger {
	traceId := ctx.TraceId()
	if len(traceId) != 0 {
		r.WithString("traceId", traceId)
	}
	return r
}

func (r *sLogger) setUser(ctx *contextplus.Context) *sLogger {
	userId := ctx.User.Id()
	if userId != uuid.Nil {
		r.WithString("userId", userId.String())
	}
	userPhoneNumber := ctx.User.PhoneNumber()
	if len(userPhoneNumber) != 0 {
		r.WithString("userPhoneNumber", userPhoneNumber)
	}
	return r
}

func (r *sLogger) logger(ctx *contextplus.Context) *sLogger {
	return r.setRequestId(ctx).setTraceId(ctx).setUser(ctx)
}

func (r *sLogger) Debug(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Debug(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Info(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Info(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Warn(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Warn(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Error(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Error(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) DPanic(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).DPanic(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Panic(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Panic(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) Fatal(ctx *contextplus.Context, message string) {
	r.logger(ctx).zapLogger.With(zap.Namespace("[Details]")).Fatal(message, r.fields...)
	r.fields = nil
}

func (r *sLogger) GormLogger() gormLogger.Interface {
	return r.gormLogger
}

func (r *sLogger) Sync() error {
	return r.zapLogger.Sync()
}
