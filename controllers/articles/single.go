package articles

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
)

// Single will show a single article
func Single(c *gin.Context) {
	id := c.Param("id")
	var a models.Article
	// var user models.User
	db := models.DBConn()
	if err := db.Where("id = ?", id).Preload("User").Preload("Categories").First(&a).Error; err != nil {
		fmt.Print(err)
		c.JSON(500, gin.H{
			"Error": err,
		})
		return
	}
	// convert to singleArticle.
	sa := a.ArticleToSingleArticle()
	c.JSON(200, sa)

}
