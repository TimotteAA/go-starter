package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TimotteAA/go-starter/config"
	"github.com/TimotteAA/go-starter/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(config *config.Config) {
	err := initDB(config)
	if err != nil {
		global.SysLogger.Errorf("数据库连接失败 %v", err)
		return
	}
}

func initDB(config *config.Config) error {
	dsn := buildDSN(config)
	db, err := gorm.Open(postgres.Open(dsn), buildConfig())
	if err != nil {
		return err
	}

	global.DB = db;
	global.SysLogger.Info("数据库连接成功!")

	err = setupDBConnectionPool()
	if err != nil {
		return err
	}
	return nil
}

func buildDSN(config *config.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", 
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBPort,
		config.AppName,
		config.AppTimeZone,
	)
}

func buildConfig() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:              2 * time.Second,   // Slow SQL threshold
				LogLevel:                   logger.Silent, // Log level
				IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,           // Don't include params in the SQL log
				Colorful:                  false,          // Disable color
			},
		),
	}
}

// 设置gorm底层的数据库配置
func setupDBConnectionPool() error {
	SqlDB, err := global.DB.DB()
	if err != nil {
		global.SysLogger.Errorf("配置数据库连接池失败 %v", err)
		return err
	}

	SqlDB.SetMaxOpenConns(100)
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetConnMaxLifetime(10 * time.Second)

	return nil
}