package Application

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

type shareResources interface {
	share()
}

func (r *Request) share() {}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		connectToDB(&request)
		return request
	}
}

func NewRequest(c *gin.Context) Request {
	req := req()
	request := req(c)
	return request
}

func (r Request) response(code int, body interface{}) {
	CloseConnection(&r)
	r.Context.JSON(code, body)
}
