package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	Role         string             `bson:"role" json:"role"`
	FirstName    string             `bson:"firstName" json:"firstName"`
	LastName     string             `bson:"lastName" json:"lastName"`
	Password     string             `bson:"password" json:"password"`
	Email        string             `bson:"email" json:"email"`
}
