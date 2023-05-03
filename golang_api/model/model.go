package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fname  string
	Lname  string
	Email  string
	Avatar string
}
