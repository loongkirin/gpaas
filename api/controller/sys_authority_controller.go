package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/loongkirin/gpaas/api/core"
	dto "github.com/loongkirin/gpaas/api/dto"
	app "github.com/loongkirin/gpaas/app"
	oauth "github.com/loongkirin/gpaas/domain/oauth"
	repoImpl "github.com/loongkirin/gpaas/domain/repository/implement"
	service "github.com/loongkirin/gpaas/service"
	serviceImpl "github.com/loongkirin/gpaas/service/implement"
	util "github.com/loongkirin/gpaas/util"
	"github.com/mojocn/base64Captcha"
)

// var store = base64Captcha.DefaultMemStore

var store = util.NewDefaultCaptchaRedisStore()

type SystemAuthorityController struct {
	userService service.UserService
	oauthMaker  oauth.OAuthMaker
}

func NewSystemAuthorityController() *SystemAuthorityController {
	userReop := repoImpl.NewUserRepository(app.AppContext.APP_DbContext)
	userService := serviceImpl.NewUserService(userReop)
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("NewSystemAuthorityController error")
	}
	return &SystemAuthorityController{
		userService: userService,
		oauthMaker:  oauthMaker,
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
	fmt.Printf(l.Mobile)

	// r, err := t.userService.Login(&l)
	// if err != nil {
	// 	response.Fail(c, err.Message, map[string]interface{}{})
	// } else {
	// 	t.TokenNext(c, r)
	// }

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		r, err := t.userService.Login(&l)
		if err != nil {
			response.Fail(c, err.Message, map[string]interface{}{})
		} else {
			t.TokenNext(c, r)
		}
	} else {
		response.Fail(c, "验证码错误", map[string]interface{}{})
	}
}

func (t *SystemAuthorityController) TokenNext(c *gin.Context, r *dto.LoginResponse) {
	token, _, err := t.oauthMaker.GenerateAccessToken(r.Mobile, r.UserName)
	if err != nil {
		response.Fail(c, "获取access token失败", map[string]interface{}{})
		return
	}
	r.AccessToken = token
	token, _, err = t.oauthMaker.GenerateRefreshToken(r.Mobile, r.UserName)
	if err != nil {
		response.Fail(c, "获取refresh token失败", map[string]interface{}{})
		return
	}
	r.RefreshToken = token
	fmt.Printf("mobile:", r.Mobile, "UserName:", r.UserName, "accessToken:", r.AccessToken, "refreshToken:", r.RefreshToken)
	response.Ok(c, "登录成功", r)
}

func (t *SystemAuthorityController) Register(c *gin.Context) {
	var l dto.RegisterRequest
	_ = c.ShouldBindJSON(&l)

	err := t.userService.Register(&l)

	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "注册成功", map[string]interface{}{})
	}
}

func (t *SystemAuthorityController) RefreshToken(c *gin.Context) {
	var l dto.RefreshTokenRequest
	_ = c.ShouldBindJSON(&l)

	token := ""

	// j := util.NewJWTUtil() // 唯一签名
	// token, err := j.RefreshToken(l.AccessToken)
	// if err != nil {
	// 	response.Unauthorized(c, err.Error(), map[string]interface{}{})
	// 	return
	// }
	r := &dto.RefreshTokenRequest{
		RefreshToken: token,
	}
	response.Ok(c, "Refresh token success", r)
}
