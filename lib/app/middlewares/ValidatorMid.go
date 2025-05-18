package middlewares

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"goblog/lib/core/responses"
)

func ValidatorMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		modelType := reflect.TypeOf(model)
		if modelType.Kind() != reflect.Ptr {
			modelValue := reflect.New(reflect.TypeOf(model))
			model = modelValue.Interface()
		}
		
		if err := c.ShouldBindBodyWith(model, binding.JSON); err != nil {
			responses.ValidationErrorResponse(c, err)
			c.Abort()
			return
		}
		c.Set("validatedData", model)
		c.Next()
	}
}
