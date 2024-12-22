package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hype08/go-proj/internal/errorh"
	"github.com/hype08/go-proj/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	SearchUsers(ctx context.Context, params *models.UserSearchParams) ([]*models.User, error)
	UpdateUser(ctx context.Context, params *models.UserUpdateParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db sqlx.ExtContext
}

func NewUserRepository(db sqlx.ExtContext) UserRepository {
	if db == nil {
		log.Fatal().Err(errorh.ErrNilPointer).Send()
	}
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, `
		INSERT INTO users (
			id,
			name,
			email,
			address,
			created_at,
			modified_at
		) VALUES (
			$1, $2, $3, $4, NOW(), NOW()
		) RETURNING id`,
		uuid.New(),
		params.Name,
		params.Email,
		params.Address,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &id, nil
}

func (r *userRepository) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	err := sqlx.GetContext(ctx, r.db, user, `
		SELECT 
			id,
			name,
			email,
			address,
			created_at,
			modified_at
		FROM users 
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *userRepository) SearchUsers(ctx context.Context, params *models.UserSearchParams) ([]*models.User, error) {
	query := `
		SELECT 
			id,
			name,
			email,
			address,
			created_at,
			modified_at
		FROM users
		WHERE 1=1`

	args := []interface{}{}
	argCount := 0

	if params.ID != nil {
		argCount++
		query += fmt.Sprintf(" AND id = $%d", argCount)
		args = append(args, params.ID)
	}

	if params.Email != nil {
		argCount++
		query += fmt.Sprintf(" AND email = $%d", argCount)
		args = append(args, params.Email)
	}

	if params.Text != nil {
		argCount++
		query += fmt.Sprintf(" AND (name ILIKE $%d OR address ILIKE $%d)", argCount, argCount)
		searchPattern := "%" + *params.Text + "%"
		args = append(args, searchPattern)
	}

	// Add pagination
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", params.Count, params.Page*params.Count)

	var users []*models.User
	err := sqlx.SelectContext(ctx, r.db, &users, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, params *models.UserUpdateParams) error {
	query := `
		UPDATE users 
		SET 
			modified_at = NOW()`

	args := []interface{}{params.ID}
	argCount := 1

	if params.Name != nil {
		argCount++
		query += fmt.Sprintf(", name = $%d", argCount)
		args = append(args, *params.Name)
	}
	if params.Email != nil {
		argCount++
		query += fmt.Sprintf(", email = $%d", argCount)
		args = append(args, *params.Email)
	}
	if params.Address != nil {
		argCount++
		query += fmt.Sprintf(", address = $%d", argCount)
		args = append(args, *params.Address)
	}

	query += ` WHERE id = $1 RETURNING id`

	err := r.db.QueryRowxContext(ctx, query, args...).Scan(&params.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := r.db.QueryRowxContext(ctx, `
		DELETE FROM users 
		WHERE id = $1
		RETURNING id`,
		id,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
