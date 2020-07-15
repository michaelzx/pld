package pld_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/michaelzx/pld/pld_errs"
	"github.com/michaelzx/pld/pld_lang"
	"github.com/michaelzx/pld/pld_validator"
	"reflect"
	"strings"
)

func CheckDTO(c *gin.Context, reqPtr interface{}) {
	langTag := GetLangTag(c)
	if bindErr := c.ShouldBindJSON(reqPtr); bindErr != nil {
		if vErrs, ok := bindErr.(validator.ValidationErrors); ok {
			for _, vErr := range vErrs {
				e := vErr.(validator.FieldError)
				var errMsg string
				if langTag == pld_lang.Cn {
					errMsg = e.Translate(pld_validator.TransCn)
				} else {
					errMsg = e.Translate(pld_validator.TransEn)
				}
				if filed, ok := reflect.TypeOf(reqPtr).Elem().FieldByName(e.StructField()); ok {
					var filedName string
					if langTag == pld_lang.Cn {
						filedName = filed.Tag.Get("cn")
					} else {
						filedName = filed.Tag.Get("en")
					}
					errMsg = strings.ReplaceAll(errMsg, e.Field(), filedName)
				}
				panic(pld_errs.CommonBadRequest(errMsg))
			}
		} else {
			panic(pld_errs.CommonBadRequest(bindErr.Error()))
		}
	}
}
