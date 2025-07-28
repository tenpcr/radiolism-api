package models

import (
	"os"
	"radiolism_api/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Banner struct {
	Name      string    `json:"name" bson:"name"`
	Image     string    `json:"image" bson:"image"`
	DateStart time.Time `json:"date_start" bson:"date_start"`
	DateEnd   time.Time `json:"date_end" bson:"date_end"`
	Url       string    `json:"url" bson:"url"`
}

func BannerCollection() *mongo.Collection {
	client := config.MongoClient
	return client.Database(os.Getenv("MONGODB_NAME")).Collection("banners")
}
