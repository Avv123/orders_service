package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DB *mongo.Client

func ConnectDatabase() error {
	log.Println("Starting database connection process...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v\n", err)
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URI is empty or not set in environment")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Attempting to connect to MongoDB...")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	log.Println("Connection established, attempting to ping database...")
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	DB = client
	return nil
}
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		panic("Database connection is not initialized. Call ConnectDatabase first.")
	}
	return DB.Database("user_oms").Collection(collectionName)
}
