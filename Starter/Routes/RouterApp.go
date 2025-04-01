package Routes

import "go-bilal-starter/Application"

type RouterApp struct {
	*Application.Application
}

func (routerApp RouterApp) Routes() {
	routerApp.visitorsRoutes()
	routerApp.authRoutes()
	routerApp.adminRoutes()
}
