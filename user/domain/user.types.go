package domain

import (
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user UserInput) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}


type UserInput struct {
    Email  string  `json:"email"`
    Password string `json:"password"`
}

type User struct {
    ID    int
    Name  string
    Email string
    Password string
}