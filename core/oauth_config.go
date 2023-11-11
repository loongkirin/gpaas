package core

type OAuthConfig struct {
	SecretKey   string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	ExpiresTime int64  `mapstructure:"expires_time" json:"expires_time" yaml:"expires_time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
