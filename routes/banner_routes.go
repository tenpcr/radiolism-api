package routes

import (
	"radiolism_api/controllers"

	"github.com/gin-gonic/gin"
)

func BannerRoutes(router *gin.Engine) {
	banner := router.Group("/banner")
	{
		banner.GET("/list", controllers.GetBannerList)
	}

}
