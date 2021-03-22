package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cavdy-play/go_mongo/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/Produces", controllers.GetAllProduces)
	router.POST("/Produce", controllers.CreateProduce)
	router.GET("/Produce/:PRODUCEID", controllers.GetSingleProduce)
	router.PUT("/Produce/:PRODUCEID", controllers.UpdateProduce)
	router.DELETE("/Produce/:PRODUCEID", controllers.DeleteProduce)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
