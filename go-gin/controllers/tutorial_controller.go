package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-gin/configs"
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func InsertTutorials() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func UpdateTutorials() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// Determine update target by PK and update
		c.IndentedJSON(http.StatusOK, "OK")
		println("updateTutorials")
	}
}

func DeleteTutorials() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Determine detele target by PK and dlete

		// get parameter id from request parameter
		id := "644f29bc21eb646280e59c84"

		// change id to ObjectId
		// objId, err := primitive.ObjectIDFromHex(id)
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		//	return
		//}

		// delete document
		client := configs.ConnectDB()
		collection := client.Database("tutorial").Collection("tutorial_collection")

		_, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.IndentedJSON(http.StatusOK, "OK")
		println("deleteTutorials %s", id)
	}
}
