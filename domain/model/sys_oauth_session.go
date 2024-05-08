package model

import (
	db "github.com/loongkirin/gpaas/pkg/database"
)

type OAuthSession struct {
	db.DbBaseModel
	UserId       string `json:"user_id" gorm:"size:32"`
	Mobile       string `json:"mobile" gorm:"size:100"`
	UserName     string `json:"user_name"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIp     string `json:"client_ip"`
	IsBlocked    bool   `json:"is_blocked"`
	ExpiredAt    int64  `json:"expired_at"`
}

func (entity *OAuthSession) TableName() string {
	return "gpaas_sys_oauth_session"
}
