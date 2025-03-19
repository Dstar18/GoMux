package main

import (
	"GoMux/config"
	"GoMux/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// connect db and migrate
	config.InitDB()
	config.InitMigrate()

	// run server
	r := mux.NewRouter()
	routes.InitRoutes(r)

	log.Println("Server Run ::3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
