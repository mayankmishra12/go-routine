package Controller

import (
	"UsaProject/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	newresp := resp.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(newresp.StatusCode)
	fmt.Println(newresp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	assert.Equal(t, resp.Code, 200)
}
func TestGetCustomerByLegalEntityID(t *testing.T) {
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

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	var respData map[string]model.Customer
	 json.Unmarshal([]byte(string(body)),&respData)
	 customerDataFromResp, ok := respData["data"]
	 if !ok {
	 	t.Error("invalid response")
	 }
	legalEntityID := customerDataFromResp.LegalEntityID
 legalEntityIDToStr  := strconv.Itoa(legalEntityID)
   fmt.Println(respData["data"].LegalEntityID)


	getReq, err := http.NewRequest("GET", "/api/customer/get",nil)
	if err != nil {
		fmt.Println(err)
	}
   getResp := httptest.NewRecorder()
	q := getReq.URL.Query()
	q.Add("legal_entity_id", legalEntityIDToStr)
	getReq.URL.RawQuery = q.Encode()
	testRouter.ServeHTTP(getResp, getReq)

	assert.Equal(t, getResp.Code, 200)
}


func TestSearchCustomer(t *testing.T) {
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

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	var respData map[string]model.Customer
	json.Unmarshal([]byte(string(body)),&respData)

	serchCustomerData := &model.Customer{
		CompanyName:            "test",

	}
	searchdata, _:= json.Marshal(serchCustomerData)
	searchReq, err := http.NewRequest("POST", "/api/customer/search", bytes.NewBuffer(searchdata))
	if err != nil {
		fmt.Println(err)
	}

	searchResp := httptest.NewRecorder()
	testRouter.ServeHTTP(searchResp, searchReq)
	assert.Equal(t, searchResp.Code, 200)
}

func  TestUpdateCustomerByLegalEntityID(t *testing.T) {
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

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	var respData map[string]model.Customer
	json.Unmarshal([]byte(string(body)),&respData)
	customerDataFromResp, ok := respData["data"]
	if !ok {
		t.Error("invalid response")
	}
	legalEntityID := customerDataFromResp.LegalEntityID
	legalEntityIDTOStr  := strconv.Itoa(legalEntityID)
	fmt.Println(respData["data"].LegalEntityID)
	updataedCustomer := &model.Customer{
		CompanyName:            "test",
		FirstName:              "updatedFirstName",
		LastName:               "updatedSecondName",
	}
	udata, _:= json.Marshal(updataedCustomer)
	updateReq, err := http.NewRequest("PUT", "/api/customer/", bytes.NewBuffer(udata))
	if err != nil {
		fmt.Println(err)
	}
	q := updateReq.URL.Query()
	q.Add("legal_entity_id", legalEntityIDTOStr)

	updateReq.URL.RawQuery = q.Encode()
	updatedResp := httptest.NewRecorder()

	testRouter.ServeHTTP(updatedResp, updateReq)
	assert.Equal(t,updatedResp.Code, 200)
}
func TestDeleteCustomerByByLegalEntityID(t *testing.T) {
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

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	var respData map[string]model.Customer
	json.Unmarshal([]byte(string(body)),&respData)
	customerDataFromResp, ok := respData["data"]
	if !ok {
		t.Error("invalid response")
	}
	legalEntityID := customerDataFromResp.LegalEntityID
	legalEntityIDToStr  := strconv.Itoa(legalEntityID)
	fmt.Println(respData["data"].LegalEntityID)
	deletereq, err := http.NewRequest("DELETE", "/api/customer/",nil)
	if err != nil {
		fmt.Println(err)
	}
	q := deletereq.URL.Query()
	q.Add("legal_entity_id", legalEntityIDToStr)

	deletereq.URL.RawQuery = q.Encode()
	deleteResp := httptest.NewRecorder()

	testRouter.ServeHTTP(deleteResp, deletereq)
	assert.Equal(t, deleteResp.Code, 200)
}


func setUpRoutes () *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/customer")
	{
		v1.POST("/", InsertCustomer)
		v1.POST("/search", SearchCustomer)
		v1.GET("/get", GetCustomerByLegalEntityID)
		v1.PUT("/", UpdateCustomerByLegalEntityID)
		v1.DELETE("/", DeleteCustomerByByLegalEntityID)
	}
	return router
}
