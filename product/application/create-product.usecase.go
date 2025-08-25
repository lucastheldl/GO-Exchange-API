package application

import (
	"context"
	"go-api/product/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, product domain.ProductInput) error
}

func CreateProductUseCase(ctx context.Context, input domain.ProductInput, repo ProductRepository) (domain.ProductInput, error) {
	product := domain.ProductInput{
		Name:        input.Name,
		Description: input.Description,
		ImgUrl:      input.ImgUrl,
		UserId:      input.UserId,
	}

	err := repo.Save(ctx, product)
	return product, err
}