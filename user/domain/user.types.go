package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	Save(ctx context.Context, user UserInput) (*User, error)
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
    CreatedAt time.Time `json:"created_at"`
}