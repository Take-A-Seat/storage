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

type Restaurant struct {
	Id              primitive.ObjectID `bson:"_id" json:"id" form:"id"`
	Name            string             `bson:"name" json:"name" form:"name"`
	Description     string             `bson:"description" json:"description" form:"description"`
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
	VisibleOnline   bool               `bson:"visibleOnline" json:"visibleOnline" form:"visibleOnline"`
	DeleteAt        time.Time          `bson:"deleteAt" json:"deleteAt"`
}

type Program struct {
	StartAt string `json:"startAt" bson:"startAt" form:"startAt" schema:"startAt"`
	EndAt   string `json:"endAt" bson:"endAt" form:"endAt" schema:"endAt"`
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
	Tables         []Table            `json:"tables"`
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
	RestaurantId primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	Pages        []Pages            `bson:"pages" json:"pages"`
}

type SpecificRestaurant struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}

type SpecificRestaurantRelation struct {
	Id                   primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantId         primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	SpecificRestaurantId primitive.ObjectID `bson:"specificRestaurantId" json:"specificRestaurantId"`
}

type TypeRestaurant struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}

type TypeRestaurantRelation struct {
	Id               primitive.ObjectID `bson:"_id" json:"id"`
	RestaurantId     primitive.ObjectID `bson:"restaurantId" json:"restaurantId"`
	TypeRestaurantId primitive.ObjectID `bson:"typeRestaurantId" json:"typeRestaurantId"`
}

type AvailableDataReservation struct {
	TimeString string    `json:"timeString"`
	DateTime   time.Time `json:"dateTime"`
}

type ProductReservation struct {
	ItemMenu
	Status string `bson:"status" json:"status"`
}

type Reservation struct {
	Id                   primitive.ObjectID   `bson:"_id" json:"id"`
	Persons              int                  `bson:"persons" json:"persons" `
	StartReservationDate time.Time            `bson:"startReservationDate" json:"startReservationDate"`
	EndReservationDate   time.Time            `bson:"endReservationDate" json:"endReservationDate"`
	RestaurantId    primitive.ObjectID   `bson:"restaurantId" json:"restaurantId"`
	Phone           string               `bson:"phone" json:"phone"`
	FirstName       string               `bson:"firstName" json:"firstName"`
	LastName        string               `bson:"lastName" json:"lastName"`
	Email           string               `bson:"email" json:"email"`
	Details         string               `bson:"details" json:"details"`
	Status          string               `bson:"status" json:"status"`
	TableId         []primitive.ObjectID `bson:"tableId" json:"tableId"`
	MessageToClient string               `bson:"messageToClient" json:"messageToClient"`
	Products        []ProductReservation `bson:"products" json:"products"`
	TotalToPay      float64              `bson:"totalToPay" json:"totalToPay"`
	NeedAssistance  bool                 `bson:"needAssistance" json:"needAssistance"`
	Code            string               `bson:"code" json:"code"`
}

type RestaurantWithDetails struct {
	RestaurantDetails Restaurant                   `json:"restaurantDetails"`
	ListSpecifics     []SpecificRestaurantRelation `json:"listSpecifics"`
	ListTypes         []TypeRestaurantRelation     `json:"listTypes"`
}

type ChartData struct {
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Avg  float64 `json:"avg"`
	Name string  `json:"name"`
}

type CharWithValue struct {
	Name string `json:"name"`
	Value float64 `json:"value"`
}

type StatisticReservations struct {
	Persons              []ChartData     `json:"persons"`
	TotalPay             []ChartData     `json:"totalPay"`
	TotalMoneyReceived   []CharWithValue     `json:"totalMoneyReceived"`
	NumberReservations   []CharWithValue `json:"numberReservations"`
	NumberPeopleReturned []CharWithValue `json:"numberPeopleReturned"`
	Declined             []CharWithValue `json:"declined"`
	Finished             []CharWithValue `json:"finished"`
}
