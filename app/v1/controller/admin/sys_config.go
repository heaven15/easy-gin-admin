package admin

import (
	dto "github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/service"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type SysConfigController struct {
	SysConfigService service.SysConfigService
}

// Create 创建系统配置
func (s *SysConfigController) Create(ctx *gin.Context) {
	request := dto.SysConfigInfoReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysConfig{
		RealName:       request.RealName,
		AppKey:         request.AppKey,
		AppVal:         request.AppVal,
		PermissionCode: request.PermissionCode,
		Remark:         request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err := s.SysConfigService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新系统配置
func (s *SysConfigController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysConfigInfoReq{}
	if err = ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysConfig{
		ID:             i,
		RealName:       request.RealName,
		AppKey:         request.AppKey,
		AppVal:         request.AppVal,
		PermissionCode: request.PermissionCode,
		Remark:         request.Remark,
		//获取登录时的用户id
		CreateUser: 0,
	}
	if err = s.SysConfigService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Config 根据Key获取系统配置
func (s *SysConfigController) Config(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		utils.FailWithMessage("ConfigKeyNotEmpty", ctx)
	}
	data, err := s.SysConfigService.QueryKey(key)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithDetailed(&data, ctx)
	return
}

// Delete 删除系统配置
func (s *SysConfigController) Delete(ctx *gin.Context) {
	requestForms := dto.IdsReq{}
	if err := ctx.ShouldBind(&requestForms); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if err := s.SysConfigService.Delete(requestForms.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 系统配置列表
func (s *SysConfigController) List(ctx *gin.Context) {
	pageForms := dto.PageReq{}
	if err := ctx.ShouldBind(&pageForms); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	data, err := s.SysConfigService.List(&pageForms)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
