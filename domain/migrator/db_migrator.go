package migrator

import (
	"fmt"

	model "github.com/loongkirin/gpaas/domain/model"
	db "github.com/loongkirin/gpaas/pkg/database"
)

func MigrateDb(dbContext db.DbContext) {
	db := dbContext.GetDb()
	if db == nil {
		return
	}
	fmt.Println("migrate db start......")
	// db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.OAuthSession{})
	db.AutoMigrate(&model.Tenant{})
	fmt.Println("migrate db end.....")
}
