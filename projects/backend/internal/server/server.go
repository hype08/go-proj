package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hype08/go-proj/internal/config"
	"github.com/rs/cors"
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
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", s.getPing).Methods("GET")

	handler := cors.AllowAll().Handler(router)
	return http.ListenAndServe(address, handler)
}
