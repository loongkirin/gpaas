package model

import (
	db "github.com/loongkirin/gpaas/pkg/database"
)

type Tenant struct {
	db.DbBaseModel
	Name     string `json:"name" gorm:"size:500;not null;unique"`
	Tel      string `json:"tel" gorm:"size:100"`
	PostCode string `json:"post_code" gorm:"size:100"`
	Address  string `json:"address" gorm:"size:1000"`
	Email    string `json:"email" gorm:"size:100"`
	Status   string `json:"status" gorm:"size:100"`
}

func (entity *Tenant) TableName() string {
	return "gpaas_sys_tenant"
}
