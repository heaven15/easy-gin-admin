package model

type SysDictionary struct {
	ID         int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:字典ID" json:"id"`
	RealName   string     `gorm:"column:real_name;not null;type:varchar(30);comment:字典名称" json:"real_name"`
	Code       string     `gorm:"column:code;index:idx_code;unique;not null;type:varchar(50);comment:字典标识" json:"code"`
	IsType     string     `gorm:"column:is_type;not null;type:varchar(50);comment:字典类型" json:"is_type"`
	Remark     string     `gorm:"column:remark;not null;type:varchar(1000);comment:备注" json:"remark"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

func (s SysDictionary) TableName() string {
	return "sys_dictionary"
}
