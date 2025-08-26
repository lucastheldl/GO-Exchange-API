package infra

import (
	"context"
	"go-api/user/domain"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	Conn *pgx.Conn
}

func NewProductRepo(conn *pgx.Conn) *UserRepo {
	return &UserRepo{Conn: conn}
}

func (h UserRepo) Save(ctx context.Context, p domain.UserRepository) error{


}