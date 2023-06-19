package handler

import (
	"auth/config"
	"auth/repository"
	"auth/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
)

func Start() *httptest.Server {
	err := godotenv.Load("../../.env.testing")
	if err != nil {
		log.Fatal("Error loading .env.testing file")
	}

	cfg := config.NewConfig()

	userRepo := repository.NewUserRepositoryMock()
	tokenService := service.NewTokenService(cfg)

	authHandler := NewAuthHandler(cfg)
	userHandler := NewUserHandler(userRepo, tokenService)

	mux := http.NewServeMux()
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/profile", userHandler.GetProfile)

	return httptest.NewServer(mux)
}
