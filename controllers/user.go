package controllers

import (
	"api/logics"
	"api/models"
	"api/response"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService logics.IUserService
	response    response.IResponse
}

type IUserController interface {
	Get() http.HandlerFunc
	GetByID() http.HandlerFunc
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
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
			return
		}

		c.response = response.Response{Data: getUsers}
		c.response.Success(w)
	}
}

func (c UserController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		getUsers, err := c.userService.GetByID(uint(utils.StringToInt(id)))

		if err != nil {
			c.response = response.Response{}
			c.response.Error(w, err.Error())
			return
		}

		c.response = response.Response{Data: getUsers}
		c.response.Success(w)
	}
}

func (c UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		json.NewDecoder(r.Body).Decode(&user)

		err := c.userService.Create(&user)

		if err != nil {
			c.response = response.Response{}
			c.response.Error(w, err.Error())
			return
		}

		c.response = response.Response{}
		c.response.Success(w)
	}
}

func (c UserController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		user := models.User{}
		json.NewDecoder(r.Body).Decode(&user)

		err := c.userService.Update(
			uint(utils.StringToInt(id)),
			&user,
		)

		if err != nil {
			c.response = response.Response{}
			c.response.Error(w, err.Error())
			return
		}

		c.response = response.Response{}
		c.response.Success(w)
	}
}

func (c UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := c.userService.Delete(uint(utils.StringToInt(id)))
		if err != nil {
			c.response = response.Response{}
			c.response.Error(w, err.Error())
			return
		}

		c.response = response.Response{}
		c.response.Success(w)
	}
}
