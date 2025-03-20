package routes

import (
	"GoMux/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go Mux!"))
	})

	return r
}
