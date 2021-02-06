package middleware

import (
	"errors"
	"fitness/global"
	"fitness/pkg/code"
	"fitness/pkg/util/jwtToken"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("api-token")
		if token == "" {
			response.Failed(code.NoLoginOrIllegalAccess, nil)
			c.Abort()
			return
		}

		j := jwtToken.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Failed(code.AuthorizationHasExpired, err.Error())
			c.Abort()
			return
		}

		//检测token是否过期
		if time.Now().Unix() < claims.ExpiresAt {
			claims.StandardClaims.ExpiresAt = time.Now().Unix() + int64(global.TokenExpireDuration.Seconds())
			newToken, _ := j.GenToken(claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
