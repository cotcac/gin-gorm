package articles

import (
	"strconv"

	"../../models"
	"github.com/gin-gonic/gin"
)

// List of all article
func List(c *gin.Context) {

	articles := make([]models.Article, 0)
	results := make([]models.Articles, 0)

	db := models.DBConn()
	perPage := 5
	pageStr := c.DefaultQuery("page", "1")
	var err error
	page, err := strconv.Atoi(pageStr) // the err is necessary
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err,
		})
		return
	}
	var skip int

	if page > 0 {
		skip = (page - 1) * perPage
	} else {
		skip = 0
	}
	// var user models.User
	if err := db.Preload("User").Preload("Categories").Find(&articles).Offset(skip).Limit(perPage + 1).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	//  fmt.Println(articles)

	var nextPage bool
	articlesLen := len(articles)
	if articlesLen > perPage {
		nextPage = true
		articles = articles[:len(articles)-1] // remote the last item.
	} else {
		nextPage = false
	}

	for _, i := range articles {
		results = append(results, i.ArticleToArticles())
	}

	c.JSON(200, gin.H{
		"page":   page,
		"result": results,
		"next":   nextPage,
	})
}
