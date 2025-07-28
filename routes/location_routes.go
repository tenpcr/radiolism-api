package routes

import (
	"radiolism_api/controllers"

	"github.com/gin-gonic/gin"
)

func LocationRoutes(router *gin.Engine) {
	location := router.Group("/location")
	{
		location.GET("/country", controllers.GetLocationCountry)
		location.GET("/state", controllers.GetLocationState)
	}

}
