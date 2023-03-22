package app

import (
	"api/app/config"
	"api/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

func StartApp() {
	config.SetupConfig()
	Connection()
	registerRoutes()

	//start server
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
			middleware.JsonContentTypeMiddleware(router),
		),
	)
}
