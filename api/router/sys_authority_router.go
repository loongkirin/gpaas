package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/loongkirin/gpaas/api/controller"
)

type SystemAuthorityRouter struct{}

func (s *SystemAuthorityRouter) InitSystemAuthorityRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := router.Group("auth")
	authApi := controller.NewSystemAuthorityController()
	authRouter.GET("captcha", authApi.Captcha)
	authRouter.POST("login", authApi.Login)
	authRouter.POST("register", authApi.Register)
	authRouter.POST("refresh_token", authApi.RefreshToken)

	return authRouter
}
