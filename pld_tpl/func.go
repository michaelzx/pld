package pld_tpl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/michaelzx/pld/pld_lang"
	"html/template"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Nl2br(s string) string {
	reg := regexp.MustCompile(`\n|\r|\n\r|\r\n`)
	s = reg.ReplaceAllString(s, `<br>`)
	return s
}
func Key(key string, m map[string]interface{}) interface{} {
	if val, exist := m[key]; exist {
		return val
	}
	return nil
}
func Type(v interface{}) string {
	if v == nil {
		return "nil"
	} else {
		return reflect.TypeOf(v).String()
	}
}
func FtoI64(v float64) int64 {
	return int64(v)
}

func Unescape(s string) template.HTML {
	return template.HTML(s)
}

func LangSwitch(langTag pld_lang.Tag) func(langVs ...string) string {
	return func(langVs ...string) string {
		switch langTag {
		case pld_lang.Cn:
			if len(langVs) >= 1 {
				return langVs[0]
			} else {
				return ""
			}
		case pld_lang.En:
			if len(langVs) >= 2 {
				return langVs[1]
			} else {
				return ""
			}
		default:
			return ""
		}
	}
}

func Add(b, a interface{}) (interface{}, error) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	switch av.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Int() + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Int() + int64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Int()) + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(av.Uint()) + bv.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Uint() + bv.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return float64(av.Uint()) + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	case reflect.Float32, reflect.Float64:
		switch bv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return av.Float() + float64(bv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return av.Float() + float64(bv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return av.Float() + bv.Float(), nil
		default:
			return nil, fmt.Errorf("add: unknown type for %q (%T)", bv, b)
		}
	default:
		return nil, fmt.Errorf("add: unknown type for %q (%T)", av, a)
	}
}
func Json(i interface{}) (string, error) {
	result, err := json.Marshal(i)
	if err != nil {
		return "", fmt.Errorf("toJSON", err)
	}
	return string(bytes.TrimSpace(result)), err
}


type QueryItem struct {
	K string
	V interface{}
}

func NewQueryItem(k string, v interface{}) *QueryItem {
	return &QueryItem{K: k, V: v}
}

func NewQuery(queryValues url.Values) func(newItems ...QueryItem) string {
	return func(newItems ...QueryItem) string {
		values := make(url.Values)
		for queryK, queryV := range queryValues {
			values[queryK] = queryV
		}
		for _, item := range newItems {
			vStr := ""
			switch item.V.(type) {
			case string:
				vStr = item.V.(string)
			case int32:
				vStr = strconv.FormatInt(int64(item.V.(int32)), 10)
			case int64:
				vStr = strconv.FormatInt(item.V.(int64), 10)
			case int:
				vStr = strconv.Itoa(item.V.(int))
			default:
				panic(errors.New("NewQuery not support type of " + reflect.TypeOf(item.V).String()))
			}
			// logger.Debug(item.K, item.V)
			if vStr == "" || vStr == "0" {
				values.Del(item.K)
			} else {
				values.Set(item.K, vStr)
			}
		}
		items := make([]string, 0, 0)
		for k, vs := range values {
			vStr := strings.Join(vs, ",")
			items = append(items, k+"="+vStr)
		}
		var queryStr string
		// if len(items) > 0 {
		queryStr = "?" + strings.Join(items, "&")
		// }
		// logger.Debug("NewQuery", queryStr)
		return queryStr
	}
}