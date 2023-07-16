package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TutorialRoute(router *gin.Engine) {
	router.GET("/api/tutorials", controllers.GetAllTutorials())
	router.PUT("/api/tutorials", controllers.InsertTutorials())
	router.POST("/api/tutorials/:id", controllers.UpdateTutorials())
	router.DELETE("/api/tutorials/:id", controllers.DeleteTutorials())
}
