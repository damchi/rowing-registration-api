package club

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rowing-registation-api/api/models"
	"rowing-registation-api/api/services"
	"rowing-registation-api/api/validations"
	"rowing-registation-api/pkg/apierror"
	headers "rowing-registation-api/pkg/header"
	"rowing-registation-api/pkg/logger"
)

func RegisterClub(c *gin.Context) {
	lang := headers.GetAcceptLanguage(c)
	var param models.ClubRegistrationParam

	err := c.BindJSON(&param)
	ctx := c.Request.Context()
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("RegisterClub BindJSON: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, apierror.CreateError(apierror.CodeValidationFailed, apierror.MsgValidationFailed, "en"))
		return
	}

	errMsg := validations.CreateClubValidation(ctx, param, lang)
	if errMsg.HasError {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": errMsg})
		return
	}

	result, err := services.GetClubService(ctx, nil).SaveClub(param)
	created := false
	if result == 1 {
		created = true
	}
	c.JSON(http.StatusOK, gin.H{"created": created})
}
