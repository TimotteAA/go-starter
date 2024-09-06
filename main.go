package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/TimotteAA/go-starter/config"
	"github.com/TimotteAA/go-starter/logger"
	"github.com/TimotteAA/go-starter/middleware"
	"github.com/TimotteAA/go-starter/redis"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var mode = "development"
	flag.StringVar(&mode, "mode", "development", "Set the mode of app, development or production")

	flag.Parse()

	config, err := config.InitConfig(mode)
	if err != nil {
		log.Fatal("Fail to load config")
	}

	// 创建logger
	logger.InitLogger(config)

	// 连接各种中间件
	redis.New(config)


	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	// 注册中间件
	middleware.InitMiddleware(app)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.JSON("hello fiber")
	})


	app.Listen(fmt.Sprintf(":%s", config.AppPort))
}