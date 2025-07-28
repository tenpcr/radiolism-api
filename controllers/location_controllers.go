package controllers

import (
	"context"
	"net/http"
	"radiolism_api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLocationCountry(c *gin.Context) {

	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	pipeline := mongo.Pipeline{
		bson.D{{"$sort", bson.D{{"name", -1}}}},
	}

	cursor, err := models.LocationCountryCollection().Aggregate(ctx, pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch radio list"})
		return
	}

	defer cursor.Close(ctx)

	var location []models.LocationCountry

	if err := cursor.All(ctx, &location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse location country list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": location,
	})
}

func GetLocationState(c *gin.Context) {
	countryQuery := c.Query("country")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{}
	if countryQuery != "" {
		filter = bson.D{{"country", countryQuery}}
	}

	pipeline := mongo.Pipeline{
		bson.D{{"$match", filter}},
		bson.D{{"$sort", bson.D{{"label", 1}}}},
	}

	cursor, err := models.LocationStateCollection().Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch location states"})
		return
	}
	defer cursor.Close(ctx)

	var locations []models.LocationState // สมมติ struct สำหรับ State

	if err := cursor.All(ctx, &locations); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse location state list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locations,
	})
}
