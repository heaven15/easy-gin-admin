package model

type SysFile struct {
	ID             int64            `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:文件ID" json:"id"`
	RealName       string           `gorm:"column:real_name;not null;type:varchar(30);comment:文件名称" json:"real_name"`
	Path           string           `gorm:"column:path;not null;type:varchar(300);comment:文件路径" json:"path"`
	PermissionList []*SysPermission `gorm:"many2many:sys_file_permission" json:"permission_list"`
	Remark         string           `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status         int8             `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser     int64            `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime       `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime       `gorm:"comment:更新时间" json:"-"`
}

func (s SysFile) TableName() string {
	return "sys_file"
}
