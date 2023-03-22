package app

import (
	"api/constants"
	"api/controllers"
)

type controllerRoutes struct {
	userController *controllers.UserController
}

func initControllers() *controllerRoutes {
	return &controllerRoutes{
		userController: controllers.InitUserController(nil, nil),
	}
}

func registerRoutes() {
	var (
		controllerList = initControllers()
	)

	noAuthRouter(controllerList)
}

func noAuthRouter(c *controllerRoutes) {
	// CRUD User
	router.HandleFunc("/users", c.userController.Get()).Methods(constants.GET)
	router.HandleFunc("/users/{id}", c.userController.GetByID()).Methods(constants.GET)
	router.HandleFunc("/users", c.userController.Create()).Methods(constants.POST)
	router.HandleFunc("/users/{id}", c.userController.Update()).Methods(constants.PUT)
	router.HandleFunc("/users/{id}", c.userController.Delete()).Methods(constants.DELETE)
}
