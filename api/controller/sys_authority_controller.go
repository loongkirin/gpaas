package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	response "github.com/loongkirin/gpaas/api/core"
	dto "github.com/loongkirin/gpaas/api/dto"
	app "github.com/loongkirin/gpaas/app"
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
}

func NewSystemAuthorityController() *SystemAuthorityController {
	userReop := repoImpl.NewUserRepository(app.AppContext.APP_DbContext)
	userService := serviceImpl.NewUserService(userReop)
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
	_ = c.ShouldBindJSON(&l)
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
	j := util.NewJWTUtil() // 唯一签名
	claims := j.CreateClaims(r.Mobile, r.UserName)
	token, err := j.GenerateToken(claims)
	if err != nil {
		response.Fail(c, "获取token失败", map[string]interface{}{})
		return
	}
	r.AccessToken = token
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
	var l dto.RefreshToken
	_ = c.ShouldBindJSON(&l)

	j := util.NewJWTUtil() // 唯一签名
	token, err := j.RefreshToken(l.AccessToken)
	if err != nil {
		response.Unauthorized(c, err.Error(), map[string]interface{}{})
		return
	}
	r := &dto.RefreshToken{
		AccessToken: token,
	}
	response.Ok(c, "Refresh token success", r)
}
