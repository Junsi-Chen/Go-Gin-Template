package testctl

import (
	"github.com/gin-gonic/gin"
)

// GetTest 测试接口
// @Summary 测试接口
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/test [get]
func GetTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
