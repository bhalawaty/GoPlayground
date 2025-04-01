package Visitors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bilal-starter/Application"
	"go-bilal-starter/Models"
)

func CreateUser(c *gin.Context) {
	request := Application.NewRequest(c)
	fmt.Println("check the connection in NewRequest :", request.Connection.Ping())

	user := Models.User{
		FirstName: "bilal",
		LastName:  "elhalawaty",
		Age:       28,
		Email:     "bilal@gmail.com",
		Password:  "123456",
	}

	result := request.DB.Create(&user)

	if result.Error != nil {
		fmt.Println("❌ Failed to create user:", result.Error)
	} else {
		fmt.Println("✅ User created with ID:", user.ID)
	}

	request.Created(user)
}
