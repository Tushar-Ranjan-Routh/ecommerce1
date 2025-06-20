package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"order-service/kafka"
	"order-service/models"
	"time"

	"github.com/google/uuid"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()

	go kafka.PublishOrder(order)

	log.Println("✅ Order received & published:", order.ID)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(order)
}
