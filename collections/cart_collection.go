package collections

import (
	"context"
	"errors"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (CartCollection *CartCollection) CreateCart(cart models.UserCart) error {
	result, err := CartCollection.collection.InsertOne(context.Background(), cart)
	if err != nil {
		return err
	}
	cart.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (CartCollection *CartCollection) GetCartItems(username string) ([]models.Cart, error) {
	var cartItems models.UserCart
	filter := bson.M{"username": username}
	err := CartCollection.collection.FindOne(context.Background(), filter).Decode(&cartItems)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// doc not found
			return []models.Cart{}, nil
		}
		return nil, err
	}
	return cartItems.Cart, nil
}

func (CartCollection *CartCollection) GetProductByID(username string, productId string) (models.Cart, error) {

	var product models.Cart
	var cartItems models.UserCart
	userFilter := bson.M{"username": username}
	err := CartCollection.collection.FindOne(context.Background(), userFilter).Decode(&cartItems)
	if err != nil {
		return models.Cart{}, nil
	}
	for _, item := range cartItems.Cart {
		if item.ID == productId {
			product = item
			break
		}
	}
	return product, nil
}
func (CartCollection *CartCollection) AddProductToCart(username string, product models.Cart) (models.UserCart, error) {
	userFilter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"cart": product}}
	_, err := CartCollection.collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return models.UserCart{}, err
	}

	updated := models.UserCart{}
	err = CartCollection.collection.FindOne(context.Background(), userFilter).Decode(&updated)
	if err != nil {
		return models.UserCart{}, err
	}

	return updated, nil
}

func (CartCollection *CartCollection) DeleteProductFromCart(username string, productId string) error {
	userFilter := bson.M{"username": username}
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	update := bson.M{"$pull": bson.M{"cart": bson.M{"_id": objId}}}
	updated, err := CartCollection.collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return err
	}
	if updated.ModifiedCount == 0 {
		return errors.New("Product not found in cart")
	}
	return nil

}

func (CartCollection *CartCollection) ClearCart(username string) error {
	userFilter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"cart": []models.Cart{}}}
	_, err := CartCollection.collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return err
	}
	return nil
}

func (CartCollection *CartCollection) IncrementProductQuantity(username string, productId string) error {
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	userFilter := bson.M{"username": username, "cart._id": objId}

	update := bson.M{"$inc": bson.M{"cart.$.quantity": 1}}
	updated, err := CartCollection.collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return err
	}
	if updated.ModifiedCount == 0 {
		return errors.New("Product not found in cart")
	}
	return nil
}

func (CartCollection *CartCollection) DecrementProductQuantity(username string, productId string) error {
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}
	userFilter := bson.M{"username": username, "cart._id": objId}

	update := bson.M{"$inc": bson.M{"cart.$.quantity": -1}}
	updated, err := CartCollection.collection.UpdateOne(context.Background(), userFilter, update)
	if err != nil {
		return err
	}
	if updated.ModifiedCount == 0 {
		return errors.New("Product not found in cart")
	}
	return nil
}
