package pld_gin

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/michaelzx/pld/pld_logger"
	"github.com/michaelzx/pld/pld_mw"
	"github.com/michaelzx/pld/pld_validator"
	"net/http"
)

type RouterRegisterFunc func(*gin.Engine)

var router *gin.Engine

func IntRouter(register RouterRegisterFunc, debug bool) *gin.Engine {
	// gin.DebugPrintRouteFunc = nil
	// gin.DisableConsoleColor()
	binding.Validator = &pld_validator.DefaultValidator{}
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	rRoot := gin.New()
	replaceGinLogger(rRoot)
	rRoot.Use(gin.Recovery())
	rRoot.Use(pld_mw.Recovery())
	register(rRoot)
	// printRouteMapping(rRoot)
	router = rRoot
	return rRoot
}

func GetRouter() *gin.Engine {
	if router == nil {
		pld_logger.Fatal("db 未初始化")
	}
	return router
}

func replaceGinLogger(g *gin.Engine) {
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		pld_logger.Debug(fmt.Sprintf("[%d] [%s] %s | %s", param.StatusCode, param.Method, param.Path, param.Latency.String()))
		return "" // 替换gin自带的日志
	}))
}
func printRouteMapping(g *gin.Engine) {
	var buf bytes.Buffer
	for _, route := range g.Routes() {
		buf.Reset()
		buf.WriteString("")
		buf.WriteString("[")
		buf.WriteString(route.Method)
		buf.WriteString("]")
		if route.Method == http.MethodGet {
			buf.WriteString(" ")
		}
		buf.WriteByte(' ')
		buf.WriteString(route.Path)
		buf.WriteString("  ----> ")
		// cutStart := strings.LastIndex(route.Handler, "/")
		// buf.WriteString(utils.Str.SubStringFromStart(route.Handler, cutStart+1))
		pld_logger.Debug(buf.String())
	}
}
