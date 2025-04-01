package Routes

import "go-bilal-starter/Controllers/Visitors"

func (application RouterApp) visitorsRoutes() {

	application.Gin.GET("/users/create", Visitors.CreateUser)
}
