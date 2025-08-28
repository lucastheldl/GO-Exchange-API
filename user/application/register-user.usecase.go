package application

import (
	"context"
	"fmt"
	"go-api/user/domain"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserUseCase(ctx context.Context, input domain.UserInput, repo domain.UserRepository )(*domain.User, error){

	existingUser,err := repo.GetUserByEmail(ctx,input.Email);

	if err != nil && err != pgx.ErrNoRows {
        return nil, err
    }

	if existingUser != nil{
		return nil, fmt.Errorf("user with email %s already exists", input.Email)
	}

	//register user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	userData := domain.UserInput{
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	
	user, err := repo.Save(ctx, userData)
	if err != nil {
		return nil, fmt.Errorf("error registering user: %w", err)
	}
	
	return user, nil
}