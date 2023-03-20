package admin

import (
	v1 "github.com/develop-kevin/easy-gin-vue-admin/app/v1/controller/admin"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) {
	r := router.Group("base")
	captcha := v1.SysCaptchaController{}
	{
		r.GET("captcha", captcha.Captcha)
	}
}
