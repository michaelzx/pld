package pld_sql

import (
	"bytes"
	"github.com/michaelzx/pld/pld_reflect"
	"github.com/michaelzx/pld/pld_logger"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

var tmpl = template.New("")

// Resolver 使用模板引擎解析sql模板及参数，返回最终的sql和参数
type Resolver struct {
	sql    string
	values []interface{}
}

func NewResolver(sqlTplStr string, values interface{}) (*Resolver, error) {
	// 先简单格式化一下sql字符串，方便在日志中查看
	r := &Resolver{}
	err := r.resolve(sqlTplStr, values)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Resolver) GetSql() string {
	return s.sql
}

func (s *Resolver) GetValues() []interface{} {
	return s.values
}

func (s *Resolver) getTpl(sqlTplStr string) (*template.Template, error) {
	// newMd5 := md5.New()
	// newMd5.Write([]byte(sqlTplStr))
	// sqlTplStrMd5 := hex.EncodeToString(newMd5.Sum(nil))
	// 将这个MD5作为模板名称，如果存在直接返回
	t := tmpl.Lookup(sqlTplStr)
	if t != nil {
		pld_logger.Debug("找到sql模板")
		return t, nil
	}
	// 如果不存在，则创建后返回
	pld_logger.Debug("创建sql模板")
	t = tmpl.New(sqlTplStr)
	t, err := t.Parse(sqlTplStr)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Resolver) resolve(sqlTplStr string, queryParams interface{}) error {
	if !pld_reflect.IsPtr(queryParams) {
		return errors.New("params 必须是指针类型")
	}
	t, err := s.getTpl(sqlTplStr)
	if err != nil {
		return errors.Wrap(err, "获取模板失败")
	}
	var sql bytes.Buffer
	if err := t.Execute(&sql, queryParams); err != nil {
		return err
	}
	sqlStr := sql.String()
	// 正则找出，模板分析出中所有的 #{xxxx}
	tplParamsRegexp := regexp.MustCompile(`#{(?P<param>\w*)}`)
	tplParams := tplParamsRegexp.FindAllStringSubmatch(sqlStr, -1)
	// 校验 tplParams 是否 都在params中
	// 仅支持 Params 及 struct
	if p, ok := queryParams.(*map[string]interface{}); ok {
		s.values = make([]interface{}, 0, len(tplParams))
		pMap := *p
		for _, tplParam := range tplParams {
			full := tplParam[0]
			short := tplParam[1]
			v, exist := pMap[short]
			if !exist {
				panic(full + "：不在queryParams中")
			}
			sqlStr = strings.Replace(sqlStr, full, "?", 1)
			s.values = append(s.values, v)
		}
	} else if pld_reflect.IsStruct(queryParams) {
		rt := reflect.TypeOf(queryParams)
		rve := reflect.ValueOf(queryParams).Elem()
		fim := pld_reflect.StructFieldIdxMap(rt)
		for _, tplParam := range tplParams {
			full := tplParam[0]
			short := tplParam[1]
			fi, exist := fim[short]
			if !exist {
				panic(full + "：不在queryParams中")
			}
			v := rve.Field(fi).Interface()
			sqlStr = strings.Replace(sqlStr, full, "?", 1)
			s.values = append(s.values, v)
		}
	} else {
		return errors.New("queryParams 仅支持 map 及 struct")
	}
	// 去掉换行、大于1个的空格
	sqlStr = strings.ReplaceAll(sqlStr, "\n", " ")
	re, _ := regexp.Compile(`\s{2,}`)
	sqlStr = re.ReplaceAllString(sqlStr, " ")
	s.sql = sqlStr
	return nil
}
