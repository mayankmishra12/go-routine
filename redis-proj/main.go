package main

import (
	"github.com/gin-gonic/gin"
	"go-routine/redis-proj/controller"
)
func main() {
	router:= setUpRoutes()



	router.Run()
}

func setUpRoutes () *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/customer")
	{
		v1.POST("/", controller.AddData)

	}
	return router

}