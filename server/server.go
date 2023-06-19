package server

import (
	"auth/config"
	"auth/repository"
	"auth/server/handler"
	"auth/service"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {

	userRepo := repository.NewUserRepository()
	tokenService := service.NewTokenService(cfg)

	authHandler := handler.NewAuthHandler(cfg)
	userHandler := handler.NewUserHandler(userRepo, tokenService)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
