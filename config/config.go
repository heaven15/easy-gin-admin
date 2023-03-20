package config

type Config struct {
	System        System        `mapstructure:"System" json:"System"`
	GormDB        GormDB        `mapstructure:"GormDB" json:"GormDB"`
	Redis         Redis         `mapstructure:"Redis" json:"Redis"`
	Elasticsearch Elasticsearch `mapstructure:"Elasticsearch" json:"Elasticsearch"`
	JWT           JWT           `mapstructure:"JWT" json:"JWT"`
	Zap           Zap           `mapstructure:"Zap" json:"Zap"`
}
