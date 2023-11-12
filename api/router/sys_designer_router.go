package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/loongkirin/gpaas/api/middleware"
	app "github.com/loongkirin/gpaas/app"
	oauth "github.com/loongkirin/gpaas/domain/oauth"
)

type SystemDesignerRouter struct{}

func (s *SystemDesignerRouter) InitSystemDesignerRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	designerRouter := router.Group("desinger")
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)
	if err != nil {
		panic("InitSystemDesignerRouter error")
	}

	designerRouter.Use(middleware.OAuthMiddleware(oauthMaker))
	designerRouter.GET("/desinger", func(c *gin.Context) {
		c.JSON(http.StatusOK, "disinger ok")
	})
	return designerRouter
}
