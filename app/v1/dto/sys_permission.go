package dto

type SysPermissionInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"字典名称"`
	Code     string `json:"code" binding:"required,min=1,max=50,alpha" label:"字典标识"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}
