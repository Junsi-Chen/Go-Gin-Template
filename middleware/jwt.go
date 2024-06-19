package middleware

import (
	"github.com/gin-gonic/gin"
	"template/tool/response"
	"template/tool/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = response.SUCCESS
		// 在头部获取token
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token") // 如果头部没有，那么在参数中获取token
		}
		// 有token
		claims, err := util.ParseToken(token)
		// 验证token
		if err != nil {
			code = response.Unauthorized
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = response.AuthCheckTokenTimeOut
		}
		// 如果不成功，那么返回错误信息
		if code != response.SUCCESS {
			response.CommonResponse(c, code)
			c.Abort()
			return
		}
		// 保存上下文用户id
		c.Set(util.CtxUserId, claims.Id)
		c.Next()
	}
}
