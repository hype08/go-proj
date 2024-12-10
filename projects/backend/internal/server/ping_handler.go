package server

import "net/http"

func (s *Server) getPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
