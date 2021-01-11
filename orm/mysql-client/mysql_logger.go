/*
** implement logger.Interface
	// Interface logger interface
	type Interface interface {
		LogMode(LogLevel) Interface
		Info(context.Context, string, ...interface{})
		Warn(context.Context, string, ...interface{})
		Error(context.Context, string, ...interface{})
		Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
	}
*/
package mysqlclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	. "github.com/luckyweiwei/base/logger"
	"gorm.io/gorm/logger"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type SqlLogger struct {
	SqlDebug int
}

func (s *SqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	return s
}

func (s *SqlLogger) Info(ctx context.Context, msg string, data ...interface{}) {}

func (s *SqlLogger) Warn(ctx context.Context, msg string, data ...interface{}) {}

func (s *SqlLogger) Error(ctx context.Context, msg string, data ...interface{}) {}

func (s *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if s.SqlDebug != 1 {
		return
	}

	elapsed := time.Since(begin)
	sql, sqlAffected := fc()

	ignorestr := `/*no print*/`
	if strings.Contains(sql, ignorestr) {
		return
	}

	msg := fmt.Sprintf("\n Sql:%v%v%v\n RowsAffected:%v%v%v\n Cost:%v%v%v\n",
		Blue, sql, Reset, Yellow, sqlAffected, Reset, Green, elapsed, Reset)
	Log.Info(msg)

	if err != nil {
		Log.Error(err)
	}
}
