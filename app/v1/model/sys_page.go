package model

type SysPage struct {
	ID             int64          `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:页面ID" json:"id"`
	RealName       string         `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:页面名称" json:"real_name"`
	Url            string         `gorm:"column:url;type:varchar(1000);null;comment:文件路径" json:"url"`
	Sort           int32          `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	IsType         int8           `gorm:"column:is_type;not null;type:int;comment:页面类型：1.菜单 2.iframe 3.外链 4.按钮" json:"is_type"` //菜单类型
	Remark         string         `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	PermissionID   int64          `gorm:"column:permission_id;type:int;comment:权限ID" json:"permission_id"`
	PermissionInfo *SysPermission `gorm:"foreignKey:permission_id;references:id" json:"permission_info"`
	Status         int8           `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser     int64          `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt      *LocalTime     `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt      *LocalTime     `gorm:"comment:更新时间" json:"-"`
}

func (s SysPage) TableName() string {
	return "sys_page"
}
