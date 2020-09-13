package services

import (
	Property "user_management.be_webapi/internal/property"

	Params "user_management.be_webapi/pkg/params"
)

type UserService struct{}

func (e *UserService) Search(params *Params.UserParams) *Property.ResultDataProperty {
	resultData := new(Property.ResultDataProperty)

	var userList []*Params.UserParams

	user := new(Params.UserParams)
	user.Id = "1"
	user.Name = "Adisak Boonmark"
	userList = append(userList, user)

	user = new(Params.UserParams)
	user.Id = "2"
	user.Name = "Arisa Boonmark"
	userList = append(userList, user)

	resultData.Status = true
	resultData.Code = 200

	resultData.Data = userList
	resultData.UserMessage = "Search"
	resultData.DeveloperMessage = "Search"

	return resultData
}

func (e *UserService) GetDetail(params *Params.UserParams) *Property.ResultDataProperty {
	resultData := new(Property.ResultDataProperty)
	return resultData
}
