package main

import (
	"log"
	"user-profile-service/routes"
	"shared/utils"

	"github.com/gin-gonic/gin"
)


func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	config := utils.LoadConfig()
	utils.ConnectDB(config)
	
	routes.SetupRoutes(r)
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
