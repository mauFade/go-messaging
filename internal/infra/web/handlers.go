package web

import (
	"encoding/json"
	"net/http"

	"github.com/mauFade/go-messaging/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase
}

func NewProductsHandlers(createProductuseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductuseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(writer http.ResponseWriter, request *http.Request) {
	var input usecase.CreateProductInputDTO

	err := json.NewDecoder(request.Body).Decode(&input)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductUseCase.Execute(input)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusCreated)

	// Retorna o json
	json.NewEncoder(writer).Encode(output)
}

func (p *ProductHandlers) ListProductHandler(writer http.ResponseWriter, request *http.Request) {
	output, err := p.ListProductsUseCase.Execute()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)

	// Retorna o json
	json.NewEncoder(writer).Encode(output)
}
