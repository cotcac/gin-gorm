package articles

import (
	"github.com/gin-gonic/gin"
	"../../models"
	"strconv"
	"fmt"
)

func List(c *gin.Context) {

	articles := make([]models.Article,0)

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
	// var user models.User
	if err := db.Preload("User").Find(&articles).Offset(skip).Limit(perPage + 1).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }
	 fmt.Println(articles)
	var nextPage bool
	articlesLen := len(articles)
	if articlesLen > perPage {
		nextPage = true
		articles = articles[:len(articles)-1]// remote the last item.
	} else {
		nextPage = false
	}

	c.JSON(200,gin.H{
		"page":page,
		"result":articles,
		"next": nextPage,
	})
	defer db.Close()
}