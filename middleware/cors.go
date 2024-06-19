package middleware

import "github.com/gin-gonic/gin"

var corsUrl = "*"

func InitCorsUrl(url string) {
	if url != "" {
		corsUrl = url
	}
}

// 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With, token
		//content-type,Version,Device,Unique-Device,Drives,Web-Token
		c.Writer.Header().Set("Access-Control-Allow-Origin", corsUrl)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Version,Device,Unique-Device,"+
			"Drives,Web-Token,Device-Model,screen-size,Channel-Name,Language-Code,Version-Name,Origin-Id,IDFA")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}
