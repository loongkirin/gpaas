package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// migrator "github.com/loongkirin/gpaas/domain/migrator"

	repo "github.com/loongkirin/gpaas/domain/repository"
	repoImpl "github.com/loongkirin/gpaas/domain/repository/implement"
)

func Run() {
	fmt.Println("gpaas server start......")
	InitAppContext()

	// fmt.Println(AppContext.APP_DbContext.GetDb() == nil)
	// fmt.Println(AppContext.APP_REDIS == nil)

	// router := initializeRouter()
	// router.Run(":8081")

	// migrator.MigrateDb(AppContext.APP_DbContext)

	var userRepo repo.UserRepository
	userRepo = repoImpl.NewUserRepository(AppContext.APP_DbContext)
	user, err := userRepo.FindById("1")
	if err != nil {
		fmt.Printf(err.Message)
	} else {
		fmt.Printf(user.Id)
	}
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
