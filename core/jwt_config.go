package core

type JWTConfig struct {
	SecretKey   string `mapstructure:"scret_key" json:"scret_key" yaml:"scret_key"`
	ExpiresTime int64  `mapstructure:"expires_time" json:"expires_time" yaml:"expires_time"`
	BufferTime  int64  `mapstructure:"buffer_time" json:"buffer_time" yaml:"buffer_time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
