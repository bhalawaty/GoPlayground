package main

import "github.com/gin-gonic/gin"

func (r Request) ok(body interface{}) {
	r.response(200, body)
}

func (r Request) created(body interface{}) {
	r.Context.JSON(201, body)
}

func (r Request) notAuth() {
	r.response(402, gin.H{
		"message": "Not Authorized",
	})
}
