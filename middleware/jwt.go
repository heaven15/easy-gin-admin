package middleware

import (
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		token := strings.Replace(authorization, "Bearer ", "", 1)
		if token == "" {
			utils.FailWithMessage("PleaseLogIn", ctx)
			ctx.Abort()
			return
		}
		jwt := utils.NewJWT()
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				utils.FailWithMessage("AuthorizationHasExpired", ctx)
				ctx.Abort()
				return
			}
			utils.FailWithMessage("NotloggedIn", ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.EGVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := jwt.RefreshToken(token)
			newClaims, _ := jwt.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
