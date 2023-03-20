package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SysMenu struct {
	ID         int64      `gorm:"column:id;primary_key;auto_increment;primary_key;unique;comment:菜单ID" json:"id"`
	ParentID   int64      `gorm:"column:parent_id;type:int;null;comment:父类菜单ID" json:"parent_id"`
	Children   []*SysMenu `gorm:"foreignKey:parent_id;references:id" json:"children,omitempty"`
	RealName   string     `gorm:"column:real_name;index:idx_real_name;unique;not null;type:varchar(30);comment:菜单名称" json:"real_name"`
	Redirect   string     `gorm:"column:redirect;type:varchar(50);null;comment:重定向" json:"redirect"`
	PageID     int64      `gorm:"column:page_id;type:int;comment:页面ID" json:"page_id"` //页面ID
	PageInfo   SysPage    `gorm:"foreignKey:page_id;references:id" json:"page_info"`
	Meta       MetaJson   `gorm:"column:meta;type:varchar(500);comment:元数据" json:"meta"`
	Sort       int32      `gorm:"column:sort;not null;type:int;comment:排序" json:"sort"`
	Status     int8       `gorm:"column:status;not null;type:tinyint(1);default:1;comment:状态：1启用，2禁用" json:"status"`
	CreateUser int64      `gorm:"column:create_user;not null;default:0;comment:创建用户" json:"create_user"`
	CreatedAt  *LocalTime `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt  *LocalTime `gorm:"comment:更新时间" json:"-"`
}

type MetaJson struct {
	Icon      string `gorm:"comment:菜单图标" json:"icon"`      // 菜单图标
	Active    string `gorm:"comment:菜单高亮" json:"active"`    //菜单高亮
	Color     string `gorm:"comment:颜色" json:"color"`       //菜单颜色
	FullPage  int8   `gorm:"comment:整页路由" json:"full_page"` // 整页路由
	Tag       string `gorm:"comment:标签" json:"tab"`         // 标签
	Component string `gorm:"comment:组件" json:"component"`   //组件
}

func (c MetaJson) Value() (driver.Value, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (c *MetaJson) Scan(value interface{}) error {
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

func (s SysMenu) TableName() string {
	return "sys_menu"
}
