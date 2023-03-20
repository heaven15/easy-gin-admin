package admin

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/dto"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/service"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SysAuthController struct {
	SysUserService service.SysUserService
}

// Login 授权登录
func (s *SysAuthController) Login(ctx *gin.Context) {
	request := dto.SysUserLoginReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	var err error
	//判断验证码
	if !store.Verify(request.CaptchaId, request.Code, true) {
		utils.FailWithMessage("CaptchaCodeError", ctx)
		return
	}
	data := &model.SysUser{}
	token := &vo.AccessTokenVo{}
	if request.LoginType == "username" {
		//处理用户名
		m := model.SysUser{
			UserName: request.UserName,
			PassWord: request.PassWord,
			Ip:       ctx.ClientIP(),
		}
		data, err = s.SysUserService.QueryName(&m)
		if err != nil {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
		if ok := utils.BcryptCheck(fmt.Sprintf("%s%s", request.PassWord, data.Salt), data.PassWord); !ok {
			utils.FailWithMessage("SysUserPassWordError", ctx)
			return
		}
	}
	if request.LoginType == "mobile" {
		//短信验证码校验
		m := model.SysUser{
			Mobile: request.Mobile,
			Ip:     ctx.ClientIP(),
		}
		data, err = s.SysUserService.QueryMobile(&m)
		if err != nil {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if data.Status == 2 {
		utils.FailWithMessage("SysUserAccountIsDisabled", ctx)
		return
	}
	//生成token
	token, err = s.GenerateSysToken(data)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(token, ctx)
	return
}

// Register 授权注册
func (s *SysAuthController) Register(ctx *gin.Context) {
	request := dto.SysUserRegisterReq{}
	if err := ctx.ShouldBind(&request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(errors, ctx)
			return
		} else {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	var err error
	m := model.SysUser{
		UserName: request.UserName,
		NickName: request.AccountName,
		PassWord: request.PassWord,
		Mobile:   request.Mobile,
		Email:    request.Email,
		Avatar:   request.Avatar,
		Ip:       ctx.ClientIP(),
		Remark:   request.Remark,
	}
	if request.LoginType == "username" {
		if err = s.SysUserService.Create(&m); err != nil {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if request.LoginType == "mobile" {
		//短信验证码校验

	}
	utils.Success(ctx)
	return
}

// GenerateSysToken 生成系统Token
func (s *SysAuthController) GenerateSysToken(u *model.SysUser) (*vo.AccessTokenVo, error) {
	baseClaims := utils.BaseClaims{
		ID:       u.ID,
		UserName: u.UserName,
		Mobile:   u.Mobile,
	}
	jwt := utils.NewJWT()
	claims := jwt.CreateClaims(baseClaims)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		return nil, err
	}
	data := vo.AccessTokenVo{
		AccessToken: token,
		ExpireAt:    claims.StandardClaims.ExpiresAt,
	}
	return &data, nil
}
