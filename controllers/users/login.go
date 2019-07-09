package users

import (
	"fmt"
	"time"

	"../../models"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// MyCustomClaims is ...
type MyCustomClaims struct {
	Usernname string `json:"username"`
	jwt.StandardClaims
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var mySigningKey = []byte("secret")

//Login func
func Login(c *gin.Context) {
	db := models.DBConn()
	var json models.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	fmt.Print(json)
	if err := db.Where("email = ?", json.Email).First(&user).Error; err != nil {
		c.JSON(422, gin.H{"Error": err.Error()})
		return
	}
	// compare password

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		c.JSON(422, gin.H{"Error": err.Error()})
		return
	}
	Username := user.Name
	claims := MyCustomClaims{
		Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySigningKey)
	fmt.Println(tokenString, err)
	c.JSON(200, gin.H{
		"success": true,
		"token":   tokenString,
	})

}
