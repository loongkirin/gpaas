package model

import (
	db "github.com/loongkirin/gpaas/pkg/database"
)

type ObjectDescribe struct {
	db.DbBaseModel
	Module         string `json:"module" gorm:"size:32"`
	ObjectApiName  string `json:"object_api_name" gorm:"size:100;not null"`
	ObjectName     string `json:"object_name" gorm:"size:200;not null"`
	Description    string `json:"description" gorm:"size:1000"`
	StoreTableName string `json:"store_table_name" gorm:"size:100;not null"`
	DefineType     string `json:"define_type" gorm:"size:32"`
	PrimaryApiName string `json:"primary_api_name" gorm:"size:100"`
	IsUserDefine   bool   `json:"is_user_define"`
	IsActive       bool   `json:"is_active"`
	Remark         string `json:"remark" gorm:"size:1000"`
}

func (entity *ObjectDescribe) TableName() string {
	return "gpaas_sys_object_describe"
}
