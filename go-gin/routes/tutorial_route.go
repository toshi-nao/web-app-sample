package routes

import (
	"go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TutorialRoute(router *gin.Engine) {
	router.GET("/api/tutorials", controllers.GetAllTutorials())
	router.POST("/api/tutorials", controllers.CreateTutorial())
	router.PUT("/api/tutorials/:id", controllers.UpdateTutorial())
	router.DELETE("/api/tutorials/:id", controllers.DeleteTutorial())
}
