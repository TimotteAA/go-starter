package redis

import (
	"time"

	"github.com/TimotteAA/go-starter/config"
	"github.com/TimotteAA/go-starter/global"
	"github.com/gomodule/redigo/redis"
)

func New(cfg *config.Config) {
	global.Redis = defaultPool(cfg).Get()
	global.SysLogger.Errorf("redis连接成功")
}

func defaultPool(cfg *config.Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                 // 最大空闲连接数
        MaxActive:   10000,                // 最大活跃连接数，0表示没有限制
        IdleTimeout: 240 * time.Second,  // 空闲连接的超时时间
        Dial: func() (redis.Conn, error) {  // 创建连接的函数
			db := cfg.RedisDB
            c, err := redis.Dial(
				"tcp", 
				cfg.RedisHost,
				redis.DialDatabase(db),
				redis.DialPassword(cfg.RedisPassword),
				// 写入延迟
				redis.DialWriteTimeout(2 * time.Second),
				// 读延迟
				redis.DialReadTimeout(1 * time.Second),
				// 连接延迟
				redis.DialConnectTimeout(10 * time.Second),
			)
            if err != nil {
                return nil, err
            }
            return c, nil
        },
	}
}