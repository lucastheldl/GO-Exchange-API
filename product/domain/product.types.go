package domain

import (
	"context"
)

type ProductRepository interface {
	Save(ctx context.Context, product ProductInput) error
	GetAll(ctx context.Context) ([]*Product,error)
}

type ProductInput struct {
    Name  string  `json:"name"`
    Description string `json:"description"`
    ImgUrl string `json:"img_url"`
    UserId int64 `json:"user_id"`
}
type Product struct {
    ID  string  `json:"id"`
    Name  string  `json:"name"`
    Description string `json:"description"`
    ImgUrl string `json:"img_url"`
    UserId int64 `json:"user_id"`
    CreatedAt int64 `json:"created_at"`
    UpdatedAt int64 `json:"updated_at"`
}