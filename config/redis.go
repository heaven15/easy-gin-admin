package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DataBase int    `mapstructure:"database" json:"database"`
}
