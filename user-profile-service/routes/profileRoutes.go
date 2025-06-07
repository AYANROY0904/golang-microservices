package routes

import (
	"user-profile-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/profile", controllers.HandleProfile)
}
