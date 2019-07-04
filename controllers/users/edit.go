package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Edit(c *gin.Context) {
	id := c.Param("id")
	db:= models.DBConn()
	type User struct {
		ID   int
		Name string
	  }
	type EditUser struct {
		Name string `json:"name" binding:"required,min=3,max=15"`
	}
	var json EditUser
    
	if err:= c.ShouldBindJSON(&json); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return 
	}
	fmt.Println(json)
	var user User
	if err := db.First(&user,id).Error; err != nil {
		c.AbortWithStatus(404)
		return
	 }
	 user.Name = json.Name
	 if err := db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
	c.JSON(200, gin.H{
		"Edited": user,
	})	
}
