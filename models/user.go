package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username        string             `json:"username" bson:"username,omitempty"`
	Email           string             `json:"email" bson:"email,omitempty"`
	Password        string             `json:"password" bson:"password,omitempty"`
	ConfirmPassword string             `json:"confirmPassword" bson:"confirm_password,omitempty"`
}
