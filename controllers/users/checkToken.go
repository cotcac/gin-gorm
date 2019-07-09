package users

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//ValidateToken is ...
func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Usernname
}

// CheckToken is ...
func CheckToken(c *gin.Context) {

	reqToken := c.Request.Header.Get("Authorization")
	// Split at the space
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	IsTokenValid, _ := ValidateToken(reqToken)
	if !IsTokenValid {
		c.Abort()
		c.JSON(301, gin.H{
			"message": "unAuth",
		})
		return
	}
	c.Next()
}
