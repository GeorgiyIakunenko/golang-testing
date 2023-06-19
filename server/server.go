package server

import (
	"auth/config"
	"auth/repository"
	"auth/server/handler"
	"auth/service"
	"fmt"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {

	userRepo := repository.NewUserRepository()
	tokenService := service.NewTokenService(cfg)

	authHandler := handler.NewAuthHandler(cfg)
	userHandler := handler.NewUserHandler(userRepo, tokenService)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)

	fmt.Printf("server is running")
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
