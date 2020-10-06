package Controller

import (
	"UsaProject/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestInsertCustomer(t *testing.T) {
	testRouter := setUpRoutes()
	custmorData := &model.Customer{
		CompanyName:            "test",
		FirstName:              "firstnametest",
		LastName:               "lastnametest",
		LegalEntityStage:       "first",
		LegalEntityType:        "asset",
	}
	data, _:= json.Marshal(custmorData)
	req, err := http.NewRequest("POST", "/api/customer/", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
func TestGetCustomerByLegalEntityID(t *testing.T) {
	testRouter := setUpRoutes()
	req, err := http.NewRequest("GET", "/api/customer/",nil)
	if err != nil {
		fmt.Println(err)
	}
   resp := httptest.NewRecorder()
	q := req.URL.Query()
	q.Add("legal_entity_id", "1")
	req.URL.RawQuery = q.Encode()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}


func TestSearchCustomer(t *testing.T) {
	testRouter := setUpRoutes()
	custmorData := &model.Customer{
		CompanyName:            "test",

	}
	data, _:= json.Marshal(custmorData)
	req, err := http.NewRequest("POST", "/api/customer/search", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp, 200)
}

func  TestUpdateCustomerByLegalEntityID(t *testing.T) {
	testRouter := setUpRoutes()
	custmorData := &model.Customer{
		CompanyName:            "test",
		FirstName:              "updatedFirstName",
		LastName:               "updatedSecondName",
	}
	data, _:= json.Marshal(custmorData)
	req, err := http.NewRequest("PUT", "/api/customer/", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("legal_entity_id", "1")

	req.URL.RawQuery = q.Encode()
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
func TestDeleteCustomerByByLegalEntityID(t *testing.T) {
	testRouter := setUpRoutes()

	req, err := http.NewRequest("DELETE", "/api/customer/",nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("legal_entity_id", "1")

	req.URL.RawQuery = q.Encode()
	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}


func setUpRoutes () *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/customer")
	{
		v1.POST("/", InsertCustomer)
		v1.POST("/search", SearchCustomer)
		v1.GET("/", GetCustomerByLegalEntityID)
		v1.PUT("/", UpdateCustomerByLegalEntityID)
		//	v1.DELETE("/:id", deleteTodo)
	}
	return router
}
