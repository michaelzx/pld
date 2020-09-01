package pld_gin

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func NginxInternalResponse(gc *gin.Context, filename, redirect string) {
	filename = url.QueryEscape(filename) // 防止中文乱码
	filename = strings.Replace(filename, "+", "%20", -1)
	if strings.HasSuffix(strings.ToLower(filename), ".pdf") {
		// pdf可以在线预览
		gc.Header("Content-Type", "application/pdf")
		gc.Header("Content-Disposition", "filename="+filename)
	} else {
		gc.Header("Content-Type", "application/octet-stream")
		gc.Header("Content-Disposition", "attachment; filename="+filename)
	}
	gc.Header("Cache-Control", "no-cache")
	gc.Header("X-Accel-Buffering", " no")
	// gc.Header("X-Accel-Limit-Rate", "102400") // 速度限制 102400 Byte/s = 100KB/s
	// header("Accept-Ranges: none");//单线程 限制多线程
	gc.Header("X-Accel-Redirect", redirect)

}
