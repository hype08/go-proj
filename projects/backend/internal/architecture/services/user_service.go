package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/hype08/go-proj/internal/models"
)

func (s *Services) CreateUser(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error) {
	var userID *uuid.UUID

	err := s.txm.Transaction(ctx, "CreateUser", func(ctx context.Context) (err error) {
		userID, err = s.d.Users.CreateUser(ctx, params)
		return
	})
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (s *Services) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user *models.User

	err := s.txm.Transaction(ctx, "GetUser", func(ctx context.Context) (err error) {
		user, err = s.d.Users.GetUser(ctx, id)
		return
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Services) SearchUsers(ctx context.Context, params *models.UserSearchParams) ([]*models.User, error) {
	var users []*models.User

	err := s.txm.Transaction(ctx, "SearchUsers", func(ctx context.Context) (err error) {
		users, err = s.d.Users.SearchUsers(ctx, params)
		return
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Services) UpdateUser(ctx context.Context, params *models.UserUpdateParams) error {
	return s.txm.Transaction(ctx, "UpdateUser", func(ctx context.Context) error {
		return s.d.Users.UpdateUser(ctx, params)
	})
}

func (s *Services) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.txm.Transaction(ctx, "DeleteUser", func(ctx context.Context) error {
		return s.d.Users.DeleteUser(ctx, id)
	})
}
