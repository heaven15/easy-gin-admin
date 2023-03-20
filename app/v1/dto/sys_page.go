package dto

type SysPageInfoReq struct {
	RealName     string `json:"real_name" binding:"required,min=1,max=30" label:"页面名称"`
	Url          string `json:"url" binding:"required" label:"文件路径"`
	Sort         int32  `json:"sort" binding:"omitempty,numeric" label:"排序"`
	IsType       int8   `json:"is_type,default=0" binding:"omitempty,oneof=1 2 3 4" label:"页面类型"`
	Status       int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark       string `json:"remark" binding:"omitempty" label:"备注"`
	PermissionID int64  `json:"permission_id" binding:"required,numeric" label:"权限ID"`
}
