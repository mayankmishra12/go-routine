package main


import (

	"github.com/gin-gonic/gin"
	"go-routine/JsonFileProj/controller"
)
func main() {
	router:= setUpRoutes()

	router.Run()
}

func setUpRoutes () *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/jsonData")
	{
		v1.GET("/", controller.GetJsonFile)
	}
	return router
}