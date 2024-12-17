package main

import (
	"fmt"

	"github.com/hype08/go-proj/internal/config"
	"github.com/hype08/go-proj/internal/database"
	"github.com/hype08/go-proj/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.NewServerConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config.")
	}

	server := server.NewServer(config)

	config.Log().Apply()

	db, err := database.NewDatabase(config.DatabaseUrl())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database.")
	}

	err = database.MigrateUp(db, config.MigrationDir())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database.")
	}

	txm := database.NewTxManager(db.Pool)
	log.Printf("Tx manager initialized.: %#v", txm)

	// hook up domains

	address := fmt.Sprintf(":%d", config.Port())
	log.Printf("Starting server on %s", address)
	err = server.Run(address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server.")
	}
}
