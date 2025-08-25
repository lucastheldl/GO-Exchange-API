package application

import (
	"go-api/product/domain"
)

type ProductRepository interface {
    Save(product domain.ProductInput) error
}
func CreateProductUseCase(input domain.ProductInput, repo ProductRepository) (domain.ProductInput, error){
 
	product := domain.ProductInput{
        Name:		 input.Name,
        Description: input.Description,
        ImgUrl:      input.ImgUrl,
        UserId:      input.UserId,
    }

    err := repo.Save(product)
    return product, err
}