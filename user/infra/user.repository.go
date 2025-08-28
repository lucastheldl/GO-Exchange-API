package infra

import (
	"context"
	"go-api/user/domain"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	Conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) *UserRepo {
	return &UserRepo{Conn: conn}
}

func (h UserRepo) Save(ctx context.Context, input domain.UserInput) (*domain.User, error) {
    var user domain.User
    err := h.Conn.QueryRow(ctx, `
        INSERT INTO users (email, password) 
        VALUES ($1, $2) 
        RETURNING id, email, password, created_at;
    `, input.Email, input.Password).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
    
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}

func (h UserRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error){
	var user domain.User


	row := h.Conn.QueryRow(ctx, `
	SELECT id, email, password FROM users u WHERE u.email = $1
	`, email)

	err := row.Scan(&user.ID, &user.Email,&user.Password)

	if err != nil {
		return nil, err 
	}

	return &user,nil

}