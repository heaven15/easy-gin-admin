package admin

import (
	v1 "github.com/develop-kevin/easy-gin-vue-admin/app/v1/controller/admin"
	"github.com/develop-kevin/easy-gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSystemRouter 系统路由
func InitSystemRouter(router *gin.RouterGroup) {
	r := router.Group("system").Use(middleware.JWT())
	sysConfig := v1.SysConfigController{}
	{
		r.POST("config", sysConfig.Create)
		r.PUT("config/:id", sysConfig.Update)
		r.GET("config/:key", sysConfig.Config)
		r.DELETE("config", sysConfig.Delete)
		r.GET("config", sysConfig.List)
	}
	sysPost := v1.SysPostController{}
	{
		r.POST("post", sysPost.Create)
		r.PUT("post/:id", sysPost.Update)
		r.DELETE("post", sysPost.Delete)
		r.GET("post", sysPost.List)
	}
	sysDepartment := v1.SysDepartmentController{}
	{
		r.POST("department", sysDepartment.Create)
		r.PUT("department/:id", sysDepartment.Update)
		r.POST("department/join", sysDepartment.Join)
		r.DELETE("department", sysDepartment.Delete)
		r.GET("department", sysDepartment.List)
	}
	sysDictionary := v1.SysDictionaryController{}
	{
		r.POST("dictionary", sysDictionary.Create)
		r.PUT("dictionary/:id", sysDictionary.Update)
		r.DELETE("dictionary", sysDictionary.Delete)
		r.GET("dictionary", sysDictionary.List)
	}
	sysRank := v1.SysRankController{}
	{
		r.POST("rank", sysRank.Create)
		r.PUT("rank/:id", sysRank.Update)
		r.DELETE("rank", sysRank.Delete)
		r.GET("rank", sysRank.List)
	}
	sysGroup := v1.SysGroupController{}
	{
		r.POST("group", sysGroup.Create)
		r.PUT("group/:id", sysGroup.Update)
		r.DELETE("group", sysGroup.Delete)
		r.GET("group", sysGroup.List)
	}
	sysOperation := v1.SysOperationController{}
	{
		r.POST("operation", sysOperation.Create)
		r.PUT("operation/:id", sysOperation.Update)
		r.DELETE("operation", sysOperation.Delete)
		r.GET("operation", sysOperation.List)
	}
	sysFile := v1.SysFileController{}
	{
		r.POST("file", sysFile.Create)
		r.PUT("file/:id", sysFile.Update)
		r.DELETE("file", sysFile.Delete)
		r.GET("file", sysFile.List)
	}
	sysPage := v1.SysPageController{}
	{
		r.POST("page", sysPage.Create)
		r.PUT("page/:id", sysPage.Update)
		r.DELETE("page", sysPage.Delete)
		r.GET("page", sysPage.List)
	}
	sysPermission := v1.SysPermissionController{}
	{
		r.POST("permission", sysPermission.Create)
		r.PUT("permission/:id", sysPermission.Update)
		r.DELETE("permission", sysPermission.Delete)
		r.GET("permission", sysPermission.List)
	}
	sysMenu := v1.SysMenuController{}
	{
		r.POST("menu", sysMenu.Create)
		r.PUT("menu/:id", sysMenu.Update)
		r.DELETE("menu", sysMenu.Delete)
		r.GET("menu", sysMenu.List)
	}
	sysRole := v1.SysRoleController{}
	{
		r.POST("role", sysRole.Create)
		r.PUT("role/:id", sysRole.Update)
		r.PUT("role/assignAuthority", sysRole.AssignAuthority)
		r.DELETE("role", sysRole.Delete)
		r.GET("role", sysRole.List)
	}
	sysUser := v1.SysUserController{}
	{
		r.POST("user", sysUser.Create)
		r.PUT("user/:id", sysUser.Update)
		r.PUT("user/assignPostRank", sysUser.AssignPostRank)
		r.DELETE("user", sysUser.Delete)
		r.GET("user", sysUser.List)
	}
	sysManager := v1.SysManagerController{}
	{
		r.GET("manager/info", sysManager.GetInfo)
		r.GET("manager/menu", sysManager.GetMenu)
		r.PUT("manager/permission", sysManager.GetPermission)
	}
}
