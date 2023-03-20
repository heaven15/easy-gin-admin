package model

type SysElement struct {
	ID             int64            `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:页面元素ID" json:"id"`
	RealName       string           `gorm:"column:real_name;not null;type:varchar(30);comment:页面名称" json:"real_name"`
	Code           string           `gorm:"column:code;index:idx_code;unique;not null;type:varchar(50);comment:页面元素标识" json:"code"`
	PageID         int64            `gorm:"column:page_id;type:int;comment:页面ID" json:"page_id"` //页面ID
	PageInfo       SysPage          `gorm:"foreignKey:page_id;references:id" json:"page_info"`
	PermissionList []*SysPermission `gorm:"many2many:sys_element_permissions" json:"permission_list"`
	Remark         string           `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status         int8             `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser     int64            `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime       `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime       `gorm:"comment:更新时间" json:"-"`
}

func (s SysElement) TableName() string {
	return "sys_element"
}
