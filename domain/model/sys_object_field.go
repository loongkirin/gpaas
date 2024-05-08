package model

import (
	db "github.com/loongkirin/gpaas/pkg/database"
)

type ObjectField struct {
	db.DbBaseModel
	FieldApiName   string `json:"field_api_name" gorm:"size:100;not null"`
	ObjectApiName  string `json:"object_api_name" gorm:"size:100;not null"`
	FieldType      string `json:"field_type" gorm:"size:32;not null"`
	FieldLabel     string `json:"field_label" gorm:"size:200;not null"`
	Description    string `json:"description" gorm:"size:1000"`
	HelpText       string `json:"help_text" gorm:"size:1000"`
	DefineType     string `json:"define_type" gorm:"size:32"`
	FieldLength    int    `json:"field_length"`
	FieldMaxLength int    `json:"field_max_length"`
	IsUserDefine   bool   `json:"is_user_define"`
	IsRequired     bool   `json:"is_required"`
	IsUnique       bool   `json:"is_unique"`
	IsActive       bool   `json:"is_active"`
	DefaultValue   string `json:"default_value" gorm:"size:100"`
}

func (entity *ObjectField) TableName() string {
	return "gpaas_sys_object_field"
}
