package pld_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_errs"
	"strconv"
)

func ParamToBool(c *fiber.Ctx, paramName string) bool {
	return c.Params(paramName) == "true"
}
func ParamInt64(c *fiber.Ctx, paramName string) int64 {
	str := c.Params(paramName)
	if str == "" {
		panic(pld_errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(pld_errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	return i64
}
func ParamInt64Default(c *fiber.Ctx, paramName string, defaultValue int64) int64 {
	str := c.Params(paramName)
	if str == "" {
		return defaultValue
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i64
}
