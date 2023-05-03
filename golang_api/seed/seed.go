package main

import (
	"crud/go-orm-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:4000)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "Karn", Lname: "Yong", Email: "karn.yong@mecallapi.com", Avatar: "https://www.mecallapi.com/users/1.png"})
	db.Create(&model.User{Fname: "Ivy", Lname: "Cal", Email: "ivy.cal@mecallapi.com", Avatar: "https://www.mecallapi.com/users/2.png"})

}
