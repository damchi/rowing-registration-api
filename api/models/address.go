package models

import (
	"context"
	"gorm.io/gorm"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
	"strings"
)

type Address struct {
	gorm.Model
	Address1       string `gorm:"size:100;not null;" json:"address1"`
	Address2       string `gorm:"size:100;" json:"address2,omitempty"`
	Address3       string `gorm:"size:100;" json:"address3,omitempty"`
	AddressCity    string `gorm:"size:100;not null;" json:"address_city"`
	AddressState   string `gorm:"size:100;not null;" json:"address_state"`
	AddressCountry string `gorm:"size:100;not null;" json:"address_country"`
	PostCode       string `gorm:"size:100;not null;" json:"post_code"`
}

type AddressManager struct {
	db      *gorm.DB
	Context context.Context
}

func (Address) TableName() string {
	return "address"
}

func GetAddressManager(ctx context.Context, db *gorm.DB) AddressManager {
	if db != nil {
		return AddressManager{db, ctx}
	}
	return AddressManager{mysqlgorm.GetConnection(), ctx}
}

func (m AddressManager) LoadFromParam(param ClubRegistrationParam) (Address, error) {
	var address Address

	address.Address1 = strings.Trim(param.Address1, " ")
	address.Address2 = strings.Trim(param.Address2, " ")
	address.Address3 = strings.Trim(param.Address3, " ")
	address.AddressCity = strings.Trim(param.AddressCity, " ")
	address.AddressState = strings.Trim(param.AddressState, " ")
	address.AddressCountry = strings.Trim(param.AddressCountry, " ")
	address.PostCode = strings.Trim(param.PostCode, " ")

	return address, nil
}
