package routers

import (
	"LATIHAN1/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/order/:id", controllers.GetOneOrder)
	router.GET("/order", controllers.GetAllOrder)
	router.POST("/order", controllers.CreateOrder)
	router.PUT("/order/:id", controllers.UpdateOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)
	
	return router
}