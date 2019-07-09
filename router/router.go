package router

import (
	"../controllers/articles"
	"../controllers/categories"
	"../controllers/users"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		// USERS END POINT
		router.GET("/users/ping", users.CheckToken, users.Ping)
		router.GET("/users/", users.List)
		router.GET("/users/single/:id", users.Single)
		router.POST("/users/", users.Insert)
		router.POST("/users/login", users.Login)
		router.PATCH("/users/edit/:id", users.Edit)
		router.DELETE("/users/delete/:id", users.Delete)
		// ARTICLE ENDPOINT

		// router.GET("/users/ping", users.Ping)
		router.GET("/articles/", articles.List)
		router.GET("/articles/single/:id", articles.Single)
		router.POST("/articles/", articles.Insert)
		// router.PATCH("/users/edit/:id", users.Edit)
		// router.DELETE("/users/delete/:id", users.Delete)

		// CATEGORY ENDPOINT.

		router.POST("/categories/", categories.Insert)
		router.GET("/categories/", categories.List)

	}
	return r
}

func Router() *gin.Engine {
	return setupRouter()
}
