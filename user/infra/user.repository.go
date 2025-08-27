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

func (h UserRepo) Save(ctx context.Context, input domain.UserInput) error{
	_, err := h.Conn.Exec(ctx, `
	INSERT INTO users (email, password) 
	VALUES ($1, $2, $3, $4);
`, input.Email, input.Password)

return err

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