package pld_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_lang"
)

func GetLangTag(c *fiber.Ctx) pld_lang.Tag {
	tagStr := c.Get(pld_lang.GinContextKey, "")
	return pld_lang.TagFromString(tagStr)
}
