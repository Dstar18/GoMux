package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/", Homehandler).Methods("GET")
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Selamat datang di Golang Mux!")
}
