package v1

import (
	"goblog/lib/app/controllers"
	"goblog/lib/app/middlewares"
	"goblog/lib/app/validations/example"

	"github.com/gin-gonic/gin"
)

func DefultRoutes(r *gin.RouterGroup) {
	group := r.Group("/")
	group.GET("/", controllers.Index)
	group.GET("/ping", controllers.Ping)
	group.POST("/validationTest", middlewares.ValidatorMiddleware(example.CreateUserRequest{}), controllers.ValidatonTest)
}
