package global

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	// 日志文件
	LogFile *os.File
	// 系统日志，用于第三方组件的日志记录
	SysLogger *logrus.Logger
)