package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SysUser struct {
	ID          int64      `gorm:"primary_key;auto_increment;primary_key;unique;comment:用户ID" json:"id"`
	AccountName string     `gorm:"column:account_name;index:idx_account_name;unique;not null;type:varchar(50);comment:账号名称" json:"account_name"`
	UserName    string     `gorm:"column:username;index:idx_username;unique;not null;type:varchar(50);comment:用户姓名" json:"username"`
	NickName    string     `gorm:"column:nickname;not null;type:varchar(50);comment:用户昵称" json:"nickname"`
	PassWord    string     `gorm:"column:password;not null;type:varchar(150);comment:密码" json:"-"` //bcrypt加密
	Gender      int        `gorm:"column:gender;not null;type:tinyint(1);default:1;comment:性别：1男，2女，3保密" json:"gender"`
	Birthday    string     `gorm:"column:birthday;not null;type:varchar(30);comment:出生日期" json:"birthday"`
	Email       string     `gorm:"column:email;not null;type:varchar(50);comment:邮箱" json:"email"`
	Mobile      string     `gorm:"column:mobile;index:idx_mobile;not null;type:varchar(11);comment:用户手机号" json:"mobile"`
	Salt        string     `gorm:"column:salt;not null;type:varchar(20);comment:干扰码" json:"salt"`
	PostID      int64      `gorm:"column:post_id;not null;comment:岗位ID" json:"post_id"`
	PostInfo    SysPost    `gorm:"foreignKey:post_id;references:id" json:"post_info"`
	RankID      int64      `gorm:"column:rank_id;not null;comment:职级ID" json:"rank_id"`
	RankInfo    SysRank    `gorm:"foreignKey:rank_id;references:id" json:"rank_info"`
	GroupID     int64      `gorm:"column:group_id;not null;comment:用户分组ID" json:"group_id"`
	GroupInfo   SysGroup   `gorm:"foreignKey:group_id;references:id" json:"group_info"`
	Avatar      string     `gorm:"column:avatar;not null;type:varchar(1000);comment:头像" json:"avatar"`
	City        CityJson   `gorm:"column:city;type:varchar(250);comment:城市json数据" json:"city"`
	Address     string     `gorm:"column:address;not null;type:varchar(150);comment:地址详情" json:"address"`
	Status      int        `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	Ip          string     `gorm:"column:ip;not null;type:varchar(20);comment:ip地址" json:"ip"`
	Remark      string     `gorm:"column:remark;not null;type:varchar(60);comment:备注" json:"remark"`
	RoleList    []*SysRole `gorm:"many2many:sys_user_role" json:"role_list"`
	CreateUser  int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt   *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt   *LocalTime `gorm:"comment:更新时间" json:"-"`
}

type CityJson struct {
	Province string `json:"province" gorm:"comment:省份"` // 省
	City     string `json:"city" gorm:"comment:市级"`     // 市
	Area     string `json:"area" gorm:"comment:县区"`     // 区
}

func (c CityJson) Value() (driver.Value, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (c *CityJson) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("不匹配")
	}
	err := json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}

func (s SysUser) TableName() string {
	return "sys_user"
}
