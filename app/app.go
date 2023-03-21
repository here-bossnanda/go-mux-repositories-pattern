package app

import (
	"api/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
}

func StartApp() {
	setupConfig()
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

func setupConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	envVariable := godotenv.Load(environmentPath)
	if envVariable != nil {
		log.Fatal("Error loading .env file")
	}
}
