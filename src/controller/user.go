package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kakuzops/api-code-go/src/config/logger"
	"github.com/kakuzops/api-code-go/src/config/rest_err"
	validation "github.com/kakuzops/api-code-go/src/config/validate"
	request "github.com/kakuzops/api-code-go/src/controller/model"
	model "github.com/kakuzops/api-code-go/src/model"
	"github.com/kakuzops/api-code-go/src/model/service"
	"go.uber.org/zap"
)

func FingById(c *gin.Context) {
	err := rest_err.NewNotFoundError("Usuario não localizado")
	c.JSON(err.Code, err)
}

func CreateUser(c *gin.Context) {
	logger.Info("Init User Controller",
		zap.String("Jorney", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created Controller",
		zap.String("Jorney", "CreateUser"),
	)

	c.JSON(http.StatusOK, "")
}

func DeleteUser(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}

func GetAllUser(c *gin.Context) {}
