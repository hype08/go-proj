package models

import (
	"github.com/google/uuid"
)

type UserCreateParams struct {
	Email   string `validate:"required,email,max=50"`
	Name    string `validate:"required,max=50"`
	Address string `validate:"required,max=100"`
}

type UserUpdateParams struct {
	ID      uuid.UUID `validate:"required"`
	Email   *string   `validate:"omitempty,email,max=50"`
	Name    *string   `validate:"omitempty,max=50"`
	Address *string   `validate:"omitempty,max=100"`
}

type UserSearchParams struct {
	Page   int        `validate:"min=0,max=1000"`
	Count  int        `validate:"min=1,max=100"`
	ID     *uuid.UUID
	Email  *string    `validate:"omitempty,email,max=50"`
	Text   *string    `validate:"omitempty,max=50"`
}