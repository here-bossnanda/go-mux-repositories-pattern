package app

import (
	"api/controllers"
	"api/middleware"
	"log"
	"net/http"
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
	//create router

	router.HandleFunc("/users", c.userController.Get()).Methods("GET")
	// router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	// router.HandleFunc("/users", createUser(db)).Methods("POST")
	// router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	// router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	//start server
	log.Fatal(http.ListenAndServe(":8001", middleware.JsonContentTypeMiddleware(router)))
}
