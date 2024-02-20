package model

import "github.com/kakuzops/api-code-go/src/config/rest_err"

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser(UserDomain) *rest_err.RestErr
	UpdateUser(string, UserDomain) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
