package mysql_gorm

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Config struct {
	Host                  string
	Port                  int
	Schema                string
	Username              string
	Password              string
	ConnectionMaxLifetime time.Duration
	MaxIdleConnection     int
	MaxOpenConnection     int
	SslEnabled            bool
}

var gormDb *gorm.DB

func InitMysqlGorm(cfg Config) (*gorm.DB, error) {
	var err error

	var databaseUrlFormat string
	if cfg.SslEnabled {
		databaseUrlFormat = "%s:%s@tcp(%s:%d)/%s?tls=skip-verify&parseTime=True"
	} else {
		databaseUrlFormat = "%s:%s@tcp(%s:%d)/%s&parseTime=True"
	}

	databaseUrl := fmt.Sprintf(databaseUrlFormat, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema)
	gormDb, err = gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDb, nil
}

func GetConnection() *gorm.DB {
	return gormDb
}

func CheckConnection(ctx context.Context) bool {
	sqlDb, err := gormDb.DB()

	if err != nil {
		log.Println(err.Error())
		return false
	}
	err = sqlDb.Ping()

	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
