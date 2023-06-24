package collections

import (
	"context"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type CartCollection struct {
	collection *mongo.Collection
}

func CartCollectionInit(database *mongo.Database) *CartCollection {
	return &CartCollection{
		collection: database.Collection("cart"),
	}
}

func (CartCollection *CartCollection) GetCartItems() ([]models.Cart, error) {
	var cartItems []models.Cart
	cursor, err := CartCollection.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &cartItems)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
