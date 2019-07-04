package router

import (
	"../controllers/users"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r:= gin.Default()
	router := r.Group("/api") 
	{
		router.GET("/users/ping", users.Ping)
		router.GET("/users/", users.List)
		router.POST("/users/", users.Insert)


	}
	return r
}

func Router() *gin.Engine {
	return setupRouter()
}