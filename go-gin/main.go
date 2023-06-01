package main

import (
	"go-gin/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{
		"http://localhost:8081",
	}

	//run database
	configs.ConnectDB()

	router.Use(cors.New(conf))

	router.Run("localhost:8080")
}
