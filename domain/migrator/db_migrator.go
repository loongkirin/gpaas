package migrator

import (
	"fmt"

	db "github.com/loongkirin/gdk/database/gorm"
	model "github.com/loongkirin/gpaas/domain/model"
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
