package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"rowing-registation-api/api"
	"rowing-registation-api/pkg/logger"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
	"rowing-registation-api/pkg/translator"
	"strconv"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = initGormMysql()
	if err != nil {
		log.Fatal(logger.ERROR, fmt.Sprintf("Fail to connect to gor Mysql: %v", err))
		return
	}

	transConfig := translator.Config{
		TranslationFolder: "data/translations",
	}
	translator.InitTranslator(transConfig)

	startWebServer()
}
func startWebServer() {
	r := gin.New()
	r.Use(gin.Recovery())

	logger.CustomLogFormat(r)
	api.RegisterRoutes(r)
	go logger.InitLogger(os.Getenv("LOG_LEVEL"))
	defer close(logger.LogCh)
	err := r.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
	if err != nil {
		return
	}
}

func initGormMysql() (*gorm.DB, error) {
	port, _ := strconv.Atoi(os.Getenv("DB_CONFIG_PORT"))
	cfg := mysqlgorm.Config{
		Host:                  os.Getenv("DB_CONFIG_HOST"),
		Port:                  port,
		Username:              os.Getenv("DB_CONFIG_USER"),
		Password:              os.Getenv("DB_CONFIG_PASSWORD"),
		Schema:                os.Getenv("DB_CONFIG_NAME"),
		ConnectionMaxLifetime: time.Minute * 5,
		MaxOpenConnection:     5,
		MaxIdleConnection:     0,
		SslEnabled:            true,
	}
	gormDb, err := mysqlgorm.InitMysqlGorm(cfg)
	if err != nil {
		return nil, err
	}
	return gormDb, nil
}
