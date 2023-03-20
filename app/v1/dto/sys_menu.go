package dto

type SysMenuInfoReq struct {
	ParentID  int64           `json:"parent_id" binding:"omitempty,min=1" label:"父类的菜单ID"`
	RealName  string          `json:"real_name" binding:"required,min=1,max=30" label:"菜单名称"`
	Redirect  string          `json:"redirect" binding:"omitempty,min=1,max=30" label:"重定向"`
	Component string          `json:"component" binding:"omitempty,min=1,max=50" label:"组件"`
	Path      string          `json:"path" binding:"omitempty,min=1,max=250" label:"路径"`
	Meta      MetaJsonInfoReq `json:"meta" binding:"required" label:"元数据"`
	Sort      int32           `json:"sort" binding:"omitempty,numeric" label:"排序"`
	Status    int8            `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
}

type MetaJsonInfoReq struct {
	Icon      string `json:"icon" binding:"required" label:"菜单图标"`       // 菜单图标
	Active    string `json:"active" binding:"omitempty" label:"菜单高亮"`    //菜单高亮
	Color     string `json:"color" binding:"omitempty" label:"颜色"`       //菜单颜色
	FullPage  int8   `json:"full_page" binding:"omitempty" label:"整页路由"` // 整页路由
	Tag       string `json:"tab" binding:"omitempty" label:"标签"`         // 标签
	Component string `json:"component" binding:"omitempty" label:"组件"`   //组件
}
