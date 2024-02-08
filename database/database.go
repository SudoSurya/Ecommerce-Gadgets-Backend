package database

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	Database *mongo.Database
}

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	//database = os.Getenv("DB_DATABASE")
)

func NewDatabaseConnection() (*DatabaseConnection, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database("test")
	log.Print("Connected to MongoDB!")

	return &DatabaseConnection{
		Database: database,
	}, nil
}
