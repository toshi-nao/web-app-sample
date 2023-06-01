package main

import (
	"go-gin/configs"
	"go-gin/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//run database
	configs.ConnectDB()

	router := gin.Default()
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{
		"http://localhost:8081",
	}
	router.Use(cors.New(conf))

	//routes
	routes.TutorialRoute(router)
	router.Run("localhost:8080")
}
