package repositories

import (
	config2 "UsaProject/config"
	"UsaProject/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"path/filepath"
)
var conFilePath = "../config.yaml"
const (
	configDir = ""
	configFile  = "config.yaml"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "user@1234"
	dbname   = "customer"
)
var db *gorm.DB
func init() {
	cfile := filepath.Join(configDir,configFile)
	config, err := config2.NewConfig(cfile)
	if err != nil {
		errMsg := fmt.Sprintf("error while reading config path")
		panic(errMsg)
		os.Exit(1)
	}
	//open a db connection
	fmt.Println(config)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.Database.Host, config.Database.User, config.Database.Databasename, config.Database.Password)
	db, err = gorm.Open("postgres",dbUri )
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&model.Customer{})
}

func InsertCustomer (customer *model.Customer) (*model.Customer,error){
	result := db.Create(customer)
	if result.Error != nil{
		fmt.Println("error while inserting customer details")
		return nil, result.Error
	}
	return customer,nil
}

func GetByLegalEntityID( legalEntityId int) (*model.Customer, error){
	customer := model.Customer{}
	res := db.Where("legal_entity_id = ?" , legalEntityId).First(&customer)
	if res.Error != nil{
		fmt.Println(res.Error)
		return nil, res.Error
	}

	return &customer, nil
}

func GetCustomerData (customerFilterData *model.Customer) ([]*model.Customer,error){
	customer := []*model.Customer{}
	db.Where(&customerFilterData).Find(&customer)
	return customer, nil

}

func UpdateCustomerDatalegalEntityId(legalEntityId int, customerData *model.Customer) (*model.Customer, error) {
	customer := model.Customer{}
	res := db.Where("legal_entity_id = ?" , legalEntityId).First(&customer)
	if res.Error != nil{
		fmt.Println(res.Error)
		return nil, res.Error
	}
	customer = getUpdateCustomerData(customer,*customerData)
	res =db.Model(&model.Customer{}).Where(`legal_entity_id =?`,legalEntityId).Update(&customerData)

	if res.Error != nil {
		return nil,res.Error
	}
	return customerData, nil
}
func DeleteCustomerDataByEntityId(legalEntityId int) ( error) {
	customer := model.Customer{}
	res := db.Where("legal_entity_id = ?" , legalEntityId).First(&customer)
	if res.Error != nil{
		fmt.Println(res.Error)
		return  res.Error
	}
	resp  := db.Delete(&customer);
 if resp.Error !=nil{
		return resp.Error
	}
	return nil
}
func getUpdateCustomerData(existingCustomerData model.Customer, updatedCustomerData model.Customer) model.Customer {
	if updatedCustomerData.BankrupcyIndicatorFlag != nil {
		existingCustomerData.BankrupcyIndicatorFlag = updatedCustomerData.BankrupcyIndicatorFlag
	}
	if updatedCustomerData.CompanyName != "" {
		existingCustomerData.CompanyName = updatedCustomerData.CompanyName
	}
	if updatedCustomerData.DateOfBirth != nil {
		existingCustomerData.DateOfBirth = updatedCustomerData.DateOfBirth
	}
	if updatedCustomerData.FirstName != "" {
		existingCustomerData.FirstName = updatedCustomerData.FirstName
	}
	if updatedCustomerData.LegalEntityStage != ""{
		existingCustomerData.LegalEntityStage = updatedCustomerData.LegalEntityStage
	}
	if updatedCustomerData.LegalEntityType != ""{
		existingCustomerData.LegalEntityType = ""
	}
	return existingCustomerData
}