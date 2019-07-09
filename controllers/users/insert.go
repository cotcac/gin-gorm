package users

import (
	"fmt"

	"../../models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Insert user
func Insert(c *gin.Context) {
	db := models.DBConn()
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	hash, _ := HashPassword(json.Password)                                  // ignore error for the sake of simplicity
	user := models.User{Name: json.Name, Email: json.Email, Password: hash} //db.Create(&User{Name:"Kinny 123"})
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(user)
	c.JSON(200, gin.H{
		"success": user,
	})
}
