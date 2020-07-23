package pld_tpl

import (
	"regexp"
	"strings"
)

// RenderStyles 判断html中的静态资源，替换成相应的目录
func replaceSiteRes(html string, tagStr string, targetStr string, data DataMap) string {
	if v, ok := data["ThemeDir"].(string); ok {
		newStr := strings.ReplaceAll(tagStr, targetStr, "/"+v+"/"+targetStr)
		html = strings.ReplaceAll(html, tagStr, newStr)
	}
	return html
}

// styleProcessor 样式处理器
func styleProcessor(html string, data DataMap) string {
	styleReg := regexp.MustCompile(`<link\s+.*href=["']([^>]+)["'].*>`)
	styles := styleReg.FindAllStringSubmatch(html, -1)
	for _, style := range styles {
		if strings.Contains(style[0], "stylesheet") {
			if !strings.HasPrefix(style[1], "http") {
				html = replaceSiteRes(html, style[0], style[1], data)
			}
		}
	}
	return html
}

func jsProcessor(html string, data DataMap) string {
	jsReg := regexp.MustCompile(`<script\s.*src=["']([^>]+)["']\s?>.*</script>`)
	jss := jsReg.FindAllStringSubmatch(html, -1)
	for _, js := range jss {
		if !strings.HasPrefix(js[1], "http") {
			html = replaceSiteRes(html, js[0], js[1], data)
		}
	}
	return html
}
func imgProcessor(html string, data DataMap) string {
	// 2020-04-07 michael 避免和vue属性冲突
	jsReg := regexp.MustCompile(`<img.*\s+[:]{0}src=["']([^>]+)["']\s?>`)
	jss := jsReg.FindAllStringSubmatch(html, -1)
	for _, js := range jss {
		if !strings.HasPrefix(js[1], "http") && !strings.HasPrefix(js[1], "[[") && !strings.HasPrefix(js[1], "/") {
			html = replaceSiteRes(html, js[0], js[1], data)
		}
	}
	return html
}
