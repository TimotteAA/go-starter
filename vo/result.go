package vo

import (
	"time"

	biz_error "github.com/TimotteAA/go-starter/error"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Result struct {
	*biz_error.BizError
	Data interface{} `json:"data"`
	RequestId interface{} `json:"requestId"`
	Timestamp int64 `json:"timestamp"`
}

func Success(data interface{}, c *fiber.Ctx) Result {
	return Result{
		BizError: biz_error.New(biz_error.SUCCESS),
		Data: data,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
		Timestamp: time.Now().UnixMilli(),
	}
}

func Error(bizErr *biz_error.BizError, c *fiber.Ctx) Result {
	return Result{
		BizError: bizErr,
		Data: nil,
		RequestId: c.Locals(requestid.ConfigDefault.ContextKey),
		Timestamp: time.Now().UnixMilli(),
	}
}