package db

import (
	"context"
	"errors"
	"time"

	"github.com/TimotteAA/go-starter/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const slowQueryThreshold = 2 * time.Second

type dbLog struct {
	LogLevel logger.LogLevel
}

func newLogger() *dbLog {
	return new(dbLog)
}

func (l *dbLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	global.SysLogger.WithContext(ctx).Debugf(msg, data...)
}
func (l *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	global.SysLogger.WithContext(ctx).Warnf(msg, data...)
}

func (l *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	global.SysLogger.WithContext(ctx).Errorf(msg, data...)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// sql语句和行数
	sql, rows := fc()
	// 用时
	elapsed := time.Since(begin)

	// 加上字段
	logger := global.SysLogger.WithFields(
		logrus.Fields{
			"sql": sql,
			"time": elapsed.Microseconds(),
			"rows": rows,
		},
	)

	if err != nil {
		// 记录没找到，用Warn
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("Database ErrRecordNotFound")
		} else {
			// 其他错误使用 error 等级
			logger.WithFields(logrus.Fields{
				"error": err,
			}).Error("Database error")
		}
	}

	// 慢查询日志
	if elapsed > slowQueryThreshold {
		logger.Warn("Slow sql query")
	} else {
		// 其他的sql查询用debug
		logger.Debug("Database query")
	}
}
