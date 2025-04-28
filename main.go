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
	oneRouter := router.PathPrefix("/api/one").Subrouter()
	twoRouter := router.PathPrefix("/api/two").Subrouter()

	routers.OneRoutes(oneRouter)
	routers.TwoRoutes(twoRouter)

	http.ListenAndServe(":8080", router)
}