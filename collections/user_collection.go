package collections

import (
	"context"
	"errors"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserCollection struct {
	collection *mongo.Collection
}

func UserCollectionInit(database *mongo.Database) *UserCollection {
	return &UserCollection{
		collection: database.Collection("users"),
	}
}

func (userCollection *UserCollection) CreateUser(user models.User) (models.User, error) {
	result, err := userCollection.collection.InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (UserCollection *UserCollection) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := UserCollection.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// doc not found
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (UserCollection *UserCollection) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := UserCollection.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// doc not found
			return models.User{}, nil
		}
		return models.User{}, nil
	}
	return user, nil
}

func (UserCollection *UserCollection) GetUserByID(userId string) (models.User, error) {
	var user models.User
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.User{}, err
	}
	filter := bson.M{"_id": objId}
	err = UserCollection.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
