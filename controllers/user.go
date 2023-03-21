package controllers

import (
	"api/logics"
	"api/response"
	"api/utils"
	"net/http"
)

type UserController struct {
	userService logics.IUserService
	response    response.IResponse
}

type IUserController interface {
	Get() http.HandlerFunc
}

func InitUserController(userService logics.IUserService) *UserController {
	if utils.IsNil(userService) {
		userService = logics.InitUserService(nil)
	}

	return &UserController{
		userService: userService,
	}
}

func (c UserController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUsers, err := c.userService.Get()

		if err != nil {
			c.response = response.Response{}
			c.response.Error(w, err.Error())
		}

		c.response = response.Response{Data: getUsers}
		c.response.Success(w)
	}
}
