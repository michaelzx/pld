package pld_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_errs"
	"github.com/michaelzx/pld/pld_logger"
	"go.uber.org/zap"
)

// New creates a new middleware handler
func Recover() fiber.Handler {
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Catch panics
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(*pld_errs.BadRequest); ok {
					_ = c.SendStatus(e.Status)
					_ = c.JSON(e.BizErr)
				} else if e, ok := err.(*pld_errs.Unauthorized); ok {
					_ = c.SendStatus(e.Status)
					_ = c.JSON(e.BizErr)
				} else if e, ok := err.(*pld_errs.Forbidden); ok {
					_ = c.SendStatus(e.Status)
					_ = c.JSON(e.BizErr)
					return
				} else if e, ok := err.(*pld_errs.NotFound); ok {
					_ = c.SendStatus(e.Status)
					_ = c.JSON(e.BizErr)
				} else {
					unknown := pld_errs.NewUnknown("服务器繁忙")
					_ = c.SendStatus(unknown.Status)
					_ = c.JSON(unknown.BizErr)
					pld_logger.Error("get error from recover", zap.Any("unknown error", err))
				}
			}
		}()
		// Return err if exist, else move to next handler
		return c.Next()
	}
}
