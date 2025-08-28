package application

import (
	"context"
	"fmt"
	"go-api/user/domain"

	"go-api/utils"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
    User  *domain.User `json:"user"`
    Token string       `json:"token"`
}


func LoginUserUseCase(ctx context.Context, input domain.UserInput, repo domain.UserRepository )(*LoginResponse, error){

	existingUser,err := repo.GetUserByEmail(ctx,input.Email);

	if err != nil && err != pgx.ErrNoRows {
        return nil, err
    }

	if existingUser == nil{
		return nil, fmt.Errorf("user with email %s dont exists", input.Email)
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
	
	if passwordErr != nil{
		return nil, fmt.Errorf("invalid credentials");
	}
	//log user 

	token,tokenErr := utils.GenerateJWTToken(existingUser.ID)

	if tokenErr != nil{
		return nil, fmt.Errorf("error generating jwt token");
	}
    
    return &LoginResponse{
        User:  existingUser,
        Token: token,
    }, nil
}