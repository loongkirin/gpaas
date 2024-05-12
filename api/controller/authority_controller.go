package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/loongkirin/gdk/cache"
	"github.com/loongkirin/gdk/captcha"
	"github.com/loongkirin/gdk/gin/api"
	// "github.com/loongkirin/gpaas/app"
)

// var store = base64Captcha.DefaultMemStore
// var cpCache = cache.NewRedisStore(app.AppContext.APP_REDIS, "cpatcha_", time.Minute*3)
var cpCache = cache.NewInMemoryStore(time.Minute * 3)
var store = captcha.NewCaptchaStore(cpCache, time.Minute*1)
var cp = captcha.NewCaptcha((store))

type AuthorityController struct {
}

func NewAuthorityController() *AuthorityController {

	return &AuthorityController{}
}

func (t *AuthorityController) Captcha(c *gin.Context) {
	if id, b64s, _, err := cp.GenerateDigitCaptcha(); err != nil {
		api.Fail(c, "验证码获取失败", map[string]interface{}{})
	} else {
		api.Ok(c, "验证码获取成功", gin.H{
			"captcha_id":     id,
			"pic_path":       b64s,
			"captcha_length": 4,
		})
	}
}
