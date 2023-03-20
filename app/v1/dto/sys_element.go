package dto

type SysElementInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"页面名称"`
	Code     string `json:"code" binding:"required,min=1,max=50,alpha" label:"页面元素标识"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}
