package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func somethingFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("something route"))
}

func somethingElseFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("something else route"))
}

func OneRoutes(r *mux.Router) {
	r.HandleFunc("/something", somethingFunc)
	r.HandleFunc("/something-else", somethingElseFunc)
}
