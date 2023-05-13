package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tutorial struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   bool   `json:"published"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	ID          string `json:"id"`
}

var tutorialList = []tutorial{
	{Title: "Tutorial A", Description: "This is tutorial A", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29bc21eb646280e59c84"},
	{Title: "Tutorial B", Description: "This is tutorial B", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f29c2acefe3e62a866731"},
	{Title: "Tutorial C", Description: "This is tutorial C", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "644f2a0eacefe3e62a866735"},
}

var tutorialPostList = []tutorial{
	{Title: "Tutorial Post A", Description: "This is tutorial Post A", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "1"},
	{Title: "Tutorial Post B", Description: "This is tutorial Post B", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "2"},
	{Title: "Tutorial Post C", Description: "This is tutorial Post C", Published: false,
		CreatedAt: "2023-05-01T02:53:48.690Z", UpdatedAt: "2023-05-01T02:53:48.690Z", ID: "3"},
}

func main() {
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

	var results []tutorial
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

	// /*
	//         List databases
	// */
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Println(databases)

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
	c.IndentedJSON(http.StatusOK, tutorialList)
	println("getAllTutorials")
}

func insertTutorials(c *gin.Context) {
	// TODO
	// Get Request date from body and Insert them to collenction
	c.IndentedJSON(200, "OK")
	println("insertTutorials")
}

func updateTutorials(c *gin.Context) {
	// TODO
	// Determine update target by PK and update
	c.IndentedJSON(200, "OK")
	println("updateTutorials")
}

func deleteTutorials(c *gin.Context) {
	// TODO
	// Determine detele target by PK and dlete
	c.IndentedJSON(200, "OK")
	println("deleteTutorials")
}
