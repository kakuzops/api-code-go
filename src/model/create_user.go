package model

import (
	"github.com/kakuzops/api-code-go/src/config/logger"
	"github.com/kakuzops/api-code-go/src/config/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	return nil
}
