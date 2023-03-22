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

func InitUserController(userService logics.IUserService, res response.IResponse) *UserController {
	if utils.IsNil(userService) {
		userService = logics.InitUserService(nil)
	}

	if utils.IsNil(res) {
		res = response.InitResponse(response.Response{})
	}

	return &UserController{
		userService: userService,
		response:    res,
	}
}

func (c UserController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUsers, err := c.userService.Get()
		c.response = response.Response{Data: getUsers}
		c.response.Response(w, err)
	}
}

func (c UserController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		getUser, err := c.userService.GetByID(uint(utils.StringToInt(id)))
		c.response = response.Response{Data: getUser}
		c.response.Response(w, err)
	}
}

func (c UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		json.NewDecoder(r.Body).Decode(&user)

		err := c.userService.Create(&user)
		c.response.Response(w, err)
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

		c.response.Response(w, err)
	}
}

func (c UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := c.userService.Delete(uint(utils.StringToInt(id)))
		c.response.Response(w, err)
	}
}
