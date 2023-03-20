package utils

import (
	"errors"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 常量
var (
	TokenExpired     = errors.New("Token Expired")
	TokenNotValidYet = errors.New("Token Not Active Yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Mobile   string `json:"mobile"`
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return global.EGVA_CONFIG.JWT.SecretKey
}

func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	jwtConf := global.EGVA_CONFIG.JWT
	bf, _ := ParseDuration(jwtConf.BufferTime)
	ep, _ := ParseDuration(jwtConf.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwtgo.StandardClaims{
			Issuer:    jwtConf.IsSuer,            //签名的发行者
			NotBefore: time.Now().Unix() - 1000,  //签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(), //过期时间7天 配置文件
		},
	}
	return claims
}

// GenerateToken 生成一个token
func (j *JWT) GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed //令牌解析错误
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired //令牌过期
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet //令牌无效
			} else {
				return nil, TokenInvalid //令牌无效
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid //令牌无效
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.GenerateToken(*claims)
	}
	return "", TokenInvalid //令牌无效
}

func GetTokenUserId(ctx *gin.Context) (int64, error) {
	claims, ok := ctx.Get("claims")
	if !ok {
		return 0, errors.New("TokenInvalidation")
	}
	return claims.(*CustomClaims).ID, nil
}
