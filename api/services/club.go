package services

import (
	"context"
	"gorm.io/gorm"
	"rowing-registation-api/api/models"
	"strings"
)

type ClubService struct {
	Ctx            context.Context
	ClubManager    models.ClubManager
	UserManager    models.UserManager
	AddressManager models.AddressManager
}

func GetClubService(ctx context.Context, db *gorm.DB) ClubService {
	clubManager := models.GetClubManager(ctx, db)
	userManager := models.GetUserManager(ctx, db)
	addressManager := models.GetAddressManager(ctx, db)
	return ClubService{
		ctx,
		clubManager,
		userManager,
		addressManager,
	}
}

func (cs ClubService) SaveClub(param models.ClubRegistrationParam) (int64, error) {
	user, err := cs.UserManager.LoadFromParam(param)
	if err != nil {
		return 0, err
	}
	address, err := cs.AddressManager.LoadFromParam(param)
	club := models.Club{
		Name:    strings.Trim(param.ClubName, " "),
		Phone:   strings.Trim(param.Phone, " "),
		Address: address,
		Owner:   user,
	}

	result, err := cs.ClubManager.SaveClub(club)
	if err != nil {
		return 0, err
	}

	return result, nil
}
