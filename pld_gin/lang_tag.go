package pld_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_lang"
)

func GetLangTag(gc *gin.Context) pld_lang.Tag {
	tag, exist := gc.Get(pld_lang.GinContextKey)
	if !exist {
		return pld_lang.None
	}
	v, ok := tag.(pld_lang.Tag)
	if !ok {
		return pld_lang.None
	}
	return v
}
