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

func newProductsHandlers(createProductuseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductuseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProducthandler(writer http.ResponseWriter, request *http.Request) {
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

	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(http.StatusCreated)

	// Retorna o json
	json.NewEncoder(writer).Encode(output)
}

func (p *ProductHandlers) ListProducthandler(writer http.ResponseWriter) {
	output, err := p.ListProductsUseCase.Execute()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("content-type", "application/json")

	writer.WriteHeader(http.StatusOK)

	// Retorna o json
	json.NewEncoder(writer).Encode(output)
}
