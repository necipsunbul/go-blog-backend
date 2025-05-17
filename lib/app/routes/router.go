package routes

import (
	v1 "goblog/lib/app/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	v1.BuildRoutes(api)
}
