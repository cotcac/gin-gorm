package categories

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"strconv"
)

func List(c *gin.Context) {

	categories := make([]models.Category,0)
	db:= models.DBConn()
	perPage := 5
	pageStr := c.DefaultQuery("page", "1")
	var err error
	page, err := strconv.Atoi(pageStr)// the err is necessary
	if err!=nil {
		c.JSON(500, gin.H{
			"Error":err,
		})
		return 
	}
	var skip int

	if page >0 {
		skip = (page - 1)* perPage
	}else {
		skip = 0
	}
	if err := db.Offset(skip).Limit(perPage + 1).Find(&categories).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
	var nextPage bool
	categoriesLen := len(categories)
	if categoriesLen > perPage {
		nextPage = true
		categories = categories[:len(categories)-1]// remote the last item.
	} else {
		nextPage = false
	}
	c.JSON(200,gin.H{
		"page":page,
		"result":categories,
		"next": nextPage,
	})
}