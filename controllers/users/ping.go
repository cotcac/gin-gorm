package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Ping(c *gin.Context) {
	db:= models.DBConn()
	db.DB().Ping()
	c.JSON(200, gin.H{
		"message": "pung",
	})
}
