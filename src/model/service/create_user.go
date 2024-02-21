package service

import (
	"fmt"

	"github.com/kakuzops/api-code-go/src/config/logger"
	"github.com/kakuzops/api-code-go/src/config/rest_err"
	"github.com/kakuzops/api-code-go/src/model"
	"go.uber.org/zap"
)

var (
	UserDomain model.UserDomainInterface
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println()
	return nil
}
