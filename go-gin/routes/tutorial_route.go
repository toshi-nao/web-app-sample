package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TutorialRoute(router *gin.Engine) {
	router.GET("/api/tutorials", controllers.GetAllTutorials())
	router.PUT("/api/tutorials", controllers.CreateTutorial())
	router.POST("/api/tutorials/:id", controllers.UpdateTutorial())
	router.DELETE("/api/tutorials/:id", controllers.DeleteTutorial())
}
