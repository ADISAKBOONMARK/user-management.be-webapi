package services

import (
	Property "user_management.be_webapi/internal/property"

	Params "user_management.be_webapi/pkg/params"
)

type RegisterService struct{}

func (e *UserService) Register(params *Params.UserParams) *Property.ResultDataProperty {
	resultData := new(Property.ResultDataProperty)
	return resultData
}