package categories

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"fmt"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	var json models.Category
	if err:= c.ShouldBindJSON(&json); err!=nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	cate := models.Category{Title:json.Title}  
	if err := db.Create(&cate).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
    fmt.Println(cate)
	c.JSON(200, gin.H{
		"success": cate,
	})
}
