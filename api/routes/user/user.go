package user

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

func RegisterUser(c *gin.Context) {
	lang := headers.GetAcceptLanguage(c)
	var param models.User

	err := c.BindJSON(&param)
	ctx := c.Request.Context()
	if err != nil {
		logger.Log(logger.ERROR, fmt.Sprintf("RegisterUser BindJSON: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, apierror.CreateError(apierror.CodeValidationFailed, apierror.MsgValidationFailed, "en"))
		return
	}

	errMsg := validations.CreateUserValidation(ctx, param, lang)
	if errMsg.HasError {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": errMsg})
		return
	}

	result, err := services.GetUserService(ctx, nil).SaveUser(param)
	created := false
	if result == 1 {
		created = true
	}
	c.JSON(http.StatusOK, gin.H{"created": created})
}



func Login(c *gin.Context) {
	lang := headers.GetAcceptLanguage(c)
	var param models.UserLoginParam

	errBind := c.BindJSON(&param)
	if errBind != nil {
		logger.Log(logger.WARNING, errBind)
		c.JSON(http.StatusBadRequest, apierror.CreateError(http.StatusBadRequest, apierror.MsgValidJsonBody, lang))
		return
	}

	errMsg := validations.LoginUserValidation(param, lang)
	if errMsg.HasError {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": errMsg})
		return
	}

	ctx := c.Request.Context()

	user, token, err := services.GetUserService(ctx, nil).LoginUser(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierror.CreateError(http.StatusBadRequest, apierror.MsgCredentialNotValid, lang))
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}
