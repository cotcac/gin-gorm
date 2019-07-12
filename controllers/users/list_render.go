package users

import (
	"strconv"

	"../../models"
	"github.com/gin-gonic/gin"
)

// ListRender is ...
func ListRender(c *gin.Context) {
	users := make([]models.User, 0)
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
	if err := db.Offset(skip).Limit(perPage + 1).Find(&users).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	var nextPage bool
	usersLen := len(users)
	if usersLen > perPage {
		nextPage = true
		users = users[:len(users)-1] // remote the last item.
	} else {
		nextPage = false
	}
	results := make([]models.ListUsers, 0)
	for _, i := range users {
		results = append(results, i.ToListUsers())
	}

	c.HTML(200, "list-users.html", gin.H{
		"page":   page,
		"result": results,
		"next":   nextPage,
	})
}
