package main

import (
	"log"
	"login-service/routes"
	"shared/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	utils.InitSentry("login-service")
	config := utils.LoadConfig()
	utils.ConnectDB(config)

	r := gin.Default()

	routes.AuthRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
