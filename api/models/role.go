package models

import (
	"context"
	"gorm.io/gorm"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
)

const ClubOwner = 1
const Coach = 2
const Athelete = 3
const Customer = 4

type Role struct {
	gorm.Model
	Name string `gorm:"size:255;not null;" json:"name"`
}

type RoleManager struct {
	db      *gorm.DB
	Context context.Context
}

func (Role) TableName() string {
	return "role"
}

func GetRoleManager(ctx context.Context, db *gorm.DB) RoleManager {
	if db != nil {
		return RoleManager{db, ctx}
	}
	return RoleManager{mysqlgorm.GetConnection(), ctx}
}
