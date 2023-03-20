package dto

type SysConfigInfoReq struct {
	RealName       string `json:"real_name" binding:"required,min=1,max=30" label:"配置名称"`
	AppKey         string `json:"app_key" binding:"required,alpha,min=1,max=30" label:"配置标识"`
	AppVal         string `json:"app_val" binding:"required" label:"配置值"`
	PermissionCode string `json:"permission_code" binding:"required,alpha,min=1,max=30" label:"权限标识"`
	Remark         string `json:"remark" binding:"omitempty" label:"备注"`
}
