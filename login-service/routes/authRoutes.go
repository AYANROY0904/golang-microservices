package routes

import (
	"login-service/controllers"
	"shared/utils"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	r.Use(utils.SentryMiddleware())

	r.POST("/login", controllers.Login)
	r.POST("/verify-otp", controllers.VerifyOTP)
}
