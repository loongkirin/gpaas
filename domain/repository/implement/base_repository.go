package implement

import (
	core "github.com/loongkirin/gpaas/core"
	"gorm.io/gorm"
)

type BaseRepository struct {
	dbContext core.DbContext
}

func (r *BaseRepository) getDb() (*gorm.DB, *core.AppError) {
	db := r.dbContext.GetDb()
	if db == nil {
		return nil, core.NewUnexpectedError("database not initialized error")
	}
	return db, nil
}
