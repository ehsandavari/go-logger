package logger

import (
	"context"
	contextplus "github.com/ehsandavari/go-context-plus"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type SPostgresLogger struct {
	iLogger ILogger
	config  gormLogger.Config
}

func NewSPostgresLogger(iLogger ILogger, slowThreshold time.Duration, colorful, ignoreRecordNotFoundError, parameterizedQueries bool, logLevel gormLogger.LogLevel) *SPostgresLogger {
	return &SPostgresLogger{
		iLogger: iLogger,
		config: gormLogger.Config{
			SlowThreshold:             slowThreshold,
			Colorful:                  colorful,
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError,
			ParameterizedQueries:      parameterizedQueries,
			LogLevel:                  logLevel,
		},
	}
}

func (r *SPostgresLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := r
	newLogger.config.LogLevel = level
	return newLogger
}

func (r *SPostgresLogger) Info(ctx context.Context, str string, args ...any) {
	if r.config.LogLevel < gormLogger.Info {
		return
	}
	r.iLogger.WithAny("args", args).Debug(contextplus.NewContext(ctx), str)
}

func (r *SPostgresLogger) Warn(ctx context.Context, str string, args ...any) {
	if r.config.LogLevel < gormLogger.Warn {
		return
	}
	r.iLogger.WithAny("args", args).Warn(contextplus.NewContext(ctx), str)
}

func (r *SPostgresLogger) Error(ctx context.Context, str string, args ...any) {
	if r.config.LogLevel < gormLogger.Error {
		return
	}
	r.iLogger.WithAny("args", args).Error(contextplus.NewContext(ctx), str)
}

func (r *SPostgresLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if r.config.LogLevel <= gormLogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && r.config.LogLevel >= gormLogger.Error && (!r.config.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		r.iLogger.WithError(err).WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Error(contextplus.NewContext(ctx), "trace")
	case r.config.SlowThreshold != 0 && elapsed > r.config.SlowThreshold && r.config.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		r.iLogger.WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Warn(contextplus.NewContext(ctx), "trace")
	case r.config.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		r.iLogger.WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Debug(contextplus.NewContext(ctx), "trace")
	}
}

func (r *SPostgresLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if r.config.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}
