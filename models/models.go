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
	Name         string `bson:"name" json:"name" form:"name"`
	OriginalName string `bson:"originalName" json:"originalName" form:"originalName"`
	Path         string `bson:"path" json:"path" form:"path"`
}

type Coordinates struct {
}

type Restaurant struct {
	Id              primitive.ObjectID `bson:"_id" json:"id" form:"id"`
	Name            string             `bson:"name" json:"name" form:"name"`
	Description     string             `bson:"description" json:"description" form:"description"`
	Address         string             `bson:"address" json:"address" form:"address"`
	Phone           string             `bson:"phone" json:"phone" form:"phone"`
	Program         []Program          `bson:"program" json:"program" form:"program" schema:"program"`
	PostCode        int                `bson:"postCode"  json:"postCode" form:"postCode"`
	Logo            File               `bson:"logo" json:"logo" form:"logo"`
	Country         string             `bson:"country" json:"country" form:"country"`
	City            string             `bson:"city" json:"city" form:"city"`
	Email           string             `bson:"email" json:"email" form:"email"`
	Website         string             `json:"website" bson:"website" form:"website"`
	Facebook        string             `bson:"facebook" json:"facebook" form:"facebook"`
	Instagram       string             `bson:"instagram" json:"instagram" form:"instagram"`
	Twitter         string             `bson:"twitter" json:"twitter" form:"twitter"`
	Province        string             `bson:"province" json:"province" form:"province"`
	StreetAndNumber string             `bson:"streetAndNumber" json:"streetAndNumber" form:"streetAndNumber"`
	Lng             float64            `bson:"lng" json:"lng" form:"lng"`
	Lat             float64            `bson:"lat" json:"lat" form:"lat"`
	DeleteAt        time.Time          `bson:"deleteAt" json:"delete_at"`
}

type Program struct {
	StartAt time.Time `json:"startAt" bson:"startAt" form:"startAt" schema:"startAt"`
	EndAt   time.Time `json:"endAt" bson:"endAt" form:"endAt" schema:"endAt"`
	Day     int       `json:"day" bson:"day" form:"day" schema:"day"`
	Close   bool      `bson:"close" json:"close" form:"close" schema:"close"`
}

type ManagerDetails struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Restaurant Restaurant         `bson:"restaurant" json:"restaurant"`
	User       User               `bson:"user" json:"user"`
}

type Area struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantId   primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	Name           string             `bson:"name" json:"name"`
	DisplayName    string             `bson:"displayName" json:"displayName"`
	Priority       int                `bson:"priority" json:"priority"`
	OnlineCapacity int                `bson:"onlineCapacity" json:"onlineCapacity"`
	NumberTables   int                `json:"numberTables"`
	MinPartySize   int                `bson:"minPartySize" json:"minPartySize"`
	Capacity       string             `json:"capacity"`
	MaxPartySize   int                `bson:"maxPartySize" json:"maxPartySize"`
	DeleteAt       time.Time          `bson:"deleteAt" json:"deleteAt"`
}

type Table struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	TableGroupId    primitive.ObjectID `bson:"tableGroupId" json:"tableGroupId"`
	AreaId          primitive.ObjectID `bson:"areaId" json:"areaId"`
	Number          int                `bson:"number" json:"number"`
	Priority        int                `bson:"priority" json:"priority"`
	AvailableOnline bool               `bson:"availableOnline" json:"availableOnline"`
	AvailableNow    bool               `bson:"availableNow" json:"availableNow"`
	MinPeople       int                `bson:"minPeople" json:"minPeople"`
	MaxPeople       int                `bson:"maxPeople" json:"maxPeople"`
	DeleteAt        time.Time          `bson:"deleteAt" json:"deleteAt"`
}

type GroupTables struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	DeleteAt time.Time          `bson:"deleteAt" json:"deleteAt"`
}

type ItemMenu struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Ingredients string             `bson:"ingredients" json:"ingredients"`
	Price       float64             `bson:"price" json:"price"`
}

type SectionMenu struct {
	TitleSection string             `json:"titleSection" bson:"titleSection"`
	Products     []ItemMenu         `json:"products" json:"products"`
}

type Pages struct {
	Sections []SectionMenu `bson:"sections" json:"sections"`
	Number   int           `bson:"number" json:"number"`
}

type Menu struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantId primitive.ObjectID             `bson:"restaurantId" json:"restaurantId"`
	Pages        []Pages            `bson:"pages" json:"pages"`
}
