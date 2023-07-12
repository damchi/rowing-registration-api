package services

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"rowing-registation-api/api/models"
	"rowing-registation-api/pkg/logger"
)

type UserService struct {
	Ctx         context.Context
	UserManager models.UserManager
}

func GetUserService(ctx context.Context, db *gorm.DB) UserService {
	userManager := models.GetUserManager(ctx, db)
	return UserService{
		ctx,
		userManager,
	}
}

func (us UserService) FindByEmail(email string) bool {

	user, err := us.UserManager.FindByEmail(email)
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("Find By email: %v", err))
		return false
	}
	if user.Email != "" {
		return true
	}
	return false
}

func (us UserService) SaveUser(param models.User) (int64, error) {

	result, err := us.UserManager.SaveUser(param)
	if err != nil {
		return 0, err
	}

	return result, nil
}
