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
		client := configs.ConnectDB()
		collection := client.Database("tutorial").Collection("tutorial_collection")

		ID := c.Param("id")
		println(ID)
		var tutorial models.Tutorial
		// defer cancel()
		objId, _ := primitive.ObjectIDFromHex(ID)

		// println("objId", objId)
		if err := c.BindJSON(&tutorial); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		println("title:", tutorial.Title)
		println("description:", tutorial.Description)

		// //validate the request body
		// if err := c.BindJSON(&tutorial); err != nil {
		//     c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		//     return
		// }

		// //use the validator library to validate required fields
		// if validationErr := validate.Struct(&user); validationErr != nil {
		//     c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		//     return
		// }

		// update := bson.M{"title": tutorial.Title, "description": tutorial.Description}
		// println("Updated:", update)
		// coll.FindOne(context.TODO(), filter).Decode(&result)
		var result models.Tutorial
		filter := bson.M{"_id": objId}
		err := collection.FindOne(c, filter).Decode(&result)
		// result, err := collection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			// c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			panic(err)
			// return
		}
		println(result.Title)
		println(result.Description)
		// res := map[string]interface{}{"data": result}

		//get updated user details
		// var updatedTutorial models.Tutorial
		// println("MC:" , result.MatchedCount)
		// if result.MatchedCount == 1 {
		//     err := collection.FindOne(c, bson.M{"id": objId}).Decode(&updatedTutorial)
		//     if err != nil {
		//         // c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		// 		panic(err)
		//         // return
		//     }
		// }

		// c.IndentedJSON(http.StatusOK, gin.H{"Title": updatedTutorial.Title, "Description": updatedTutorial.Description, "ID": updatedTutorial.ID})
		// output, err := json.MarshalIndent(updatedTutorial, "", "    ")
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// fmt.Printf("%s\n", output)
	}
}

func DeleteTutorials() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// Determine detele target by PK and dlete
		c.IndentedJSON(http.StatusOK, "OK")
		println("deleteTutorials")
	}
}
