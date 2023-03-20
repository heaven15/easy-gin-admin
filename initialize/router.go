package initialize

import (
	"fmt"
	middleware2 "github.com/develop-kevin/easy-gin-vue-admin/middleware"
	"github.com/develop-kevin/easy-gin-vue-admin/router/v1/admin"
	"github.com/gin-gonic/gin"
	"time"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//配置跨域
	// LoggerWithFormatter 中间件会写入日志到 gin.DefaultWriter
	// 默认 gin.DefaultWriter = os.Stdout
	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	Router.Use(gin.Recovery())
	Router.Use(middleware2.GinI18nLocalize())
	Router.Use(middleware2.Cors())
	AdminGroupV1 := Router.Group("v1/admin")
	admin.InitBaseRouter(AdminGroupV1)
	admin.InitAuthRouter(AdminGroupV1)
	admin.InitSystemRouter(AdminGroupV1)
	return Router
}
