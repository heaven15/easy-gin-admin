package model

// SysRank 系统职级表
type SysRank struct {
	ID         int64      `gorm:"primary_key;auto_increment;primary_key;unique;comment:职级ID" json:"id"`
	RealName   string     `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:职级名称" json:"real_name"`
	Sort       int32      `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	Remark     string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysRank) TableName() string {
	return "sys_rank"
}
