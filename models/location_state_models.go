package models

import (
	"os"
	"radiolism_api/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type LocationState struct {
	Label string `json:"label" bson:"label"`
	Value string `json:"value" bson:"value"`
}

func LocationStateCollection() *mongo.Collection {
	client := config.MongoClient
	return client.Database(os.Getenv("MONGODB_NAME")).Collection("location_states")
}
