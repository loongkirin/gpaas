package migrator

import (
	"fmt"

	core "github.com/loongkirin/gpaas/core"
	model "github.com/loongkirin/gpaas/domain/model"
)

func MigrateDb(dbContext core.DbContext) {
	db := dbContext.GetDb()
	if db == nil {
		return
	}
	fmt.Println("migrate db start......")
	db.AutoMigrate(&model.User{})
	fmt.Println("migrate db end.....")
}
