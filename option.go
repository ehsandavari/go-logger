package logger

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"time"
)

type Option interface {
	apply(*sLogger)
}

type optionFunc func(*sLogger)

func (f optionFunc) apply(log *sLogger) {
	f(log)
}

func WithElk(endpoint string, timeoutSecond byte) Option {
	return optionFunc(func(logger *sLogger) {
		conn, err := net.DialTimeout("udp", endpoint, time.Duration(timeoutSecond)*time.Second)
		if err != nil {
			log.Fatalln("connecting to logstash failed", err)
		}

		loggerConfig := logger.config()
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

		logger.cores = append(logger.cores, ecszap.WrapCore(zapcore.NewCore(
			zapcore.NewJSONEncoder(loggerConfig.EncoderConfig),
			zapcore.AddSync(conn),
			zap.NewAtomicLevelAt(logger.getLevel()),
		)))
	})
}

func WithGormLogger(slowThreshold time.Duration, colorful, ignoreRecordNotFoundError, parameterizedQueries bool) Option {
	return optionFunc(func(logger *sLogger) {
		level, exist := gormLoggerLevelMap[logger.sConfig.level]
		if !exist {
			log.Fatalln("logger level is not valid")
		}
		logger.gormLogger = &sGormLogger{
			sLogger:                   logger,
			SlowThreshold:             slowThreshold,
			Colorful:                  colorful,
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError,
			ParameterizedQueries:      parameterizedQueries,
			LogLevel:                  level,
		}
	})
}
