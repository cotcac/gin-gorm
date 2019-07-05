package articles

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Single(c *gin.Context) {
	id := c.Param("id")
	var a models.Article
	// var user models.User
	db:= models.DBConn()
	if err := db.Where("id = ?", id).Preload("User").First(&a).Error; err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{
			"Error":err,
		})
		return
	 } else {
		c.JSON(200, a)
	 }
}
