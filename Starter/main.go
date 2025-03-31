package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app().Gin //run the application
	application.GET("/ping", func(c *gin.Context) {
		request := getRequest(c)
		fmt.Println("ping before close ", request.Connection.Ping())
		request.Connection.Close()
		fmt.Println("ping after close ", request.Connection.Ping())
		request.Context.JSON(http.StatusOK, gin.H{
			"message": "pooooong",
		})
	})
	err := application.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
