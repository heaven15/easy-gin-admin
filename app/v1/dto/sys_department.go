package dto

type SysDepartmentInfoReq struct {
	ParentID int64  `json:"parent_id" binding:"omitempty,min=1" label:"父类部门ID"`
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"部门名称"`
	FullName string `json:"full_name" binding:"required,min=1,max=50" label:"部门全称"`
	IsType   int    `json:"is_type,default=0" binding:"omitempty,oneof=1 2 3 4" label:"部门类型"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
	Sort     int32  `json:"sort" binding:"omitempty,numeric" label:"排序"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
}

type SysUserJoinDepartmentReq struct {
	DepartmentID int64 `json:"department_id" binding:"required,min=1,numeric" label:"部门ID"`
	UserList     []int `json:"user_list" binding:"required" label:"用户数组ID"`
}
