package Auth

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bilal-starter/Application"
	"go-bilal-starter/Models"
)

func Register(c *gin.Context) {
	request := Application.NewRequest(c)

	// binding request
	var user Models.User
	request.Context.ShouldBindJSON(&user)

	// validate request
	err := validation.Errors{
		"firstName": validation.Validate(user.FirstName, validation.Required, validation.Length(2, 50)),
		"email":     validation.Validate(user.Email, validation.Required, validation.Length(2, 50)),
		"password":  validation.Validate(user.Password, validation.Required, validation.Length(2, 50)),
	}.Filter()

	if err != nil {
		request.BadRequest(err)
		return
	}
	request.Ok("okay")
}
