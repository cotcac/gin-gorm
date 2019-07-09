package users

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Edit user
func Edit(c *gin.Context) {
	id := c.Param("id")
	db := models.DBConn()
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(json)
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	user.Name = json.Name
	if err := db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"Edited": user,
	})
}
