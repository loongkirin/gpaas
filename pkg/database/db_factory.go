package database

import (
	cfg "github.com/loongkirin/gpaas/pkg/config"
	pg "github.com/loongkirin/gpaas/pkg/database/postgres"
)

func CreateDbContext(cfg cfg.DbConfig) DbContext {
	var dbcontext DbContext
	switch cfg.DbType {
	case "postgres":
		dbcontext = pg.NewPostgresDbContext(&cfg)
	}

	return dbcontext
}
