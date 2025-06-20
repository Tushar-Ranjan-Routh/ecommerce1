package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"order-service/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB connects to PostgreSQL and creates the orders table if it doesn't exist.
func InitDB() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "orders")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	for i := 1; i <= 10; i++ {
		DB, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("❌ Failed to connect to DB (attempt %d/10): %v", i, err)
			time.Sleep(time.Duration(i*2) * time.Second)
			continue
		}

		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(25)
		DB.SetConnMaxLifetime(5 * time.Minute)

		err = DB.Ping()
		if err == nil {
			break
		}
		log.Printf("❌ Ping DB error (attempt %d/10): %v", i, err)
		time.Sleep(time.Duration(i*2) * time.Second)
	}
	if err != nil {
		log.Fatalf("❌ Could not connect to DB after 10 attempts: %v", err)
	}

	createTable()
	log.Println("✅ Connected to PostgreSQL and table ready.")
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL,
		product_id TEXT NOT NULL,
		quantity INTEGER NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Failed to create orders table: %v", err)
	}
}

// SaveOrder inserts an order into the database
func SaveOrder(o models.Order) error {
	_, err := DB.Exec(
		`INSERT INTO orders (id, user_id, product_id, quantity, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		o.ID, o.UserID, o.ProductID, o.Quantity, o.CreatedAt,
	)
	return err
}

// getEnv is a helper to fetch environment variables with fallback
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
