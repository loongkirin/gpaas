package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	router "github.com/loongkirin/gpaas/api/router"
	app "github.com/loongkirin/gpaas/app"
)

func Run() {
	fmt.Println("gpaas server start......")
	app.InitAppContext()

	// fmt.Println(AppContext.APP_DbContext.GetDb() == nil)
	// fmt.Println(AppContext.APP_REDIS == nil)

	fmt.Println("migrate database start......")
	// migrator.MigrateDb(app.AppContext.APP_DbContext)
	fmt.Println("migrate database end......")

	fmt.Println("migrate database start......")
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
	v1Group := Router.Group("v1")
	ginRouterEntry := router.Entry{}
	ginRouterEntry.InitAllRouter(v1Group)

	return Router
}
