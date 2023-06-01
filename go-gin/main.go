package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/models"
	"net/http"

	"go-gin/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	router.GET("/api/tutorials", getAllTutorials)
	router.PUT("/api/tutorials", insertTutorials)
	router.POST("/api/tutorials", updateTutorials)
	router.DELETE("/api/tutorials", deleteTutorials)
	router.Run("localhost:8080")
}

func getAllTutorials(c *gin.Context) {

	client := configs.ConnectDB()
	collection := client.Database("tutorial").Collection("tutorial_collection")
	filter := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	var results []models.Tutorial
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var tutorialList []models.Tutorial
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
		tutorialList = results
	}

	c.IndentedJSON(http.StatusOK, tutorialList)
	println("getAllTutorials")
}

func insertTutorials(c *gin.Context) {
	// TODO
	// Get Request date from body and Insert them to collenction

	// Parse JSON from Request Body
	var json models.Tutorial
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Echo Reqested Data
	println(json.Description)

	c.IndentedJSON(http.StatusOK, gin.H{"Title": json.Title, "Description": json.Description, "ID": json.ID})

	println("insertTutorials")
}

func updateTutorials(c *gin.Context) {
	// TODO
	// Determine update target by PK and update
	c.IndentedJSON(http.StatusOK, "OK")
	println("updateTutorials")
}

func deleteTutorials(c *gin.Context) {
	// TODO
	// Determine detele target by PK and dlete
	c.IndentedJSON(http.StatusOK, "OK")
	println("deleteTutorials")
}
