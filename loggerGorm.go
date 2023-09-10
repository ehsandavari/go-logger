package logger

import (
	"context"
	"errors"
	contextplus "github.com/ehsandavari/go-context-plus"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type sGormLogger struct {
	iLogger                   ILogger
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
	LogLevel                  gormLogger.LogLevel
}

var gormLoggerLevelMap = map[string]gormLogger.LogLevel{
	"debug":  gormLogger.Silent,
	"info":   gormLogger.Info,
	"warn":   gormLogger.Warn,
	"error":  gormLogger.Error,
	"dPanic": gormLogger.Error,
	"panic":  gormLogger.Error,
	"fatal":  gormLogger.Error,
}

func (r *sGormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := r
	newLogger.LogLevel = level
	return newLogger
}

func (r *sGormLogger) Info(ctx context.Context, str string, args ...any) {
	if r.LogLevel < gormLogger.Info {
		return
	}
	r.iLogger.WithAny("args", args).Debug(contextplus.NewContext(ctx), str)
}

func (r *sGormLogger) Warn(ctx context.Context, str string, args ...any) {
	if r.LogLevel < gormLogger.Warn {
		return
	}
	r.iLogger.WithAny("args", args).Warn(contextplus.NewContext(ctx), str)
}

func (r *sGormLogger) Error(ctx context.Context, str string, args ...any) {
	if r.LogLevel < gormLogger.Error {
		return
	}
	r.iLogger.WithAny("args", args).Error(contextplus.NewContext(ctx), str)
}

func (r *sGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if r.LogLevel <= gormLogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && r.LogLevel >= gormLogger.Error && (!r.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		r.iLogger.WithError(err).WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Error(contextplus.NewContext(ctx), "trace")
	case r.SlowThreshold != 0 && elapsed > r.SlowThreshold && r.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		r.iLogger.WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Warn(contextplus.NewContext(ctx), "trace")
	case r.LogLevel >= gormLogger.Info:
		sql, rows := fc()
		r.iLogger.WithDuration("elapsed", elapsed).WithInt64("rows", rows).WithString("sql", sql).Debug(contextplus.NewContext(ctx), "trace")
	}
}

func (r *sGormLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if r.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}
