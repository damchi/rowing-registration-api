package api

import (
	"github.com/gin-gonic/gin"
	"rowing-registation-api/api/middlewares"
	"rowing-registation-api/api/routes/club"
	"rowing-registation-api/api/routes/health"
	"rowing-registation-api/api/routes/user"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		//HEALTH
		api.GET("/health", health.CheckHealth)
		api.GET("/health/report", health.CheckHealthReport)

		api.POST("/register-club", middlewares.Language(), club.RegisterClub)
		api.POST("/register-user", middlewares.Language(), user.RegisterUser)
	}
}
