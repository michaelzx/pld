package pld_tpl

import (
	"bytes"
	"github.com/michaelzx/pld/pld_fs"
	"github.com/pkg/errors"
	"html/template"
	"io/ioutil"
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

func RenderHtml(themeName, tplName string, funcMap FuncMap, dataMap DataMap) (html string, err error) {
	// 解析出正确的模板路径
	tplPath := filepath.Join(pld_fs.WebPath, pld_fs.ThemeUrl, themeName, tplName)
	// pld_logger.Debug(tplPath)
	themeDir := filepath.Join("theme", themeName)
	// 全局变量，所有模板都可调用 TODO 待定
	// funcMap["TplName"] = tplName
	// funcMap["ThemeName"] = themeName
	// dataMap["ThemeDir"] = filepath.Join("theme", themeName)
	// funcMap["Path"] = gc.Request.URL.Path
	// funcMap["LangTag"] = langTag

	// 加载模板引擎
	// start := time.Now()
	// pld_logger.Debug("开始")
	t, err := loadTemplateEngine(tplPath, funcMap)
	// fmt.Println(err)
	if err != nil {
		// pld_logger.Error("模板引擎加载失败", err)
		err = errors.Wrap(err, "模板引擎加载失败")
		return
	}
	var tplBuffer bytes.Buffer
	err = t.Execute(&tplBuffer, dataMap)
	if err != nil {
		// pld_logger.Error("模板引擎加载失败", err)
		err = errors.Wrap(err, "模板引擎加载失败")
		return
	}
	htmlStr := tplBuffer.String()
	htmlStr = styleProcessor(htmlStr, themeDir)
	htmlStr = jsProcessor(htmlStr, themeDir)
	htmlStr = imgProcessor(htmlStr, themeDir)
	htmlStr, err = m.String("text/html", htmlStr)
	if err != nil {
		// pld_logger.Error("模板引擎加载失败", err)
		err = errors.Wrap(err, "模板引擎加载失败")
		return
	}

	// pld_logger.Debug("结束", time.Since(start).String())
	html = htmlStr
	return
}

// loadTemplateEngine 创建一个模板引擎对象
func loadTemplateEngine(path string, funcMap FuncMap) (t *template.Template, err error) {
	var html string
	html, err = LoadHtmlFile(path)
	if err != nil {
		return
	}
	t, err = template.New(path).Funcs(template.FuncMap(funcMap)).Parse(html)
	if err != nil {
		return
	}
	return
}

// LoadHtmlFile 读取模板的html
func LoadHtmlFile(path string) (html string, err error) {

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		err = errors.Wrap(err, "模板文件不存在:"+path)
		return
	}

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
				// pld_logger.Debug("模板缓存命中", path)
				cacheHit = true
			}
		}
	}
	// 如果没有命中有效缓存，则从文件中取，并保存到缓存
	if !cacheHit {
		// pld_logger.Debug("模板缓未存命中", path)
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
	// fmt.Println(html)
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
		nestedHtml, err = LoadHtmlFile(nestedPath)
		if err != nil {
			return
		}
		newHtml = strings.ReplaceAll(newHtml, group[0], nestedHtml)
		nestedPathMap[nestedPath] = struct{}{}
	}
	return
}
