package Controller

import (
	"UsaProject/model"
	"UsaProject/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InsertCustomer(c *gin.Context){
var customer model.Customer
c.BindJSON(&customer)
resp, err :=repositories.InsertCustomer(&customer)
if err != nil {
	fmt.Println("error while inserting customer data into database",err)
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
	return
}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})

}

func SearchCustomer(c *gin.Context) {
	var customer model.Customer
	c.BindJSON(&customer)
	resp, err :=repositories.GetCustomerData(&customer)
	if err != nil {
		fmt.Println("error while inserting customer data into database",err)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})

}

func GetCustomerByLegalEntityID(c * gin.Context){
	LegalEntityID := c.Query("legal_entity_id")
	if LegalEntityID == "" {
		fmt.Println("please pass the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return
	}
	LegalEntityIDToInt, err := strconv.Atoi(LegalEntityID)
	if err != nil{
		fmt.Println("please pass correct the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return

	}
	resp, err :=repositories.GetByLegalEntityID(LegalEntityIDToInt)
	if err != nil {
		fmt.Println("error while inserting customer data into database",err)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})
}

func UpdateCustomerByLegalEntityID(c * gin.Context){
	var customer model.Customer
	c.BindJSON(&customer)
	LegalEntityID := c.Query("legal_entity_id")
	if LegalEntityID == "" {
		fmt.Println("please pass the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return
	}
	LegalEntityIDToInt, err := strconv.Atoi(LegalEntityID)
	if err != nil{
		fmt.Println("please pass correct the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return

	}
	resp, err := repositories.UpdateCustomerDatalegalEntityId(LegalEntityIDToInt,&customer)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})
}

func DeleteCustomerByByLegalEntityID(c * gin.Context) {

	LegalEntityID := c.Query("legal_entity_id")
	if LegalEntityID == "" {
		fmt.Println("please pass the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return
	}
	LegalEntityIDToInt, err := strconv.Atoi(LegalEntityID)
	if err != nil{
		fmt.Println("please pass correct the legal_entity_id in query parameter")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": "please pass the legal_entity_id "})
		return

	}
	 err = repositories.DeleteCustomerDataByEntityId(LegalEntityIDToInt)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}