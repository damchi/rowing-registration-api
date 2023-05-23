package services

import (
	"context"
	"gorm.io/gorm"
	"rowing-registation-api/api/models"
)

type AddressService struct {
	Ctx            context.Context
	AddressManager models.AddressManager
}

func GetAddressService(ctx context.Context, db *gorm.DB) AddressService {
	addressManager := models.GetAddressManager(ctx, db)
	return AddressService{
		ctx,
		addressManager,
	}
}
