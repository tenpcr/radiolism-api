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
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBannerList(c *gin.Context) {

	queryLimitStr := c.Query("limit")
	limit := 1
	currentTime := time.Now()

	if queryLimitStr != "" {
		parsedLimit, err := strconv.Atoi(queryLimitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterMongo := bson.D{
		{"$match", bson.D{
			{"date_start", bson.D{{"$lte", currentTime}}},
			{"date_end", bson.D{{"$gte", currentTime}}},
		}},
	}

	pipeline := mongo.Pipeline{
		filterMongo,
		bson.D{{"$limit", limit}},
	}

	cursor, err := models.BannerCollection().Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch radio list"})
		return
	}

	var banners []models.Banner

	if err := cursor.All(ctx, &banners); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse radio list"})
		return
	}

	defer cursor.Close(ctx)

	for i, banner := range banners {
		if banner.Image != "" {
			banners[i].Image = fmt.Sprintf("%s/%s", os.Getenv("ENDPOINT_IMG_UPLOAD"), banner.Image)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": banners})

}
