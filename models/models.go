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
	Facebook  string `bson:"facebook" json:"facebook" form:"facebook"`
	Instagram string `bson:"instagram" json:"instagram" form:"instagram"`
	Twitter   string `bson:"twitter" json:"twitter" form:"twitter"`
}

type Restaurant struct {
	Id          primitive.ObjectID `bson:"_id" json:"id" form:"id"`
	Name        string             `bson:"name" json:"name" form:"name"`
	Description string             `bson:"description" json:"description" form:"description"`
	Address     string             `bson:"address" json:"address" form:"address"`
	Phone       string             `bson:"phone" json:"phone" form:"phone"`
	Program     []Program          `bson:"program" json:"program" form:"program"`
	PostCode    int                `bson:"postCode"  json:"postCode" form:"postCode"`
	Logo        File               `bson:"logo" json:"logo" form:"logo"`
	Country     string             `bson:"country" json:"country" form:"country"`
	Email       string             `bson:"email" json:"email" form:"email"`
	Website     string             `json:"website" bson:"website" form:"website"`
	SocialProfile
	Latitude  float64   `bson:"latitude" json:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude"`
	DeleteAt  time.Time `bson:"deleteAt" json:"delete_at"`
}

type Program struct {
	StartAt time.Time `json:"startAt" bson:"startAt" form:"startAt"`
	EndAt   time.Time `json:"endAt" bson:"endAt" form:"endAt"`
	Day     int       `json:"day" bson:"day" form:"day"`
}

type ManagerDetails struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Restaurant Restaurant         `bson:"restaurant" json:"restaurant"`
	User       User          `bson:"user" json:"user"`
}