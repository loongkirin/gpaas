package core

type AppConfig struct {
	JWTConfig   JWTConfig   `mapstructure:"jwtconfig" json:"jwtconfig" yaml:"jwtconfig"`
	RedisConfig RedisConfig `mapstructure:"redisconfig" json:"redisconfig" yaml:"redisconfig"`
	DbConfig    DbConfig    `mapstructure:"dbconfig" json:"dbconfig" yaml:"dbconfig"`
}
