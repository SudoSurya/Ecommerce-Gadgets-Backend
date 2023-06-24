package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ID          string             `json:"id" bson:"id"`
	Name        string             `json:"name" bson:"name"`
	Brand       string             `json:"brand" bson:"brand"`
	Price       float64            `json:"price" bson:"price"`
	Rating      float64            `json:"rating" bson:"rating"`
	Type        string             `json:"type" bson:"type"`
	Image       string             `json:"image" bson:"image"`
	Description string             `json:"description" bson:"description"`
	Quantity    int64              `json:"quantity" bson:"quantity"`
}

type UserCart struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Cart     []Cart             `json:"cart" bson:"cart"`
}
