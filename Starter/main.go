package main

import (
	"fmt"
	"go-bilal-starter/Application"
	"go-bilal-starter/Models"
	"go-bilal-starter/Routes"
)

func main() {
	app := Application.NewApp()
	migrationError := app.DB.AutoMigrate(&Models.User{})
	fmt.Println("check the App connection in migration :", app.Connection.Ping())

	if migrationError != nil {
		println("migration error: ", migrationError.Error())
	}

	Application.CloseConnection(&app)

	routerApp := Routes.RouterApp{Application: &app}
	routerApp.Routes()
	app.Gin.Run()

}
