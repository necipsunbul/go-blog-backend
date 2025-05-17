package main

import (
	"goblog/lib/app/routes"
	"goblog/lib/core/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	gin.SetMode(config.BuildAppMode())
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run()
}
