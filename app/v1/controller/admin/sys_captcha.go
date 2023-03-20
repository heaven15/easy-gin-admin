/*
 * @Author: null 1060236395@qq.com
 * @Date: 2023-03-20 16:32:53
 * @LastEditors: null 1060236395@qq.com
 * @LastEditTime: 2023-03-20 16:39:29
 * @FilePath: \easy-gin-vue-admin\app\v1\controller\admin\sys_captcha.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package admin

import (
	"image/color"

	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/service"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/vo"
	"github.com/develop-kevin/easy-gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type SysCaptchaController struct{}

var store = service.NewDefaultRedisStore()

// Captcha 生成验证码
func (s *SysCaptchaController) Captcha(ctx *gin.Context) {
	var driver base64Captcha.Driver
	//driver = base64Captcha.NewDriverDigit(50, 200, 4, 0.1, 100)
	driverString := base64Captcha.DriverString{
		Height:          30,
		Width:           90,
		NoiseCount:      0,
		ShowLineOptions: 1 | 2,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 255, G: 255, B: 255, A: 0},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx))
	id, b64s, _, err := cp.Generate()
	if err != nil {
		utils.FailWithMessage(err.Error(), ctx)
		return
	}
	utils.SuccessWithData(vo.CaptchaVo{
		CaptchaId:   id,
		Base64Image: b64s,
	}, ctx)
	return
}
