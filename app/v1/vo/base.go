/*
 * @Author: null 1060236395@qq.com
 * @Date: 2023-03-20 16:32:53
 * @LastEditors: null 1060236395@qq.com
 * @LastEditTime: 2023-03-21 16:43:37
 * @FilePath: \easy-gin-vue-admin\app\v1\vo\base.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package vo

// CaptchaVo 验证码返回的结构
type CaptchaVo struct {
	CaptchaId   string `json:"captcha_id"`
	Base64Image string `json:"base64_image"`
}

// AccessTokenVo 生成用户token返回的结构
type AccessTokenVo struct {
	Token    string `json:"token" label:"token数据"`
	ExpireAt int64  `json:"expire_at" label:"token过期时间"`
}

// PageDataVo 接口分页统一返回结构
type PageDataVo struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

// AllDataVo 接口全部数据统一返回结构
type AllDataVo struct {
	Data interface{} `json:"data"`
}
