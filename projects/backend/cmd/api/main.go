package main

import (
	"fmt"
	"log"

	"github.com/hype08/go-proj/internal/config"
	"github.com/hype08/go-proj/internal/database"
	"github.com/hype08/go-proj/internal/server"
)

func main() {
	config, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("Failed to load server config.")
	}

	server := server.NewServer(config)

	db, err := database.NewDatabase(config.DatabaseUrl())
	if err != nil {
		log.Fatalf("Failed to connect database. %v\nDatabase URL: %s", err, config.DatabaseUrl())
	}
	log.Println("Successfully connected to database.")

	txm := database.NewTxManager(db.Pool)
	log.Printf("Tx manager initialized.: %#v", txm)

	address := fmt.Sprintf(":%d", config.Port())
	log.Printf("Starting server on %s", address)
	err = server.Run(address)
	if err != nil {
		log.Fatal("Failed to start server.")
	}
}
