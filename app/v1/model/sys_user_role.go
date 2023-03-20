package model

type SysUserRole struct {
	SysUserID int64 `gorm:"column:sys_user_id;not null;comment:用户ID" json:"sys_user_id"`
	SysRoleID int64 `gorm:"column:sys_role_id;not null;comment:角色ID" json:"sys_role_id"`
}
