package main

import (
	"GoMux/config"
	"GoMux/handler"
	"GoMux/repository"
	"GoMux/routes"
	"GoMux/usecase"
	"log"
	"net/http"
)

func main() {
	// initialize log
	config.InitLogger()

	// connect db and migrate
	config.InitDB()
	config.InitMigrate()

	// initialize routes
	userRepo := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	r := routes.InitRoutes(userHandler)

	// run server
	port := config.AppConfig.Server.Port
	config.Logger.Info("starting server on port: ", port)
	log.Println("starting server on port:", port)
	http.ListenAndServe(":"+port, r)
}
