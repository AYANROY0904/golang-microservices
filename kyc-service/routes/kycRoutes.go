package routes

import (
	"kyc-service/controllers"
	"shared/utils"
	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(utils.SentryMiddleware())

	r.POST("/kyc", controllers.HandleKYC)
}
