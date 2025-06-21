package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stoppieboy/gfs/routers"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []byte("Shivam")
	w.Header().Set("Content-Type", "application/json")
	w.Write(users)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	users := []byte("Ayushi")
	w.Header().Set("Content-Type", "application/json")
	w.Write(users)
}

func main() {
	router := mux.NewRouter()
	ApiRouter := router.PathPrefix("/api").Subrouter()
	WebServiceRouter := router.PathPrefix("/web").Subrouter()

	routers.ApiRoutes(ApiRouter)
	routers.WebServiceRoutes(WebServiceRouter)

	http.ListenAndServe(":3000", router)
}