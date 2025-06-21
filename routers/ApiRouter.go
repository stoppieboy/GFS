package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("something route"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("something else route"))
}

func ApiRoutes(r *mux.Router) {
	r.HandleFunc("/upload", uploadHandler)
	r.HandleFunc("/download", downloadHandler)
}
