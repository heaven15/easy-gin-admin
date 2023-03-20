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

type SysRankController struct {
	SysRankService service.SysRankService
}

// Create 创建职级
func (s *SysRankController) Create(ctx *gin.Context) {
	request := dto.SysRankInfoReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysRank{
		RealName: request.RealName,
		Sort:     request.Sort,
		Status:   request.Status,
		Remark:   request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err := s.SysRankService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新职级
func (s *SysRankController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysRankInfoReq{}
	if err = ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysRank{
		ID:       i,
		RealName: request.RealName,
		Sort:     request.Sort,
		Status:   request.Status,
		Remark:   request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err = s.SysRankService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Delete 删除职级
func (s *SysRankController) Delete(ctx *gin.Context) {
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
	if err := s.SysRankService.Delete(request.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 职级列表
func (s *SysRankController) List(ctx *gin.Context) {
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
	data, err := s.SysRankService.List(&request)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
