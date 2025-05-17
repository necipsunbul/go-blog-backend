package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func BuildAppMode() string {
	if os.Getenv("GIN_MODE") == "release" {
		log.Println("Running in release mode")
		return gin.ReleaseMode
	}
	log.Println("Running in debug mode")
	return gin.DebugMode
}
