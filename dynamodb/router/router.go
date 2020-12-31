package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"dynamodb-crud/curd"
	"dynamodb-crud/middleware"
)


func Initroute() *mux.Router {
	router := 	mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	middleware.Routes(router)
	apirouter := router.PathPrefix("/api/v1").Subrouter()
	apirouter.Use(middleware.AlreadyLoggedIn)
	curd.Routescall(apirouter)
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to API")
}