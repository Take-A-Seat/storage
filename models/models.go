package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Role      string             `bson:"role" json:"role"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Password  string             `bson:"password" json:"password"`
	Email     string             `bson:"email" json:"email"`
}

type File struct {
	Name         string `bson:"name" json:"name"`
	OriginalName string `bson:"originalName" json:"originalName"`
	Path         string `bson:"path" json:"path"`
}

type SocialProfile struct {
	Facebook  string `bson:"facebook" json:"facebook"`
	Instagram string `bson:"instagram" json:"instagram"`
	Twitter   string `bson:"twitter" json:"twitter"`
}

type Restaurant struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Address     string             `bson:"address" json:"address"`
	Phone       string             `bson:"phone" json:"phone"`
	Program     []Program          `bson:"program" json:"program"`
	PostCode    int                `bson:"postCode"  json:"postCode"`
	Logo        File               `bson:"logo" json:"logo"`
	Country     string             `bson:"country" json:"country"`
	Email       string             `bson:"email" json:"email"`
	Website     string             `json:"website" bson:"website"`
	SocialProfile
	DeleteAt time.Time `bson:"deleteAt" json:"delete_at"`
}

type Program struct {
	StartAt time.Time
	EndAt   time.Time
	Day     int
}

type ManagerDetails struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Restaurant Restaurant         `bson:"restaurant" json:"restaurant"`
	User       User          `bson:"user" json:"user"`
}