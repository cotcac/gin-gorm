package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	db:= models.DBConn()
	type User struct {
		ID   int
		Name string
	  }
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		return
	 }  
	 fmt.Print(user)
	 if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
	c.JSON(200, gin.H{
		"Deleted": user,
	})
}