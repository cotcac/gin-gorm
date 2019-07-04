package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	type User struct {
		ID   int
		Name string
	  }
	type CreateUser struct {
		Name string `json:"name" binding:"required"`
	}  
	var json CreateUser
	if err:= c.ShouldBindJSON(&json); err!=nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user := User{Name:json.Name}  //db.Create(&User{Name:"Kinny 123"})
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
    fmt.Println(user)
	c.JSON(200, gin.H{
		"success": user,
	})
}
