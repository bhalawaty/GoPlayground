package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"type:varchar(50)"`
	LastName  string `json:"lastName" gorm:"type:varchar(50)"`
	Age       int8   `json:"age"`
	Email     string `json:"email" gorm:"type:varchar(50)"`
	Password  string `json:"password" gorm:"type:varchar(50)"`
	Token     string `json:"token" gorm:"type:varchar(100)"`
	Group     string `json:"group" gorm:"type:varchar(20)"`
}
