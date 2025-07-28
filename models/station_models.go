package models

import (
	"os"
	"radiolism_api/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Station struct {
	ObjectId  primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Detail    string             `json:"detail" bson:"detail"`
	StreamUrl string             `json:"stream_url" bson:"stream_url"`
	ImageUrl  string             `json:"image" bson:"image"`
	Frequency string             `json:"frequency" bson:"frequency"`
}

func StationCollection() *mongo.Collection {
	client := config.MongoClient
	return client.Database(os.Getenv("MONGODB_NAME")).Collection("stations")
}
