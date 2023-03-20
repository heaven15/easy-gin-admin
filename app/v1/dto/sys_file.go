package dto

type SysFileInfoReq struct {
	RealName string `json:"real_name" binding:"required,min=1,max=30" label:"文件名称"`
	Path     string `json:"path" binding:"required,min=1,max=300" label:"文件路径"`
	Status   int8   `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string `json:"remark" binding:"omitempty" label:"备注"`
}
