package processor

import (
	"log"
	"shipping-service/models"
)

func ProcessOrder(order models.Order) {
	log.Printf("📦 Processing shipment for OrderID: %s (Product: %s, Qty: %d)\n",
		order.ID, order.ProductID, order.Quantity)

	// Simulate shipping logic — Save to DB or update status
	log.Println("✅ Shipment created and scheduled!")
}
