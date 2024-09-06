package global

import (
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

var (
	// 日志文件
	LogFile *os.File
	// 系统日志，用于第三方组件的日志记录
	SysLogger *logrus.Logger
	// Redis链接
	Redis redis.Conn
)