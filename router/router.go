package router

import (
	"../controllers/articles"
	"../controllers/categories"
	"../controllers/home"
	"../controllers/users"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")

	// home
	r.GET("/", home.Home)

	// Api auth required
	auth := r.Group("/api", users.CheckToken)
	{
		// users

		auth.PATCH("/users/edit/:id", users.Edit)
		auth.DELETE("/users/delete/:id", users.Delete)

		// article
		auth.POST("/articles/", articles.Insert)

		// category
		auth.POST("/categories/", categories.Insert)

	}
	router := r.Group("/api")
	{
		// router.Use(users.CheckToken)

		// USERS END POINT
		router.POST("/users/", users.Insert)
		router.POST("/users/login", users.Login)
		router.GET("/users/ping", users.CheckToken, users.Ping)
		router.GET("/users/", users.List)
		router.GET("/users/single/:id", users.Single)

		// ARTICLE ENDPOINT

		router.GET("/articles/", articles.List)
		router.GET("/articles/single/:id", articles.Single)

		// CATEGORY ENDPOINT.

		router.GET("/categories/", categories.List)

	}
	return r
}

// Router is ...
func Router() *gin.Engine {
	return setupRouter()
}
