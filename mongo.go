package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(username string, password string, host string, database string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + "@" + host + "/" + database + "?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Couldn't connect to MongoDB!")
		return client, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func DisconnectFromDatabase(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Couldn't disconnect from MongoDB!")
		return
	}

	fmt.Println("Disconnected from MongoDB!")
	return
}
