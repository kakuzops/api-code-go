package service

import (
	"github.com/kakuzops/api-code-go/src/config/rest_err"
	"github.com/kakuzops/api-code-go/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
