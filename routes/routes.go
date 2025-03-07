package routes

import (
	newsInfra "main/newspaper/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	newsDeps := newsInfra.Init()
	newsGroup := router.Group("/news")
	newsInfra.Routes(newsGroup, newsDeps)
	return router

}
