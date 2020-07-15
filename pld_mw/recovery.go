package pld_mw

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_errs"
	"github.com/michaelzx/pld/pld_logger"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(*pld_errs.BadRequest); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*pld_errs.Unauthorized); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*pld_errs.Forbidden); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				if e, ok := err.(*pld_errs.NotFound); ok {
					c.JSON(e.Status, e.BizErr)
					c.AbortWithStatus(e.Status)
					return
				}
				unknown := pld_errs.NewUnknown("服务器繁忙")
				c.JSON(unknown.Status, unknown.BizErr)
				pld_logger.Error("get error from recover", zap.Any("unknown error", err))
				c.AbortWithStatus(unknown.Status)
				return
			}
		}()
		c.Next()
	}
}
