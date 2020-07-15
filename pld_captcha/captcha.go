package pld_captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"strings"
	"time"
)

var captchaStore base64Captcha.Store

func init() {
	expiration := time.Duration(30) * time.Second
	store := base64Captcha.NewMemoryStore(999, expiration)
	captchaStore = store
}

func Generate() (id, b64s string, err error) {
	driver := base64Captcha.NewDriverString(70, 200, 0,
		base64Captcha.OptionShowHollowLine, 4,
		base64Captcha.TxtNumbers+base64Captcha.TxtSimpleCharaters,
		&color.RGBA{
			R: 240,
			G: 248,
			B: 254,
			A: 254,
		}, []string{
			// "wqy-microhei.ttc",
			// "3Dumb.ttf",
			// "ApothecaryFont.ttf",
			// "Comismsh.ttf",
			// "DENNEthree-dee.ttf",
			// "DeborahFancyDress.ttf",
			// "Flim-Flam.ttf",
			"chromohv.ttf",
			// "actionj.ttf",
			// "RitaSmith.ttf",
		})
	driver.Source = "346789ABCDEFGHJKLMNPQRTUVWXY" // 排除容易混淆的字符：0125IOSZ
	c := base64Captcha.NewCaptcha(driver, captchaStore)
	return c.Generate()
}
func GenerateNum() (id, b64s string, err error) {
	driver := base64Captcha.NewDriverString(70, 200, 0,
		base64Captcha.OptionShowHollowLine, 4,
		base64Captcha.TxtNumbers,
		&color.RGBA{
			R: 243,
			G: 251,
			B: 254,
			A: 0,
		}, []string{
			// "wqy-microhei.ttc",
			// "3Dumb.ttf",
			// "ApothecaryFont.ttf",
			// "Comismsh.ttf",
			// "DENNEthree-dee.ttf",
			// "DeborahFancyDress.ttf",
			"Flim-Flam.ttf",
			"chromohv.ttf",
			// "actionj.ttf",
			"RitaSmith.ttf",
		})
	c := base64Captcha.NewCaptcha(driver, captchaStore)
	return c.Generate()
}

// base64Captcha verify http handler
func Verify(id, code string) bool {
	storeCode := captchaStore.Get(id, true)
	// logger.Debug("验证码验证-系统中存的", storeCode)
	// logger.Debug("验证码验证-用户输入的", code)
	return strings.ToLower(code) == strings.ToLower(storeCode)
}
