package core

import (
	"gorm.io/gorm"
)

type DbContext interface {
	DSN() string
	GetDb() *gorm.DB
}
