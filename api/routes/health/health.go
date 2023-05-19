package health

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
	mysqlgorm "rowing-registation-api/pkg/mysql-gorm"
	"time"
)

const (
	StatusOk   string = "OK"
	StatusDown string = "DOWN"
)

type Manifest struct {
	Version string `json:"version"`
}

type HealthBasic struct {
	AppName           string `json:"APP_NAME"`
	CurrentSystemTime string `json:"CURRENT_SYSTEM_TIME"`
	Message           string `json:"MESSAGE"`
}

type HealthServices struct {
	Mysql string `json:"MYSQL"`
}

type HealthAdvanced struct {
	AppName           string         `json:"APP_NAME"`
	CurrentSystemTime string         `json:"CURRENT_SYSTEM_TIME"`
	Status            HealthServices `json:"STATUS"`
}

// CheckHealth API health
func CheckHealth(c *gin.Context) {
	ctx := c.Request.Context()
	statusCode := 200
	message := StatusOk

	if !checkConnectionToDatabase(ctx) {
		statusCode = 500
		message = StatusDown
	}

	c.JSON(statusCode, HealthBasic{
		AppName:           os.Getenv("APP_NAME"),
		CurrentSystemTime: time.Now().Format("2006-01-02 15:04:05"),
		Message:           message,
	})
}

// CheckHealthReport return a more detailled response about status of the services
func CheckHealthReport(c *gin.Context) {
	ctx := c.Request.Context()

	databaseStatus := StatusDown
	if checkConnectionToDatabase(ctx) {
		databaseStatus = StatusOk
	}

	c.JSON(200, HealthAdvanced{
		AppName:           os.Getenv("APP_NAME"),
		CurrentSystemTime: time.Now().Format("2006-01-02 15:04:05"),
		Status: HealthServices{
			Mysql: databaseStatus,
		},
	})
}

func checkConnectionToDatabase(ctx context.Context) bool {
	return mysqlgorm.CheckConnection(ctx)
}
