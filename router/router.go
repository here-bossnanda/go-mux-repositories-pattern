package router

import (
	"api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	//create router
	router := mux.NewRouter()

	//start server
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(router)))
}
