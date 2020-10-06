package main


import (
	_ "UsaProject/repositories"
	"github.com/gin-gonic/gin"
	"UsaProject/Controller"
)
func main() {
	router:= setUpRoutes()
	router.Run()
}

func setUpRoutes () *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/customer")
	{
		v1.POST("/", Controller.InsertCustomer)
		v1.POST("/search", Controller.SearchCustomer)
		v1.GET("/get", Controller.GetCustomerByLegalEntityID)
		v1.PUT("/", Controller.UpdateCustomerByLegalEntityID)
		v1.DELETE("/", Controller.DeleteCustomerByByLegalEntityID)
	}
	return router
}