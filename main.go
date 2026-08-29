package main

import (
	"log"
	"os"
	"radiolism_api/config"
	"radiolism_api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()
	config.ConnectMongo()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	routes.BannerRoutes(router)
	routes.StationRoutes(router)
	routes.LocationRoutes(router)


	log.Printf("Server is running on port %s", os.Getenv("PORT"))
	
	port := os.Getenv("PORT")

	if port == "" {
    port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
    log.Fatal(err)
	}
}
