package routes

import (
	"radiolism_api/controllers"

	"github.com/gin-gonic/gin"
)

func StationRoutes(router *gin.Engine) {
	radio := router.Group("/stations")
	{
		radio.GET("/list", controllers.GetRadioList)
		radio.GET("/category/:category_id/list", controllers.GetRadioListByCategoryId)
	}
}
