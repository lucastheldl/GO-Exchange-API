package application

import (
	"context"
	"fmt"
	"go-api/user/domain"

	"github.com/jackc/pgx/v5"
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
}