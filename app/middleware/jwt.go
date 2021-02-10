package middleware

import (
	"fitness/global"
	"fitness/pkg/code"
	"fitness/pkg/util/jwt"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("api-token")
		if token == "" {
			response.Failed(c,code.NoLoginOrIllegalAccess, nil)
			c.Abort()
			return
		}

		j := jwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Failed(c,code.AuthorizationHasExpired, err.Error())
			c.Abort()
			return
		}

		//检测token刷新
		c.Header("TokenUpdate", "no")
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			expiredTime := time.Now().Unix() + global.Config.JWT.ExpiresTime
			claims.StandardClaims.ExpiresAt = expiredTime
			newToken, _ := j.GenToken(claims)
			c.Header("TokenUpdate", "yes")
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(expiredTime, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
