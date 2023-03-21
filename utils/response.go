/*
 * @Author: null 1060236395@qq.com
 * @Date: 2023-03-20 16:32:53
 * @LastEditors: null 1060236395@qq.com
 * @LastEditTime: 2023-03-21 12:31:27
 * @FilePath: \easy-gin-vue-admin\utils\response.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"net/http"

	"github.com/develop-kevin/easy-gin-vue-admin/global"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		msgI18n = msg
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
	Result(ERROR, map[string]interface{}{}, ginI18n.MustGetMessage(message), ctx)
	ctx.Abort()
}

// ValidateError 校验数据错误
func ValidateError(errors validator.ValidationErrors, ctx *gin.Context) {
	err := errors.Translate(global.EGVA_TRANS)
	for _, validationErr := range err {
		FailWithMessage(validationErr, ctx)
		return
	}
}
