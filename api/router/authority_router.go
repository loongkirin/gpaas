package router

import (
	"github.com/gin-gonic/gin"
	controller "github.com/loongkirin/gpaas/api/controller"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := router.Group("auth")
	authApi := controller.NewAuthorityController()
	authRouter.GET("captcha", authApi.Captcha)

	return authRouter
}
