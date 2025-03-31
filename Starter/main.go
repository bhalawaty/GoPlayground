package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	application := app().Gin //run the application
	application.GET("/ping", func(c *gin.Context) {
		request := newRequest(c)
		request.ok(gin.H{
			"message": "pong",
		})
	})
	err := application.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
