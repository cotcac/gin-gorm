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

	// home
	r.GET("/", home.Home)

	// Api auth required
	auth := r.Group("/api", users.CheckToken)
	{
		// users
		auth.POST("/users/", users.Insert)
		auth.POST("/users/login", users.Login)
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
