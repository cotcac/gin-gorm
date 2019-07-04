package articles

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	type CreateArticle struct {
		Title string `json:"title" binding:"required,min=3,max=15"`
		UserID int `json:"user_id" binding:"required"`
	}  
	var json CreateArticle
	if err:= c.ShouldBindJSON(&json); err!=nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	article := models.Article{Title:json.Title, UserID:json.UserID}  //db.Create(&User{Name:"Kinny 123"})
	if err := db.Create(&article).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
    fmt.Println(article)
	c.JSON(200, gin.H{
		"success": article,
	})
}
