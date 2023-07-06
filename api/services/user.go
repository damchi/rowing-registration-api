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

func (us UserService) LoginUser(param models.UserLoginParam) (*models.User, string, error) {
	return us.UserManager.Login(param)
}
