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
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	SearchUsers(ctx context.Context, params *models.UserSearchParams) ([]*models.User, error)
	UpdateUser(ctx context.Context, params *models.UserUpdateParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type domain struct {
	users repositories.UserRepository
}

func New(users repositories.UserRepository) Domain {
	if users == nil {
		log.Fatal().Err(errorh.ErrNilPointer).Send()
	}

	return &domain{
		users,
	}
}

func (d *domain) CreateUser(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error) {
	userID, err := d.users.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (d *domain) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := d.users.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *domain) SearchUsers(ctx context.Context, params *models.UserSearchParams) ([]*models.User, error) {
	users, err := d.users.SearchUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *domain) UpdateUser(ctx context.Context, params *models.UserUpdateParams) error {
	return d.users.UpdateUser(ctx, params)
}

func (d *domain) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return d.users.DeleteUser(ctx, id)
}
