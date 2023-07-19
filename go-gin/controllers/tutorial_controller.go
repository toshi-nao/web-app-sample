package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/configs"
	"go-gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTutorials() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		}
		tutorialList = results
		c.IndentedJSON(http.StatusOK, tutorialList)
		println("getAllTutorials")
	}
}

func CreateTutorial() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := configs.ConnectDB()
		collection := client.Database("tutorial").Collection("tutorial_collection")

		var tutorial models.Tutorial

		// get tutorial data from request
		if err := c.BindJSON(&tutorial); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		newTutorial := bson.M{
			"title":       tutorial.Title,
			"description": tutorial.Description,
			"published":   tutorial.Published,
			"createdAt":   time.Now().Format(time.RFC3339),
			"updatedAt":   time.Now().Format(time.RFC3339),
		}

		//Insert a single tutorial
		_, err := collection.InsertOne(c, newTutorial)
		if err != nil {
			panic(err)
		}

	}
}

func UpdateTutorial() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := configs.ConnectDB()
		collection := client.Database("tutorial").Collection("tutorial_collection")

		ID := c.Param("id")
		println(ID)
		var tutorial models.Tutorial
		// defer cancel()
		objId, _ := primitive.ObjectIDFromHex(ID)

		if err := c.BindJSON(&tutorial); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		filter := bson.M{"_id": objId}
		update := bson.M{
			"$set": bson.M{
				"title":       tutorial.Title,
				"description": tutorial.Description,
				"published":   tutorial.Published,
				"updatedAt":   time.Now().Format(time.RFC3339),
			},
		}

		//update a single tutorial by _id
		updateResult, err := collection.UpdateOne(c, filter, update)
		if err != nil {
			panic(err)
		}

		//get updated tutorial details
		var updatedTutorial models.Tutorial
		if updateResult.MatchedCount == 1 {
			err := collection.FindOne(c, bson.M{"_id": objId}).Decode(&updatedTutorial)
			if err != nil {
				panic(err)
			}
		}

		output, err := json.MarshalIndent(updatedTutorial, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

func DeleteTutorial() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := configs.ConnectDB()
		collection := client.Database("tutorial").Collection("tutorial_collection")

		ID := c.Param("id")
		objId, _ := primitive.ObjectIDFromHex(ID)

		_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tutorial"})
			return
		}

		c.IndentedJSON(http.StatusOK, "OK")
	}
}
