package main

import (
	Providers "user_management.be_webapi/internal/providers"

	Controllers "user_management.be_webapi/pkg/controllers"
)

func main() {

	server := new(Providers.Server)
	app := server.App()

	userController := new(Controllers.UserController)
	userController.Run(app)

	server.Listen(app)
}
