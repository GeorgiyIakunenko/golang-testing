package main

import (
	"auth/config"
	"auth/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.NewConfig()

	server.Start(cfg)
}
