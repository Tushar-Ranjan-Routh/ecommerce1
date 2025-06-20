package main

import (
	"log"
	"net/http"
	"order-service/handlers"
	"order-service/kafka"
)

func main() {
	kafka.InitKafka()

	http.HandleFunc("/place-order", handlers.PlaceOrder)
	log.Println("🚀 Order Service running on :8082")
	log.Fatal(http.ListenAndServe("0.0.0.0:8082", nil))
}
