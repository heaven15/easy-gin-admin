package model

// SysRole 系统角色
type SysRole struct {
	ID             int64            `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:角色ID" json:"id"`
	RealName       string           `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:角色名称" json:"real_name"`
	Code           string           `gorm:"column:code;index:idx_code;unique;not null;type:varchar(50);comment:角色标识" json:"code"`
	Remark         string           `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Sort           int32            `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	PermissionList []*SysPermission `gorm:"many2many:sys_role_permission" json:"permission_list"`
	Status         int8             `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser     int64            `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime       `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime       `gorm:"comment:更新时间" json:"-"`
}

func (s SysRole) TableName() string {
	return "sys_role"
}
