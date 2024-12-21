package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/hype08/go-proj/internal/architecture/repositories"
	"github.com/hype08/go-proj/internal/errorh"
	"github.com/hype08/go-proj/internal/models"
	"github.com/rs/zerolog/log"
)

type Domain interface {
	CreateUser(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error)
}

type domain struct {
	users repositories.UserRepository
}

func New(
	users repositories.UserRepository,
) Domain {
	if users == nil {
		log.Fatal().Err(errorh.ErrNilPointer).Send()
	}

	return &domain{
		users,
	}
}

func (d *domain) CreateUser(
	ctx context.Context,
	params *models.UserCreateParams,
) (*uuid.UUID, error) {
	userID, err := d.users.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	return userID, nil
}
