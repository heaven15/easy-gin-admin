package config

type GormDB struct {
	Host        string `mapstructure:"host" json:"host"`                      // 服务器地址:端口
	Port        string `mapstructure:"port" json:"port"`                      //:端口
	Config      string `mapstructure:"config" json:"config"`                  // 高级配置
	DataBase    string `mapstructure:"database" json:"database"`              // 数据库名
	UserName    string `mapstructure:"username" json:"username"`              // 数据库用户名
	PassWord    string `mapstructure:"password" json:"password"`              // 数据库密码
	Prefix      string `mapstructure:"prefix" json:"prefix"`                  //全局表前缀，单独定义TableName则不生效
	Singular    bool   `mapstructure:"singular" json:"singular"`              //是否开启全局禁用复数，true表示开启
	Engine      string `mapstructure:"engine" json:"engine" default:"InnoDB"` //数据库引擎，默认InnoDB
	MaxIdleCons int    `mapstructure:"max_idle_cons" json:"max_idle_cons"`    // 空闲中的最大连接数
	MaxOpenCons int    `mapstructure:"max_open_cons" json:"max_open_cons"`    // 打开到数据库的最大连接数
	LogMode     string `mapstructure:"log_mode" json:"log_mode"`              // 是否开启Gorm全局日志
	LogZap      bool   `mapstructure:"log_zap" json:"log_zap"`                // 是否通过zap写入日志文件
}
