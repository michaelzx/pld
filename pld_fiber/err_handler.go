package pld_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_errs"
	"github.com/michaelzx/pld/pld_logger"
	"go.uber.org/zap"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	pld_logger.Debug("ErrorHandler", ctx, err)
	if e, ok := err.(*pld_errs.BadRequest); ok {
		return ctx.Status(e.Status).JSON(e.BizErr)
	} else if e, ok := err.(*pld_errs.Unauthorized); ok {
		return ctx.Status(e.Status).JSON(e.BizErr)
	} else if e, ok := err.(*pld_errs.Forbidden); ok {
		return ctx.Status(e.Status).JSON(e.BizErr)
	} else if e, ok := err.(*pld_errs.NotFound); ok {
		return ctx.Status(e.Status).JSON(e.BizErr)
	} else {
		unknown := pld_errs.NewUnknown("服务器繁忙")
		pld_logger.Error("get error from recover", zap.Any("unknown error", err))
		return ctx.Status(unknown.Status).JSON(unknown.BizErr)
	}
}
