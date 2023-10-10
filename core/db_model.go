package core

type DbBaseModel struct {
	Id           string `gorm:"primaryKey" json:"id"`
	Create_Time  int64  `gorm:"autoCreateTime:nano" json:"create_time"`
	Update_Time  int64  `gorm:"autoUpdateTime:nano" json:"update_time"`
	Data_Version int    `json:"data_version"`
	Data_Status  int    `json:"data_status"`
}
