package main

import (
	"fmt"
	"log"

	"github.com/hype08/go-proj/internal/config"
	"github.com/hype08/go-proj/internal/server"
)

func main() {
	config, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("Failed to load server config.")
	}

	server := server.NewServer(config)

	address := fmt.Sprintf(":%d", config.Port())
	err = server.Run(address)
	if err != nil {
		log.Fatal("Failed to start server.")
	}
}
