package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteByParam(DbUsername string, DbPassword string, DbHost string, filterId string, dbName string, collectionName string, deletedField string) error {
	client, err := Connect(DbUsername, DbPassword, DbHost)
	if err != nil {
		return err
	}

	fltId, _ := primitive.ObjectIDFromHex(filterId)
	filter := bson.M{
		"_id": bson.M{
			"$eq": fltId,
		},
	}
	update := bson.M{
		"$set": bson.M{
			deletedField: time.Now(),
		},
	}

	countersCollection := client.Database(dbName).Collection(collectionName)
	result := countersCollection.FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
