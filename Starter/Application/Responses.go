package Application

import "github.com/gin-gonic/gin"

func (r Request) Ok(body interface{}) {
	r.response(200, body)
}

func (r Request) Created(body interface{}) {
	r.Context.JSON(201, body)
}

func (r Request) NotAuth() {
	r.response(402, gin.H{
		"message": "Not Authorized",
	})
}

func (r Request) BadRequest(err interface{}) {
	r.response(422, err)
}
