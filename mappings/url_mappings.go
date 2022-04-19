package mappings

import (
	"learning/controllers"
	"learning/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine
var RouterAuth *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(middlewares.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		v1.GET("/users/", controllers.GetUser)
		v1.GET("/user/:id", controllers.GetUserADetail)
		v1.POST("/login/", controllers.Login)
	}

	// v2 of the API
	RouterAuth = gin.Default()
	Router.Use(middlewares.AuthorizeJWT())
	v2 := Router.Group("/v2")
	{
		v2.GET("/users/", controllers.GetUser)
		v2.GET("/user/:id", controllers.GetUserADetail)
	}
}
