package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

var (
	basePath = ""
	jsonFilesFolderName = "storage"
)
func GetJsonFile(c * gin.Context){
	fileName := c.Query("filename")
	fmt.Println(fileName)
	path := filepath.Join(basePath,jsonFilesFolderName)
	filepath := filepath.Join(path,fileName)
	filepathwithextension := fmt.Sprintf("%s.json",filepath)
	fmt.Println(filepath)
	files,err := os.Open(filepathwithextension)
	if err!=nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
		fmt.Println("return error ")
	}
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.FileAttachment(filepathwithextension,filepath)
	fmt.Println(files)
}