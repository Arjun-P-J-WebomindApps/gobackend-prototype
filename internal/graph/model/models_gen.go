// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Role     int32  `json:"role"`
	IsActive bool   `json:"isActive"`
}

type CreateUserOTPInput struct {
	UserID uuid.UUID `json:"user_id"`
}

type CreateUserSessionInput struct {
	UserID uuid.UUID `json:"user_id"`
}

type Mutation struct {
}

type Query struct {
}
