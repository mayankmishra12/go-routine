package model

import (
	"time"
)

type Customer struct {
	BankrupcyIndicatorFlag *bool   `json:"bankrupcyIndicatorFlag"`
	CompanyName            string `json:"companyName"`
	DateOfBirth            *string `json:"dateOfBirth"`
	FirstName              string `json:"firstName"`
	LastName               string `json:"lastName"`
	LegalEntityStage       string `json:"legalEntityStage"`
	LegalEntityType        string `json:"legalEntityType"`
	LegalEntityID         int    `gorm:"column:legal_entity_id;not null;AUTO_INCREMENT; primaryKey;"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}