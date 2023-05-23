package models

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"rowing-registation-api/pkg/logger"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
	"strings"
)

type User struct {
	gorm.Model
	Role      Role `json:"role"`
	RoleId    uint64
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255;not null;" json:"last_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

type UserManager struct {
	db      *gorm.DB
	Context context.Context
}

func (User) TableName() string {
	return "user"
}

func GetUserManager(ctx context.Context, db *gorm.DB) UserManager {
	if db != nil {
		return UserManager{db, ctx}
	}
	return UserManager{mysqlgorm.GetConnection(), ctx}
}

func (m UserManager) FindByEmail(email string) (*User, error) {
	var user User

	result := m.db.Where("state = ?", 1).Where("email = ?", email).Find(&user)

	if result.Error != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("Fail to retrieve user: %v", result.Error))
		return nil, result.Error
	}
	return &user, nil
}

func (m UserManager) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (m UserManager) LoadFromParam(param ClubRegistrationParam) (User, error) {
	var user User
	hashedPassword, err := m.Hash(param.Password)
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("RegisterClub Failed Hash Password: %v", err))
		return user, err
	}
	user.FirstName = strings.Trim(param.FirstName, " ")
	user.LastName = strings.Trim(param.LastName, " ")
	user.Password = string(hashedPassword)
	user.Email = strings.Trim(param.Email, " ")
	user.Role.ID = ClubOwner

	return user, nil
}
