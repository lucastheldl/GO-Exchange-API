package application

import (
	"context"
	"go-api/product/domain"
)


func CreateProductUseCase(ctx context.Context, input domain.ProductInput, repo domain.ProductRepository) (domain.ProductInput, error) {
	product := domain.ProductInput{
		Name:        input.Name,
		Description: input.Description,
		ImgUrl:      input.ImgUrl,
		UserId:      input.UserId,
	}

	err := repo.Save(ctx, product)
	return product, err
}
func GetAllProductsUseCase(ctx context.Context, repo domain.ProductRepository) ([]*domain.Product, error) {
	
	products,err := repo.GetAll(ctx)

	return products, err
}