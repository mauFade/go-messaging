package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/mauFade/go-messaging/internal/infra/akafka"
	"github.com/mauFade/go-messaging/internal/infra/repository"
	"github.com/mauFade/go-messaging/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *kafka.Message)

	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewRepository(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDTO{}

		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			fmt.Print(err)
		}

		_, err = createProductUseCase.Execute(dto)
	}

}
