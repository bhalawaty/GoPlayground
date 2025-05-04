package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) share() {}

func App() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectToDB(&application)
		err := gotrans.InitLocales("./Public/Lang")
		if err != nil {
			fmt.Println("Error For Init Lang")
		}
		return application
	}
}

func NewApp() Application {
	newApp := App()
	return newApp()
}
