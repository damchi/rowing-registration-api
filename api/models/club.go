package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"rowing-registation-api/pkg/logger"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
)

type Club struct {
	gorm.Model
	OwnerId   uint64  `json:"owner_id"`
	AddressId uint64  `json:"address_id"`
	Address   Address `gorm:"not null;foreignKey:AddressId;" json:"address"`
	Owner     User    `gorm:"not null;foreignKey:OwnerId; reference:UserId;" json:"user"`
	Phone     string  `gorm:"size:20;not null" json:"phone"`
	Name      string  `gorm:"size:255;not null;" json:"name"`
}

func (Club) TableName() string {
	return "club"
}

type ClubManager struct {
	db      *gorm.DB
	Context context.Context
}

func GetClubManager(ctx context.Context, db *gorm.DB) ClubManager {
	if db != nil {
		return ClubManager{db, ctx}
	}
	return ClubManager{mysqlgorm.GetConnection(), ctx}
}

func (m ClubManager) SaveClub(club Club) (int64, error) {

	result := m.db.Create(&club)

	if result.Error != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("Save club : %v", result.Error))
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
