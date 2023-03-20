package model

type SysGroup struct {
	ID         int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:用户组ID" json:"id"`
	ParentID   int64      `gorm:"column:parent_id;type:int;null;comment:父类用户组ID" json:"parent_id"`
	RealName   string     `gorm:"column:real_name;not null;type:varchar(30);comment:用户组名称" json:"real_name"`
	Remark     string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Sort       int32      `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	RoleList   []*SysRole `gorm:"many2many:sys_group_role" json:"role_list"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysGroup) TableName() string {
	return "sys_group"
}
