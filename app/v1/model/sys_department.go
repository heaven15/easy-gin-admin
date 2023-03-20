package model

// SysDepartment 系统部门表
type SysDepartment struct {
	ID         int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:部门ID" json:"id"`
	ParentID   int64      `gorm:"column:parent_id;type:int;null;comment:父类部门ID" json:"parent_id"`
	RealName   string     `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:部门名称" json:"real_name"`
	FullName   string     `gorm:"column:full_name;not null;type:varchar(50);comment:部门全称" json:"full_name"`
	IsType     int        `gorm:"column:is_type;not null;type:tinyint(1);default:0;comment:部门类型：1公司，2子公司，3.部门，4.小组" json:"is_type"`
	Remark     string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Sort       int32      `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	UserList   []*SysUser `gorm:"many2many:sys_department_user" json:"user_list"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysDepartment) TableName() string {
	return "sys_department"
}
