package users

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Single user
func Single(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	db := models.DBConn()
	if err := db.Where("id = ?", id).First(&user).Association("Article").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Print(err)
	} else {
		c.JSON(200, user)
	}
}
