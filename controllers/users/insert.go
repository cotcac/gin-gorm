package users

import (
	"github.com/gin-gonic/gin"
	"../../models"
)

func Insert(c *gin.Context) {
	db:= models.DBConn()
	db.DB().Ping()
	type User struct {
		ID   int
		Name string
	  }
	db.Create(&User{Name:"Kinny"})
	if err := db.Create(&User{Name:"Kinny 123"}).Error; err != nil {
		c.JSON(500, gin.H{
			"message":err,
		})
		return 
	 }

	c.JSON(200, gin.H{
		"message": "Inserted",
	})
}
