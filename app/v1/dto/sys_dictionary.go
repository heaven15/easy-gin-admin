package dto

type SysDictionaryInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"字典名称"`
	Code     string `json:"code" binding:"required,min=1,max=50,alpha" label:"字典标识"`
	IsType   string `json:"is_type,default=0" binding:"omitempty" label:"字典类型"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
}
