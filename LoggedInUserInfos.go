package storage

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetails struct {
	UserId    string `json:"UserId" bson:"_id"`
	Email     string `json:"Email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}
func doRequestToIsAuthenticated(c *gin.Context, apiUrl string) (UserDetails, error) {
	var user UserDetails

	headers := Headers{}
	if c.Request.Header["Authorization"] != nil {
		headers = Headers{Authorization: c.Request.Header["Authorization"][0]}
	}
	requestObject := Request{
		Url:     apiUrl + "/auth/isAuthenticated",
		Method:  "GET",
		Headers: headers,
		//Body: nil
	}
	var responseBody interface{}

	_, err := MakeHttpRequest(requestObject, &responseBody)
	if err != nil {
		return user, errors.New("Err: getLoggedInUserId: " + err.Error())
	}

	data, _ := json.Marshal(responseBody)
	err = json.Unmarshal(data, &user)

	if err != nil {
		return user, errors.New("Err: getLoggedInUserId: " + err.Error())
	}
	return user, nil
}

func GetLoggedInUserId(c *gin.Context, apiUrl string) (primitive.ObjectID, error) {
	user, err := doRequestToIsAuthenticated(c, apiUrl)
	if err != nil {
		return primitive.ObjectID{}, err
	}

 	objIdUser,err :=primitive.ObjectIDFromHex(user.UserId)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return objIdUser, nil
}
