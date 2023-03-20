package dto

type SysPostInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"岗位名称"`
	Sort     int32  `json:"sort" binding:"omitempty,numeric" label:"排序"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}
