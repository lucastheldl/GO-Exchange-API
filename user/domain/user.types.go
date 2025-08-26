package domain

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user UserInput) error
}


type UserInput struct {
    Email  string  `json:"email"`
    Password string `json:"password"`
}