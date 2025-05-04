package Routes

import (
	"go-bilal-starter/Controllers/Auth"
)

func (application RouterApp) authRoutes() {
	application.Gin.POST("/auth/register", Auth.Register)

}
