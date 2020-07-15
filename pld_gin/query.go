package pld_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_errs"
	"strconv"
)

func QueryToBool(c *gin.Context, paramName string) bool {
	return c.Query(paramName) == "true"
}
func QueryInt64(c *gin.Context, paramName string) int64 {
	str := c.Query(paramName)
	if str == "" {
		panic(pld_errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(pld_errs.ParamsErr.Suffix(paramName + "，必须是数字"))
	}
	return i64
}
func QueryInt64Default(c *gin.Context, paramName string, defaultValue int64) int64 {
	str := c.Query(paramName)
	if str == "" {
		return defaultValue
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return i64
}
