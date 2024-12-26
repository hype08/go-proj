package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hype08/go-proj/internal/config"
	"github.com/hype08/go-proj/internal/errorh"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Config *config.ServerConfig
}

func NewServer(config *config.ServerConfig) *Server {
	if config == nil {
		log.Fatal().Err(errorh.ErrNilPointer).Send()
	}

	return &Server{
		Config: config,
	}
}

func (s *Server) Bootstrap() error {
	log.Info().Msg("Finished bootstrapping.")
	return nil
}

func (s *Server) Run(address string) error {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", s.getPing).Methods("GET")

	handler := cors.AllowAll().Handler(router)
	return http.ListenAndServe(address, handler)

	// TODO: setup graph handler
}
