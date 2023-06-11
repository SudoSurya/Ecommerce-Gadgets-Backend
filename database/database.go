package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	Database *mongo.Database
}

func NewDatabaseConnection() (*DatabaseConnection, error) {

	clientOption := options.Client().ApplyURI("mongodb+srv://suryanarayana7826:pass@cluster0.jf2vfhi.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		return nil, err
	}
	database := client.Database("test")
	log.Print("Connected to MongoDB!")
	return &DatabaseConnection{
		Database: database,
	}, nil
}
