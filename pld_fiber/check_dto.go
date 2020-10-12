package pld_fiber

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_errs"
	"github.com/michaelzx/pld/pld_lang"
	"github.com/michaelzx/pld/pld_validator"
	"reflect"
	"strings"
)

func CheckDTO(c *fiber.Ctx, reqPtr interface{}) {
	langTag := GetLangTag(c)
	if langTag == pld_lang.None {
		langTag = pld_lang.Cn
	}
	parseErr := c.BodyParser(reqPtr)
	if parseErr != nil {
		panic(pld_errs.CommonBadRequest(parseErr.Error()))
	}
	bindErr := validate.ValidateStruct(reqPtr)
	if bindErr != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := bindErr.(*validator.InvalidValidationError); ok {
			fmt.Println(bindErr)
			return
		}
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

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
