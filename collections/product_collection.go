package collections

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ProductCollection struct {
	collection *mongo.Collection
}

func ProductCollectionInit(database *mongo.Database) *ProductCollection {
	return &ProductCollection{
		collection: database.Collection("products"),
	}
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	Type        string  `json:"type"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}

func (ProductCollection *ProductCollection) GetAllProducts() ([]Product, error) {
	var products []Product
	cursor, err := ProductCollection.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
