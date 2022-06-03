package api

import (
	"easy-admin/app/api"
	"github.com/gin-gonic/gin"
)

// accountRouter
type accountRouter struct{}

// InitAccount 初始化
func (r *accountRouter) InitAccount(router *gin.RouterGroup) {
	apiRouter := router.Group("api")
	apiRouterApi := api.Api
	{
		apiRouter.GET("/test", apiRouterApi.Test)        // 测试
		apiRouter.POST("/decrypt", apiRouterApi.Decrypt) // 创建
	}
}
