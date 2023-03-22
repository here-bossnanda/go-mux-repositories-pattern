package app

import (
	"api/controllers"
)

type controllerRoutes struct {
	userController *controllers.UserController
}

func initControllers() *controllerRoutes {
	return &controllerRoutes{
		userController: controllers.InitUserController(nil),
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
	router.HandleFunc("/users", c.userController.Get()).Methods("GET")
	router.HandleFunc("/users/{id}", c.userController.GetByID()).Methods("GET")
	router.HandleFunc("/users", c.userController.Create()).Methods("POST")
	router.HandleFunc("/users/{id}", c.userController.Update()).Methods("PUT")
	router.HandleFunc("/users/{id}", c.userController.Delete()).Methods("DELETE")
}
