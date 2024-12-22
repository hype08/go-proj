package services

import (
	"github.com/hype08/go-proj/internal/architecture/domains"
	"github.com/hype08/go-proj/internal/errorh"
	"github.com/rs/zerolog/log"
)

type Services struct {
	txm TxManager
	d   *domains.Domains
}

func NewServices(
	txm TxManager,
	domains *domains.Domains,
) *Services {
	if txm == nil || domains == nil {
		log.Fatal().Err(errorh.ErrNilPointer)
	}

	return &Services{
		txm,
		domains,
	}
}
