package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app().Gin //run the application
	application.GET("/ping", func(c *gin.Context) {
		request := getRequest(c)
		request.Context.JSON(http.StatusOK, gin.H{
			"message": "pooooong",
		})
	})
	err := application.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
