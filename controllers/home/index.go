package home

import (
	"github.com/gin-gonic/gin"
)

// Home is ....
func Home(c *gin.Context) {

	c.HTML(200, "home.html", gin.H{
		"title": "Main website",
	})
}
