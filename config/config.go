package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// 应用名称
	AppName string
	// 端口
	AppPort string
	// 时区
	AppTimeZone string
	// 数据库host
	DBHost string
	// 数据库Port
	DBPort string
	// 数据库用户
	DBUser string
	// 数据库用户密码
	DBPassword string
	// Redis Host
	RedisHost string
	// 用哪个redis
	RedisDB int
	// redis密码
	RedisPassword string
	// 日志文件存储地址
	LogFilePath     string
	// 日志文件名称
	LogFileName     string
	// 日志文件时间戳的格式
	LogTimestampFmt string
	// 日志最多保留时间，天数
	LogMaxAge       int
	// 日志最多保存数量
	LogMaxBackups  int
	// 日志最多保留大小，MB
	LogMaxSize int
	// 日志级别
	LogLevel        string
}

var (
	config *Config
)

func InitConfig(mode string) (*Config, error) {
	var (
		configFilePath = ".env." + strings.Trim(mode, ".")
	)

	err := godotenv.Load(configFilePath)
	if err != nil {
		log.Fatalf("fail to load config file %s", configFilePath)
		return nil, err
	}

	config = &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
		AppTimeZone: os.Getenv("APP_TIMEZONE"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		LogFilePath: os.Getenv("LOG_FILE_PATH"),
		LogFileName: os.Getenv("LOG_FILE_NAME"),
		LogTimestampFmt: os.Getenv("LOG_TIMESTAMP_FMT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}


	config.RedisDB, err = strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Fail to parse redis db");
		return nil, err
	}

	config.LogMaxAge, err = strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		log.Fatalf("Fail to parse log max age");
		return nil, err
	}

	config.LogMaxBackups, err = strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		log.Fatalf("Fail to parse log rotation time");
		return nil, err
	}

	config.LogMaxSize, err = strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		log.Fatalf("Fail to parse log max size");
		return nil, err
	}

	return config, nil
}
