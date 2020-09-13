package models

import (
	Property "user_management.be_webapi/internal/property"

	Params "user_management.be_webapi/pkg/params"
	Services "user_management.be_webapi/pkg/services"
)

type UserModel struct{}

func (e *UserModel) Search(params *Params.UserParams) *Property.ResultDataProperty {
	service := new(Services.UserService)
	resultData := service.Search(params)
	return resultData
}

func (e *UserModel) GetDetail(params *Params.UserParams) *Property.ResultDataProperty {
	service := new(Services.UserService)
	resultData := service.GetDetail(params)
	return resultData
}