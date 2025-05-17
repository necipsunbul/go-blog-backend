package v1

import (
	"goblog/lib/app/controllers"

	"github.com/gin-gonic/gin"
)

func DefultRoutes(r *gin.RouterGroup) {
	group := r.Group("/")
	group.GET("/", controllers.Index)
	group.GET("/ping", controllers.Ping)
}
