package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"go-bilal-starter/Models"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	User       *Models.User
	IsAuth     bool
	IsAdmin    bool
	Lang       string
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
		setLang(&request)
		return request
	}
}

func setLang(request *Request) {
	lang := gotrans.DetectLanguage(request.Context.GetHeader("Accept-Language"))
	err := gotrans.SetDefaultLocale(lang)
	if err != nil {
		fmt.Println("Error While Set Default Lang")
	}
	request.Lang = lang
}

func NewRequest(c *gin.Context) Request {
	req := req()
	return req(c)
}

func (r Request) response(code int, body interface{}) {
	CloseConnection(&r)
	r.Context.JSON(code, body)
}

func (r Request) Auth() Request {
	r.IsAuth = false

	requestHeader := r.Context.GetHeader("Authorization")
	fmt.Println("Authorization Header => ", requestHeader)

	if requestHeader != "" {
		r.DB.Where("token = ?", requestHeader).First(&r.User)
		if r.User.ID != 0 {
			r.IsAuth = true
		}

	}

	return r

}
