package home

import (
	"github.com/gin-gonic/gin"
)

// Init ....
func Home(c *gin.Context) {

	c.JSON(200, gin.H{
		"success": "hello world",
	})
}
