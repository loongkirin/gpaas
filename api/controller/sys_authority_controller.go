package controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/loongkirin/gpaas/api/core"
	dto "github.com/loongkirin/gpaas/api/dto"
	service "github.com/loongkirin/gpaas/service"
	serviceImpl "github.com/loongkirin/gpaas/service/implement"
	util "github.com/loongkirin/gpaas/util"
	"github.com/mojocn/base64Captcha"
)

// var store = base64Captcha.DefaultMemStore

var store = util.NewDefaultCaptchaRedisStore()

type SystemAuthorityController struct {
	userService service.UserService
}

func NewSystemAuthorityController() *SystemAuthorityController {
	userService := serviceImpl.NewUserService()

	return &SystemAuthorityController{
		userService: userService,
	}
}

func (t *SystemAuthorityController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithContext(c))
	if id, b64s, err := cp.Generate(); err != nil {
		response.Fail(c, "验证码获取失败", map[string]interface{}{})
	} else {
		response.Ok(c, "验证码获取成功", dto.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: 4,
		})
	}
}

func (t *SystemAuthorityController) Login(c *gin.Context) {
	var l dto.LoginRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		r, err := t.userService.Login(c, &l)
		if err != nil {
			response.Fail(c, err.Message, map[string]interface{}{})
		}

		response.Ok(c, "登录成功", r)
	} else {
		response.Fail(c, "验证码错误", map[string]interface{}{})
	}
}

func (t *SystemAuthorityController) Register(c *gin.Context) {
	var l dto.RegisterRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
	}

	err := t.userService.Register(c, &l)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "注册成功", map[string]interface{}{})
	}
}

func (t *SystemAuthorityController) RefreshToken(c *gin.Context) {
	var l dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
	}

	r, err := t.userService.RefreshToken(c, &l)
	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "Refresh token success", r)
	}
}
