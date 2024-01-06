package log

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/j23063519/clean_architecture/pkg/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// GormLogger accomplish gormlogger.Interface
type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

// NewGormLogger for external use, new GormLogger
func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,                 // use global logger.Logger
		SlowThreshold: 200 * time.Millisecond, // Slow query log threshold, in thousandths of a second
	}
}

// LogMode accomplish gormlogger.Interface's LogMode
func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

// Info accomplish gormlogger.Interface's Info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)

}

// Warn accomplish gormlogger.Interface's Warn
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}

// Error accomplish gormlogger.Interface's Error
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args...)
}

// Trace accomplish gormlogger.Interface's Trace
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// excuting time
	elapsed := time.Since(begin)
	// sql and rows
	sql, rows := fc()

	// common word
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", util.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}

	// Gorm error message
	if err != nil {
		// warning
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			// error level
			l.logger().Error("Database Error", logFields...)
		}
	}

	// slow log
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	// record all sql query
	l.logger().Debug("Database Query", logFields...)
}

// Logger auxiliary method to ensure the accuracy of Zap’s built-in information Caller
func (l GormLogger) logger() *zap.Logger {
	// skip gorm default setting
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// Subtract one package and add zap.AddCallerSkip(1) to the logger initialization.
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// Returns a new zap logger with skipped line numbers
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return l.ZapLogger
}
