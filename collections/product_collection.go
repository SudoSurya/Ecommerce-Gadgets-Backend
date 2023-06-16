package collections

import (
	"context"
	"errors"
	"math"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Brand       string             `json:"brand"`
	Price       float64            `json:"price"`
	Rating      float64            `json:"rating"`
	Type        string             `json:"type"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
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

func (ProductCollection *ProductCollection) GetProductById(productId string) (models.Product, error) {
	var product models.Product
	objID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return models.Product{}, err
	}

	filter := bson.M{"_id": objID}
	err = ProductCollection.collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// doc not found
			return models.Product{}, errors.New("Product not found")
		}
		return models.Product{}, err
	}
	return product, nil
}

func (ProductCollection *ProductCollection) GetProductByPage(page int, pageSize int) ([]Product, error) {
	mockProducts, err := ProductCollection.GetAllProducts()
	if err != nil {
		return nil, err
	}

	startIndex := (page - 1) * pageSize
	endIndex := math.Min(float64(startIndex+pageSize), float64(len(mockProducts)))

	return mockProducts[startIndex:int(endIndex)], nil

}
