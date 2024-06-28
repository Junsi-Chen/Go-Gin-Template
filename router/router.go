package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "template/docs"
	"template/middleware"
	"template/router/api/testctl"
)

// InitRouter initialize routing information
func InitRouter(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api")
	api.Use(middleware.CORSMiddleware())
	testRoute(api.Group("/test"))
}

func testRoute(rg *gin.RouterGroup) {
	rg.GET("", testctl.GetTest)
}
