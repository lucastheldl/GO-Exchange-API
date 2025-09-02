package application

import (
	"context"
	"go-api/product/domain"
)


func GetAllProductsUseCase(ctx context.Context, repo domain.ProductRepository) ([]*domain.Product, error) {
	
	products,err := repo.GetAll(ctx)

	return products, err
}