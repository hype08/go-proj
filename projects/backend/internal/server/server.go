package server

import (
	"log"

	"github.com/hype08/go-proj/internal/config"
)

type Server struct {
	Config *config.ServerConfig
}

func NewServer(config *config.ServerConfig) *Server {
	if config == nil {
		log.Fatal("Server config nil pointer")
	}

	return &Server{
		Config: config,
	}
}

func (s *Server) Bootstrap() error {
	log.Println("Server bootstrapped")
	return nil
}

func (s *Server) Run(address string) error {
	return nil
}
