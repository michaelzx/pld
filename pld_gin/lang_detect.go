package pld_gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/pld/pld_lang"
	"github.com/michaelzx/pld/pld_logger"
	"net/http"
	"strings"
	"time"
)

const LangCookieName = "lang"

// 网站语言侦测
func LangDetectMiddleware(defaultLangTag pld_lang.Tag, enableUrlPrefixDetect bool) gin.HandlerFunc {
	return func(gc *gin.Context) {
		// 获取语言侦测方式
		// 1、URI前缀
		// 2、非URI前缀：参数[l] / cookie[langTag] / header[Accept-Language]
		langTag := pld_lang.None

		if enableUrlPrefixDetect {
			uri := strings.Split(gc.Request.RequestURI, "/")
			firstFolder := uri[1]
			if firstFolder != "" {
				if t := pld_lang.TagFromString(firstFolder); t != pld_lang.None {
					langTag = t
					gc.Request.RequestURI = "/" + strings.Join(uri[2:], "/")
				}
			}
		}
		// ----------------------------------------------------------------------------
		// 1、从URL参数中获取语言参数
		// ----------------------------------------------------------------------------
		if langTag == pld_lang.None {
			langQuery := gc.Query("l")
			langTag = pld_lang.TagFromString(langQuery)
		}
		if langTag != pld_lang.None { // 如果有，则保存到cookie中，供下次使用
			now := time.Now()
			maxAge := now.Add(time.Duration(365*24) * time.Hour)
			CookieAdd(gc, LangCookieName, langTag.String(), int(maxAge.Unix()))
		}
		// ----------------------------------------------------------------------------
		// 2、从COOKIE中获取语言参数
		// ----------------------------------------------------------------------------
		if langTag == pld_lang.None {
			if cookieValue, err := gc.Cookie(LangCookieName); err == nil {
				langTag = pld_lang.TagFromString(cookieValue)
			}
		}
		// ----------------------------------------------------------------------------
		// 3、从请求头中的[Accept-Language]获取
		// ----------------------------------------------------------------------------
		if langTag == pld_lang.None {
			acceptLanguageStr := gc.GetHeader("Accept-Language")
			acceptLanguageArr := strings.Split(acceptLanguageStr, ",")
			for _, s := range acceptLanguageArr {
				if !strings.HasPrefix(s, "q=") {
					t := pld_lang.TagFromString(s)
					if t != pld_lang.None {
						langTag = t
						break
					}
				}
			}
		}

		// ----------------------------------------------------------------------------
		// 如果都没有，则有可能是其他语言的用户，默认显示中文
		// ----------------------------------------------------------------------------
		if langTag == pld_lang.None {
			if defaultLangTag == pld_lang.None {
				langTag = pld_lang.Cn
			} else {
				langTag = defaultLangTag
			}
		}
		if cv, err := gc.Cookie(LangCookieName); cv != langTag.String() || errors.Is(err, http.ErrNoCookie) {
			now := time.Now()
			maxAge := now.Add(time.Duration(365*24) * time.Hour)
			CookieAdd(gc, LangCookieName, langTag.String(), int(maxAge.Unix()))
		}
		pld_logger.Debug("lang_detect", langTag)
		gc.Set(pld_lang.GinContextKey, langTag)
		gc.Next()
	}
}
