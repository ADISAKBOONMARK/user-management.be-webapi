package controllers

import (
	"github.com/gin-gonic/gin"

	Params "user_management.be_webapi/pkg/params"
	Models "user_management.be_webapi/pkg/models"
)

type UserController struct{}

func (e *UserController) Run(app *gin.Engine) {

	app.GET("/user/search", func(c *gin.Context) {

		params := new(Params.UserParams)

		model := new(Models.UserModel)

		resultData := model.Search(params)

		c.JSON(200, resultData)
	})

	app.GET("/user/get-detail", func(c *gin.Context) {

		params := new(Params.UserParams)

		model := new(Models.UserModel)

		resultData := model.GetDetail(params)

		c.JSON(200, resultData)
	})
}
