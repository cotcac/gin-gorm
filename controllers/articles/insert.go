package articles

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Insert is ...
func Insert(c *gin.Context) {
	db := models.DBConn()

	var json models.Article
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	article := models.Article{Title: json.Title, Content: json.Content, UserID: json.UserID, Categories: json.Categories}
	if err := db.Create(&article).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(article)
	c.JSON(200, gin.H{
		"success": article.ID,
	})
}
