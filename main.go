package main

import (
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-rest-api/domain"
	"github.com/go-rest-api/infrastructure/database"
	"github.com/go-rest-api/infrastructure/kafka"
	"github.com/go-rest-api/infrastructure/persistence"
	"github.com/go-rest-api/routes"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	database.ConnectDataBase()
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go com Kafka")
	go routes.HandleRequest()

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "customers",
		"auto.offset.reset": "earliest",
	}
	topics := []string{"customers"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	fmt.Println("Kafka consumer has started")
	go consumer.Consume(msgChan)
	kafkaWorker(msgChan)
}

func kafkaWorker(msgChan chan *ckafka.Message) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChan {
		var customer domain.Customer
		err := json.Unmarshal(msg.Value, &customer)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		repo := persistence.NewCustomerRepositoryPostgres(database.DB)
		err = repo.CreateCustomer(&customer)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Customer Created:", customer)
		}
	}
}
