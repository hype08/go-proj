package graph

import (
	"github.com/hype08/go-proj/internal/architecture/services"
	"github.com/hype08/go-proj/internal/errorh"
	"github.com/rs/zerolog/log"
)

type Resolver struct {
	s *services.Services
}

func NewResolver(
	services *services.Services,
) *Resolver {
	if services == nil {
		log.Fatal().Err(errorh.ErrNilPointer).Send()
	}

	return &Resolver{
		services,
	}
}
