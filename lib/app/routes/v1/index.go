package v1

import (
	"github.com/gin-gonic/gin"
)

func BuildRoutes(r *gin.RouterGroup) {
	V1path := r.Group("/v1")
	DefultRoutes(V1path)
}
