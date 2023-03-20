package admin

import (
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/service"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type SysPostController struct {
	SysPostService service.SysPostService
}

// Create 创建岗位
func (s *SysPostController) Create(ctx *gin.Context) {
	request := dto.SysPostInfoReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysPost{
		RealName: request.RealName,
		Sort:     request.Sort,
		Status:   request.Status,
		Remark:   request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err := s.SysPostService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新岗位
func (s *SysPostController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysPostInfoReq{}
	if err = ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysPost{
		ID:       i,
		RealName: request.RealName,
		Sort:     request.Sort,
		Status:   request.Status,
		Remark:   request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err = s.SysPostService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Delete 删除岗位
func (s *SysPostController) Delete(ctx *gin.Context) {
	request := dto.IdsReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if err := s.SysPostService.Delete(request.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 岗位列表
func (s *SysPostController) List(ctx *gin.Context) {
	request := dto.PageReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	data, err := s.SysPostService.List(&request)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
