package logger

import "time"

type (
	ILogger interface {
		Debug(args ...any)
		Debugf(template string, args ...any)
		Info(args ...any)
		Infof(template string, args ...any)
		Warn(args ...any)
		Warnf(template string, args ...any)
		WarnErrMsg(msg string, err error)
		Error(args ...any)
		Errorf(template string, args ...any)
		Err(msg string, err error)
		DPanic(args ...any)
		DPanicf(template string, args ...any)
		Fatal(args ...any)
		Fatalf(template string, args ...any)
		Sync() error
		Printf(template string, args ...any)
		Named(name string)
		HttpMiddlewareAccessLogger(method string, uri string, status int, size int64, time time.Duration)
		GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error)
		GrpcMiddlewareAccessLoggerErr(method string, time time.Duration, metaData map[string][]string, err error)
		GrpcClientInterceptorLogger(method string, req any, reply any, time time.Duration, metaData map[string][]string, err error)
		GrpcClientInterceptorLoggerErr(method string, req, reply any, time time.Duration, metaData map[string][]string, err error)
		KafkaProcessMessage(topic string, partition int, message []byte, workerID int, offset int64, time time.Time)
		KafkaLogCommittedMessage(topic string, partition int, offset int64)
		KafkaProcessMessageWithHeaders(topic string, partition int, message []byte, workerID int, offset int64, time time.Time, headers map[string]any)
	}
)
