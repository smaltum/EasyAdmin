package api

import "github.com/gin-gonic/gin"

var RouterGroup = new(routerGroup)

type routerGroup struct {
	accountRouter
}

func (r *routerGroup) InitPublic(router *gin.RouterGroup) {
}

func (r *routerGroup) InitPrivate(router *gin.RouterGroup) {
	r.InitAccount(router)
}
