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

type SysOperationController struct {
	SysOperationService service.SysOperationService
}

// Create 创建功能操作
func (s *SysOperationController) Create(ctx *gin.Context) {
	request := dto.SysOperationInfoReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysOperation{
		ParentID:   request.ParentID,
		RealName:   request.RealName,
		Code:       request.Code,
		Url:        request.Url,
		Status:     request.Status,
		Remark:     request.Remark,
		CreateUser: 0,
	}
	if err := s.SysOperationService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新功能操作
func (s *SysOperationController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysOperationInfoReq{}
	if err = ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysOperation{
		ID:         i,
		ParentID:   request.ParentID,
		RealName:   request.RealName,
		Code:       request.Code,
		Url:        request.Url,
		Status:     request.Status,
		Remark:     request.Remark,
		CreateUser: 0,
	}
	if err = s.SysOperationService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Delete 删除功能操作
func (s *SysOperationController) Delete(ctx *gin.Context) {
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
	if err := s.SysOperationService.Delete(request.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 功能操作列表
func (s *SysOperationController) List(ctx *gin.Context) {
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
	data, err := s.SysOperationService.List(&request)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
