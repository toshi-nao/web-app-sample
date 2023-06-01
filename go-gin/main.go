package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tutorialList = []models.Tutorial{
	{Title: "Tutorial A", Description: "This is tutorial A", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29bc21eb646280e59c84"},
	{Title: "Tutorial B", Description: "This is tutorial B", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29c2acefe3e62a866731"},
	{Title: "Tutorial C", Description: "This is tutorial C", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f2a0eacefe3e62a866735"},
}

var tutorialPostList = []models.Tutorial{
	{Title: "Tutorial Post A", Description: "This is tutorial Post A", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "1"},
	{Title: "Tutorial Post B", Description: "This is tutorial Post B", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "2"},
	{Title: "Tutorial Post C", Description: "This is tutorial Post C", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "3"},
}

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8081",
	}
	router.Use(cors.New(config))
	router.GET("/api/tutorials", getAllTutorials)
	router.PUT("/api/tutorials", insertTutorials)
	router.POST("/api/tutorials", updateTutorials)
	router.DELETE("/api/tutorials", deleteTutorials)
	router.Run("localhost:8080")
}

func getAllTutorials(c *gin.Context) {

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    "admin",
		Username:      "root",
		Password:      "brHZ-!_rHAZF4xR2-EsRKx9e",
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:28001").SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
		tutorialList = results
	}

	defer client.Disconnect(ctx)
	println("main method")
	println(tutorialList[0].Description)

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
