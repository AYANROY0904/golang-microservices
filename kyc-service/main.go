package main

import (
	"kyc-service/routes"
	"log"
	"shared/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//utils.InitSentry("kyc-service")
	r := gin.Default()

	config := utils.LoadConfig()
	utils.ConnectDB(config)

	routes.SetupRoutes(r)
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
