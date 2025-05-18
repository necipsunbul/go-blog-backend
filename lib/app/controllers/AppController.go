package controllers

import (
	"log"

	"github.com/gin-gonic/gin"

	"goblog/lib/core/responses"
)

func Index(c *gin.Context) {
	responses.OkResponse(c, gin.H{
		"message": "Hello World",
	})
}

func Ping(c *gin.Context) {
	responses.OkResponse(c, gin.H{
		"message": "pong",
	})
}

func ValidatonTest(c *gin.Context) {
	log.Println("validatedData => ",c.MustGet("validatedData"))
	responses.OkResponse(c,  c.MustGet("validatedData"))
}
