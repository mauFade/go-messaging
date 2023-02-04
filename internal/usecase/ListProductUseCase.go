package usecase

import "github.com/mauFade/go-messaging/internal/entity"

type ListProductsOutputDTO struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{ProductRepository: productRepository}
}

func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDTO, error) {
	products, err := u.ProductRepository.Find()

	if err != nil {
		return nil, err
	}

	var productOutput []*ListProductsOutputDTO

	for _, product := range products {
		productOutput = append(productOutput, &ListProductsOutputDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productOutput, nil
}
