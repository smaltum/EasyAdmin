package admin

import "github.com/gin-gonic/gin"

var RouterGroup = new(routerGroup)

type routerGroup struct {
	apiRouter
}

func (r *routerGroup) Init(router *gin.RouterGroup) {

	r.InitApi(router)

}
