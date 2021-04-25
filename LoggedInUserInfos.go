package storage

import (
	"encoding/json"
	"errors"
	"github.com/Take-A-Seat/storage/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func doRequestToIsAuthenticated(c *gin.Context, apiUrl string) (models.User, error) {
	var user models.User

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

	return user.Id, nil
}
