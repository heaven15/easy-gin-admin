package model

type SysOperation struct {
	ID             int64            `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:功能操作ID" json:"id"`
	ParentID       int64            `gorm:"column:parent_id;type:int;null;comment:父类功能操作ID" json:"parent_id"`
	RealName       string           `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:功能操作名称" json:"real_name"`
	Code           string           `gorm:"column:code;index:idx_code;unique;not null;type:varchar(50);comment:功能操作标识" json:"code"`
	Url            string           `gorm:"column:url;not null;type:varchar(300);comment:拦截URL前缀" json:"url"`
	PermissionList []*SysPermission `gorm:"many2many:sys_operation_permission" json:"permission_list"`
	Remark         string           `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status         int8             `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser     int64            `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime       `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime       `gorm:"comment:更新时间" json:"-"`
}

func (s SysOperation) TableName() string {
	return "sys_operation"
}
