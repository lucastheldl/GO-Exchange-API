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

func (r *ProductRepo) GetAll(ctx context.Context) ([]*domain.Product, error) {
    rows, err := r.Conn.Query(ctx, `SELECT * FROM products`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*domain.Product
    for rows.Next() {
        var p domain.Product
        err := rows.Scan(&p.ID, &p.Name,&p.CreatedAt,&p.Description,&p.ImgUrl,&p.UpdatedAt,&p.UserId) 
        if err != nil {
            return nil, err
        }
        products = append(products, &p)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return products, nil
}