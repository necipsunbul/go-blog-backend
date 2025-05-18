package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}


func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}


func ErrorResponse(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}


func ValidationErrorResponse(c *gin.Context, err error) {
	errorResponse := Response{
		Success: false,
		Message: "Validation failed",
	}
	
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorResponse.Errors = buildValidationErrors(validationErrors)
	} else {
		errorResponse.Errors = gin.H{"general": err.Error()}
	}
	
	c.JSON(http.StatusBadRequest, errorResponse)
}


func buildValidationErrors(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for _, err := range errors {
		fieldName := err.Field()
		
		switch err.Tag() {
		case "required":
			result[fieldName] = fieldName + " is required"
		case "email":
			result[fieldName] = fieldName + " must be a valid email address"
		case "min":
			result[fieldName] = fieldName + " must be at least " + err.Param() + " characters long"
		case "max":
			result[fieldName] = fieldName + " must be at most " + err.Param() + " characters long"
		case "gte":
			result[fieldName] = fieldName + " must be greater than or equal to " + err.Param()
		case "lte":
			result[fieldName] = fieldName + " must be less than or equal to " + err.Param()
		default:
			result[fieldName] = err.Error()
		}
	}
	return result
}


func OkResponse(c *gin.Context, data interface{}) {
	SuccessResponse(c, http.StatusOK, "Ok", data)
}

func CreatedResponse(c *gin.Context, data interface{}) {
	SuccessResponse(c, http.StatusCreated, "Resource created successfully", data)
}

func BadRequestResponse(c *gin.Context, message string, errors interface{}) {
	ErrorResponse(c, http.StatusBadRequest, message, errors)
}

func UnauthorizedResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusUnauthorized, "Unauthorized access", nil)
}

func ForbiddenResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusForbidden, "Access forbidden", nil)
}

func NotFoundResponse(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

func ServerErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusInternalServerError, "Internal server error", gin.H{"error": err.Error()})
} 