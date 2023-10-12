package model

import (
	core "github.com/loongkirin/gpaas/core"
)

type ObjectDescribe struct {
	core.DbBaseModel
	ApiName        string `json:"api_name" gorm:"size:100;not null;unique"`
	Description    string `json:"mobile" gorm:"size:1000"`
	StoreTableName string `json:"store_table_name" gorm:"size:100;not null"`
	DefineType     string `json:"defint_type" gorm:"size:100"`
}

func (entity *ObjectDescribe) TableName() string {
	return "gpaas_sys_object_describe"
}
