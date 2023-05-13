package logger

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sLogger struct {
	level       string
	mode        string
	encoder     string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

func NewLogger(level string, mode string, encoder string) ILogger {
	logger := &sLogger{
		level:   level,
		mode:    mode,
		encoder: encoder,
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
	level, exist := loggerLevelMap[r.level]
	if !exist {
		log.Fatalln("logger level is not valid")
	}
	return level
}

func (r *sLogger) config(logLevel zapcore.Level) {
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if r.mode == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else if r.mode == "production" {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		log.Fatalln("logger mode is not valid")
	}

	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	var encoder zapcore.Encoder
	if r.encoder == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else if r.encoder == "json" {
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		log.Fatalln("logger encoder is not valid")
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	r.logger = zapLogger
	r.sugarLogger = zapLogger.Sugar()
}

// Named add logger microservice name
func (r *sLogger) Named(name string) {
	r.logger = r.logger.Named(name)
	r.sugarLogger = r.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (r *sLogger) Debug(args ...any) {
	r.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (r *sLogger) Debugf(template string, args ...any) {
	r.sugarLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message
func (r *sLogger) Info(args ...any) {
	r.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (r *sLogger) Infof(template string, args ...any) {
	r.sugarLogger.Infof(template, args...)
}

// Printf uses fmt.Sprintf to log a templated message
func (r *sLogger) Printf(template string, args ...any) {
	r.sugarLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (r *sLogger) Warn(args ...any) {
	r.sugarLogger.Warn(args...)
}

// WarnErrMsg log error message with warn level.
func (r *sLogger) WarnErrMsg(msg string, err error) {
	r.logger.Warn(msg, zap.String("error", err.Error()))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (r *sLogger) Warnf(template string, args ...any) {
	r.sugarLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (r *sLogger) Error(args ...any) {
	r.sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (r *sLogger) Errorf(template string, args ...any) {
	r.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (r *sLogger) Err(msg string, err error) {
	r.logger.Error(msg, zap.Error(err))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the logger then panics. (See DPanicLevel for details.)
func (r *sLogger) DPanic(args ...any) {
	r.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the logger then panics. (See DPanicLevel for details.)
func (r *sLogger) DPanicf(template string, args ...any) {
	r.sugarLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (r *sLogger) Panic(args ...any) {
	r.sugarLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (r *sLogger) Panicf(template string, args ...any) {
	r.sugarLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (r *sLogger) Fatal(args ...any) {
	r.sugarLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (r *sLogger) Fatalf(template string, args ...any) {
	r.sugarLogger.Fatalf(template, args...)
}

func (r *sLogger) Sync() error {
	return r.sugarLogger.Sync()
}

func (r *sLogger) HttpMiddlewareAccessLogger(method, uri string, status int, size int64, time time.Duration) {
	r.logger.Info(
		Http,
		zap.String(Method, method),
		zap.String(Uri, uri),
		zap.Int(Status, status),
		zap.Int64(Size, size),
		zap.Duration(Time, time),
	)
}

func (r *sLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	r.logger.Info(
		Grpc,
		zap.String(Method, method),
		zap.Duration(Time, time),
		zap.Any(MetaData, metaData),
		zap.Any(Error, err),
	)
}

func (r *sLogger) GrpcMiddlewareAccessLoggerErr(method string, time time.Duration, metaData map[string][]string, err error) {
	r.logger.Error(
		Grpc,
		zap.String(Method, method),
		zap.Duration(Time, time),
		zap.Any(MetaData, metaData),
		zap.Any(Error, err),
	)
}

func (r *sLogger) GrpcClientInterceptorLogger(method string, req, reply any, time time.Duration, metaData map[string][]string, err error) {
	r.logger.Info(
		Grpc,
		zap.String(Method, method),
		zap.Any(Request, req),
		zap.Any(Reply, reply),
		zap.Duration(Time, time),
		zap.Any(MetaData, metaData),
		zap.Any(Error, err),
	)
}

func (r *sLogger) GrpcClientInterceptorLoggerErr(method string, req, reply any, time time.Duration, metaData map[string][]string, err error) {
	r.logger.Error(
		Grpc,
		zap.String(Method, method),
		zap.Any(Request, req),
		zap.Any(Reply, reply),
		zap.Duration(Time, time),
		zap.Any(MetaData, metaData),
		zap.Any(Error, err),
	)
}

func (r *sLogger) KafkaProcessMessage(topic string, partition int, message []byte, workerID int, offset int64, time time.Time) {
	r.logger.Debug(
		"(Processing Kafka message)",
		zap.String(Topic, topic),
		zap.Int(Partition, partition),
		zap.Int(MessageSize, len(message)),
		zap.Int(WorkerID, workerID),
		zap.Int64(Offset, offset),
		zap.Time(Time, time),
	)
}

func (r *sLogger) KafkaLogCommittedMessage(topic string, partition int, offset int64) {
	r.logger.Debug(
		"(Committed Kafka message)",
		zap.String(Topic, topic),
		zap.Int(Partition, partition),
		zap.Int64(Offset, offset),
	)
}

func (r *sLogger) KafkaProcessMessageWithHeaders(topic string, partition int, message []byte, workerID int, offset int64, time time.Time, headers map[string]any) {
	r.logger.Debug(
		"(Processing Kafka message)",
		zap.String(Topic, topic),
		zap.Int(Partition, partition),
		zap.Int(MessageSize, len(message)),
		zap.Int(WorkerID, workerID),
		zap.Int64(Offset, offset),
		zap.Time(Time, time),
		zap.Any(KafkaHeaders, headers),
	)
}
