package dto

type SysRoleInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"角色名称"`
	Code     string `json:"code" binding:"required,min=1,max=50,alpha" label:"角色标识"`
	Sort     int32  `json:"sort" binding:"omitempty,numeric" label:"排序"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}

type SysRolePermissionReq struct {
	RoleID         int64 `json:"role_id" binding:"required,min=1" label:"角色ID"`
	PermissionList []int `json:"permission_list" binding:"required" label:"权限数组ID"`
}
