package config

type JWT struct {
	SecretKey   string `mapstructure:"secret_key" json:"secret_key"`
	ExpiresTime string `mapstructure:"expires_time" json:"expires_time"`
	BufferTime  string `mapstructure:"buffer_time" json:"buffer_time"`
	IsSuer      string `mapstructure:"issuer" json:"issuer"`
}
