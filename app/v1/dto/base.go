package dto

type IdReq struct {
}

type IdsReq struct {
	Ids []int `json:"ids" binding:"required"`
}

type PageReq struct {
	Page     int    `form:"page" binding:"omitempty,numeric" label:"页码"`
	PageSize int    `form:"page_size" binding:"omitempty,numeric" label:"每页数量"`
	Keyword  string `form:"keyword" binding:"omitempty" label:"关键字"`
	Order    string `form:"order" binding:"omitempty" label:"字段"`
	Sort     string `form:"sort" binding:"omitempty" label:"排序"`
	Status   int8   `form:"status" binding:"omitempty,oneof=1 2" label:"状态"`
}
