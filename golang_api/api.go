package main

import (
	"crud/go-orm-api/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:4000)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var users model.User
		db.First(&users, id)
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user model.User
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		result := db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"RowsAffected": result.RowsAffected})
	})

	r.PUT("/users", func(c *gin.Context) {
		var user model.User
		var updateUser model.User
		// in this case proper binding will be automatically selected
		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		db.First(&updateUser, user.ID)
		updateUser.Fname = user.Fname
		updateUser.Lname = user.Lname
		updateUser.Email = user.Email
		updateUser.Avatar = user.Avatar
		db.Save(&updateUser)
		c.JSON(http.StatusOK, gin.H{"RowsAffected": updateUser})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var users model.User
		db.First(&users, id)
		db.Delete(&users)
		c.JSON(http.StatusOK, users)
	})

	config := cors.DefaultConfig()
	r.Use(cors.New(config))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
