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

type SysUserController struct {
	SysUserService service.SysUserService
}

// Create 创建用户
func (s *SysUserController) Create(ctx *gin.Context) {
	request := dto.SysUserInfoReq{}
	if err := ctx.BindJSON(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysUser{
		AccountName: request.AccountName,
		UserName:    request.UserName,
		NickName:    request.NickName,
		PassWord:    request.PassWord,
		Gender:      request.Gender,
		Birthday:    request.Birthday,
		Email:       request.Email,
		Mobile:      request.Mobile,
		//PostID:      request.PostID,
		//RankID:      request.RankID,
		//GroupID:     request.GroupID,
		Avatar:     request.Avatar,
		City:       request.City,
		Address:    request.Address,
		Status:     request.Status,
		Ip:         ctx.ClientIP(),
		Remark:     request.Remark,
		RoleList:   request.RoleList,
		CreateUser: 0,
	}
	if err := s.SysUserService.Create(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Update 更新用户
func (s *SysUserController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	request := dto.SysUserInfoReq{}
	if err := ctx.BindJSON(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysUser{
		ID:          i,
		AccountName: request.AccountName,
		UserName:    request.UserName,
		NickName:    request.NickName,
		PassWord:    request.PassWord,
		Gender:      request.Gender,
		Birthday:    request.Birthday,
		Email:       request.Email,
		Mobile:      request.Mobile,
		//PostID:      request.PostID,
		//RankID:      request.RankID,
		//GroupID:     request.GroupID,
		Avatar:     request.Avatar,
		City:       request.City,
		Address:    request.Address,
		Status:     request.Status,
		Ip:         ctx.ClientIP(),
		Remark:     request.Remark,
		RoleList:   request.RoleList,
		CreateUser: 0,
	}
	if err = s.SysUserService.Update(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// Delete 删除用户
func (s *SysUserController) Delete(ctx *gin.Context) {
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
	if err := s.SysUserService.Delete(request.Ids); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// AssignPostRank 分配岗位和职级
func (s *SysUserController) AssignPostRank(ctx *gin.Context) {
	request := dto.SysUserAssignPostAndRankReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysUser{
		ID:     request.UserId,
		PostID: request.PostID,
		RankID: request.RankID,
	}
	if err := s.SysUserService.AssignPostRank(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// ResetPassWord 重置用户密码
func (s *SysUserController) ResetPassWord(ctx *gin.Context) {
	request := dto.SysUserPassWordReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	m := model.SysUser{
		ID:       0,
		PassWord: request.PassWord,
	}
	if err := s.SysUserService.ResetPassWord(&m); err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.Success(ctx)
	return
}

// List 用户列表
func (s *SysUserController) List(ctx *gin.Context) {
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
	data, err := s.SysUserService.List(&request)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(data, ctx)
	return
}
