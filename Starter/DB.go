package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeConnection() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/GoLang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("make Connection issue :", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("return Connection issue :", err.Error())
	}
	return connection
}

func connectToDB(share shareResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.DB = makeConnection()
		app.Connection = returnConnection(app.DB)
	case *Request:
		request := share.(*Request)
		request.DB = makeConnection()
		request.Connection = returnConnection(request.DB)
	}

}

func closeConnection(share shareResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.Connection.Close()
	case *Request:
		request := share.(*Request)
		request.Connection.Close()
	}

}
