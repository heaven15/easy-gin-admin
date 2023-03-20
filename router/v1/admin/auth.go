package admin

import (
	v1 "github.com/develop-kevin/easy-gin-vue-admin/app/v1/controller/admin"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(router *gin.RouterGroup) {
	r := router.Group("auth")
	auth := v1.SysAuthController{}
	{
		r.POST("register", auth.Register)
		r.POST("login", auth.Login)
	}
}
