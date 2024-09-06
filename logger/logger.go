package logger

import (
	// 日志库

	"log"
	"os"
	"path"

	"github.com/TimotteAA/go-starter/config"
	"github.com/TimotteAA/go-starter/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rifflock/lfshook"
	logrus "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(cfg *config.Config) {
	logger := logrus.New()

	var (
		// 日志存储的目录
		logFilePath = cfg.LogFilePath
		// 日志文件名称
		logFileName = cfg.LogFileName
		fileName = path.Join(logFilePath, logFileName)
	)
	
	// 确保文件存在
	// linux的文件管理分为两种，rwd的模式和数字的模式，这里用的数字模式
	// 7 = 4 (读) + 2 (写) + 1 (执行) = rwx，用户权限
	// 5 = 4 (读) + 1 (执行) = r-x 同组只能读
	// 5 = 4 (读) + 1（执行）= r-x，非同组的只能读
	// 0 = 无任何权限 = ---
	err := os.MkdirAll(logFilePath, 0755)
	if err != nil {
		log.Fatalf("创建日志目录失败 %v", err)
	}

	global.LogFile, err = os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
	if err != nil {
		log.Fatalf("创建日志存储文件失败 %v", err)
	}
	// 指定logger写入的文件
	logger.Out = global.LogFile
	// 设定时间戳格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: cfg.LogTimestampFmt,
	})
	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("解析日志level失败 %v", err)
	}
	// 设定日志级别
	logger.SetLevel(logLevel)

	// 处理日志轮转
	// 配置lumberjack日志轮转
	writer := &lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		
		MaxSize:    cfg.LogMaxSize,              // 日志文件最大尺寸（MB）
		MaxBackups: cfg.LogMaxBackups,                // 保留的旧日志文件个数
		MaxAge:     cfg.LogMaxAge,               // 日志文件最长保留天数
		Compress:   true,             // 是否压缩旧日志文件
	}

	// 配置日志级别与轮转日志的映射
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}

	hook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: cfg.LogTimestampFmt,
	})

	logger.AddHook(hook)
	global.SysLogger = logger
}

func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqId := c.Locals(requestid.ConfigDefault.ContextKey)
		// 每个请求链单独的业务日志logger
		bizLogger := global.SysLogger.WithFields(logrus.Fields{
			"requestId": reqId,
			"requestIp": c.IP(),
		})
		c.Locals("bizLogger", bizLogger)
		return c.Next()
	}
}