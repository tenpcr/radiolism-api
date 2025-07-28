package models

import (
	"os"
	"radiolism_api/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StationGroups struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"category_id"`
	StationID  primitive.ObjectID `bson:"station_id" json:"station_id"`

	Name      string `bson:"name" json:"name"`
	Detail    string `bson:"detail" json:"detail"`
	StreamURL string `bson:"stream_url" json:"stream_url"`
	Image     string `bson:"image" json:"image"`
	Frequency string `bson:"frequency" json:"frequency"`
}

func StationGroupsCollection() *mongo.Collection {
	client := config.MongoClient
	return client.Database(os.Getenv("MONGODB_NAME")).Collection("station_groups")
}
