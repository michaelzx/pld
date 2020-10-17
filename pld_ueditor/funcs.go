package pld_ueditor

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

func GetFileHeader(ctx interface{}, key string) (*multipart.FileHeader, error) {
	if _ctx, ok := ctx.(*fiber.Ctx); ok {
		return _ctx.FormFile(key)
	}
	if _ctx, ok := ctx.(*gin.Context); ok {
		return _ctx.FormFile(key)
	}
	return nil, errors.New("仅支持：*fiber.Ctx、*gin.Context")
}
