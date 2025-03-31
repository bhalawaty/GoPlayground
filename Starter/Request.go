package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		request.DB, request.Connection = connectToDB()

		return request
	}
}

func newRequest(c *gin.Context) Request {
	req := req()
	request := req(c)
	return request
}

func connectToDB() (*gorm.DB, *sql.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/GoLang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("there is a connect issue :", err.Error())
	}

	connection, err := db.DB()

	if err != nil {
		fmt.Println("there is a connection issue :", err.Error())
	}
	return db, connection
}

func (r Request) response(code int, body interface{}) {
	r.closeConnection()
	r.Context.JSON(code, body)
}

func (r Request) closeConnection() {
	r.Connection.Close()
}
