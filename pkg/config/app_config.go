package config

import (
	"github.com/loongkirin/gdk/cache"
	db "github.com/loongkirin/gdk/database"
	"github.com/loongkirin/gdk/oauth"
)

type AppConfig struct {
	OAuthConfig oauth.OAuthConfig `mapstructure:"oauthconfig" json:"oauthconfig" yaml:"oauthconfig"`
	RedisConfig cache.RedisConfig `mapstructure:"redisconfig" json:"redisconfig" yaml:"redisconfig"`
	DbConfig    db.DbConfig       `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
}
