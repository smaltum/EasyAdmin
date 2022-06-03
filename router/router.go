package router

import "easy-admin/router/admin"

var RouterGroup = new(routerGroup)

type routerGroup struct {
	Admin admin.RouterGroup
}
