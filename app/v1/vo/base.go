package vo

// CaptchaVo 验证码返回的结构
type CaptchaVo struct {
	CaptchaId   string `json:"captcha_id"`
	Base64Image string `json:"Base64_image"`
}

// AccessTokenVo 生成用户token返回的结构
type AccessTokenVo struct {
	AccessToken string `json:"access_token" label:"access_token数据"`
	ExpireAt    int64  `json:"expire_at" label:"access_token过期时间"`
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
