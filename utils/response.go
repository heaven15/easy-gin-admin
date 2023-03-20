package utils

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

const (
	ERROR = iota
	SUCCESS
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Result 数据返回格式
func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	msgI18n := ""
	if msg != "" {
		str := strings.Split(msg, "_")
		if len(str) == 0 {
			msgI18n = ginI18n.MustGetMessage(msg)
		} else {
			msgI18n = fmt.Sprintf(ginI18n.MustGetMessage(str[0]), str[1])
		}
	} else {
		msgI18n = ginI18n.MustGetMessage("DataIsUndefined")
	}
	if gin.Mode() == "release" {
		msgI18n = ginI18n.MustGetMessage("SystemError")
	}
	ctx.JSONP(http.StatusOK, Response{
		code,
		data,
		msgI18n,
	})
}

// Success 操作成功
func Success(ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, ginI18n.MustGetMessage("SuccessFullOperation"), ctx)
}

// SuccessWithMessage 成功信息
func SuccessWithMessage(message string, ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}

// SuccessWithData 查询成功
func SuccessWithData(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, data, ginI18n.MustGetMessage("QuerySuccess"), ctx)
}

// SuccessWithDetailed 成功查询单条数据
func SuccessWithDetailed(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, data, ginI18n.MustGetMessage("QuerySuccess"), ctx)
}

// Fail 操作失败
func Fail(ctx *gin.Context) {
	Result(ERROR, map[string]interface{}{}, ginI18n.MustGetMessage("FailFullOperation"), ctx)
}

// FailWithMessage 失败信息
func FailWithMessage(message string, ctx *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, ctx)
}

// ValidateError 校验数据错误
func ValidateError(errors validator.ValidationErrors, ctx *gin.Context) {
	err := errors.Translate(global.EGVA_TRANS)
	for _, validationErr := range err {
		FailWithMessage(validationErr, ctx)
		return
	}
}
