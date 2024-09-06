package middleware

import (
	"github.com/TimotteAA/go-starter/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitMiddleware(app *fiber.App) {
	// request id 中间件
	app.Use(requestid.New())
	// 日志中间件
	app.Use(logger.New())
}