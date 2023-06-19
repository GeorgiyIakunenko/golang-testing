package main

import (
	"auth/config"
	"auth/server"
)

func main() {
	cfg := config.NewConfig()

	server.Start(cfg)
}
