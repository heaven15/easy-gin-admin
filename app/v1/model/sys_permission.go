package model

// SysPermission 系统数据权限
type SysPermission struct {
	ID         int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:权限ID" json:"id"`
	RealName   string     `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:权限名称" json:"real_name"`
	Code       string     `gorm:"column:code;index:idx_code;unique;not null;type:varchar(50);comment:权限标识" json:"code"`
	Remark     string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysPermission) TableName() string {
	return "sys_permission"
}
