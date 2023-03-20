package model

type SysConfig struct {
	ID             int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:配置ID" json:"id"`
	RealName       string     `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:配置名称" json:"real_name"`
	AppKey         string     `gorm:"column:app_key;index:idx_app_key;unique;not null;type:varchar(30);comment:配置标识" json:"app_key"`
	AppVal         string     `gorm:"column:app_val;not null;type:longtext;comment:配置值" json:"app_val"`
	PermissionCode string     `gorm:"column:permission_code;not null;type:varchar(30);comment:权限标识"`
	Remark         string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	CreateUser     int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysConfig) TableName() string {
	return "sys_config"
}
