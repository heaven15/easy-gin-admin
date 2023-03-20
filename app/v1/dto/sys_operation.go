package dto

type SysOperationInfoReq struct {
	ParentID int64  `json:"parent_id" binding:"omitempty,min=1" label:"父类功能操作ID"`
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"功能操作名称"`
	Code     string `json:"code" binding:"required,min=1,max=50,alpha" label:"功能操作标识"`
	Url      string `json:"url" binding:"omitempty" label:"拦截URL前缀"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}
