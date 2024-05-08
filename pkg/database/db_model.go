package database

type DbBaseModel struct {
	Id          string `json:"id" gorm:"primaryKey;size:32"`
	TenantId    string `json:"tenant_id" gorm:"size:32"`
	DataVersion int64  `json:"data_version"`
	DataStatus  int    `json:"data_status"`
	CreateTime  int64  `json:"create_time" gorm:"autoCreateTime:milli"`
	UpdateTime  int64  `json:"update_time" gorm:"autoUpdateTime:milli"`
}

func NewDbBaseModel(tenantId string, id string) DbBaseModel {
	return DbBaseModel{
		Id:          id,
		TenantId:    tenantId,
		DataVersion: 1,
		DataStatus:  1,
	}
}
