package pld_tpl

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_fs"
	"github.com/pkg/errors"
	"html/template"
	"path/filepath"
)

func Render404(gc *gin.Context, themeName string) {
	RenderErr(gc, themeName, 404, "not found")
}
func Render500(gc *gin.Context, themeName, message string) {
	RenderErr(gc, themeName, 500, message)
}
func RenderErr(gc *gin.Context, themeName string, code int, message string) error {
	// 解析出正确的模板路径
	tplPath := "." + filepath.Join(pld_fs.WebDir, pld_fs.ThemeUrl, themeName, "error") + ".gohtml"
	// 加载模板引擎
	t, err := loadTemplateEngine(tplPath, template.FuncMap{})
	if err != nil {
		return errors.Wrap(err, "模板引擎加载失败")
	}
	var doc bytes.Buffer
	err = t.Execute(&doc, DataMap{
		"Code":    code,
		"Message": message,
	})
	if err != nil {
		return errors.Wrap(err, "模板引擎执行失败")
	}
	htmlStr := doc.String()
	htmlStr, err = m.String("text/html", htmlStr)
	if err != nil {
		return err
	}
	gc.Data(code, "text/html; charset=utf-8", []byte(htmlStr))
	return nil
}
