package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/mauFade/go-messaging/internal/infra/akafka"
	"github.com/mauFade/go-messaging/internal/infra/repository"
	"github.com/mauFade/go-messaging/internal/infra/web"
	"github.com/mauFade/go-messaging/internal/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewRepository(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productsHandlers := web.NewProductsHandlers(createProductUseCase, listProductsUseCase)

	router := chi.NewRouter()

	// Rotas
	router.Post("/products", productsHandlers.CreateProductHandler)
	router.Get("/products", productsHandlers.ListProductHandler)

	go http.ListenAndServe(":8081", router)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDTO{}

		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			fmt.Print(err)
		}

		_, err = createProductUseCase.Execute(dto)
	}

}
