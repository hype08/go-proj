// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserInput struct {
	Address *int   `json:"address,omitempty"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateUserInput struct {
	Address *int    `json:"address,omitempty"`
	Email   *string `json:"email,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   *int      `json:"address,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
