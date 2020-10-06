package repositories

import "UsaProject/model"

type repositories interface {
	InsertCustomer (customer *model.Customer) (*model.Customer,error)
	GetByLegalEntityID( legalEntityId int) (*model.Customer, error)
	GetCustomerData (customerFilterData *model.Customer) ([]*model.Customer,error)
	UpdateCustomerDatalegalEntityId(legalEntityId int, customerData *model.Customer) (*model.Customer, error)
	DeleteCustomerDataByEntityId(legalEntityId int) ( error)
}
