package dto

import (
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
)

type SysUserInfoReq struct {
	AccountName string `json:"account_name" binding:"required,min=2,max=50" label:"账号名称"`
	UserName    string `json:"username" binding:"required,min=2,max=50" label:"用户姓名"`
	NickName    string `json:"nickname" binding:"required,min=1,max=50" label:"用户昵称"`
	PassWord    string `json:"password" binding:"required,min=5,max=30" label:"密码"`
	Gender      int    `json:"gender,default=3" binding:"omitempty,oneof=1 2 3" label:"性别"`
	Birthday    string `json:"birthday" binding:"omitempty" label:"出生日期"`
	Email       string `json:"email" binding:"omitempty,email" label:"邮箱"`
	Mobile      string `json:"mobile" binding:"required,len=11,mobile" label:"用户手机号"`
	//PostID      int64            `json:"post_id" binding:"required" label:"岗位ID"`
	//RankID      int64            `json:"rank_id" binding:"required" label:"职级ID"`
	//GroupID     int64            `json:"group_id" binding:"omitempty" label:"用户分组ID"`
	Avatar   string           `json:"avatar" binding:"omitempty" label:"头像"`
	City     model.CityJson   `json:"city" binding:"omitempty" label:"城市json数据"`
	Address  string           `json:"address" binding:"omitempty" label:"地址详情"`
	Status   int              `json:"status" binding:"omitempty,oneof=1 2" label:"状态"`
	Remark   string           `json:"remark" binding:"omitempty" label:"备注"`
	RoleList []*model.SysRole `json:"role_list" binding:"omitempty" label:"角色列表"`
}

type SysUserAssignPostAndRankReq struct {
	UserId int64 `json:"user_id" binding:"required,min=1" label:"用户ID"`
	PostID int64 `json:"post_id" binding:"required" label:"岗位ID"`
	RankID int64 `json:"rank_id" binding:"required" label:"职级ID"`
}

type SysUserJoinRoleReq struct {
	UserId   int64            `json:"user_id" binding:"required,min=1" label:"用户ID"`
	RoleList []*model.SysRole `json:"role_list" binding:"required" label:"角色列表"`
}

type SysUserPassWordReq struct {
	PassWord string `json:"password" binding:"required,min=5,max=30" label:"密码"`
}

type SysUserLoginReq struct {
	UserName  string `json:"username" binding:"required,min=2,max=30" label:"用户名"` // 用户名
	Mobile    string `json:"mobile" binding:"omitempty,len=11,mobile" label:"用户手机号"`
	PassWord  string `json:"password" binding:"required,min=5,max=30" label:"密码"` // 密码
	LoginType string `json:"login_type" binding:"required" label:"登录方式"`          // 登录方式
	CaptchaId string `json:"captcha_id" binding:"required" label:"验证码ID"`
	Code      string `json:"code"  binding:"required" label:"验证码"`
}

type SysUserRegisterReq struct {
	AccountName string `json:"account_name" binding:"required,min=2,max=50" label:"账号名称"`
	UserName    string `json:"username" binding:"required,min=2,max=50" label:"用户姓名"`
	NickName    string `json:"nickname" binding:"required,min=1,max=50" label:"用户昵称"`
	PassWord    string `json:"password" binding:"required,min=5,max=30" label:"密码"`
	Gender      int    `json:"gender,default=3" binding:"omitempty,oneof=1 2 3" label:"性别"`
	Birthday    string `json:"birthday" binding:"omitempty" label:"出生日期"`
	Email       string `json:"email" binding:"omitempty,email" label:"邮箱"`
	Mobile      string `json:"mobile" binding:"required,len=11,mobile" label:"用户手机号"`
	Avatar      string `json:"avatar" binding:"omitempty" label:"头像"`
	Remark      string `json:"remark" binding:"omitempty" label:"备注"`
	LoginType   string `json:"login_type" binding:"required" label:"登录方式"` // 登录方式
}
