package admin

import (
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/service"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

type SysManagerController struct {
	SysUserService service.SysUserService
	SysMenuService service.SysMenuService
	SysRoleService service.SysRoleService
}

// GetInfo 获取用户信息
func (s *SysManagerController) GetInfo(ctx *gin.Context) {
	id, err := utils.GetTokenUserId(ctx)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	m := model.SysUser{
		ID: id,
	}
	data, err := s.SysUserService.Detail(&m)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithDetailed(data, ctx)
	return
}

// GetMenu 获取菜单列表
func (s *SysManagerController) GetMenu(ctx *gin.Context) {
	//TODO 暂时不获取用户的角色对应的菜单
	data := s.SysMenuService.Tree()
	utils.SuccessWithData(data, ctx)
	return
}

func (s *SysManagerController) GetPermission(ctx *gin.Context) {
	id, err := utils.GetTokenUserId(ctx)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	m := model.SysUser{
		ID: id,
	}
	data, err := s.SysUserService.Detail(&m)
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	var userRoleData map[string]interface{}
	userRoleData["user_id"] = data.ID
	userRoleData["account_name"] = data.AccountName
	userRoleData["username"] = data.UserName
	userRoleData["nickname"] = data.NickName
	userRoleData["avatar"] = data.Avatar
	userRoleData["city"] = data.City
	userRoleData["address"] = data.Address
	userRoleData["mobile"] = data.Mobile
	userRoleData["post_info"] = data.PostInfo
	userRoleData["rank_info"] = data.RankInfo
	userRoleData["remark"] = data.Remark
	var roleList []string
	var permissionList []string
	if len(data.RoleList) > 0 {
		var roleIds []int
		for _, v := range data.RoleList {
			roleIds = append(roleIds, int(v.ID))
		}
		permissionData, err := s.SysRoleService.GetPermission(roleIds)
		if err != nil {
			utils.FailWithMessage(err.Error(), ctx)
			return
		}
		if len(permissionData) > 0 {
			for _, v1 := range permissionData {
				roleList = append(roleList, v1.RealName)
				if len(v1.PermissionList) > 0 {
					for _, v2 := range v1.PermissionList {
						permissionList = append(permissionList, v2.Code)
					}
				}
			}
		}
	}
	userRoleData["user_id"] = permissionList
	utils.SuccessWithDetailed(userRoleData, ctx)
	return
}
