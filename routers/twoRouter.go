package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func TwoRoutes(r *mux.Router) {
	r.HandleFunc("/nothing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nothing route"))
	})
}
