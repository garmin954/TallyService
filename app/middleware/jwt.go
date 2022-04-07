package middleware

import (
	"errors"
	"net/http"
	"smg/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.Configs.Jwt.Key),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	Uid      int    `json:"uid"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     error = errors.New("Token已过期,请重新登录")
	TokenNotValidYet error = errors.New("Token无效,请重新登录")
	TokenMalformed   error = errors.New("Token不正确,请重新登录")
	TokenInvalid     error = errors.New("这不是一个token,请重新登录")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		code = utils.ERROR_TOKEN_EXIST
		if tokenHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")

		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  utils.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// 解析token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": utils.AUTHORIZATION_HAS_EXPIRED,
					"msg":  utils.GetErrMsg(utils.AUTHORIZATION_HAS_EXPIRED),
					"data": nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"code": utils.AUTHORIZATION_HAS_EXPIRED,
				"msg":  err.Error(),
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Set("userinfo", claims)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}
