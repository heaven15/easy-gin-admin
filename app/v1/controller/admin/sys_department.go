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

type SysDepartmentController struct {
	SysDepartmentService service.SysDepartmentService
}

// Create 创建部门
func (s *SysDepartmentController) Create(ctx *gin.Context) {
	request := dto.SysDepartmentInfoReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysDepartment{
		ParentID:   request.ParentID,
		RealName:   request.RealName,
		FullName:   request.FullName,
		IsType:     request.IsType,
		Remark:     request.Remark,
		Sort:       request.Sort,
		Status:     request.Status,
		CreateUser: 0,
	}
	if err := s.SysDepartmentService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新部门
func (s *SysDepartmentController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysDepartmentInfoReq{}
	if err = ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysDepartment{
		ID:         i,
		ParentID:   request.ParentID,
		RealName:   request.RealName,
		FullName:   request.FullName,
		IsType:     request.IsType,
		Remark:     request.Remark,
		Sort:       request.Sort,
		Status:     request.Status,
		CreateUser: 0,
	}
	if err = s.SysDepartmentService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Delete 删除部门
func (s *SysDepartmentController) Delete(ctx *gin.Context) {
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
	if err := s.SysDepartmentService.Delete(request.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Join 加入部门
func (s *SysDepartmentController) Join(ctx *gin.Context) {
	request := dto.SysUserJoinDepartmentReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if err := s.SysDepartmentService.Join(&request); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 部门列表
func (s *SysDepartmentController) List(ctx *gin.Context) {
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
	data, err := s.SysDepartmentService.List(&request)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
