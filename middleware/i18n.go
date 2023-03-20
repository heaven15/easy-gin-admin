package middleware

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func GinI18nLocalize() gin.HandlerFunc {
	dir, file := utils.GetRoot()
	return ginI18n.Localize(
		ginI18n.WithBundle(&ginI18n.BundleCfg{
			RootPath:         fmt.Sprintf("%s%slanguage", dir, file),
			AcceptLanguage:   []language.Tag{language.Chinese, language.English},
			DefaultLanguage:  language.Chinese,
			UnmarshalFunc:    toml.Unmarshal,
			FormatBundleFile: "toml",
		}),
		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lang := context.GetHeader("Accept-Language")
				if lang == "" {
					return defaultLng
				}
				return lang
			},
		),
	)
}
