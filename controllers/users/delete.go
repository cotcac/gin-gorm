package users

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Delete user
func Delete(c *gin.Context) {
	id := c.Param("id")
	db := models.DBConn()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}
	fmt.Print(user)
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"Deleted": user,
	})
}
