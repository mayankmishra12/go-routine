package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-routine/redis-proj/client"
	"net/http"
)

var (
	basePath = ""
	jsonFilesFolderName = "storage"
	rclient =  &client.Client{}
)
func AddData(c * gin.Context){
	cutomerName := c.Query("customer_name")
	radisClient:=  rclient.NewRedisClient()
	var value interface{}
	c.BindJSON(value)
	status := radisClient.Set(context.Background(),cutomerName,value,0)

	if status.Err() != nil{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": status.Err().Error()})
		return
	}

 c.JSON(http.StatusOK,gin.H{"status": http.StatusNotFound,"messager":"added data to the radis"})
}

