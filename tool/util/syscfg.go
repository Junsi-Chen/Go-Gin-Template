package util

import (
	"github.com/gin-gonic/gin"
)

const (
	CtxUserId = "ctx_user_id"
)

func UserId(ctx *gin.Context) int64 {
	if id, ok := ctx.Get(CtxUserId); ok {
		return id.(int64)
	}
	return 0
}
