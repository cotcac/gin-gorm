package home

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Home is ....
func Home(c *gin.Context) {

	c.HTML(200, "home.html", gin.H{
		"title": "Main website",
	})
}

// Header is ...
func Header(c *gin.Context) {

	c.HTML(200, "header.html", gin.H{
		"title": "Main website",
	})
}

// Body is ...
func Body(c *gin.Context) {
	time.Sleep(time.Second * 5)
	c.HTML(200, "body.html", gin.H{
		"title": "Main website",
	})
}

// Footer is ...
func Footer(c *gin.Context) {

	c.HTML(200, "footer.html", gin.H{
		"title": "Main website",
	})
}
