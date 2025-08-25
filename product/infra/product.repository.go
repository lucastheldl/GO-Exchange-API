package infra

import (
	"context"
	"go-api/product/domain"

	"github.com/jackc/pgx/v5"
)

type ProductRepo struct {
	Conn *pgx.Conn
}

func NewProductRepo(conn *pgx.Conn) *ProductRepo {
	return &ProductRepo{Conn: conn}
}

func (r *ProductRepo) Save(ctx context.Context, p domain.ProductInput) error {
	_, err := r.Conn.Exec(ctx, `
		INSERT INTO products (name, description, img_url, user_id) 
		VALUES ($1, $2, $3, $4);
	`, p.Name, p.Description, p.ImgUrl, p.UserId)

	return err
}