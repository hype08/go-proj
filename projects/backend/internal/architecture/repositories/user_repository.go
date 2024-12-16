package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hype08/go-proj/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, params *models.UserUpdateParams) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	// query and exec with context
	db sqlx.ExtContext
}

func NewUserRepository(db sqlx.ExtContext) UserRepository {
	if db == nil {
		log.Fatal("Database nil pointer")
	}
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, params *models.UserCreateParams) (*uuid.UUID, error) {
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
		return nil, err
	}
	return &id, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
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
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
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
		WHERE email = $1`,
		email,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, params *models.UserUpdateParams) error {
	query := `
		UPDATE users 
		SET 
			modified_at = NOW()`

	args := []interface{}{params.ID}
	argCount := 1

	if params.Name != nil {
		argCount++
		query += `, name = $` + string(argCount)
		args = append(args, *params.Name)
	}
	if params.Email != nil {
		argCount++
		query += `, email = $` + string(argCount)
		args = append(args, *params.Email)
	}
	if params.Address != nil {
		argCount++
		query += `, address = $` + string(argCount)
		args = append(args, *params.Address)
	}

	query += ` WHERE id = $1 RETURNING id`

	return r.db.QueryRowxContext(ctx, query, args...).Scan(&params.ID)
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.QueryRowxContext(ctx, `
		DELETE FROM users 
		WHERE id = $1
		RETURNING id`,
		id,
	).Scan(&id)
}

