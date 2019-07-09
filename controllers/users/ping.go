package users

import (
	"../../models"
	"github.com/gin-gonic/gin"
)

// Ping check
func Ping(c *gin.Context) {
	db := models.DBConn()
	db.DB().Ping()
	c.JSON(200, gin.H{
		"message": "pung",
	})
}
