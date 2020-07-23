package pld_tpl

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_fs"
	"github.com/michaelzx/pld/pld_gin"
	"github.com/michaelzx/pld/pld_logger"
	"github.com/pkg/errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var tplCache sync.Map

type tplCacheItem struct {
	LastModifyTime int64
	Html           string
}

// func LangDictFunc(langTag, k string) string {
// 	langMap := basic.GetLangCache(langTag)
// 	if v, exist := langMap[k]; exist {
// 		return v
// 	} else {
// 		return ""
// 	}
// }

func RenderHtml(gc *gin.Context, themeName, tplName string, funcMap template.FuncMap, data DataMap) error {
	langTag := pld_gin.GetLangTag(gc)
	// 解析出正确的模板路径
	tplPath := "." + filepath.Join(pld_fs.WebDir, pld_fs.ThemeUrl, themeName, tplName) + ".gohtml"
	pld_logger.Debug(tplPath)
	// 最终的模板函数
	finalFunc := make(template.FuncMap)
	for k, v := range funcMap {
		finalFunc[k] = v
	}
	finalFunc["LangSwitch"] = LangSwitch(langTag)
	for k, v := range DefaultFuncMap {
		finalFunc[k] = v
	}

	// finalFunc["LangDict"] = LangDictFunc
	// 最终的数据
	finalData := NewDataMap()
	for k, v := range data {
		finalData[k] = v
	}
	// 全局变量，所有模板都可调用
	finalData["TplName"] = tplName
	finalData["ThemeName"] = themeName
	finalData["ThemeDir"] = filepath.Join("theme", themeName)
	finalData["Path"] = gc.Request.URL.Path
	finalData["LangTag"] = langTag

	// 加载模板引擎
	t, err := loadTemplateEngine(tplPath, finalFunc)
	if err != nil {
		return errors.Wrap(err, "模板引擎加载失败")
	}
	var doc bytes.Buffer
	err = t.Execute(&doc, finalData)
	if err != nil {
		return errors.Wrap(err, "模板引擎执行失败")
	}
	htmlStr := doc.String()
	htmlStr = styleProcessor(htmlStr, finalData)
	htmlStr = jsProcessor(htmlStr, finalData)
	htmlStr = imgProcessor(htmlStr, finalData)
	htmlStr, err = m.String("text/html", htmlStr)
	if err != nil {
		return err
	}
	gc.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlStr))
	return nil
}

// loadTemplateEngine 创建一个模板引擎对象
func loadTemplateEngine(path string, funcMap template.FuncMap) (t *template.Template, err error) {
	var html string
	html, err = LoadTemplateHtml(path)
	if err != nil {
		return
	}
	t, err = template.New(filepath.Base(path)).Funcs(funcMap).Parse(html)
	if err != nil {
		return
	}
	return
}

// LoadTemplateHtml 读取模板的html
func LoadTemplateHtml(path string) (html string, err error) {

	f, err := os.Open(path)
	if err != nil {
		err = errors.Wrap(err, "模板文件不存在:"+path)
		return
	}
	defer f.Close()

	fStat, err := f.Stat()
	if err != nil {
		return
	}
	cacheHit := false
	// 先从有效缓存里面拿
	if cacheItemInMap, ok := tplCache.Load(path); ok {
		if cacheItem, ok := cacheItemInMap.(tplCacheItem); ok {
			if cacheItem.LastModifyTime == fStat.ModTime().Unix() {
				// 缓存中记录的最后修改时间与文件一致，则加载缓存
				html = cacheItem.Html
				cacheHit = true
			}
		}
	}
	// 如果没有命中有效缓存，则从文件中取，并保存到缓存
	if !cacheHit {
		var fd []byte
		fd, err = ioutil.ReadAll(f)
		if err != nil {
			return
		}
		html = string(fd)
		tplCache.Store(path, tplCacheItem{
			LastModifyTime: fStat.ModTime().Unix(),
			Html:           html,
		})
	}
	html, err = parseNested(html, path)
	if err != nil {
		return
	}
	return
}

// parseNested 解析include标签
func parseNested(html, path string) (newHtml string, err error) {
	dir := filepath.Dir(path)
	reg := regexp.MustCompile(`\{\{\s*include\s+"(?P<filename>.+\.gohtml)"\s*}}`)
	result := reg.FindAllStringSubmatch(html, -1)
	nestedPathMap := make(map[string]struct{})
	newHtml = html
	for _, group := range result {
		nestedPath := filepath.Join(dir, group[1])
		if _, ok := nestedPathMap[nestedPath]; ok {
			continue
		}
		var nestedHtml string
		nestedHtml, err = LoadTemplateHtml(nestedPath)
		if err != nil {
			return
		}
		newHtml = strings.ReplaceAll(newHtml, group[0], nestedHtml)
	}
	return
}
