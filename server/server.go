package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("gpaas server start......")
	InitAppContext()

	// fmt.Println(AppContext.APP_DbContext.GetDb() == nil)
	// fmt.Println(AppContext.APP_REDIS == nil)

	router := initializeRouter()
	router.Run(":8081")
}

// 初始化gin总路由
func initializeRouter() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return Router
}
