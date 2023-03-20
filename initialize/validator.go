package initialize

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func InitTrans(local string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("label")
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备选语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(zhT, zhT, enT)
		global.EGVA_TRANS, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("初始翻译错误：%s", local)
		}
		//注册验证器
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = v.RegisterValidation("mobile", utils.ValidateMobile)
			_ = v.RegisterTranslation("mobile", global.EGVA_TRANS, func(ut ut.Translator) error {
				return ut.Add("mobile", "{0}无效号码", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
		}
		switch local {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.EGVA_TRANS)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.EGVA_TRANS)
		default:
			_ = zh_translations.RegisterDefaultTranslations(v, global.EGVA_TRANS)
		}
		return
	}
	return
}
