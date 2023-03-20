package config

type System struct {
	Port   int    `mapstructure:"port" json:"port"`
	DbType string `mapstructure:"db_type" json:"db_type"`
}
