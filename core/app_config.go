package core

type AppConfig struct {
	OAuthConfig OAuthConfig `mapstructure:"oauthconfig" json:"oauthconfig" yaml:"oauthconfig"`
	RedisConfig RedisConfig `mapstructure:"redisconfig" json:"redisconfig" yaml:"redisconfig"`
	DbConfig    DbConfig    `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
}
