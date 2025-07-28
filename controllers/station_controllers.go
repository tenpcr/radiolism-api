package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"radiolism_api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRadioList(c *gin.Context) {

	queryLimitStr := c.Query("limit")
	queryPageStr := c.Query("page")
	limit := 1
	page := 0

	if queryLimitStr != "" {
		parsedLimit, err := strconv.Atoi(queryLimitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if queryPageStr != "" {
		parsedPage, err := strconv.Atoi(queryPageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	queryCountryStr := c.Query("country")
	queryStateStr := c.Query("state")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filters := bson.A{}

	if queryCountryStr != "" {
		filters = append(filters, bson.D{{"place.country", queryCountryStr}})
	}

	if queryStateStr != "" {
		filters = append(filters, bson.D{{"place.state", queryStateStr}})
	}

	matchStage := bson.D{}
	if len(filters) > 0 {
		matchStage = bson.D{{"$match", bson.D{{"$and", filters}}}}
	} else {
		matchStage = bson.D{{"$match", bson.D{}}}
	}

	skip := (page - 1) * limit

	pipeline := mongo.Pipeline{
		matchStage,
		bson.D{{"$skip", skip}},
		bson.D{{"$limit", limit}},
	}

	cursor, err := models.StationCollection().Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch radio list"})
		return
	}
	defer cursor.Close(ctx)

	var stations []models.Station
	if err := cursor.All(ctx, &stations); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse radio list"})
		return
	}

	for i, station := range stations {
		if station.ImageUrl != "" {
			stations[i].ImageUrl = fmt.Sprintf("%s/%s", os.Getenv("ENDPOINT_IMG_UPLOAD"), station.ImageUrl)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": stations,
	})
}

func GetRadioListByCategoryId(c *gin.Context) {

	paramCategory := c.Param("category_id")
	queryLimitStr := c.Query("limit")
	queryPageStr := c.Query("page")

	page := 0
	limit := 1

	if queryLimitStr != "" {
		parsedLimit, err := strconv.Atoi(queryLimitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if queryPageStr != "" {
		parsedPage, err := strconv.Atoi(queryPageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if paramCategory == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Category ID is required"})
		return
	}

	categoryId, err := primitive.ObjectIDFromHex(paramCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	skip := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"category_id", categoryId}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "stations"},
			{"let", bson.D{{"station_id", "$station_id"}}},
			{"pipeline", bson.A{
				bson.D{{"$match", bson.D{
					{"$expr", bson.D{{"$eq", bson.A{"$_id", "$$station_id"}}}},
				}}},
				bson.D{{"$sort", bson.D{{"name", -1}}}},
			}},
			{"as", "station_data"},
		}}},
		bson.D{{"$unwind", "$station_data"}},
		bson.D{{"$replaceRoot", bson.D{
			{"newRoot", bson.D{
				{"$mergeObjects", bson.A{"$station_data", "$$ROOT"}},
			}},
		}}},
		bson.D{{"$project", bson.D{
			{"station_data", 0},
		}}},
		bson.D{{"$skip", skip}},
		bson.D{{"$limit", limit}},
	}

	cursor, err := models.StationGroupsCollection().Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch radio list"})
		return
	}
	defer cursor.Close(ctx)

	var stations []models.StationGroups
	if err := cursor.All(ctx, &stations); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse radio list"})
		return
	}

	for i, station := range stations {
		if station.Image != "" {
			stations[i].Image = fmt.Sprintf("%s/%s", os.Getenv("ENDPOINT_IMG_UPLOAD"), station.Image)
		}

		stations[i].ID = stations[i].StationID
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stations,
	})

}
